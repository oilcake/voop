package library

type cat struct {
	parent string
	read
	files []string
}

func (f *cat) What(i int) string {
	return f.files[i]
}

func NewCat(path string, supported []string) *cat {

	files := SupportedFilesFrom(path, supported)

	return &cat{
		path,
		read{
			rightNow: 0,
			size:     len(files),
		},
		files,
	}
}
