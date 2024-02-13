

## Crawler

Crawler - это пример программы парсинга URL-адресов. Функция crawlWeb принимает входной канал (для получения URL-адресов) и возвращает другой канал для результатов сканирования (тело веб-страницы в виде строки). Кроме того, мы ограничиваем объем выполняемой им работы. В любой момент времени не может быть более 8 горутин, параллельно запрашивающих страницы. Также мы имеем возможность остановить процесс сканирования в любой момент, нажав Ctrl+C (и наш код выполняет плавное завершение работы - "graceful shutdown"). Для этого мы используем отменяемый контекст (context.WithCancel). Если программа не прервана, она должна корректно остановится после обработки всех заданных URL-адресов.

Как протестировать:

В main функции циклом проходимся по статическим ссылкам:
```go
    for i := 6; i < 30; i++ {
        time.Sleep(1 * time.Second)
        urls <- fmt.Sprintf("https://rickandmortyapi.com/api/character/%d", i)
    }
```

Запускаем:
```bash
go run Crawler.go
```

Вывод:
```output
Received result: {"id":6,"name":"Abadango Cluster Princess","status":"Alive","species":"Alien","type":"","gender":"Female","origin":{"name":"Abadango","url":"https://rickandmortyapi.com/api/location/2"},"location":{"name":"Abadango","url":"https://rickandmortyapi.com/api/location/2"},"image":"https://rickandmortyapi.com/api/character/avatar/6.jpeg","episode":["https://rickandmortyapi.com/api/episode/27"],"url":"https://rickandmortyapi.com/api/character/6","created":"2017-11-04T19:50:28.250Z"}
Received result: {"id":7,"name":"Abradolf Lincler","status":"unknown","species":"Human","type":"Genetic experiment","gender":"Male","origin":{"name":"Earth (Replacement Dimension)","url":"https://rickandmortyapi.com/api/location/20"},"location":{"name":"Testicle Monster Dimension","url":"https://rickandmortyapi.com/api/location/21"},"image":"https://rickandmortyapi.com/api/character/avatar/7.jpeg","episode":["https://rickandmortyapi.com/api/episode/10","https://rickandmortyapi.com/api/episode/11"],"url":"https://rickandmortyapi.com/api/character/7","created":"2017-11-04T19:59:20.523Z"}
...
```
