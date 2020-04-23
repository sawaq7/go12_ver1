package suiri

import "fmt"

func Seisui1( a1 float64 ,a2 float64 ,p1 float64 ,omega float64 ,h float64 ) float64 {

///
///    the calculation of hydrostatic pressure　　　
///

//     IN    a1(�E�！E   : cross-sectional-area of u-tube1
//     IN    a2(�E�！E   : cross-sectional-area of u-tube2
//     IN    p(t)       : pressure21
//     IN    omega
//     IN    h(m)       : high

//    OUT    one        : pressure2
   var p2 float64

    fmt.Println ("func Seisui1 U字管断面積　",a1,"㎡" )
    fmt.Println ("func Seisui1 U字管�断面積　",a2,"㎡" )
    fmt.Println ("func Seisui1 荷重　",p1,"t" )
    fmt.Println ("func Seisui1 単位容積重量　",omega,"t")
    fmt.Println ("func Seisui1 高度差　",h,"t")

// from "p1/a1 = ( p2/a2 + omega * h )"

    p2 = (p1/a2 + omega* h ) * a1

    return p2

   }
