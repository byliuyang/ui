package graphics

type Graphics interface {
	GetX() int
	GetY() int
	MoveTo(x int, y int)
}