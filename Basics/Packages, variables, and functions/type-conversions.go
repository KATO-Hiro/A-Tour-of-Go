package main

import (
    "fmt"
    "math"
)

func main() {
    var x, y int = 3, 4
    var f float64 = math.Sqrt(float64(x*x + y*y))
    var z uint = uint(f)
    // cannot use x * x + y * y (type int) as type float64 in argument to math.Sqrt
    // var f float64 = math.Sqrt(x*x + y*y)
    // cannot use f (type float64) as type uint in assignment
    // var z uint = f
    fmt.Println(x, y, z)

    // Put more simply.
    i := 42
    f2 := float64(i)
    u := uint(f2)
    fmt.Println(i, f2, u)
}
