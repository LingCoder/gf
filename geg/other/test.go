package main

import (
    "fmt"
    "gitee.com/johng/gf/g/util/gutil"
)

func main() {
    gutil.TryCatch(func() {
        fmt.Println(1)
        panic("error")
        fmt.Println(2)
    }, func(err interface{}) {
        fmt.Println(err)
    })
}