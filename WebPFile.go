package imageserver

import (
	"os"
	"path"

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
func (output *WebPFile) Save(metaImage *MetaImage, baseName string) error {
	img := metaImage.Image

	// Resize & crop
	if img.Bounds().Dx() != output.Width || img.Bounds().Dy() != output.Height {
		img = imaging.Fill(img, output.Width, output.Height, imaging.Center, imaging.Lanczos)
	}

	// File name
	fileName := path.Join(output.Directory, baseName+".webp")

	// Convert and save in the given file
	return metaImage.ConvertToFile("webp", fileName)
}

// Delete deletes the file from the file system.
func (output *WebPFile) Delete(baseName string) error {
	return os.Remove(path.Join(output.Directory, baseName+".webp"))
}
