package main

import (
	"encoding/binary"
	"fmt"
	"log/slog"
	"os"
)

type Box struct {
	Size uint32
	Type [4]byte
}

func main() {
	file, err := os.Open("videos/sample.mp4")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	for {
		var header Box
		err := binary.Read(file, binary.BigEndian, &header)
		if err != nil {
			break
		}
		boxType := string(header.Type[:])
		if boxType == "moov" || boxType == "mdat" {
			slog.Info("Important box found", "type", boxType, "size", header.Size)
		}

		if _, err := file.Seek(int64(header.Size)-8, 1); err != nil {
			break
		}
	}
}
