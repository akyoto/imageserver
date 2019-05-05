package imageserver

import (
	"os"
	"path"
)

// WebPFile ...
type WebPFile struct {
	Directory string
	Width     int
	Height    int
	Quality   int
}

// Save writes the image in WebP format to the file system.
func (output *WebPFile) Save(metaImage *MetaImage, baseName string) error {
	fileName := path.Join(output.Directory, baseName+".webp")
	return metaImage.ConvertToFile("webp", output.Width, output.Height, output.Quality, fileName)
}

// Delete deletes the file from the file system.
func (output *WebPFile) Delete(baseName string) error {
	return os.Remove(path.Join(output.Directory, baseName+".webp"))
}
