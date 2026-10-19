package main

import(
    "fmt"
    "runtime"
)

func main() {
    fmt.Println("-----------------------------------")
    fmt.Println("|             SysInfo             |")
    fmt.Println("-----------------------------------")
    // information display
    fmt.Println("| OS:", runtime.GOOS)
    fmt.Println("| Arch:", runtime.GOARCH)
    fmt.Println("| CPUs:", runtime.NumCPU())
    fmt.Println("| GO Version:", runtime.Version())
}