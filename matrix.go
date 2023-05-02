package skin

import (
	"image"
	"math"
)

var (
	sideMatrix         matrix2x2 = rotationMatrix(degToRad(30)).Multiply(skewXMatrix(degToRad(30))).Multiply(scaleYMatrix(0.86603))
	frontMatrix        matrix2x2 = rotationMatrix(degToRad(-30)).Multiply(skewXMatrix(degToRad(-30))).Multiply(scaleYMatrix(0.86603))
	plantMatrix        matrix2x2 = rotationMatrix(degToRad(30)).Multiply(skewXMatrix(degToRad(-30))).Multiply(scaleYMatrix(0.86603))
	overlayMatrix      matrix2x2 = scaleMatrix(8.5 / 8.0)
	sideOverlayMatrix  matrix2x2 = sideMatrix.Multiply(overlayMatrix)
	frontOverlayMatrix matrix2x2 = frontMatrix.Multiply(overlayMatrix)
	plantOverlayMatrix matrix2x2 = plantMatrix.Multiply(overlayMatrix)
)

type matrix2x2 [4]float64

func (a matrix2x2) Multiply(b matrix2x2) matrix2x2 {
	return matrix2x2{
		a[0]*b[0] + a[1]*b[2],
		a[0]*b[1] + a[1]*b[3],
		a[2]*b[0] + a[3]*b[2],
		a[2]*b[1] + a[3]*b[3],
	}
}

func (a matrix2x2) Determinant() float64 {
	return a[0]*a[3] - a[1]*a[2]
}

func (a matrix2x2) Inverse() matrix2x2 {
	d := a.Determinant()

	return matrix2x2{
		a[3] * (1.0 / d),
		-a[1] * (1.0 / d),
		-a[2] * (1.0 / d),
		a[0] * (1.0 / d),
	}
}

func scaleYMatrix(a float64) matrix2x2 {
	return matrix2x2{
		1, 0,
		0, a,
	}
}

func scaleMatrix(a float64) matrix2x2 {
	return matrix2x2{
		a, 0,
		0, a,
	}
}

func skewXMatrix(a float64) matrix2x2 {
	return matrix2x2{
		1, math.Tan(a),
		0, 1,
	}
}

func rotationMatrix(a float64) matrix2x2 {
	return matrix2x2{
		math.Cos(a), -math.Sin(a),
		math.Sin(a), math.Cos(a),
	}
}

func degToRad(a float64) float64 {
	return a * (math.Pi / 180.0)
}

func transformRect(m matrix2x2, r image.Rectangle) (output image.Rectangle) {
	ps := []image.Point{
		{r.Min.X, r.Min.Y},
		{r.Max.X, r.Min.Y},
		{r.Min.X, r.Max.Y},
		{r.Max.X, r.Max.Y},
	}

	for i, p := range ps {
		sxf := float64(p.X)
		syf := float64(p.Y)
		dxi, dyi := translateCoordinatesWithMatrix(sxf, syf, m)
		dx, dy := int(math.Floor(dxi)), int(math.Floor(dyi))

		if i == 0 {
			output = image.Rectangle{
				Min: image.Point{dx + 0, dy + 0},
				Max: image.Point{dx + 1, dy + 1},
			}

			continue
		}

		if output.Min.X > dx {
			output.Min.X = dx
		}

		dx++

		if output.Max.X < dx {
			output.Max.X = dx
		}

		if output.Min.Y > dy {
			output.Min.Y = dy
		}

		dy++

		if output.Max.Y < dy {
			output.Max.Y = dy
		}
	}

	return output
}

func translateCoordinatesWithMatrix(x, y float64, b matrix2x2) (float64, float64) {
	return b[0]*x + b[1]*y, b[2]*x + b[3]*y
}
