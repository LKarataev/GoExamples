package main

import (
	"fmt"
	"reflect"
	"strings"
)

type UnknownPlant struct {
	FlowerType string
	LeafType   string
	Color      int `color_scheme:"rgb"`
}

type AnotherUnknownPlant struct {
	FlowerColor int
	LeafType    string
	Height      int `unit:"inches"`
}

type SomeUnknownPlant struct {
	FlowerName string
	LeafSize   float64
	LeafColor  string `color_scheme:"color_name"`
}

func describePlant(plant interface{}) {
	v := reflect.ValueOf(plant)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		if len(field.Tag) > 0 {
			parts := strings.SplitN(string(field.Tag), ":", 2)
			fmt.Printf("%s(%s=%s): %v\n", field.Name, parts[0], parts[1], value.Interface())
		} else {
			fmt.Printf("%s: %v\n", field.Name, value.Interface())
		}
	}
}

func main() {
	plant1 := UnknownPlant{"Rose", "Lanceolate", 10}
	plant2 := AnotherUnknownPlant{10, "Lanceolate", 15}
	plant3 := SomeUnknownPlant{"Chamomile", 0.4, "White"}

	describePlant(plant1)
	fmt.Println()
	describePlant(plant2)
	fmt.Println()
	describePlant(plant3)
	fmt.Println()
}
