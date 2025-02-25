package mp4

import (
	"encoding/binary"
	"log/slog"
	"os"
)

func parseStcoBox(file *os.File, size uint32) {
	slog.Info("parsing stco box")
	var entrycount uint32
	binary.Read(file, binary.BigEndian, &entrycount)
	slog.Info("Chunk offset count", "count", entrycount)
	for i := 0; i < int(entrycount); i++ {
		var chunkoffset uint32
		binary.Read(file, binary.BigEndian, &chunkoffset)
		slog.Info("Chunk offset", "index", i, "offset", chunkoffset)
	}
}
