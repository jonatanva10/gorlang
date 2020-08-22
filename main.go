/*
 * Books API
 *
 * This web service offers information on books
 *
 * API version: 0.1.9
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package main

import (
	"log"
    "net/http"
    "os"
	sw "github.com/jonatanva10/gorlang/go"
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
	//log.Fatal(http.ListenAndServe(":8080", router))
}