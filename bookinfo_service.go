package main

import (
    "context"
    "github.com/gofrs/uuid"
    pb "github.com/jonatanva10/booksapp"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "strconv"
)

type server struct {
    bookMap map[string]*pb.Book
}

func (s *server) AddBook(ctx context.Context, in *pb.Book) (*pb.BookID, error) {
    out, err := uuid.NewV4()
    if err != nil {
        return nil, status.Errorf(codes.Internal,
            "Error while generating Book ID", err)
    }
    in.Id = out.String()
    if s.bookMap == nil {
        s.bookMap = make(map[string]*pb.Book)
    }
    s.bookMap[in.Id] = in
    //tempObject:= s.bookMap[in.Id] // Get Object from Map
    return &pb.BookID{Value: in.Id}, status.New(codes.OK, "").Err()
    //return &pb.BookID{Value: tempObject.Language}, status.New(codes.OK, "").Err()
}

func (s *server) GetBook(ctx context.Context, in *pb.BookID) (*pb.Book, error) {
    value, exists := s.bookMap[in.Value]
    if exists {
        return value, status.New(codes.OK, "").Err()
    }
    return nil, status.Errorf(codes.NotFound, "Book does not exist.", in.Value)
}

func (s *server) GetSize(ctx context.Context, in *pb.BookID) (*pb.BookSize, error) {
    return &pb.BookSize{Size: strconv.Itoa(len(s.bookMap))}, status.New(codes.OK, "").Err()
}

func (s *server) DeleteBook(ctx context.Context, in *pb.BookID) (*pb.BookID, error) {
    _, exists := s.bookMap[in.Value]
    if exists {
        //result:=1
        delete(s.bookMap, in.Value)
        return &pb.BookID{Value: in.Value}, status.New(codes.OK, "").Err()
    }
    return nil, status.Errorf(codes.NotFound, "Error Book deleted.", in.Value)
}

func (s *server) GetAllBooks(ctx context.Context, in *pb.BookID) (*pb.BookMapList, error) {
    mapList:=make(map[string]*pb.Book)
    mapList=s.bookMap
    if  mapList == nil {
        //return nil, status.Errorf(codes.NotFound, "No Elements.", mapList)
        return nil, status.Errorf(codes.NotFound, "No Elements In Map.", mapList)
    }
    example:= &pb.BookMapList{List: mapList}
    return example, status.New(codes.OK, "").Err()
}

func (s *server) UpdateBook(ctx context.Context, param *pb.Book) (*pb.Book, error) {
    keyBook:=param.Id
    currentBook, exists := s.bookMap[keyBook]
    if exists {
        // Send Complete Fields
        newBook:= &pb.Book{}
        newBook=param
        delete(s.bookMap, keyBook)
        s.bookMap[keyBook] = newBook

        return s.bookMap[keyBook], status.New(codes.OK, "").Err()
    } else {
        return nil, status.Errorf(codes.NotFound, "Error Book Updated Not Exists.", currentBook)
    }
    return nil, status.Errorf(codes.NotFound, "Book cannot update.", param.Id)
}

/*func isNilBetter(dog *pb.Book) bool {
    var ret bool

    //v := dog
    ret = dog.Author == "" || dog.Language
    if dog.Author == "" {
        fmt.Println("s2.Property has not been set")
    }

    return ret
}*/