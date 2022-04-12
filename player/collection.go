package player

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

var (
	SupportedTypes = [...]string{".mp4", ".mpg", ".mov", ".avi", ".wmv", ".mkv"}
)

type Collection []*Media

func OpenFolder(path *string) (Collection, error) {
	files := SupportedFiles(path)
	opened := make([]*Media, len(files))
	var err error
	for i, file := range files {
		opened[i], err = NewMedia(file)
		if err != nil {
			return nil, err
		}
	}
	return opened, nil
}

func CloseFolder(c Collection) {
	for _, clip := range c {
		clip.Close()
	}
}

func SupportedFiles(path *string) (sf []string) {
	sf = make([]string, 0)
	files, err := ioutil.ReadDir(*path)
	if err != nil {
		log.Fatal("can't open folder", err)
	}
	log.Println("files total", len(files))
	for _, file := range files {
		if Supported(file.Name()) {
			f := *path + "/" + file.Name()
			sf = append(sf, f)
		}
	}
	return
}

// f := *path + "/" + file.Name()
func Supported(file string) bool {
	ext := filepath.Ext(file)
	for _, t := range SupportedTypes {
		if ext == t {
			return true
		}
	}
	return false
}
