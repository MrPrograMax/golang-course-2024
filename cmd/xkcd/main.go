package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"golang-cource-2024/pkg/database"
	"golang-cource-2024/pkg/xkcd"
	"log"
)

func main() {
	//Обработкаа аргументов
	var outputJSON bool
	var comicCount int

	flag.BoolVar(&outputJSON, "o", false, "Output JSON structure")
	flag.IntVar(&comicCount, "n", 10, "Number of comics to display")
	flag.Parse()

	log.Printf("Get flags: { Output JSON: %t } { Count of comics: %d }\n", outputJSON, comicCount)

	//Загрузка конфига
	log.Println("Uploading the config...")
	if err := initConfig(); err != nil {
		log.Fatalf("Error initializing configs: %s", err.Error())
	}
	log.Println("Upload complete.")

	var domain = viper.GetString("source_url")
	var file = viper.GetString("source_file")
	var dbFileName = viper.GetString("db_file")
	log.Printf("Received data: { Url: %s }; { File: %s } { DB File Name: %s }\n",
		domain, file, dbFileName)

	//Получение комиксов
	log.Println("Receiving comics...")
	comics, err := xkcd.RequestComics(comicCount, domain, file)
	if err != nil {
		log.Fatalf("Something went wrong: %s", err.Error())
	}
	log.Println("Receipt complete.")

	//Преобразование комиксов с сайта в нужный формат
	log.Println("Collecting comics details...")
	comicsDetails, err := database.GetComics(comics)
	if err != nil {
		log.Fatalf("Error converting comics to required format %s", err.Error())
	}
	log.Println("Collection is complete.")

	//Конвертирование деталей в JSON
	log.Println("Converting data to JSON...")
	jsonData, err := database.ConvertComicsDetailsToJson(comicsDetails)
	if err != nil {
		log.Fatalf("Error converting comics details to JSON %s", err.Error())
	}
	log.Println("Conversion complete")

	//Сохранение JSON в файл
	log.Println("Saving JSON to file...")
	if err := database.SaveComics(jsonData, dbFileName); err != nil {
		log.Fatalf("Error saving comics to file %s", err.Error())
	}
	log.Println("Save is complete")

	//Вывод по требованию
	if outputJSON {
		log.Println("Obtained JSON:")
		fmt.Println(string(jsonData))
	}
}

func initConfig() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
