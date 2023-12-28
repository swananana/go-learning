package main

import "fmt"

func FindType(x interface{}) string {
    switch x.(type) {
    case int:
        return "int"
    case float64:
        return "float64"
    default:
        return "unknown"
    }
}

func main() {

    fmt.Println("go" + "lang")

    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}