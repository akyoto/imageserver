package imageserver

import (
	"fmt"
	"image"
	"os"

	"github.com/aerogo/http/client"
)

// ServerPort is the default port
var ServerPort = "7000"

// MetaImage represents a single image with the name of the format
// and the original byte buffer that was used to create it.
type MetaImage struct {
	Image  image.Image
	Data   []byte
	Format string
}

// Extension returns the file extension of the image.
func (meta *MetaImage) Extension() string {
	switch meta.Format {
	case "jpg", "jpeg":
		return ".jpg"
	case "png":
		return ".png"
	case "gif":
		return ".gif"
	default:
		return ""
	}
}

// String returns a text representation of the format, width and height.
func (meta *MetaImage) String() string {
	return fmt.Sprint(meta.Format, " | ", meta.Image.Bounds().Dx(), "x", meta.Image.Bounds().Dy())
}

// ConvertToFile sends a request to the server and saves the resulting image in the given format.
// Format can be one of the following: png, jpeg, gif, webp
func (meta *MetaImage) ConvertToFile(format string, width int, height int, fileName string) error {
	request := client.Get("http://127.0.0.1:" + ServerPort + "/")
	request.Header("Content-Type", "image/"+meta.Format)
	request.Header("Accept-Type", "image/"+format)
	request.Header("Image-Width", fmt.Sprintf("%d", width))
	request.Header("Image-Height", fmt.Sprintf("%d", height))
	request.Body(meta.Data)
	response, err := request.End()

	if err != nil {
		return err
	}

	if !response.Ok() {
		return fmt.Errorf("Status: %d", response.StatusCode())
	}

	file, err := os.Create(fileName)

	if err != nil {
		return err
	}

	_, err = response.WriteTo(file)
	file.Close()
	return err
}
