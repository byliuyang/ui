package view

func NewRoot() View {
	return &Container{
		Base: Base{
			Width:  100,
			Height: 200,
		},
		BackgroundColor: Color{
			Red:   255,
			Blue:  255,
			Green: 255,
			Alpha: 1,
		},
	}
}
