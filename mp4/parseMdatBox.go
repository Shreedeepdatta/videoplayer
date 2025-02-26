package mp4

import (
	"log/slog"
	"os"
)

func parseMdatBox(file *os.File, size uint32) {
	slog.Info("Parsing mdat box", "size", size)
	// file.Seek(int64(size-8), 1)
}
