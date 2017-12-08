package main

import ("fmt"; "container/list"; "sort")

func f1List() {
    var myList list.List
    myList.PushBack(1)
    myList.PushBack(2)
    myList.PushBack(3)
    
    for e := myList.Front(); e != nil; e=e.Next() {
        fmt.Println(e.Value.(int))
    }
}

type Person struct {
    Name string
    Age int
}

type ByName []Person
func (ps ByName) Len() int {
    return len(ps)
}
func (ps ByName) Less(i, j int) bool {
    return ps[i].Name < ps[j].Name
}
func (ps ByName) Swap(i, j int) {
    ps[i], ps[j] = ps[j], ps[i]
}

type ByAge []Person
func (ps ByAge) Len() int {
    return len(ps)
}
func (ps ByAge) Less(i, j int) bool {
    return ps[i].Age < ps[j].Age
}
func (ps ByAge) Swap(i, j int) {
    ps[i], ps[j] = ps[j], ps[i]
}

func f2Sort() {
    kids := []Person{
        {"Jill", 9},
        {"Jack", 10},
    }
    sort.Sort(ByName(kids))
    fmt.Println(kids)
    sort.Sort(ByAge(kids))
    fmt.Println(kids)
}

func main() {
    f1List()
    f2Sort()
}