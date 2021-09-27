package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/dvoeglazyi/gospelling/templates"
	"github.com/dvoeglazyi/gospelling/wikitionary"
)

const (
	wiktionaryPagesArticlesMultistreamPath = "wiki/ruwiktionary-20210801-pages-articles-multistream.xml"
	templatesPath                          = "wiki/templates.json"
)

func main() {
	// пример использования
	// example of usage
	if err := do(); err != nil {
		fmt.Println(err)
	}
	return
}

func do() error {
	if err := loadTemplatesFromWiktionary(); err != nil {
		return err
	}
	if err := parseWiktionary(); err != nil {
		return err
	}
	return nil
}

func loadTemplatesFromWiktionary() error {
	output, err := os.OpenFile(templatesPath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer output.Close()

	file, err := os.Open(wiktionaryPagesArticlesMultistreamPath)
	if err != nil {
		return err
	}
	defer file.Close()

	parsedTemplates, err := wikitionary.ParseTemplatesSource(file)
	if err != nil {
		return err
	}

	return json.NewEncoder(output).Encode(parsedTemplates)
}

func parseWiktionary() error {
	file, err := os.Open(wiktionaryPagesArticlesMultistreamPath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileTemplates, err := os.Open(templatesPath)
	if err != nil {
		return err
	}
	defer fileTemplates.Close()
	// загрузка шаблонов
	tt, err := templates.ParseTemplates(fileTemplates)
	if err != nil {
		return err
	}
	// создание парсера
	parser := wikitionary.NewParser(tt)
	// парсинг слов из словаря (в соответствии с шаблонами)
	_, err = parser.Do(file)
	if err != nil {
		return err
	}
	return nil
}
