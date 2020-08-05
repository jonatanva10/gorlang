package main

import (
    "fmt"
    "net/http"
    "os"
	"strconv"
)

func handler(writer http.ResponseWriter, request *http.Request) {
    //num := request.URL.Path[1:]
	 anno, err := strconv.Atoi(request.URL.Path[1:])
	if err == nil {
     if calculiBisects(anno) {
             fmt.Fprintf(writer, "El año: %s es bisiesto", anno)
     }else{
             fmt.Fprintf(writer, "El año: %s no es bisiesto", anno)
          }
  }
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
