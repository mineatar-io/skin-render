package skin

// Credit: https://github.com/fogleman/gg

type matrix3 struct {
	XX, YX, XY, YY, X0, Y0 float64
}

func translateMatrix(x, y float64) matrix3 {
	return matrix3{
		1, 0,
		0, 1,
		x, y,
	}
}

func (a matrix3) Multiply(b matrix3) matrix3 {
	return matrix3{
		a.XX*b.XX + a.YX*b.XY,
		a.XX*b.YX + a.YX*b.YY,
		a.XY*b.XX + a.YY*b.XY,
		a.XY*b.YX + a.YY*b.YY,
		a.X0*b.XX + a.Y0*b.XY + b.X0,
		a.X0*b.YX + a.Y0*b.YY + b.Y0,
	}
}

func (a matrix3) Translate(x, y float64) matrix3 {
	return translateMatrix(x, y).Multiply(a)
}
