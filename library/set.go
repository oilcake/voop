package library

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"voop/clip"
	"voop/sync"
)

type Set struct {
	Read
	Kit []*clip.Media
}

func (s *Set) What(i int) interface{} {
	return s.Kit[i]
}

func NewSet(path *string, t *sync.Transport) (*Set, error) {
	files := SupportedFilesFrom(path)
	opened := make([]*clip.Media, len(files))
	var err error
	for i, file := range files {
		log.Printf("Index of file being opened %v\n", i)
		opened[i], err = clip.NewMedia(file, t)
		if err != nil {
			// return nil, err
			log.Printf("\nWarning\n %#v", err)
		}
	}
	set := &Set{Read{
		RightNow: 0,
		Size:     len(opened),
	},
		opened,
	}
	return set, nil
}

func CloseSet(s *Set) {
	for _, clip := range s.Kit {
		log.Println("\nclosing file ", clip.Name)
		clip.Close()
	}
}

// Handy functions
func SupportedFilesFrom(path *string) (sf []string) {
	sf = make([]string, 0)
	if err := AddFiles(&sf, path); err != nil {
		log.Fatal("can't open folder", err)
	}
	log.Println("files total", len(sf))
	return
}

func AddFiles(sf *[]string, path *string) error {
	files, err := ioutil.ReadDir(*path)
	if err != nil {
		return err
	}

	for _, file := range files {
		if Supported(file.Name()) {
			f := *path + "/" + file.Name()
			*sf = append(*sf, f)
		}
		if file.IsDir() {
			fp := *path + "/" + file.Name()
			err = AddFiles(sf, &fp)
			if err != nil {
				return err
			}
		}
	}
	return nil
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
