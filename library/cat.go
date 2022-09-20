package library

type cat struct {
	parent string
	read
	files []string
}

func (f *cat) What(i int) string {
	return f.files[i]
}

func NewCat(path string) *cat {

	files := SupportedFilesFrom(path)

	return &cat{
		path,
		read{
			rightNow: 0,
			size:     len(files),
		},
		files,
	}
}
