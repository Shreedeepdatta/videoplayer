package mp4

import (
	"encoding/binary"
	"log/slog"
	"os"
)

type StcoBox struct {
	EntryCount   uint32
	ChunkOffsets []uint32
}

func parseStcoBox(file *os.File, size uint32) {
	slog.Info("Parsing stco box", "size", size)
	var stco StcoBox
	binary.Read(file, binary.BigEndian, &stco.EntryCount)
	stco.ChunkOffsets = make([]uint32, stco.EntryCount)
	binary.Read(file, binary.BigEndian, &stco.ChunkOffsets)
	slog.Info("Chunk Offsets", "Entry Count", stco.EntryCount)
}
