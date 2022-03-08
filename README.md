# lern-go
lerning golang

## create folder project
````
go mod init github.com/wachira90/goredis
````

## install package
````
go get github.com/go-redis/redis/v8
````

## add mod
````
go mod download github.com/go-redis/redis/v8
````

## check error
````
go build
````

## check version

````
C:\Users\admin>go version
go version go1.17.8 windows/amd64
````


## main.go
````
package main
import ("fmt")

func main() {
  fmt.Println("Hello World!")
}
````


##  run build
````
go run .\main.go
go build .\main.go
````

## example
````
package main

import (
    "fmt"
    "github.com/NanXiao/stack"
)

func main() {
    s := stack.New()
    s.Push(0)
    s.Push(1)
    s.Pop()
    fmt.Println(s)
}
````

## go get package
````
go get github.com/NanXiao/playstack
````

## rest example 
````
package main

import (
    "fmt"
    "log"
    "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
    handleRequests()
}
````
