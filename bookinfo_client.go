package main

import (
    "context"
    "encoding/csv"
    pb "github.com/jonatanva10/booksapp"
    "google.golang.org/grpc"
    "log"
    "os"
    "time"
    "fmt"
)

type Book struct {
    Id        string `json:"id"`
    Title     string `json:"title"`
    Edition   string `json:"edition"`
    Copyright string `json:"copyright"`
    Language  string `json:"language"`
    Pages     string `json:"pages"`
    Author    string `json:"author"`
    Publisher string `json:"publisher"`
}

var books []Book

func checkError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}

func readData(filePath string) []Book {
    file, err1 := os.Open(filePath)
    checkError("Unable to read input file "+filePath, err1)
    defer file.Close()

    csvReader := csv.NewReader(file)
    records, err2 := csvReader.ReadAll()
    checkError("Unable to parse file as CSV for "+filePath, err2)
    defer file.Close()

    books = []Book{}

    for _, record := range records {
        book := Book{
            Id:        record[0],
            Title:     record[1],
            Edition:   record[2],
            Copyright: record[3],
            Language:  record[4],
            Pages:     record[5],
            Author:    record[6],
            Publisher: record[7]}
        books = append(books, book)
    }
    file.Close()

    return books
}

func addBookMain(){
    address := os.Getenv("ADDRESS")
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewBookInfoClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    addBook, err := c.AddBook(ctx, &pb.Book{
        Id:        "1",
        Title:     "Hombre Lobo",
        Edition:   "2th",
        Copyright: "2001",
        Language:  "GERMAN",
        Pages:     "20",
        Author:    "Abraham Lincoln",
        Publisher: "Heredia"})
    if err != nil {
        log.Fatalf("Could not add book: %v", err)
    }

    log.Printf("Book ID: %s added successfully", addBook.Value)
    book, err := c.GetBook(ctx, &pb.BookID{Value: addBook.Value})
    if err != nil {
        log.Fatalf("Could not get book: %v", err)
    }
    log.Printf("Book: ", book.String())
    //log.Printf("Title: ", book.Title)
    //log.Printf("Param Value: ", addBook.Value)
}

func main() {
    address := os.Getenv("ADDRESS")
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewBookInfoClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    /*
    bookOldest:=&pb.Book{
        Id:        "5",
        Title:     "Reina Margarita",
        Edition:   "6th",
        Copyright: "2020",
        Publisher: "Cartago, Costa Rica"}
    r, err := c.AddBook(ctx, bookOldest)
       if err != nil {
           log.Fatalf("Could not add book: %v", err)
       }
    */
    fmt.Println( "Example 1")

    addBook, err := c.AddBook(ctx, &pb.Book{
        Id:        "1",
        Title:     "Hombre Lobo",
        Edition:   "2th",
        Copyright: "2001",
        Language:  "GERMAN",
        Pages:     "20",
        Author:    "Abraham Lincoln",
        Publisher: "Heredia"})
    if err != nil {
        log.Fatalf("Could not add book: %v", err)
    }

    addBook2, error2 := c.AddBook(ctx, &pb.Book{
        Id:        "1",
        Title:     "Hombre Lobo",
        Edition:   "2th",
        Copyright: "2001",
        Language:  "GERMAN",
        Pages:     "20",
        Author:    "Abraham Lincoln",
        Publisher: "Heredia"})
    if error2 != nil {
        log.Fatalf("Could not add book: %v", error2)
    }

    log.Printf("Book ID: %s added successfully", addBook.Value)
    log.Printf("Book ID: %s added successfully", addBook2.Value)
    //log.Printf("Book ID to Delete: ", book.Id)

    //    keyMap:="ce96b38f-e064-47fd-8245-862eb4639bee"
    // Test Get size from map
	//m, err := c.GetSize(ctx, &pb.BookID{Value: addBook.Value})
	//log.Printf("Size: ", m.Size)
	// Test Delete Item

    //keyMap:="8e318ad1-5e7c-498f-a542-cd497ffc18db"
    d, err := c.DeleteBook(ctx, &pb.BookID{Value: addBook.Value})
    log.Printf("Deleted Id: ", d)

    // y, err := c.GetSize(ctx, &pb.BookID{Value: "Test"})
    // log.Printf("Size: ", y.Size)
    // Get All Books

    mapList, errMap := c.GetAllBooks(ctx, &pb.BookID{Value: ""})
    if errMap != nil {
        log.Fatalf("Could not get map: %v", errMap)
    }
    for key, content := range mapList.List {
        log.Printf("%s -> %s\n", key, content)
    }

    // Test Get size from map
    //quantity, err := c.GetSize(ctx, &pb.BookID{Value: addBook.Value})
    //log.Printf("Size: ", quantity.Size)
    fmt.Println( "Example 2")
     bookToUpdate, errorGetBook := c.GetBook(ctx, &pb.BookID{Value: addBook2.Value})
       if errorGetBook != nil {
           log.Fatalf("Could not get book: %v", errorGetBook)
       }
    //deletedId:= bookToUpdate.Id
    update, errorUpdate := c.UpdateBook(ctx, &pb.Book{
        Id:        bookToUpdate.Id,
        Title:     "Salsa en Linea",
        Edition:   "9th",
        Copyright: "1990",
        Language:  "ENGLISH",
        Pages:     "220",
        Author:    "Cris Badilla",
        Publisher: "Alajuela"})
    if errorUpdate != nil {
        log.Fatalf("Could not update book: %v", errorUpdate)
    }
    // Get Updated Book
    bookUpdate, errUp := c.GetBook(ctx, &pb.BookID{Value: update.Id})
    if errUp != nil {
        log.Fatalf("Could not get updated book: %v", errUp)
    }
    log.Printf("Book Updated: ", bookUpdate.String())
    // Get All Books
   /*
    mapList, errMap := c.GetAllBooks(ctx, &pb.BookID{Value: ""})
    if errMap != nil {
        log.Fatalf("Could not get map: %v", errMap)
    }
    for key, content := range mapList.List {
        log.Printf("%s -> %s\n", key, content)
    }
    */

    // Ex 3
    fmt.Println( "Example 3")
    bookList := readData("books.csv")

    for _, tempRecord := range bookList{
        addBook, err := c.AddBook(ctx, &pb.Book{
            Id:        tempRecord.Id,
            Title:     tempRecord.Title,
            Edition:   tempRecord.Edition,
            Copyright: tempRecord.Copyright,
            Language:  tempRecord.Language,
            Pages:     tempRecord.Pages,
            Author:    tempRecord.Author,
            Publisher: tempRecord.Publisher})
        if err != nil {
            log.Fatalf("Could not add book: %v", err)
        }
        log.Printf("Book ID: %s added successfully", addBook.Value)
    }
    // Get All Books
    mapListFromServer, errMap := c.GetAllBooks(ctx, &pb.BookID{Value: ""})
    if errMap != nil {
        log.Fatalf("Could not get map from server: %v", errMap)
    }
    for _, content := range mapListFromServer.List {
        log.Printf("%s\n", content)
    }

    /*addBookMain()
    mapList, errMap := c.GetAllBooks(ctx, &pb.BookID{Value: ""})
    if errMap != nil {
        log.Fatalf("Could not get map: %v", errMap)
    }
    for key, content := range mapList.List {
        log.Printf("%s -> %s\n", key, content)
    }*/
}