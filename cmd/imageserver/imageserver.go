package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"

	"github.com/chai2010/webp"
)

var port string

func init() {
	flag.StringVar(&port, "port", "7000", "Port the HTTP server should listen on")
	flag.Parse()
}

func main() {
	http.HandleFunc("/", onRequest)
	err := http.ListenAndServe(":"+port, nil)
	log.Fatal(err)
}

// onRequest will convert the image into the requested format.
func onRequest(response http.ResponseWriter, request *http.Request) {
	var img image.Image
	var err error

	// Decode
	switch request.Header.Get("Content-Type") {
	case "image/png":
		img, err = png.Decode(request.Body)

		if err != nil {
			response.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
			return
		}

	case "image/jpeg":
		img, err = jpeg.Decode(request.Body)

		if err != nil {
			response.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
			return
		}

	case "image/gif":
		img, err = gif.Decode(request.Body)

		if err != nil {
			response.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
			return
		}

	case "image/webp":
		img, err = webp.Decode(request.Body)

		if err != nil {
			response.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
			return
		}

	default:
		response.WriteHeader(http.StatusBadRequest)
		fmt.Printf("Unknown content type: %s\n", request.Header.Get("Content-Type"))
		return
	}

	// Encode
	switch request.Header.Get("Accept-Type") {
	case "image/png":
		response.Header().Set("Content-Type", "image/png")
		err = png.Encode(response, img)

		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

	case "image/jpeg":
		response.Header().Set("Content-Type", "image/jpeg")
		err = jpeg.Encode(response, img, nil)

		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

	case "image/gif":
		response.Header().Set("Content-Type", "image/gif")
		err = gif.Encode(response, img, nil)

		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

	case "image/webp":
		response.Header().Set("Content-Type", "image/webp")
		err = webp.Encode(response, img, nil)

		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
	}
}
