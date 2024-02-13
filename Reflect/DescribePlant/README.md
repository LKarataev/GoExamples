

## DescribePlant

DescribePlant - это программа, которая демонстрирует возможности рефлексии (пакет `reflect`). Функция `describePlant` принимает интерфейс `plant` в качестве параметра. Если `plant` - структура, то функция анализирует её поля. Она выводит имена полей и их значения. Если поле имеет тег, она выводит также тип тега.

Протестировать:

```bash
go run .
```

Для примера берутся следующие структуры:
```go
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
```

Затем мы создаём заполненные экземпляры этих структур:
```go
plant1 := UnknownPlant{"Rose", "Lanceolate", 10}
plant2 := AnotherUnknownPlant{10, "Lanceolate", 15}
plant3 := SomeUnknownPlant{"Chamomile", 0.4, "White"}
```

После мы передаём эти структуры в функцию `describePlant` и получаем следующий вывод:
```output
FlowerType: Rose
LeafType: Lanceolate
Color(color_scheme="rgb"): 10

FlowerColor: 10
LeafType: Lanceolate
Height(unit="inches"): 15

FlowerName: Chamomile
LeafSize: 0.4
```
