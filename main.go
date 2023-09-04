package main

import (
	"os"

	"github.com/pion/mediadevices"
	"github.com/pion/mediadevices/pkg/codec/vpx"
	_ "github.com/pion/mediadevices/pkg/driver/camera"
	"github.com/pion/mediadevices/pkg/frame"
	"github.com/pion/mediadevices/pkg/prop"
	log "github.com/sirupsen/logrus"
)

var Logger = log.StandardLogger()

func main() {
	Logger.Info("###### Getting Media Devices #######")
	os.Setenv("PION_LOG_DEBUG","all")
	vpxParams, err := vpx.NewVP8Params()
	if err != nil {
		Logger.Info(err)
	}
	vpxParams.BitRate = 500_000
	if err != nil {
		Logger.Info(err)

	}
	codecSelector := mediadevices.NewCodecSelector(
		mediadevices.WithVideoEncoders(&vpxParams),
	)
	_, err = mediadevices.GetUserMedia(mediadevices.MediaStreamConstraints{
		Video: func(c *mediadevices.MediaTrackConstraints) {
			c.FrameFormat = prop.FrameFormat(frame.FormatI420)
			c.Width = prop.Int(640)
			c.Height = prop.Int(480)
		},
		Codec: codecSelector,
	})
	if err != nil {
		Logger.Error(err)
	}
	Logger.Info("###### Completed #######")
}