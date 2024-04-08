package xkcd

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"golang-cource-2024/pkg/database"
	"golang-cource-2024/pkg/words"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const File = "/info.0.json"

func BuildClient() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	var outputJSON bool
	var numComics int

	flag.BoolVar(&outputJSON, "o", false, "Output JSON structure")
	flag.IntVar(&numComics, "n", 100, "Number of comics to display")

	flag.Parse()

	c := http.Client{Timeout: time.Duration(1) * time.Second}
	comics := make(map[string]database.ComicDetails, 0)

	var Domain = viper.GetString("source_url")
	for i := 1; i <= numComics; i++ {
		resp, err := c.Get(fmt.Sprintf("%s%s%s", Domain, strconv.Itoa(i), File))
		if err != nil {
			fmt.Printf("Error %s", err)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response body for comic %d: %v\n", i, err)
			continue
		}

		var comic database.ComicData
		err = json.Unmarshal(body, &comic)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}

		keywords, err := words.Normalize(comic.Transcript)
		if err != nil {
			fmt.Println("Error of normalize:", err)
			return
		}

		comicDetails := database.ComicDetails{
			Url:      comic.Img,
			Keywords: keywords,
		}

		key := strconv.Itoa(i)
		comics[key] = comicDetails
	}

	jsonData, err := json.MarshalIndent(comics, "", "   ")
	if err != nil {
		fmt.Println("Ошибка при кодировании в JSON:", err)
		return
	}

	var filename = viper.GetString("db_file")

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = ioutil.WriteFile(filename, jsonData, 0644)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return
	}

	if outputJSON {
		fmt.Println(string(jsonData))
	}
}

func initConfig() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
