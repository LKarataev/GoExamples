

## ConvertRecipe

ConvertRecipe - это маленькая CLI утилита, которая читает файл с рецептами и выводит его в другом формате. Если у файла JSON формат, то вывод будет в XML формате. И наоборот, если у файла XML формат, то вывод будет в JSON формате. При этом учитывается, что имена полей могут немного различаться, например в *json: "time"*, а в *xml: "stovetime"*.

Файлы для теста:
 - recipes.xml
 - recipes.json

Для запуска вам нужно указать в опции *-f* путь до файла.

Как протестировать:

1) JSON в XML

```bash
go run ConvertRecipe.go -f recipes.json
```

Вывод:
```output
<CakeRecipes>
    <cake>
        <name>Red Velvet Strawberry Cake</name>
        <stovetime>45 min</stovetime>
        <ingredients>
            <item>
                <itemname>Flour</itemname>
                <itemcount>2</itemcount>
                <itemunit>mugs</itemunit>
            </item>
            <item>
                <itemname>Strawberries</itemname>
                <itemcount>8</itemcount>
                <itemunit></itemunit>
            </item>
            <item>
                <itemname>Coffee Beans</itemname>
                <itemcount>2.5</itemcount>
                <itemunit>tablespoons</itemunit>
            </item>
            <item>
                <itemname>Cinnamon</itemname>
                <itemcount>1</itemcount>
                <itemunit></itemunit>
            </item>
        </ingredients>
    </cake>
    <cake>
        <name>Moonshine Muffin</name>
        <stovetime>30 min</stovetime>
        <ingredients>
            <item>
                <itemname>Brown sugar</itemname>
                <itemcount>1</itemcount>
                <itemunit>mug</itemunit>
            </item>
            <item>
                <itemname>Blueberries</itemname>
                <itemcount>1</itemcount>
                <itemunit>mug</itemunit>
            </item>
        </ingredients>
    </cake>
</CakeRecipes>
```

2) XML в JSON

```bash
go run ConvertRecipe.go -f recipes.xml
```

Вывод:
```output
{
    "cake": [
        {
            "name": "Red Velvet Strawberry Cake",
            "time": "40 min",
            "ingredients": [
                {
                    "ingredient_name": "Flour",
                    "ingredient_count": "3",
                    "ingredient_unit": "cups"
                },
                {
                    "ingredient_name": "Vanilla extract",
                    "ingredient_count": "1.5",
                    "ingredient_unit": "tablespoons"
                },
                {
                    "ingredient_name": "Strawberries",
                    "ingredient_count": "7",
                    "ingredient_unit": ""
                },
                {
                    "ingredient_name": "Cinnamon",
                    "ingredient_count": "1",
                    "ingredient_unit": "pieces"
                }
            ]
        },
        {
            "name": "Blueberry Muffin Cake",
            "time": "30 min",
            "ingredients": [
                {
                    "ingredient_name": "Baking powder",
                    "ingredient_count": "3",
                    "ingredient_unit": "teaspoons"
                },
                {
                    "ingredient_name": "Brown sugar",
                    "ingredient_count": "0.5",
                    "ingredient_unit": "cup"
                },
                {
                    "ingredient_name": "Blueberries",
                    "ingredient_count": "1",
                    "ingredient_unit": "cup"
                }
            ]
        }
    ]
}
```
