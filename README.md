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

    type Person struct {
        ID      uint   // уникальный идентификатор
        Name    string // имя
        Surname string // фамилия
        Phone   string // номер телефона
    }
    
    type Student struct {
        Person                        // студент является человеком, поэтому нам нужны все поля структуры Person
        GroupID        uint           // идентификатор группы, в котором учится студент
        Grades         map[string]int // баллы по неделям
        TopicsToWorkOn []string       // темы, над которыми студенту нужно поработать
        Attendance     map[int]int    // присутствие на занятиях по дням месяца
    }

## Переменные

    Posts - слайс объектов типа Post. Содержит список тестовых данных 50 постов.
    
    Students - слайс объектов типа Student. Содержит список тестовых данных 14 студентов.
