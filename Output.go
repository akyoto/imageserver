package imageoutput

// Output represents a system that saves an image locally (in database or as a file, e.g.)
type Output interface {
	Save(img *MetaImage, baseName string) error
}
