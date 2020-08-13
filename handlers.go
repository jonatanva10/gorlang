package main

import (
    "encoding/json"
    "net/http"
    "path"
)

func find(x string) int {
    for i, book := range books {
        if x == book.Id {
            return i
        }
    }
    return -1
}

func findObject(x string) Book {
    for i, book := range books {
        if x == book.Id {
            return books[i]
        }
    }
    return Book{ Id:"Cachi", Title: "Pelicula 1"}
}

func findByTitle(x string) int {
    for i, book := range books {
        if x == book.Title {
            return i
        }
    }
    return -1
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
    id := path.Base(r.URL.Path)
    checkError("Parse error", err)
    i := find(id)
    if i == -1 {		
		getAllData, e := json.Marshal(books[1:])
		w.Header().Set("Content-Type", "application/json")
		w.Write(getAllData) 
		return e      
    }
    dataJson, err := json.Marshal(books[i])
    w.Header().Set("Content-Type", "application/json")
    w.Write(dataJson)
    return
}

func handleGetObject(w http.ResponseWriter, r *http.Request) (err error) {
    id := path.Base(r.URL.Path)
    checkError("Parse error", err)
    i := findObject(id)
    dataJson, err := json.Marshal(i)
    w.Header().Set("Content-Type", "application/json")
    w.Write(dataJson)
    return
}

func (u *Book) Modify() {
  u.Title = "Paul"
  return 
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
    len := r.ContentLength
    body := make([]byte, len)
    r.Body.Read(body)
    book := Book{}
    json.Unmarshal(body, &book)
    books = append(books, book)
    w.WriteHeader(200)
    return
}

func handlePostTemp(w http.ResponseWriter, r *http.Request) (err error) {

	//i := find(id)
	//book := Book{} //initialize empty user

	//tempBook := json.NewDecoder(r.Body).Decode(&book)
	tempBook, err := json.Marshal(r.Body)
    if tempBook != nil {
        //http.Error(w, tempBook.Error(), http.StatusBadRequest)
        return
    }
	w.Header().Set("Content-Type", "application/json")
	w.Write(tempBook) 
    return 
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	//id := path.Base(r.URL.Path)
    //checkError("Parse error", err)
    //i := find(id)   		
	//copy(books[i:], books[i+1:]) 
	//books[len(books)-1] = Book{}     
	//books = books[:len(books)-1]  
	//tempList, err := json.Marshal(books[1:])

	//initialize empty user
	user := Book{} 

	//Parse json request body and use it to set fields on user
	//Note that user is passed as a pointer variable so that it's fields can be modified
	err2 := json.NewDecoder(r.Body).Decode(&user)
	if err2 != nil{
		panic(err2)
	}

	//Set CreatedAt field on user to current local time
	user.Author = user.Title

	//Marshal or convert user object back to json and write to response 
	userJson, err := json.Marshal(user)
	if err != nil{
		panic(err2)
	}

	//Set Content-Type header so that clients will know how to read response
	w.Header().Set("Content-Type", "application/json")	
	//Write json response back to response 
	w.Write(userJson)
	books = append(books, user)
	return
	}

func RemoveIndex(list []Book, index int) []Book {
    return append(list[:index], list[index+1:]...)
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
    id := path.Base(r.URL.Path)
    checkError("Parse error", err)
    i := find(id)   		
	copy(books[i:], books[i+1:]) 
	books[len(books)-1] = Book{}     
	books = books[:len(books)-1]  
	tempList, err := json.Marshal(books[1:])
	w.Header().Set("Content-Type", "application/json")
	w.Write(tempList) 
	return 
}