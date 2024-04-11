package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang-cource-2024/pkg/words"
	"golang-cource-2024/pkg/xkcd"
	"os"
	"strconv"
)

type ComicDetails struct {
	Url      string   `json:"url"`
	Keywords []string `json:"keywords"`
}

func GetComics(comics []xkcd.ComicData) (map[string]ComicDetails, error) {
	comicsDetails := make(map[string]ComicDetails)

	for _, comic := range comics {
		keywords, err := words.Normalize(comic.Transcript)
		if err != nil {
			message := fmt.Sprintf("Error of normalize: %s", err.Error())
			return nil, errors.New(message)
		}

		comicDetails := ComicDetails{
			Url:      comic.Img,
			Keywords: keywords,
		}

		key := strconv.Itoa(comic.Num)
		comicsDetails[key] = comicDetails
	}

	return comicsDetails, nil
}

func SaveComics(jsonData []byte, dbFileName string) error {
	file, err := os.OpenFile(dbFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		message := fmt.Sprintf("Error during file creation: %s", err.Error())
		return errors.New(message)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		message := fmt.Sprintf("Error during writing data to a file: %s", err.Error())
		return errors.New(message)
	}

	return nil
}

func ConvertComicsDetailsToJson(comics map[string]ComicDetails) ([]byte, error) {
	jsonData, err := json.MarshalIndent(comics, "", " ")
	if err != nil {
		message := fmt.Sprintf("Error during JSON encoding: %s", err.Error())
		return nil, errors.New(message)
	}

	return jsonData, nil
}
