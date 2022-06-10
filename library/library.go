package library

import (
	"io/ioutil"
	"voop/sync"
)

// type Folder string

type Library struct {
	Read
	Index []*string
}

func (l *Library) What(i int) interface{} {
	return l.Index[i]
}

func NewLibrary(path *string, t *sync.Transport) (l *Library, err error) {
	d := make([]*string, 0)
	dirs, err := ioutil.ReadDir(*path)
	if err != nil {
		return nil, err
	}

	for _, dir := range dirs {
		if dir.IsDir() {
			f := *path + "/" + dir.Name()
			d = append(d, &f)
		}
	}

	l = &Library{Read{
		RightNow: 0,
		Size:     len(d),
	},
		d,
	}

	return

}
