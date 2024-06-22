package entity

type Category struct {
	Id     int
	UserId int
	Title  string
	Color  string
}

type CategoryAdaptor interface{}
