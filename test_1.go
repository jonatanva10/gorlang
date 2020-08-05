package main

import (
    "fmt"
    "net/http"
    "os"
)

func handler(writer http.ResponseWriter, request *http.Request) {
    num := request.URL.Path[1:]
    fmt.Fprintf(writer, calculiBisects(num), request.URL.Path[1:])
}

func calculiBisects(n int) string {
	d := "Sin Datos"
	if (n % 4 == 0 && n % 100 != 0) || (n % 100 == 0 && n % 400 == 0) {
		
		d  = "Año Bisiesto"
	} else {
		d  = "Año No Bisiesto"
	}

	return d
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
