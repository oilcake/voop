package library

type Folder string

type Library struct {
	Read
	Index []*string
}
