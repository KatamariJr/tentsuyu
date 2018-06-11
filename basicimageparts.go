package tentsuyu

//BasicImageParts is easy to set up basic sprite image
type BasicImageParts struct {
	Width, Height, Sx, Sy, DestWidth, DestHeight int
	Reverse                                      bool
}

//NewBasicImageParts returns a pointer to new BasicImageParts
func NewBasicImageParts(sx, sy, width, height int) *BasicImageParts {
	b := &BasicImageParts{
		Sx:         sx,
		Sy:         sy,
		Width:      width,
		Height:     height,
		DestHeight: height,
		DestWidth:  width,
	}
	return b
}

//SetDestinationDimensions can be used to set the size the image should be drawn to the screen
func (b *BasicImageParts) SetDestinationDimensions(width, height int) {
	b.DestWidth = width
	b.DestHeight = height
}

//ReverseX flips the image
func (b *BasicImageParts) ReverseX(reverse bool) {
	b.Reverse = reverse
}

//Len returns 1
func (b *BasicImageParts) Len() int {
	return 1
}

//Dst we just make it 1:1
func (b *BasicImageParts) Dst(i int) (x0, y0, x1, y1 int) {
	if b.DestHeight == 0 && b.DestWidth == 0 {
		return 0, 0, b.Width, b.Height
	}
	return 0, 0, b.DestWidth, b.DestHeight
}

//Src cuts out the specified rectangle from the source image to display the sprite
func (b *BasicImageParts) Src(i int) (x0, y0, x1, y1 int) {
	x := b.Sx
	y := b.Sy
	if b.Reverse {
		return x + b.Width, y, x, y + b.Height
	}
	return x, y, x + b.Width, y + b.Height
}