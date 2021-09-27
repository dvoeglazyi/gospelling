package wikitionary

import (
	"bufio"
	"encoding/xml"
	"io"
	"strings"
	"time"
)

type Page struct {
	Title    string `xml:"title"`
	NS       int    `xml:"ns"`
	ID       int    `xml:"id"`
	Revision struct {
		ID          int       `xml:"id"`
		ParentID    int       `xml:"parentid"`
		Timestamp   time.Time `xml:"timestamp"`
		Contributor struct {
			Username string `xml:"username"`
			ID       int    `xml:"id"`
		} `xml:"contributor"`
		Format string `xml:"format"`
		Model  string `xml:"model"`
		SHA1   string `xml:"sha1"`
		Text   string `xml:"text"`
	} `xml:"revision"`
}

func (p *Parser) Do(r io.Reader) ([]*Word, error) {
	reader := bufio.NewReader(r)
	var words []*Word
	for {
		value, err := readPage(reader)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		var page Page
		if err := xml.Unmarshal([]byte(value), &page); err != nil {
			return nil, err
		}

		word, err := p.getWord(page.Revision.Text)
		if err != nil {
			return nil, err
		} else if word == nil {
			continue
		}

		word.Value = page.Title

		// TODO: остановился где-то тут
		//variants := word.Template.GetFormVariants(word.Bases)
		//fmt.Println(word.Value, variants)
		//fmt.Println(word.Value, word.Template)
		words = append(words, word)
	}
	return words, nil
}

func readPage(r *bufio.Reader) (string, error) {
	var result string
	isPageTagOpened := false
	for {
		str, err := r.ReadString('\n')
		if err != nil {
			return "", err
		}
		if !isPageTagOpened {
			if strings.Contains(str, "<page>") {
				isPageTagOpened = true
				result = str
			}
			continue
		}
		result += "\n" + str
		if strings.Contains(str, "</page>") {
			return result, nil
		}
	}
}
