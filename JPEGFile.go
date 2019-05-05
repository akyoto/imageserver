package imageserver

import (
	"os"
	"path"
)

// JPEGFile ...
type JPEGFile struct {
	Directory string
	Width     int
	Height    int
	Quality   float64
}

// Save writes the image in JPEG format to the file system.
func (output *JPEGFile) Save(metaImage *MetaImage, baseName string) error {
	fileName := path.Join(output.Directory, baseName+".webp")
	return metaImage.ConvertToFile("jpeg", output.Width, output.Height, output.Quality, fileName)
}

// Delete deletes the file from the file system.
func (output *JPEGFile) Delete(baseName string) error {
	return os.Remove(path.Join(output.Directory, baseName+".jpg"))
}
