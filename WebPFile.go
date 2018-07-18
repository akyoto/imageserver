package imageoutput

import (
	"image"
	"os"
	"path"

	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
)

// WebPFile ...
type WebPFile struct {
	Directory string
	Width     int
	Height    int
	Quality   float32
}

// Save writes the image in WebP format to the file system.
func (output *WebPFile) Save(avatar *MetaImage, baseName string) error {
	img := avatar.Image

	// Resize & crop
	if img.Bounds().Dx() != output.Width || img.Bounds().Dy() != output.Height {
		img = imaging.Fill(img, output.Width, output.Height, imaging.Center, imaging.Lanczos)
	}

	// Write to file
	fileName := path.Join(output.Directory, baseName+".webp")
	return saveWebP(img, fileName, output.Quality)
}

// saveWebP saves an image as a file in WebP format.
func saveWebP(img image.Image, out string, quality float32) error {
	file, writeErr := os.Create(out)

	if writeErr != nil {
		return writeErr
	}

	defer file.Close()

	encodeErr := webp.Encode(file, img, &webp.Options{
		Quality:  quality,
		Lossless: quality == 100,
	})

	return encodeErr
}
