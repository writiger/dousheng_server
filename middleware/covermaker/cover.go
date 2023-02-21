package covermaker

import (
	"bytes"
	zaplog "dousheng_server/deploy/log"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"os"
	"strings"
)

func GetSnapshot(videoPath, snapshotPath string, frameNum int) (snapshotName string, err error) {
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		GlobalArgs("-loglevel", "quiet").
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		zaplog.ZapLogger.Warnf("failed when making cover err:%v", err)
		return "", err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		zaplog.ZapLogger.Warnf("failed when making cover err:%v", err)
		return "", err
	}

	err = imaging.Save(img, snapshotPath+".png")
	if err != nil {
		zaplog.ZapLogger.Warnf("failed when making cover err:%v", err)
		return "", err
	}

	names := strings.Split(snapshotPath, "\\")
	snapshotName = names[len(names)-1] + ".png"
	return
}
