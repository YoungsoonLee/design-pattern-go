package main

import (
	"fmt"

	"github.com/YoungsoonLee/design-pattern-go/pattern-study/wp/streamer"
)

func main() {
	// define the number of workers and jobs
	const numJobs = 4
	const numWorkers = 4

	// create channels for work and results
	notifyChan := make(chan streamer.ProcessingMessage, numJobs)
	defer close(notifyChan)

	videoQueue := make(chan streamer.VideoProcessingJob, numJobs)
	defer close(videoQueue)

	// Get a worker pool.
	wp := streamer.New(videoQueue, numWorkers)
	fmt.Println("wp: ", wp)

	// Srart the worker pool.
	wp.Run()

	fmt.Println("Worker pool started. Press enter to continue.")
	_, _ = fmt.Scanln()

	video1 := wp.NewVideo(1, "./input/puppy1.mp4", "./output", "mp4", notifyChan, nil)

	// make an error
	video2 := wp.NewVideo(2, "./input/bad.txt", "./output", "mp4", notifyChan, nil)

	// Create 4 videos to send to the worker pool
	ops := &streamer.VideoOptions{
		RenameOutput:    true,
		SegmentDuration: 10,
		MaxRate1080p:    "1200k",
		MaxRate720p:     "600k",
		MaxRate480p:     "400k",
	}
	video3 := wp.NewVideo(3, "./input/puppy2.mp4", "./output", "hls", notifyChan, ops)

	video4 := wp.NewVideo(4, "./input/puppy2.mp4", "./output", "mp4", notifyChan, nil)

	// Send the videos to the worker pool.
	videoQueue <- streamer.VideoProcessingJob{Video: video1}
	videoQueue <- streamer.VideoProcessingJob{Video: video2}
	videoQueue <- streamer.VideoProcessingJob{Video: video3}
	videoQueue <- streamer.VideoProcessingJob{Video: video4}

	// Print out the results.
	for i := 0; i < numJobs; i++ {
		msg := <-notifyChan
		fmt.Println("i: ", i, "msg: ", msg)
	}

	fmt.Println("done")
}
