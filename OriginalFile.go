package imageserver

import (
	"errors"
	"os"
	"path"
)

// OriginalFile ...
type OriginalFile struct {
	Directory string
	Width     int
	Height    int
	Quality   int
}

// Save writes the original meta to the file system.
func (output *OriginalFile) Save(metaImage *MetaImage, baseName string) error {
	// Determine file extension
	extension := metaImage.Extension()

	if extension == "" {
		return errors.New("Unknown format: " + metaImage.Format)
	}

	fileName := path.Join(output.Directory, baseName+extension)
	return metaImage.ConvertToFile(metaImage.Format, output.Width, output.Height, output.Quality, fileName)
}

// Delete deletes the file from the file system.
func (output *OriginalFile) Delete(baseName string) error {
	os.Remove(path.Join(output.Directory, baseName+".jpg"))
	os.Remove(path.Join(output.Directory, baseName+".png"))
	os.Remove(path.Join(output.Directory, baseName+".gif"))
	return nil
}
