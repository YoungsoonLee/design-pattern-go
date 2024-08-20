package streamer

import "fmt"

// VideoDispatcher is a struct that represents a video dispatcher
type VideoDispatcher struct {
	WorkerPool chan chan VideoProcessingJob
	maxWorkers int
	jobQueue   chan VideoProcessingJob
	Processor  Processor
}

// type videoWorker
type videoWorker struct {
	id         int
	jobQueue   chan VideoProcessingJob
	workerPool chan chan VideoProcessingJob
}

// newVideoWorker
func newVideoWorker(id int, workerPool chan chan VideoProcessingJob) videoWorker {
	fmt.Println("newVideoWorker: ", id)

	return videoWorker{
		id:         id,
		jobQueue:   make(chan VideoProcessingJob),
		workerPool: workerPool,
	}
}

// Start() starts a worker
func (w *videoWorker) start() {
	fmt.Println("start worker: ", w.id)

	go func() {
		for {
			// add worker to the worker pool
			w.workerPool <- w.jobQueue

			// vait for a job to come back
			select {
			case job := <-w.jobQueue:
				// process the video job
				w.processVideoJob(job.Video)
			}
		}
	}()
}

// Run starts the video dispatcher
func (d *VideoDispatcher) Run() {
	fmt.Println("Run")

	for i := 0; i < d.maxWorkers; i++ {
		worker := newVideoWorker(i, d.WorkerPool)
		worker.start()
	}
	go d.dispatch()
}

// dispatch()
func (d *VideoDispatcher) dispatch() {
	fmt.Println("dispatch")

	for {
		select {
		case job := <-d.jobQueue:
			go func(job VideoProcessingJob) {
				// try to obtain a worker job channel that is available
				// this will block until a worker is idle
				jobChannel := <-d.WorkerPool

				// dispatch the job to the worker job channel
				jobChannel <- job
			}(job)
		}
	}
}

// processVideoJob()
func (w videoWorker) processVideoJob(v Video) {
	// encode the video
	v.encode()
}
