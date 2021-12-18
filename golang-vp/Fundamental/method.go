package Fundamental

type Rectangle struct {
	Width, Lenght int
}

func (r Rectangle) GetWidht() int {
	// fmt.Println(r.Width)
	//instance rectangle
	return r.Width
}

func (r *Rectangle) SetWidht(Widht int) {
	r.Width = Widht
}
