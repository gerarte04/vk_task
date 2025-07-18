package utils

import (
	"errors"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"marketplace/pkg/http/request"
	"net/http"
	"time"
)

var (
	ErrBadImageURL = errors.New("failed to fetch image by specified url")
	ErrBadImageData = errors.New("failed to parse image by fetched data (supported formats are jpeg, png, gif)")
)

func GetImage(addr string) (image.Image, string, error) {
	req := request.NewRequest(http.MethodGet, addr)
	if req.Err() != nil {
		return nil, "", ErrBadImageURL
	}

	cli := http.Client{Timeout: time.Second}

	resp, err := cli.Do(req.Http())
	if err != nil {
		return nil, "", ErrBadImageURL
	}

	image, format, err := image.Decode(resp.Body)
	if err != nil {
		return nil, "", ErrBadImageData
	}

	return image, format, err
}
