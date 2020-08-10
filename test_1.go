package main

import (
    "fmt"
    "net/http"
    "os"
    "strconv"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	
	year, err := strconv.Atoi(request.URL.Path[1:])
	if err == nil {
		if calculiBisects(year) {
			fmt.Fprintf(writer, "El año: %s es bisiesto", request.URL.Path[1:])
		}else{
			fmt.Fprintf(writer, "El año: %s no es bisiesto", request.URL.Path[1:])
		}
	} else{
		fmt.Fprintf(writer, "Parametros incorrectos")
	}
}

func calculiBisects(n int) bool {
	if (n % 4 == 0 && n % 100 != 0) || (n % 100 == 0 && n % 400 == 0) {	
		return true
	} else {
		return  false
	}
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
