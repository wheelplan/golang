package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

type document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"text"`
	ID    int    `xml:"id"`
}

func main() {
	tmp, _ := loadDocument("C:\\tmp\\wiki.xml")
	res := search(tmp, "cat")
	fmt.Println(res)
}

func loadDocument(path string) ([]document, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	dec := xml.NewDecoder(f)
	dump := struct {
		Documents []document `xml:"doc"`
	}{}
	if err := dec.Decode(&dump); err != nil {
		return nil, err
	}

	docs := dump.Documents
	for i := range docs {
		docs[i].ID = i
	}
	return docs, nil
}

func search(docs []document, term string) []document {
	var r []document
	for _, doc := range docs {
		if strings.Contains(doc.Text, term) {
			r = append(r, doc)
		}
	}
	return r
}
