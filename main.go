package main

import (
	"fmt"
	"os"
	"videoplayer/mp4"
)

func main() {
	file, err := os.Open("videos/sample.mp4")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	mp4.Parsefile(file)
}
