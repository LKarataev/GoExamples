package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Ingredient struct {
	Name  string  `json:"ingredient_name" xml:"itemname"`
	Count float64 `json:"ingredient_count,string" xml:"itemcount"`
	Unit  string  `json:"ingredient_unit" xml:"itemunit"`
}

type Recipe struct {
	Name        string       `json:"name" xml:"name"`
	StoveTime   string       `json:"time" xml:"stovetime"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>item"`
}

type CakeRecipes struct {
	Cakes []Recipe `json:"cake" xml:"cake"`
}

type DBReader interface {
	Read(content []byte) (CakeRecipes, error)
}

type JSONReader struct{}

func (j *JSONReader) Read(content []byte) (CakeRecipes, error) {
	recipes := CakeRecipes{}
	err := json.Unmarshal(content, &recipes)
	return recipes, err
}

type XMLReader struct{}

func (x *XMLReader) Read(content []byte) (CakeRecipes, error) {
	recipes := CakeRecipes{}
	err := xml.Unmarshal(content, &recipes)
	return recipes, err
}

func main() {
	filename := flag.String("f", "", "XML or JSON file containing cake recipes")
	flag.Parse()

	content, err := ioutil.ReadFile(*filename)
	CheckError(err)

	format := GetFileExt(*filename)

	reader, err := GetReader(format)
	CheckError(err)

	recipes, err := reader.Read(content)
	CheckError(err)

	format = ReverseFormat(format)
	output, err := Marshal(recipes, format)
	CheckError(err)

	fmt.Printf("Recipes in %s format:\n%s\n", format, string(output))
}

func GetReader(format string) (DBReader, error) {
	switch format {
	case "json":
		return &JSONReader{}, nil
	case "xml":
		return &XMLReader{}, nil
	default:
		return nil, fmt.Errorf("unsupported file format: %s", format)
	}
}

func GetFileExt(filename string) string {
	if filename == "" {
		return ""
	}
	parts := strings.Split(filename, ".")
	if len(parts) < 2 {
		return ""
	}
	return parts[len(parts)-1]
}

func Marshal(recipes CakeRecipes, format string) ([]byte, error) {
	switch format {
	case "json":
		return json.MarshalIndent(recipes, "", "    ")
	case "xml":
		return xml.MarshalIndent(recipes, "", "    ")
	default:
		return nil, fmt.Errorf("unsupported file format: %s", format)
	}
}

func ReverseFormat(format string) string {
	if format == "json" {
		return "xml"
	} else {
		return "json"
	}
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
