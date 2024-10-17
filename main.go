package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

func main() {
	videoPath := "input.mp4"
	timeRanges := []struct{ Start, End float64 }{
		{0, 1200},
		{1200, 2800},
	}

	// Open the video file
	video, err := gocv.VideoCaptureFile(videoPath)
	if err != nil {
		fmt.Printf("Error opening video file: %v\n", err)
		return
	}
	defer video.Close()

	// Get video properties
	fps := video.Get(gocv.VideoCaptureFPS)
	width := int(video.Get(gocv.VideoCaptureFrameWidth))
	height := int(video.Get(gocv.VideoCaptureFrameHeight))

	// Function to write a segment
	writeSegment := func(start, end float64, index int) {
		outputPath := fmt.Sprintf("output_segment_%d.mp4", index)
		writer, err := gocv.VideoWriterFile(outputPath, "mp4v", fps, width, height, true)
		if err != nil {
			fmt.Printf("Error opening video writer: %v\n", err)
			return
		}
		defer writer.Close()

		// Set the start time
		video.Set(gocv.VideoCapturePosMsec, start)

		for {
			currentTime := video.Get(gocv.VideoCapturePosMsec)
			if currentTime > end {
				break
			}

			frame := gocv.NewMat()
			if ok := video.Read(&frame); !ok || frame.Empty() {
				frame.Close()
				break
			}

			writer.Write(frame)
			frame.Close()
		}
	}

	// Process each time range
	for i, tr := range timeRanges {
		writeSegment(tr.Start, tr.End, i)
	}
}
