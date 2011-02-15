//2,|gofmt
package iascii

import (
	"image"
	"io"
	"os"
)

func Encode(w io.Writer, m image.Image) (err os.Error) {
	b := m.Bounds()
	mw, mh := b.Dx(), b.Dy()
	if mw <= 0 || mh <= 0 {
		err = os.NewError("Bad image bounds")
		return
	}

	cm := []byte(".ocOGDQ@")
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			c := image.GrayColorModel.Convert(m.At(x, y)).(image.GrayColor)
			c.Y >>= 5
			_, err = w.Write(cm[c.Y : c.Y+1])
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
