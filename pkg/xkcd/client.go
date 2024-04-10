package xkcd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type ComicData struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

func RequestComics(comicCount int, domain, file string) ([]ComicData, error) {
	client := http.Client{Timeout: 1 * time.Second}
	comics := make([]ComicData, 0)

	for i := 1; i <= comicCount; i++ {
		fmt.Printf("\rОбрабатывается комикс: %d", i)
		request := fmt.Sprintf("%s/%d/%s", domain, i, file)
		resp, err := client.Get(request)
		if err != nil {
			message := fmt.Sprintf("Error sending get request : %s", err.Error())
			return nil, errors.New(message)
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			log.Printf("\rError of get comic. Comic ID: %d. Status Code: %d\n", i, resp.StatusCode)
			continue
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("\rError reading response body for comic %d: %v\n", i, err)
			continue
		}

		var comic ComicData
		err = json.Unmarshal(body, &comic)
		if err != nil {
			message := fmt.Sprintf("Error decoding JSON: %s", err.Error())
			return nil, errors.New(message)
		}

		comics = append(comics, comic)
	}
	fmt.Printf("\r")

	return comics, nil
}

/*func BuildClient() {
		var comic ComicData
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
*/
