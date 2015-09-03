package iascii

import (
	"fmt"
	"image"
	"image/color"
	"io"
)

func Encode(w io.Writer, m image.Image) (err error) {
	b := m.Bounds()
	mw, mh := b.Dx(), b.Dy()
	if mw <= 0 || mh <= 0 {
		err = fmt.Errorf("Bad image bounds")
		return
	}

	cm := []byte(".ocOGDQ@")
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			c := color.GrayModel.Convert(m.At(x, y))
			nc := c.(color.Gray)
			nc.Y >>= 5
			_, err = w.Write(cm[nc.Y : nc.Y+1])
			if err != nil {
				return
			}
		}
		_, err = w.Write([]byte("\n"))
		if err != nil {
			return
		}
	}
	return
}
