package storage

type Slice struct {
}

type Map struct {
}

func NewSliceStorage() *Slice {
	return &Slice{}
}

func NewMapStorage() *Map {
	return &Map{}
}
