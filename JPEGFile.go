package imageoutput

import (
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
	"path"

	"github.com/disintegration/imaging"
)

// JPEGFile ...
type JPEGFile struct {
	Directory string
	Width     int
	Height    int
	Quality   float32
}

// Save writes the image in JPEG format to the file system.
func (output *JPEGFile) Save(avatar *MetaImage, baseName string) error {
	img := avatar.Image
	suppliedWidth := img.Bounds().Dx()
	suppliedHeight := img.Bounds().Dy()
	fileName := path.Join(output.Directory, baseName+".jpg")

	if suppliedWidth != output.Width || suppliedHeight != output.Height {
		// Resize & crop
		img = imaging.Fill(img, output.Width, output.Height, imaging.Center, imaging.Lanczos)
	} else if avatar.Extension() == ".jpg" {
		// If it's the same image encoding and same size, save the file as is
		return ioutil.WriteFile(fileName, avatar.Data, 0644)
	}

	return saveJPEG(img, fileName, output.Quality)
}

// saveJPEG saves an image as a file in JPEG format.
func saveJPEG(img image.Image, out string, quality float32) error {
	file, writeErr := os.Create(out)

	if writeErr != nil {
		return writeErr
	}

	defer file.Close()

	encodeErr := jpeg.Encode(file, img, &jpeg.Options{
		Quality: int(quality),
	})

	return encodeErr
}
