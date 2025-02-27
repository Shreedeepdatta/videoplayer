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
	boxes, err := mp4.ParseMP4File(file)
	if err != nil {
		fmt.Printf("Error parsing MP4 file: %v\n", err)
		return
	}

	fmt.Println("MP4 Box Structure:")
	mp4.PrintBoxStructure(boxes, "")
}
