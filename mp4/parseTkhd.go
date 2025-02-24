package mp4

import (
	"encoding/binary"
	"log/slog"
	"os"
)

type Tkhbd struct {
	Version  uint8
	Flags    [3]byte
	TrackID  uint32
	Duration uint32
	Width    uint32
	Height   uint32
}

func parseTkhdBox(file *os.File, size uint32) {
	slog.Info("Parsing tkhd box", "size", size)
	var tkhd Tkhbd
	binary.Read(file, binary.BigEndian, &tkhd.Version)
	binary.Read(file, binary.BigEndian, &tkhd.Flags)
	binary.Read(file, binary.BigEndian, &tkhd.TrackID)
	file.Seek(4, 1)
	binary.Read(file, binary.BigEndian, &tkhd.Duration)
	file.Seek(40, 1)
	binary.Read(file, binary.BigEndian, &tkhd.Width)
	binary.Read(file, binary.BigEndian, &tkhd.Height)
	slog.Info("Track info", "Track ID", tkhd.TrackID, "Duration", tkhd.Duration, "Width", tkhd.Width, "height", tkhd.Height)
}
