package mp4

import (
	"encoding/binary"
	"log/slog"
	"os"
)

type StscBox struct {
	EntryCount uint32
	Entries    []struct {
		FirstChunk      uint32
		SamplesPerChunk uint32
		SampleDescIndex uint32
	}
}

func parseStscBox(file *os.File, size uint32) {
	slog.Info("Parsing Stsc Box", "size", size)
	var stsc StscBox
	binary.Read(file, binary.BigEndian, &stsc.EntryCount)
	stsc.Entries = make([]struct {
		FirstChunk      uint32
		SamplesPerChunk uint32
		SampleDescIndex uint32
	}, stsc.EntryCount)
	binary.Read(file, binary.BigEndian, stsc.Entries)
	slog.Info("Sample to chunk", "EntryCount", stsc.EntryCount)
}
