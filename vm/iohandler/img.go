package iohandler

import (
	"drylang/core"
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
)

// nearest neighbor resize
func resizeImage(src image.Image, w, h int) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	srcBounds := src.Bounds()
	srcW := srcBounds.Dx()
	srcH := srcBounds.Dy()

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			srcX := srcBounds.Min.X + (x * srcW / w)
			srcY := srcBounds.Min.Y + (y * srcH / h)
			dst.Set(x, y, src.At(srcX, srcY))
		}
	}
	return dst
}

// BuiltinImg handles img("resize", inPath, outPath, w, h)
func BuiltinImg(v core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	if len(args) < 1 || args[0].Type != core.ValString {
		return core.UnknownValue, v.Errorf("want img(method, ...args)")
	}
	method := args[0].Data.(string)

	switch method {
	case "resize":
		if len(args) != 5 {
			return core.UnknownValue, v.Errorf("img.resize wants (inPath, outPath, w, h)")
		}
		if args[1].Type != core.ValString || args[2].Type != core.ValString || args[3].Type != core.ValNumber || args[4].Type != core.ValNumber {
			return core.UnknownValue, v.Errorf("img.resize args must be (string, string, number, number)")
		}

		inPath := args[1].Data.(string)
		outPath := args[2].Data.(string)
		w := int(args[3].Data.(float64))
		h := int(args[4].Data.(float64))

		file, err := os.Open(inPath)
		if err != nil {
			return core.UnknownValue, v.Errorf("img open error: %s", err)
		}
		defer file.Close()

		src, _, err := image.Decode(file)
		if err != nil {
			return core.UnknownValue, v.Errorf("img decode error: %s", err)
		}

		dst := resizeImage(src, w, h)

		outFile, err := os.Create(outPath)
		if err != nil {
			return core.UnknownValue, v.Errorf("img create out error: %s", err)
		}
		defer outFile.Close()

		// Always output as PNG for simplicity
		err = png.Encode(outFile, dst)
		if err != nil {
			return core.UnknownValue, v.Errorf("img save error: %s", err)
		}

		return core.Value{Type: core.ValBool, Data: true}, nil

	default:
		return core.UnknownValue, v.Errorf("unknown img method: %s", method)
	}
}
