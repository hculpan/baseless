package dbf

type Dbf struct {
}

func NewDbf(filename string) (*Dbf, error) {
	return &Dbf{}, nil
}
