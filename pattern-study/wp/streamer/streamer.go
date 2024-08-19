package streamer

import (
	"fmt"
	"path"
	"strings"
)

// ProcessingMessage is a struct that represents a processing message
type ProcessingMessage struct {
	ID         int
	Successful bool
	Messsage   string
	OutputFile string
}

// VideoProcessingJob is a struct that represents a video processing job
type VideoProcessingJob struct {
	Video Video
}

// Processor is an interface that represents a processor
type Processor struct {
	Engine Encoder
}

// Video is a struct that represents a video
type Video struct {
	ID           int
	InputFile    string
	OutputDir    string
	EncodingType string
	NotifyChan   chan ProcessingMessage
	Options      *VideoOptions
	Encoder      Processor
}

// VideoOptions ...
type VideoOptions struct {
	RenameOutput    bool
	SegmentDuration int
	MaxRate1080p    string
	MaxRate720p     string
	MaxRate480p     string
}

// NewVideo creates a new video
func (vd *VideoDispatcher) NewVideo(id int, inputFile string, outputDir string, encodingType string, notifyChan chan ProcessingMessage, ops *VideoOptions) Video {
	if ops == nil {
		ops = &VideoOptions{}
	}

	fmt.Println("New video created: ", id, inputFile)

	return Video{
		ID:           id,
		InputFile:    inputFile,
		OutputDir:    outputDir,
		EncodingType: encodingType,
		NotifyChan:   notifyChan,
		Encoder:      vd.Processor,
		Options:      ops,
	}
}

func (v *Video) encode() {
	var fileName string

	switch v.EncodingType {
	case "mp4":
		// encode the video
		fmt.Println("Encoding video to mp4", v.ID)

		name, err := v.encodeToMP4()
		if err != nil {
			v.sendToNotifyChan(false, fileName, err.Error())
			return
		}

		fileName = fmt.Sprintf("%s.mp4", name)

	default:
		v.sendToNotifyChan(false, fileName, "Unknown encoding type")
		return
	}

	v.sendToNotifyChan(true, fileName, "Video encoded successfully")
}

func (v *Video) encodeToMP4() (string, error) {
	baseFileName := ""

	if !v.Options.RenameOutput {
		b := path.Base(v.InputFile)
		baseFileName = strings.TrimSuffix(b, path.Ext(b))
	} else {
		// TODO: Implement renaming
	}

	err := v.Encoder.Engine.EncodeToMP4(v, baseFileName)
	if err != nil {
		return "", err
	}

	return baseFileName, nil
}

func (v *Video) sendToNotifyChan(successful bool, fileName, message string) {
	v.NotifyChan <- ProcessingMessage{
		ID:         v.ID,
		Successful: successful,
		Messsage:   message,
		OutputFile: fileName,
	}
}

// New creates a new VideoDispatcher
func New(videoQueue chan VideoProcessingJob, numWorkers int) *VideoDispatcher {

	var e VideoEncoder
	p := Processor{Engine: &e}

	return &VideoDispatcher{
		WorkerPool: make(chan chan VideoProcessingJob, numWorkers),
		maxWorkers: numWorkers,
		jobQueue:   videoQueue,
		Processor:  p,
	}
}
