package main

import ("fmt"; "math")

func distance(x1, y1, x2, y2 float64) float64 {
    a := x2 - x1
    b := y2 - y1
    return math.Sqrt(a*a + b*b)
}

type Shape interface {
    area() float64
}

type Circle struct {
    x, y, radius float64
}
// method of a Circle object
func (pCircle Circle) area() float64 {
    return math.Pi * pCircle.radius*pCircle.radius
}

type Rectangle struct {
    x1, y1, x2, y2 float64
}
func (pRect Rectangle) area() float64 {
    l := distance(pRect.x1, pRect.y1, pRect.x1, pRect.y2)
    w := distance(pRect.x1, pRect.y1, pRect.x2, pRect.y1)
    return l * w
}

func totalArea(shapes ...Shape) float64 {
    var areaTotal float64
    for _, s := range shapes {
        areaTotal += s.area()
    }
    return areaTotal
}

func f1() {
    c := Circle{x:0, y:0, radius:5}
    r := Rectangle{0, 0, 10, 10}
    fmt.Println(totalArea(c, r))
}

type MultipleShape struct {
    shapes []Shape
}
// MultipleShape is-a Shape by having an area method
func (ms MultipleShape) area() float64 {
    var areaTotal float64
    for _, s := range ms.shapes {
        areaTotal += s.area() // calls individual shape's area()
    }
    return areaTotal
}

func f2() {
    multiShapes := MultipleShape{
        shapes: []Shape{
            Circle{0, 0, 5},
            Rectangle{0, 0, 10, 10}, 
        },
    }
    
    fmt.Println(multiShapes.area())
}

func main() {
    f1()
    f2()
}