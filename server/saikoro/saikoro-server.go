package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/", SaikoroHandler)
	http.ListenAndServe(":1234", nil)
}

func SaikoroHandler(w http.ResponseWriter, r *http.Request) {
	v := rand.Intn(6) + 1
	s := fmt.Sprintf("ダイス: %d", v)
	w.Write([]byte(s))
	println("ダイス: ", v)
}
