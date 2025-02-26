package mp4

import (
	"encoding/binary"
	"log/slog"
	"os"
)

func Parsefile(file *os.File) {
	for {
		var header Box
		err := binary.Read(file, binary.BigEndian, &header)
		if err != nil {
			break
		}
		boxType := string(header.Type[:])
		if boxType == "moov" {
			slog.Info("moov box found", "size", header.Size)
			ParseMoovBox(file, header.Size)
		} else if boxType == "mdat" {
			slog.Info("mdat box found", "size", header.Size)
			parseMdatBox(file, header.Size)
		} else {
			slog.Info("Skipping box", "type", boxType, "size", header.Size)
		}

		if _, err := file.Seek(int64(header.Size)-8, 1); err != nil {
			break
		}
	}
}
