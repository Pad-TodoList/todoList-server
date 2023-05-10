# todoList-server

![example workflow](https://github.com/Pad-TodoList/todoList-server/actions/workflows/server.yml/badge.svg?event=push)

Server using [go 1.20](https://go.dev/)

## Server information

Server deploy using [render](https://render.com/)

URL https://todolist-fullstack.onrender.com

```/about.json```to get all available routes

## Usage

### Everyone

After installing go, run :

```bash
go run src/main.go                    # run main
go build -o $BINARY_NAME src/main.go  # generate binary
```
### Linux/Mac users :

you can use the Makefile + [Docker](https://www.docker.com/)

``` bash
make start      # start docker-compose
make stop       # stop docker-compose
make restart    # stop and start the server
```
