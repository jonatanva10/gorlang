package main

import (
    "encoding/json"
    "net/http"
    "path"
	"sort"
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
	bookTemp := Book{}
    for i, book := range books {
        if x == book.Id {
			bookTemp:= books[i]
            return bookTemp
        }
    }	
    return bookTemp
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
		getAllData, e := json.Marshal(books[1:])//books[1:]
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
	id := path.Base(r.URL.Path)
	bookSelected:= findObject(id)
	
	// Remove Element Duplicated
	i := find(id)  
	copy(books[i:], books[i+1:]) 
	books[len(books)-1] = Book{}     
	books = books[:len(books)-1] 
	
	err2 := json.NewDecoder(r.Body).Decode(&bookSelected) // Replace all fields
	if err2 != nil{
		panic(err2)
	}

	userJson, err := json.Marshal(bookSelected) //user
	if err != nil{
		panic(err2)
	}

	w.Header().Set("Content-Type", "application/json")		
	w.Write(userJson)	
	
	books = append(books, bookSelected) // Add last position
	
	// Sorting Array
	sort.SliceStable(books[0:], func(i, j int) bool {
		return books[i].Id < books[j].Id
	})	
	
	headers:= books[len(books)-1]
	books = books[:len(books)-1]
	books = append([]Book{headers}, books...) // Add start position
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