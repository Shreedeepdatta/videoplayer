package mp4

import (
	"encoding/binary"
	"log/slog"
	"os"
)

func parseStblBox(file *os.File, size uint32) {
	slog.Info("Parsing stbl box", "size", size)
	endposition, _ := file.Seek(0, 1)
	endposition += int64(size - 8)
	for {
		var header Box
		err := binary.Read(file, binary.BigEndian, &header)
		if err != nil || int64(header.Size) < 0 {
			break
		}
		boxType := string(header.Type[:])
		switch boxType {
		case "stsd":
			parseStsdBox(file, header.Size)
		case "stts":
			parseSttsBox(file, header.Size)
		case "stsc":
			parseStscBox(file, header.Size)
		case "stsz":
			parseStszBox(file, header.Size)
		case "stco":
			parseStcoBox(file, header.Size)
		default:
			slog.Info("skipping box in stbl", "type", boxType, "size", header.Size)
		}

		if newposition, err := file.Seek(int64(header.Size)-8, 1); err != nil || newposition >= endposition {
			break
		}
	}
}
