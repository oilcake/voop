package player

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	"voop/library"
)

// Play functions
type Reader interface {
	Now() int
	Random()
	Next()
	Previous()
	What(index int) interface{}
}

func ChooseRandomFile(path string) (string, error) {

	files := library.SupportedFilesFrom(path)
	log.Println("files total", len(files))
	rand.Seed(time.Now().UnixNano())
	file := files[rand.Intn(len(files)-1)]
	fmt.Println()
	log.Printf("Playing file %v\n", file)
	return file, nil
}
