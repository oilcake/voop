package player

import (
	"io/ioutil"
	"log"
	"math/rand"
	"path/filepath"
	"voop/sync"
)

var (
	SupportedTypes = [...]string{".mp4", ".mpg", ".mov", ".avi", ".wmv", ".mkv"}
)

type Set struct {
	RightNow int
	Size     int
	Kit      []*Media
}

func NewSet(path *string, t *sync.Transport) (*Set, error) {
	files := SupportedFiles(path)
	opened := make([]*Media, len(files))
	var err error
	for i, file := range files {
		opened[i], err = NewMedia(file, t)
		if err != nil {
			return nil, err
		}
	}
	set := &Set{
		RightNow: 0,
		Size:     len(opened),
		Kit:      opened,
	}
	return set, nil
}

func CloseSet(s *Set) {
	for _, clip := range s.Kit {
		clip.Close()
	}
}

// Navigation
func (s *Set) Now() *Media {
	return s.Kit[s.RightNow]
}

func (s *Set) Random() {
	s.RightNow = rand.Intn(s.Size - 1)
}

func (s *Set) Next() {
	s.RightNow = (s.RightNow + 1) % s.Size
}

func (s *Set) Previous() {
	s.RightNow = (s.RightNow - 1) % s.Size
}

// Handy functions
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

func Supported(file string) bool {
	ext := filepath.Ext(file)
	for _, t := range SupportedTypes {
		if ext == t {
			return true
		}
	}
	return false
}
