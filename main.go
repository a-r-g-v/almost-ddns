package main

import (
	"fmt"
	"github.com/zenazn/goji"
	"net/http"

	"github.com/goji/httpauth"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

func main() {

	go work()
	goji.Use(httpauth.SimpleBasicAuth("a", "a"))
	goji.Get("/v1/log", LogAPIContoller)
	goji.Get("/v1/stat", StatAPIContoller)
	goji.Serve()

}

func work() {

	for {
		time.Sleep(60 * time.Second)
		fmt.Println("Yo")
	}
}
