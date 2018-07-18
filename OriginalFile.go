package imageoutput

import (
	"bytes"
	"errors"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"path"

	"github.com/disintegration/imaging"
)

// OriginalFile ...
type OriginalFile struct {
	Directory string
	Width     int
	Height    int
}

// Save writes the original avatar to the file system.
func (output *OriginalFile) Save(avatar *MetaImage, baseName string) error {
	// Determine file extension
	extension := avatar.Extension()

	if extension == "" {
		return errors.New("Unknown format: " + avatar.Format)
	}

	modified := false

	// Resize if needed
	data := avatar.Data
	img := avatar.Image

	if (output.Width != 0 || output.Height != 0) && (img.Bounds().Dx() > output.Width || img.Bounds().Dy() > output.Height) {
		img = imaging.Fill(img, output.Width, output.Height, imaging.Center, imaging.Lanczos)
		modified = true
	}

	if modified {
		buffer := new(bytes.Buffer)

		var err error
		switch extension {
		case ".jpg":
			err = jpeg.Encode(buffer, img, nil)
		case ".png":
			err = png.Encode(buffer, img)
		case ".gif":
			err = gif.Encode(buffer, img, nil)
		}

		if err != nil {
			return err
		}

		data = buffer.Bytes()
	}

	// Write to file
	fileName := path.Join(output.Directory, baseName+extension)
	return ioutil.WriteFile(fileName, data, 0644)
}
