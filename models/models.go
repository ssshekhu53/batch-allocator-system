package models

type Batch struct {
	ID              int
	Stream          string
	Capacity        int
	InitialCapacity int
	Timing          string
}

type Student struct {
	ID     int
	Name   string
	Gender string
	Stream string
	Batch  int
}

type Admin struct {
	ID     int
	Name   string
	Gender string
}
