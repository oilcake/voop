package library

import (
	"io/ioutil"
)

// type Folder string

type Library struct {
	read
	index []*cat
}

func (l *Library) What(i int) *cat {
	return l.index[i]
}

func NewLibrary(path *string) (l *Library, err error) {
	d := make([]*cat, 0)
	dirs, err := ioutil.ReadDir(*path)
	if err != nil {
		return nil, err
	}

	for _, dir := range dirs {
		if dir.IsDir() {
			f := *path + "/" + dir.Name()
			cat := NewCat(f)
			cat.Default()
			d = append(d, cat)
		}
	}

	l = &Library{read{
		size: len(d),
	},
		d,
	}
	return
}

// File methods
func (l *Library) FileDefault() (path string) {
	cat := l.index[l.now()]
	cat.Default()
	path = cat.files[cat.now()]
	return
}

func (l *Library) FileRandom() (path string) {
	cat := l.index[l.now()]
	cat.random()
	path = cat.files[cat.now()]
	return
}

func (l *Library) FileNext() (path string) {
	cat := l.index[l.now()]
	cat.next()
	path = cat.files[cat.now()]
	return
}

func (l *Library) FilePrev() (path string) {
	cat := l.index[l.now()]
	cat.previous()
	path = cat.files[cat.now()]
	return
}

func (l *Library) FileSuperRnd() (path string) {
	l.random()
	return l.FileRandom()
}

// Folder methods
func (l *Library) FldRnd() (path string) {
	l.random()
	return l.FileDefault()
}

func (l *Library) FldNext() (path string) {
	l.next()
	return l.FileDefault()
}

func (l *Library) FldPrev() (path string) {
	l.previous()
	return l.FileDefault()
}
