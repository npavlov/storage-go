package main

import (
	"fmt"
	"github.com/npavlov/storage-go/internal/storage"
	"log"
)

func main() {
	fmt.Println("Hello World")

	st := storage.NewStorage()

	file, err := st.Upload("test", []byte("HELLO"))

	if err != nil {
		log.Fatal(err)
	}

	restoredFile, err := st.GetById(file.Id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("it uploaded", restoredFile)
}
