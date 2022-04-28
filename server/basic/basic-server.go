package main
import ("net/http")

func main() {
	http.HandleFunc("/", GoHandler)
	http.ListenAndServe(":8888", nil)
}

func GoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
