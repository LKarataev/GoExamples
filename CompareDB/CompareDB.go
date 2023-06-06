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

func main() {
	old_filename := flag.String("old", "", "XML or JSON file containing old cake recipes")
	new_filename := flag.String("new", "", "XML or JSON file containing new cake recipes")
	flag.Parse()

	content, err := ioutil.ReadFile(*old_filename)
	CheckError(err)

	format := GetFileExt(*old_filename)
	old_recipes, err := Unmarshal(content, format)
	CheckError(err)

	content, err = ioutil.ReadFile(*new_filename)
	CheckError(err)

	format = GetFileExt(*new_filename)
	new_recipes, err := Unmarshal(content, format)
	CheckError(err)

	CheckCakes(new_recipes, old_recipes)
}

func CheckCakes(new_recipes CakeRecipes, old_recipes CakeRecipes) {
	for _, new_cake := range new_recipes.Cakes {
		if !FindCake(new_cake.Name, old_recipes) {
			fmt.Printf("ADDED cake \"%s\"\n", new_cake.Name)
		}
	}
	for _, origin_cake := range old_recipes.Cakes {
		if !FindCake(origin_cake.Name, new_recipes) {
			fmt.Printf("REMOVED cake \"%s\"\n", origin_cake.Name)
		} else {
			DiffCakes(origin_cake, new_recipes)
		}
	}
}

func FindCake(name string, recipes CakeRecipes) bool {
	found := false
	for _, recipe := range recipes.Cakes {
		if name == recipe.Name {
			found = true
			break
		}
	}
	return found
}

func DiffCakes(origin_cake Recipe, new_recipes CakeRecipes) {
	for _, new_cake := range new_recipes.Cakes {
		if origin_cake.Name == new_cake.Name {
			if origin_cake.StoveTime != new_cake.StoveTime {
				fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", origin_cake.Name, new_cake.StoveTime, origin_cake.StoveTime)
			}
			DiffIngredients(origin_cake.Ingredients, new_cake.Ingredients, origin_cake.Name)
			break
		}
	}
}

func DiffIngredients(origin_cake, new_cake []Ingredient, cake_name string) {
	origin_ingridients := make(map[string]Ingredient)
	for _, ingridient := range origin_cake {
		origin_ingridients[ingridient.Name] = ingridient
	}

	for _, new := range new_cake {
		origin, ingredient_flag := origin_ingridients[new.Name]
		if !ingredient_flag {
			fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", new.Name, cake_name)
			continue
		}
		delete(origin_ingridients, new.Name)
		if origin.Unit != new.Unit {
			if origin.Unit != "" && new.Unit != "" {
				fmt.Printf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", new.Name, cake_name, new.Unit, origin.Unit)
			} else if origin.Unit == "" {
				fmt.Printf("ADDED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", new.Unit, new.Name, cake_name)
			} else if new.Unit == "" {
				fmt.Printf("REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", origin.Unit, new.Name, cake_name)
			}
		}
		if origin.Count != new.Count {
			fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%.1f\" instead of \"%.1f\"\n", origin.Name, cake_name, new.Count, origin.Count)
		}
	}

	for _, removed := range origin_ingridients {
		fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", removed.Name, cake_name)
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

func Unmarshal(content []byte, format string) (CakeRecipes, error) {
	recipes := CakeRecipes{}
	switch format {
	case "json":
		return recipes, json.Unmarshal(content, &recipes)
	case "xml":
		return recipes, xml.Unmarshal(content, &recipes)
	default:
		return recipes, fmt.Errorf("unsupported file format: %s", format)
	}
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
