

## CompareTXT

CompareTXT - это маленькая CLI утилита, которая сравнивает два текстовых файла по содержащимся в них строкам.

Если в добавили строку, которой нет в другом файле, то утилита выведет:
```output
> ADDED *содержимое добавленной строки*
```

Если какая-то строка удалена:
```output
> REMOVED *содержимое удалённой строки*
```

Файлы для теста:
 - old.txt
 - new.txt

Для запуска вам нужно указать в опциях *--old* и *--new* пути до этих файлов.

Как протестировать:

```bash
go run CompareTXT.go --old old.txt --new new.txt
```

Вывод:
```output
> ADDED /Users/baker/recipes/database.xml
> ADDED /Users/baker/recipes/database_version3.yaml
> REMOVED /var/log/orders.log
> REMOVED /var/log/orders2.log
```
