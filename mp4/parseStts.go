package mp4

import (
	"encoding/binary"
	"log/slog"
	"os"
)

type SttsBox struct {
	EntryCount uint32
	Entries    []struct {
		SampleCount uint32
		SampleData  uint32
	}
}

func parseSttsBox(file *os.File, size uint32) {
	slog.Info("Parsing stts box", "size", size)
	var stts SttsBox
	binary.Read(file, binary.BigEndian, &stts.EntryCount)
	stts.Entries = make([]struct {
		SampleCount uint32
		SampleData  uint32
	}, stts.EntryCount)
	binary.Read(file, binary.BigEndian, &stts.Entries)
	slog.Info("Time to sample", "Entry Count", stts.EntryCount)
}
