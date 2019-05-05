package imageserver

import (
	"bytes"
	"errors"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path"

	"github.com/disintegration/imaging"
)

// OriginalFile ...
type OriginalFile struct {
	Directory string
	Width     int
	Height    int
}

// Save writes the original meta to the file system.
func (output *OriginalFile) Save(meta *MetaImage, baseName string) error {
	// Determine file extension
	extension := meta.Extension()

	if extension == "" {
		return errors.New("Unknown format: " + meta.Format)
	}

	modified := false

	// Resize if needed
	data := meta.Data
	img := meta.Image

	if (output.Width != 0 || output.Height != 0) && (img.Bounds().Dx() > output.Width || img.Bounds().Dy() > output.Height) {
		img = imaging.Fill(img, output.Width, output.Height, imaging.Center, imaging.Lanczos)
		modified = true
	}

	if modified {
		buffer := bytes.Buffer{}

		var err error
		switch extension {
		case ".jpg":
			err = jpeg.Encode(&buffer, img, nil)
		case ".png":
			err = png.Encode(&buffer, img)
		case ".gif":
			err = gif.Encode(&buffer, img, nil)
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

// Delete deletes the file from the file system.
func (output *OriginalFile) Delete(baseName string) error {
	os.Remove(path.Join(output.Directory, baseName+".jpg"))
	os.Remove(path.Join(output.Directory, baseName+".png"))
	os.Remove(path.Join(output.Directory, baseName+".gif"))
	return nil
}
