package main

import (
    "fmt"
    "github.com/nozo-moto/go1.14-try/b"
    "github.com/nozo-moto/go1.14-try/a"
)


type AB interface {
    a.A
    b.B
}

type ABimpl struct {
    A a.A
    B b.B
}

func NewAB(a a.A, b b.B) *ABimpl {
    return &ABimpl{
        A: a,
        B: b,
    }
}
func main() {
    ab := NewAB(
        &a.Aimpl{},
        &b.Bimpl{},
    )
    

    fmt.Println(
        ab.A.Get(),
        ab.B.Get(),
    )
}

