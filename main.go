package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

func CutVideoByMS(video *gocv.VideoCapture, start, end float64, index int, fps float64, width, height int) {
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
		if currentTime >= end {
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

func CutVideoByFrame(video *gocv.VideoCapture, start, end float64, index int, fps float64, width, height int) {
	outputPath := fmt.Sprintf("output_segment_%d.mp4", index)
	writer, err := gocv.VideoWriterFile(outputPath, "mp4v", fps, width, height, true)
	if err != nil {
		fmt.Printf("Error opening video writer: %v\n", err)
		return
	}
	defer writer.Close()

	startFrame := start * fps / 1000
	endFrame := end * fps / 1000

	// Set the start time
	video.Set(gocv.VideoCapturePosFrames, startFrame)

	for {
		currentFrame := video.Get(gocv.VideoCapturePosFrames)
		if currentFrame >= endFrame {
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

	// Process each time range
	for i, tr := range timeRanges {
		CutVideoByFrame(video, tr.Start, tr.End, i, fps, width, height)
	}
}
