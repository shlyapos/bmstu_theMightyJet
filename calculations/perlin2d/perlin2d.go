package perlin2d

import (
	"math"
	"math/rand"
	"time"
)

type Vector2D struct {
	X, Y float64
}

func NewVector2D(x, y int) Vector2D {
	return Vector2D{(float64)(x) + rand.Float64(), (float64)(y) + rand.Float64()}
}

const (
	tableSize = 1024
)

var (
	RIGHT  = Vector2D{X:  1}
	LEFT   = Vector2D{X: -1}
	TOP    = Vector2D{Y:  1}
	BOTTOM = Vector2D{Y :-1}
	gradientVectors = []Vector2D{RIGHT, LEFT, TOP, BOTTOM}
)

type Perlin2D struct {
	permutationTable []byte
}

func NewPerlin2D() (*Perlin2D, error) {
	var perlin Perlin2D

	rand.Seed(time.Now().UnixNano())
	perlin.permutationTable = make([]byte, tableSize)

	for i := range perlin.permutationTable {
		perlin.permutationTable[i] = byte(rand.Int())
	}

	return &perlin, nil
}

func pow(base, power int) int {
	return (int)(math.Floor(math.Pow((float64)(base), (float64)(power))))
}

func (perlin *Perlin2D) GetPseudoRandomGradientVector(x, y int) Vector2D {
	pseudoRandomNumber := (pow(x, 3) * pow(y, 5) + 17435) & 1023
	pseudoRandomNumber  = (int)(perlin.permutationTable[pseudoRandomNumber] & 3)

	return gradientVectors[pseudoRandomNumber]
}

func (perlin *Perlin2D) QuinticCurve(coord float64) float64 {
	return math.Pow(coord, 3.0) * (6 * math.Pow(coord, 2.0) - 15 * coord + 10)
}

func (perlin *Perlin2D) Lerp(min, max, param float64) float64 {
	return (min + max) * param - min
}

func (perlin *Perlin2D) DotProduct(a, b Vector2D) float64 {
	return a.X * b.X + a.Y * b.Y
}

func (perlin *Perlin2D) Noise(vector Vector2D) float64 {
	topLeftX := (int)(math.Floor(vector.X))
	topLeftY := (int)(math.Floor(vector.Y))

	pointInSquareX := vector.X - math.Floor(vector.X)
	pointInSquareY := vector.Y - math.Floor(vector.Y)

	topLeftGradient     := perlin.GetPseudoRandomGradientVector(topLeftX,     topLeftY)
	topRightGradient    := perlin.GetPseudoRandomGradientVector(topLeftX + 1, topLeftY)
	bottomLeftGradient  := perlin.GetPseudoRandomGradientVector(topLeftX,     topLeftY + 1)
	bottomRightGradient := perlin.GetPseudoRandomGradientVector(topLeftX + 1, topLeftY + 1)

	distanceToTopLeft     := Vector2D{pointInSquareX,     pointInSquareY}
	distanceToTopRight    := Vector2D{pointInSquareX - 1, pointInSquareY}
	distanceToBottomLeft  := Vector2D{pointInSquareX,     pointInSquareY - 1}
	distanceToBottomRight := Vector2D{pointInSquareX - 1, pointInSquareY - 1}

	pointInSquareX = perlin.QuinticCurve(pointInSquareX)
	pointInSquareY = perlin.QuinticCurve(pointInSquareY)

	topLeftDotProduct     := perlin.DotProduct(distanceToTopLeft,     topLeftGradient)
	topRightDotProduct    := perlin.DotProduct(distanceToTopRight,    topRightGradient)
	bottomLeftDotProduct  := perlin.DotProduct(distanceToBottomLeft,  bottomLeftGradient)
	bottomRightDotProduct := perlin.DotProduct(distanceToBottomRight, bottomRightGradient)

	topInterpolation    := perlin.Lerp(topLeftDotProduct,    topRightDotProduct,    pointInSquareX)
	bottomInterpolation := perlin.Lerp(bottomLeftDotProduct, bottomRightDotProduct, pointInSquareX)

	finalInterpolation  := perlin.Lerp(topInterpolation, bottomInterpolation, pointInSquareY)

	return finalInterpolation
}

func (perlin *Perlin2D) MultiOctaveNoise(vector Vector2D, octaves int, persistence float64) float64 {
	var amplitude float64 = 1
	var max float64 = 0
	var accumulator float64 = 0

	for ; octaves > 0; octaves -= 1 {
		max += amplitude
		accumulator += perlin.Noise(vector) * amplitude
		amplitude   *= persistence
		vector.X *= 2
		vector.Y *= 2
	}

	return accumulator / max
}
