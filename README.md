# gb rest full api [![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)
***
RESTful API (API передачи репрезентативного состояния) — это тип архитектуры веб-сервиса, который использует HTTP-запросы для выполнения операций CRUD (создание, чтение, обновление, удаление) над ресурсами. 

## Описание
__GB RESTfull API__ был разработан под специальцно спроектированную базу, которая имеет несколько сущеностей с ограничениями и зависимости между ними, что позволяет обеспечить целостность данных. Так же важно подметить что проект предпологает автоидексацию сущностей, при реализации вставки отствует поле id, которая в базе имеет тип __serial__. Ниже представлена диаграмма базы. 

![Диаграмма бд](https://media.discordapp.net/attachments/1217879745903988838/1217879775851446342/TOlKk6PABbM.png?ex=661816ed&is=6605a1ed&hm=eafb51c03c96c6c2b8b1da71ece4002b1ba2a98dc889083283e746cd08bb5c1f&=&format=webp&quality=lossless&width=866&height=527)

Как можно заметить в качестве тематики используется, предприятие по продаже [гроубоксов](https://en.wikipedia.org/wiki/Grow_box)-специализированных палаток для выращивания растений расщитанные на размещение в закрытых помещения, будто огромный ангар или собственная квартира. Сама база воссоздана при помощи PostgreSQL. В самом проекте реализованы следующие http-запросы:


1. __GET__ api/GrowBox/ - выдача всех гроубоксов
2. __GET__ api/GrowBox/__id__ - выдача гроубокса по индексу
3. __PATCH__ api/GrowBox/__id__ - замена по индексу
4. __POST__  api/GrowBox/ - добавление 
5. __DELETE__ api/GrowBox/__id__ -удаление

## В будущем

- Привязка к интерфесу 
- Добавление новых http-запросов
- Редактирование и исправления ошибок
- Реализация фильтрации
- Осуществление авторизации

>_В дальнейшем все выше описанное будет не раз редактироватся дополнятся_



Далее нужно настроить подключение к базе.

## Техническая составляющая

Необходимо открыть файл __database.go__ папку config:
```sh
cd config
```
И в той части когда где инициализируются константы надо задать те данные которые используются для работы с самой субд, значения могут разнится, ниже приведен пример таких значений:
```go
const (
	host     = "localhost"
	port     = "1234"
	user     = "postgres"
	password = "123456"
	dbName   = "postgres"
)
```

## Установка
Установка производится в процессе стандартной инициализации проекта при помощи команды __init__:
```go
go mod init application/gb
```

## Пакеты
При работе проекта используются определенные пакеты, при желании можно поменять на те которые более удобные пользователю. Ниже представлена таблица с их наименование и описанием
|        Название        |        Описание        |
|------------------------|:-----------------------|
| [HttpRouter](https://github.com/julienschmidt/httprouter) | маршрутизатор HTTP-запросов также называемый мультиплексором |
| [pq](https://pkg.go.dev/github.com/lib/pq@v1.10.9)|драйвер posgresql, для взаимодествия с субд|          
| [log](https://pkg.go.dev/github.com/rs/zerolog@v1.32.0/log)|используется для вывода полноформатного сообщения, сигнализирующая работу запущенного сервера|



