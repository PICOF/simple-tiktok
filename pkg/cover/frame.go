package cover

import (
	"bytes"
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"image"
	"image/jpeg"
	"os"
)

func GetCover(ctx context.Context, fileUrl string) (cover []byte, err error) {
	reader := bytes.NewBuffer(nil)
	err = ffmpeg.Input(fileUrl).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(reader, os.Stdout).
		Run()
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to get cover of video: %v", err)
		return
	}
	var img image.Image
	img, _, err = image.Decode(reader)
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to get cover of video: %v", err)
		return
	}

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, img, nil)
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to get cover of video: %v", err)
		return
	}
	cover = buf.Bytes()
	return
}
