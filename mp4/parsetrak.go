package mp4

import (
	"encoding/binary"
	"log/slog"
	"os"
)

func parseTrakBox(file *os.File, size uint32) {
	slog.Info("Parsing trak box", "size", size)
	endposition, _ := file.Seek(0, 1)
	endposition += int64(size - 8)
	for {
		var header Box
		err := binary.Read(file, binary.BigEndian, &header)
		if err != nil || int64(header.Size) < 8 {
			break
		}
		boxType := string(header.Type[:])
		if boxType == "tkhd" {
			parseTkhdBox(file, header.Size)
		} else if boxType == "minf" {
			parseMdiaBox(file, header.Size)
		} else {
			slog.Info("Skipping box in Trak", "type", boxType, "size", header.Size)
		}
		if newPos, err := file.Seek(int64(header.Size)-8, 1); err != nil || newPos >= endposition {
			break
		}
	}
}
