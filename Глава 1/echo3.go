// echo3 выводит аргументы командной строки
package main

import (
    "fmt"
    "strings"
    "os"
)

func main() {
    fmt.Println(strings.Join(os.Args[1:], " "))
}
