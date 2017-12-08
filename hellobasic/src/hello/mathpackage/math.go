package math

// compile by "go run build -o math.a"
// package methods start with uppercase character
// generate documentation by "godoc hello/mathpackage/math"

// Get average from a list of numbers
func Average(numbers []float64) float64 {
    total := float64(0)
    for _, number := range numbers {
        total += number
    }
    return total / float64(len(numbers))
}