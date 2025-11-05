package test

import (
	"encoding/json"
	"errors"
	"fmt"
	"image/jpeg"
	"net/http"
	"net/url"
	"os"
)

type Film struct {
	Name   string
	Poster string
}

func Poster(filename string) (*Film, error) {
	safeName := url.QueryEscape(filename)
	requestURL := fmt.Sprintf("%s?apikey=9f693ec9&t=%s", BaseUrl, safeName)
	resp, err := http.Get(requestURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	var film Film
	if err = json.NewDecoder(resp.Body).Decode(&film); err != nil {
		return nil, err
	}
	// 获取海报的图像
	resp, err = http.Get(film.Poster)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	img, err := jpeg.Decode(resp.Body)
	if err != nil {
		return nil, err
	}
	filePath := "./pic/" + film.Name + ".jpg"
	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if err := jpeg.Encode(file, img, &jpeg.Options{Quality: 75}); err != nil {
		return nil, err
	}
	return &film, nil
}
