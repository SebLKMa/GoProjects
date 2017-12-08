package math

import "testing"

// run "go test" from this directory
// test functions are prefixed Test...

func TestSimpleAverage(t *testing.T) {
    result := Average([]float64{1,2})
    if result != 1.5 {
        t.Error("Expected 1.5. Actual ", result)
    }
}

type testpair struct {
    values []float64
    average float64
}

var tests = []testpair{
    {[]float64{1,2}, 1.5},
    {[]float64{1,1,1,1,1,1}, 1},
    {[]float64{-1,1}, 0},
}

func TestAverages(t *testing.T) {
    for _, pair := range tests {
        result := Average(pair.values)
        if result != pair.average {
            t.Error(
                "For", pair.values,
                "expected", pair.average,
                "actual", result,
            )
        }
    }
}