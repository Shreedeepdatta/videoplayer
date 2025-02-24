package mp4

import (
	"encoding/binary"
	"log/slog"
	"os"
)

func parseMdiaBox(file *os.File, size uint32) {
	slog.Info("Parsing mdia box", "size", size)
	endpostion, _ := file.Seek(0, 1)
	endpostion += int64(size - 8)
	for {
		var header Box
		err := binary.Read(file, binary.BigEndian, &header)
		if err != nil || int64(header.Size) < 0 {
			break
		}
		boxType := string(header.Type[:])
		if boxType == "minf" {
			parseStblBox(file, header.Size)
		} else {
			slog.Info("Skipping box in mdia", "type", boxType, "size", header.Size)
		}
		if newPos, err := file.Seek(int64(header.Size)-8, 1); err != nil || newPos >= endpostion {
			break
		}

	}
}
