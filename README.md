# Golang Todos API

This is a simple CRUD application developed in Go. As I am from PHP / Laravel background, I tried to adopt some MVC styling. I aim to learn/implement the following:

Go essentials

Routing

Database operations

Pagination

Dependency Injection

JSON Handling

Go Doc

When comparing the performance of Go with PHP / Laravel, I am so impressed and wish to develop all my future projects in Go

Don't take me wrong, PHP is very easy compared to Go (20+ years experienced in PHP) and it took me a while to figure out some implementations. I am still learning Go, will refactor it as I go.

This is not a basic-level application, so I expect you to have some knowledge of Go. I have referred following:

- [Go dev](https://go.dev/)
- [Golang Programmers](https://www.golangprograms.com/)
- [Learn Go](https://www.karanpratapsingh.com/courses/go)


Before clonning, please setup (https://go.dev/doc/install) & configure GOHOME & GOPATH

I will recommend you create the project from scratch like

create a folder my-todo 

cd my-todo 

go mod init

Then create subfolders and go files and keep this application as a reference

### Endpoints

/todos - Get all todos

/todos - Add new todo

/todos/id - Update todo

/todos/id - Delete todo

/todos/id - Get todo

### Packages used

To See the packages used in this application:

go list -f "{{.ImportPath}} {{.Imports}}" ./...

Please download the following packages for this application (please check g.mod for more details):

go get -u gorm.io/gorm

go get -u gorm.io/driver/mysql

go get -u github.com/gorilla/mux

go get -u github.com/joho/godotenv

go get -u github.com/spf13/cast

go get -u gorm.io/datatypes