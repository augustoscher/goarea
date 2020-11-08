package area

import "math"

// Pi is the numbered proportion defined by the relationship between the perimeter and the diameter of a circle
const Pi = 3.1416

// CircunferenceArea calculates the area of circunference
func CircunferenceArea(radius float64) float64 {
	return math.Pow(radius, 2) * Pi
}

// RectangleArea calculates the area of rectangle
func RectangleArea(width, height float64) float64 {
	return width * height
}

// Not exported
func _EquilateralTriangleArea(width, height float64) float64 {
	return (width * height) / 2
}
