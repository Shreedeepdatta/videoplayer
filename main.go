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
func parseMdiaBox(file *os.File, size uint32) {
	slog.Info("Parsing mdia box", "size", size)

}
func parseTrakBox(file *os.File, size uint32) {
	slog.Info("Parsing trak box", "size", size)
	endposition, _ := file.Seek(0, 1)
	endposition += int64(size - 8)
	for {
		var header Box
		err := binary.Read(file, binary.BigEndian, &header)
		if err != nil || int64(header.Size) < 8 {
			break
		}
		boxType := string(header.Type[:])
		if boxType == "tkbhd" {
			parseTkhdBox(file, header.Size)
		} else if boxType == "mdia" {
			parseMdiaBox(file, header.Size)
		} else {
			slog.Info("Skipping box in Trak", "type", boxType, "size", header.Size)
		}
	}
}

func parseMoovBox(file *os.File, size uint32) {
	slog.Info("Parsing moov box", "size", size)
	endposition, _ := file.Seek(0, 1)
	endposition += int64(size - 8)
	for {
		var header Box
		err := binary.Read(file, binary.BigEndian, &header)
		if err != nil || int64(header.Size) < 8 {
			break
		}
		boxType := string(header.Type[:])
		if boxType == "trak" {
			parseTrakBox(file, header.Size)
		} else {
			slog.Info("Skipping box in moov", "type", boxType, "size", header.Size)
		}
		if newposition, err := file.Seek(int64(header.Size)-8, 1); err != nil || newposition >= endposition {
			break
		}
	}
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
		if boxType == "moov" {
			slog.Info("moov box found", "size", header.Size)
			parseMoovBox(file, header.Size)
		} else if boxType == "mdat" {
			slog.Info("mdat box found", "size", header.Size)
		} else {
			slog.Info("Skipping box", "type", boxType, "size", header.Size)
		}

		if _, err := file.Seek(int64(header.Size)-8, 1); err != nil {
			break
		}
	}
}
