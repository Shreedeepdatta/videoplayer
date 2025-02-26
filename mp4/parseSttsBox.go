package mp4

import (
	"encoding/binary"
	"log/slog"
	"os"
)

type StszBox struct {
	SampleSize  uint32
	EntryCount  uint32
	SampleSizes []uint32
}

func parseStszBox(file *os.File, size uint32) {
	slog.Info("parsing stsz box", "size", size)
	var stsz StszBox
	binary.Read(file, binary.BigEndian, &stsz.SampleSize)
	binary.Read(file, binary.BigEndian, &stsz.EntryCount)
	stsz.SampleSizes = make([]uint32, stsz.EntryCount)
	binary.Read(file, binary.BigEndian, &stsz.SampleSizes)
}
