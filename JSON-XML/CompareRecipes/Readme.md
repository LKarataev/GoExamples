

## CompareRecipes

CompareRecipes - это маленькая CLI утилита, которая сравнивает между собой два файла с кулинарными рецептами. Файлы могут быть как JSON, так и XML формата.

Предположим, у вас есть файл с рецептами, вы эксперементировали с ними и со временем вносили изменения. Также у вас остался старый файл, где ещё нет новых изменений в рецептах. И вы хотите сравнить - как поменялись ваши рецепты. Тогда такая утилита может вам помочь. 

Файлы для теста:
 - old_recipes.xml
 - new_recipes.json

Для запуска вам нужно указать в опциях *--old* и *--new* пути до этих файлов.

Как протестировать:

```bash
go run CompareRecipes.go --old old_recipes.xml --new new_recipes.json
```

Вывод:
```output
> ADDED cake "Moonshine Muffin"
> CHANGED cooking time for cake "Red Velvet Strawberry Cake" - "45 min" instead of "40 min"
> CHANGED unit for ingredient "Flour" for cake "Red Velvet Strawberry Cake" - "mugs" instead of "cups"
> CHANGED unit count for ingredient "Flour" for cake "Red Velvet Strawberry Cake" - "2.0" instead of "3.0"
> CHANGED unit count for ingredient "Strawberries" for cake "Red Velvet Strawberry Cake" - "8.0" instead of "7.0"
> ADDED ingredient "Coffee Beans" for cake "Red Velvet Strawberry Cake"
> REMOVED unit "pieces" for ingredient "Cinnamon" for cake "Red Velvet Strawberry Cake"
> REMOVED ingredient "Vanilla extract" for cake "Red Velvet Strawberry Cake"
> REMOVED cake "Blueberry Muffin Cake"
```
