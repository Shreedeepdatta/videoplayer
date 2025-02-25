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
	if entrycount > 0 {
		var entrySize uint32
		var format [4]byte
		binary.Read(file, binary.BigEndian, &entrySize)
		binary.Read(file, binary.BigEndian, &format)
		slog.Info("Sample format", "format", string(format[:]))
	}
}
