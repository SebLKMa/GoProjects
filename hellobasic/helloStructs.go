package main

import ("fmt"; "math")

func distance(x1, y1, x2, y2 float64) float64 {
    a := x2 - x1
    b := y2 - y1
    return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
    l := distance(x1, y1, x1, y2)
    w := distance(x1, y1, x2, y1)
    return l * w
}

func circleArea(x, y, radius float64) float64 {
    return math.Pi * radius*radius
}

func f1() {
    var rx1, ry1 float64 = 0, 0
    var rx2, ry2 float64 = 10, 10
    var cx, cy, cr float64 = 0, 0, 5
    
    fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
    fmt.Println(circleArea(cx, cy, cr))
}

type Circle struct {
    x, y, radius float64
}

func circleArea2(c Circle) float64 {
    return math.Pi * c.radius*c.radius
}

// method of a Circle object
func (pCircle *Circle) area() float64 {
    return math.Pi * pCircle.radius*pCircle.radius
}

type Rectangle struct {
    x1, y1, x2, y2 float64
}
func (pRect *Rectangle) area() float64 {
    l := distance(pRect.x1, pRect.y1, pRect.x1, pRect.y2)
    w := distance(pRect.x1, pRect.y1, pRect.x2, pRect.y1)
    return l * w
}

func f2() {
    c := Circle{x:0, y:0, radius:5}
    fmt.Println(circleArea2(c))
    c.radius = 6
    fmt.Println(circleArea2(c))
    fmt.Println(c.area())
}

func f3() {
    rect := Rectangle{0, 0, 10, 10}
    fmt.Println(rect.area())
}

func main() {
    f1()
    f2()
    f3()
}

