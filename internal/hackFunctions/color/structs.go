package color

type RGB struct {
	Red   int
	Green int
	Blue  int
}

type RGBA struct {
	Red   float32
	Green float32
	Blue  float32
	Alpha float32
}

type Hex string

func NewRGBA(r float32, g float32, b float32, a float32) *RGBA {
	return &RGBA{
		Red:   r,
		Green: g,
		Blue:  b,
		Alpha: a,
	}
}
