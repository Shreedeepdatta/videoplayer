package mp4

import (
	"encoding/binary"
	"log/slog"
	"os"
)

func parseStsdBox(file *os.File, size uint32) {
	slog.Info("Parsing Stbl Box", "size", size)
	var version uint8
	var flags [3]byte
	binary.Read(file, binary.BigEndian, &version)
	binary.Read(file, binary.BigEndian, &flags)
	var entrycount uint32
	binary.Read(file, binary.BigEndian, &entrycount)
	slog.Info("Sample description entry count", "count", entrycount)
	for i := 0; i < int(entrycount); i++ {
		var entrysize uint32
		var format [4]byte
		binary.Read(file, binary.BigEndian, &entrysize)
		binary.Read(file, binary.BigEndian, &format)
		codec := string(format[:])
		slog.Info("Sample format", "format", codec)
		if codec == "avc1" {
			file.Seek(78, 1)
			var spsSize uint16
			binary.Read(file, binary.BigEndian, &spsSize)
			file.Seek(int64(spsSize), 1)
			var ppSize uint16
			binary.Read(file, binary.BigEndian, &ppSize)
			file.Seek(int64(ppSize), 1)
		}
	}
}
