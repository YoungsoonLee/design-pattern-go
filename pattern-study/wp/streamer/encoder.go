package streamer

// Encoder is an interface that represents an encoder
type Encoder interface {
	EncodeToMP4(v *Video, baseFileName string) error
}

// VideoEncoder is a struct that represents a video encoder
type VideoEncoder struct{}

// EncodeToMP4 encodes a video to MP4
func (ve *VideoEncoder) EncodeToMP4(v *Video, baseFileName string) error {
	// Simulate encoding a video
	return nil
}
