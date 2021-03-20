package reader

type File struct {
	Number int    `csv:"file_number"`
	Name   string `csv:"file_name"`
}
