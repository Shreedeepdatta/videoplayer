package mp4

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

type Box struct {
	Size         uint32
	Type         string
	ExtendedSize uint64
	Data         []byte
	Children     []*Box
}

func ReadBox(reader io.ReadSeeker) (*Box, error) {
	box := &Box{}

	var sizeBytes [4]byte
	if _, err := io.ReadFull(reader, sizeBytes[:]); err != nil {
		if err == io.EOF {
			return nil, err
		}
		return nil, fmt.Errorf("failed to read box size: %v", err)
	}

	box.Size = binary.BigEndian.Uint32(sizeBytes[:])

	var typeBytes [4]byte
	if _, err := io.ReadFull(reader, typeBytes[:]); err != nil {
		return nil, fmt.Errorf("failed to read box type: %v", err)
	}

	box.Type = string(typeBytes[:])

	headerSize := uint32(8)

	if box.Size == 1 {

		var extendedSizeBytes [8]byte
		if _, err := io.ReadFull(reader, extendedSizeBytes[:]); err != nil {
			return nil, fmt.Errorf("failed to read extended size: %v", err)
		}

		box.ExtendedSize = binary.BigEndian.Uint64(extendedSizeBytes[:])
		headerSize += 8
	} else if box.Size == 0 {

		currentPos, err := reader.Seek(0, io.SeekCurrent)
		if err != nil {
			return nil, fmt.Errorf("failed to get current position: %v", err)
		}

		endPos, err := reader.Seek(0, io.SeekEnd)
		if err != nil {
			return nil, fmt.Errorf("failed to seek to end: %v", err)
		}

		box.Size = uint32(endPos - currentPos + 8)

		_, err = reader.Seek(currentPos, io.SeekStart)
		if err != nil {
			return nil, fmt.Errorf("failed to restore position: %v", err)
		}
	}

	dataSize := box.Size - headerSize

	switch box.Type {
	case "moov", "trak", "mdia", "minf", "stbl", "udta", "dinf":

		endPos := uint32(0)
		if box.Size > 0 {
			endPos = box.Size - headerSize
		}

		for endPos > 0 {
			childBox, err := ReadBox(reader)
			if err != nil {
				if err == io.EOF {
					break
				}
				return nil, err
			}

			box.Children = append(box.Children, childBox)

			if childBox.Size <= endPos {
				endPos -= childBox.Size
			} else {
				break
			}
		}
	default:

		if dataSize > 0 {
			box.Data = make([]byte, dataSize)
			if _, err := io.ReadFull(reader, box.Data); err != nil {
				return nil, fmt.Errorf("failed to read box data: %v", err)
			}
		}
	}

	return box, nil
}

func ParseMP4File(file *os.File) ([]*Box, error) {
	var boxes []*Box

	for {
		box, err := ReadBox(file)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		boxes = append(boxes, box)
	}

	return boxes, nil
}

func PrintBoxStructure(boxes []*Box, indent string) {
	for _, box := range boxes {
		fmt.Printf("%s- %s (%d bytes)\n", indent, box.Type, box.Size)

		if len(box.Children) > 0 {
			PrintBoxStructure(box.Children, indent+"  ")
		}
	}
}
