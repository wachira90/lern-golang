# lerning golang

```sh
uname -a

export GOOS=linux
export GOARCH=amd64
```

# install linux ubuntu 18.04

```sh
wget https://go.dev/dl/go1.17.8.linux-amd64.tar.gz

tar -xvf go1.17.8.linux-amd64.tar.gz

sudo mv go/ /usr/local/

nano ~/.profile

export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

source .profile
```

## create folder project

```sh
go mod init github.com/wachira90/goredis
```

## install package

```sh
go get github.com/go-redis/redis/v8
```

## add mod

```sh
go mod download github.com/go-redis/redis/v8
```

## check error

```sh
go build
```

## check version

```bat
C:\Users\admin>go version
go version go1.17.8 windows/amd64
```

## main.go

```
package main
import ("fmt")

func main() {
  fmt.Println("Hello World!")
}
```

##  run build

```sh
go run .\main.go
go build .\main.go
```

## example

```
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
```

## go get package

```sh
go get github.com/NanXiao/playstack
```

## rest example 

```
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

func api(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Welcome to the Api!")
  fmt.Println("Endpoint Hit: Api")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/api", api)
    fmt.Println("REST STARTING .... ")
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
    handleRequests()
}
```
