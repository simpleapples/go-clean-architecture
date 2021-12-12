Go  Clean Architecture Project 
===

An example go project using the clean architecture.

## Project Structure

```bash
.
├── README.md
├── adapters
│   ├── api
│   │   ├── api.go
│   │   └── http_api.go
│   └── db
│       ├── db.go
│       └── mysql_db.go
├── app
│   └── main.go
├── entities
│   └── article.go
├── go.mod
├── go.sum
├── scripts
│   ├── build.sh
│   └── run.sh
└── usecases
    ├── article.go
    └── usecases.go
```

## Clean Architecture

Please read [THIS BLOG](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) from Robert C. Martin to know Clean Arthicture. 

![](https://wx4.sinaimg.cn/mw2000/6ae0adaely1gxb4xhga07j210u0ulgro.jpg)
