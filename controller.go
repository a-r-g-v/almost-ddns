package main

import (
	_ "encoding/json"
	"fmt"
	"net/http"

	_ "github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zenazn/goji/web"
)

func LogAPIContoller(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "work")
}
func StatAPIContoller(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "work")
}
