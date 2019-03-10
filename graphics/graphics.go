package graphics

type Graphics interface {
	GetX() int
	GetY() int

	Clear()
	MoveTo(x int, y int)
	FillRect(x int, y int, width int, height int)
}
