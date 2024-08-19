package main

import (
	"fmt"

	"github.com/YoungsoonLee/design-pattern-go/pattern-study/wp/streamer"
)

func main() {
	// define the number of workers and jobs
	const numWorkers = 2
	const numJobs = 1

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

	// Create 4 videos to send to the worker pool
	video := wp.NewVideo(1, "./input/puppy1.mp4", "./output", "mp4", notifyChan, nil)

	// Send the videos to the worker pool.
	videoQueue <- streamer.VideoProcessingJob{Video: video}

	// Print out the results.
	for i := 0; i < numJobs; i++ {
		msg := <-notifyChan
		fmt.Println("i: ", i, "msg: ", msg)
	}

	fmt.Println("done")
}
