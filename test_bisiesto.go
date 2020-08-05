package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func calculiBisects(n int) bool {

	if (n % 4 == 0 && n % 100 != 0) || (n % 100 == 0 && n % 400 == 0) {
		//System.out.println("El año " + anio + " es bisiesto");
		return true
	} else {
		return false
	}

}

func plus(a int) int {
	return a * 5
}

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

func main() {
	//fmt.Println("Calculate Bisiesto")

	//fmt.Print("Enter year: ")
	//var input int
	//var otherResult bool
	//fmt.Scanln(&input)

	//otherResult = calculiBisects(input)
	//fmt.Print(otherResult)
	
	http.HandleFunc("/", handler)
    http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}


