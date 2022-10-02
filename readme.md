# Go Clean Architecture

Example that shows core principles of the Clean Architecture in Golang projects.

## Rules of Clean Architecture by Uncle Bob:

- Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
- Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
- Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
- Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
- Idependent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.

More on topic can be found <a href="https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html">here</a>.

![clean](images/clean.jpg "clean")

#### Structure:

4 Domain layers:

- Models layer
- Repository layer
- Business Logics layer
- Controller or Handler layer

## REST API:

### Routes

|                          METHOD                          |              URL               |   Description   |
| :------------------------------------------------------: | :----------------------------: | :-------------: |
|      ![GET](https://img.shields.io/badge/-GET-blue)      |  http://localhost:8081/users   |   GetAllUsers   |
|      ![GET](https://img.shields.io/badge/-GET-blue)      | http://localhost:8081/user/:id |   GetUserByID   |
| ![DELETE](https://img.shields.io/badge/-DELETE-critical) | http://localhost:8081/user:id  | DeleteUserByID  |
|   ![POST](https://img.shields.io/badge/-POST-success)    | http://localhost:8081/register |    Register     |
|   ![POST](https://img.shields.io/badge/-POST-success)    |  http://localhost:8081/login   |      Login      |
|     ![PUT](https://img.shields.io/badge/-PUT-orange)     |   http://localhost:8081/user   | Change Password |

### PAYLOAD

##### REGISTER

```
{
    "first_name" : "Aditya",
    "last_name" : "Erlangga",
    "email" : "aditya@gmail.com",
    "password" : "aditya123"
}
```

##### LOGIN

```
{
    "email" : "aditya@gmail.com",
    "password" : "aditya123"
}
```

##### CHANGE PASSWORD

```
{
    "id" : 3,
    "old_password" : "aditya123",
    "new_password" :"hi_everyone!"
}
```

## Requirements

- go 1.19.1

## Setup

Do this :

```
$ git clone https://github.com/adityaerlangga/go-clean-architecture
$ go mod download
$ go run main.go
```
