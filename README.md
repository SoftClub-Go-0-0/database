### Пакет имитирующий базу данных блог-платформы

# database

## Установка

    go get github.com/softclub-go-0-0/database

## Типы

    type Post struct {
        ID        string
        Content   string
        IsVisible bool
        Rating    int
        Tags      []string
    }
        

## Переменные

    Posts - слайс объектов типа Post. Содержит список тестовых данных 50 постов.