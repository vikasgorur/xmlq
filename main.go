package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		fmt.Printf("usage: xmlq [filename]")
		os.Exit(1)
	}

	xmlFile, err := os.Open(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)
	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}
		switch startElement := token.(type) {
		case xml.StartElement:
			fmt.Println(startElement.Name.Local)
		}
	}
}
