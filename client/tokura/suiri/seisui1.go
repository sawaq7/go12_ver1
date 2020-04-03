///    the calculation of hydraulics　　　静水圧の計箁E///    hydrostatic pressure
///    U-tube　　　　　　　　　　　　　　�E�字管の計箁E
package suiri

import "fmt"

func Seisui1( a1 float64 ,a2 float64 ,p1 float64 ,omega float64 ,h float64 ) float64 {

//     IN    a1(�E�！E   : 冁E��のピストンの断面積！E//     IN    a2(�E�！E   : 冁E��のピストンの断面積！E//     IN    p(t)       : 圧力！E//     IN omega�E�E/㎥�E�E: 単位容積重釁E//     IN    h(m)       : 高度差
//    OUT    one        : 圧力！E
   var p2 float64

    fmt.Println ("func Seisui1 U字管�E�断面積　",a1,"㎡" )
    fmt.Println ("func Seisui1 U字管�E�断面積　",a2,"㎡" )
    fmt.Println ("func Seisui1 荷重　",p1,"t" )
    fmt.Println ("func Seisui1 単位容積重量　",omega,"t")
    fmt.Println ("func Seisui1 高度差　",h,"�E�E)

// p1/a1 = ( p2/a2 + omega * h ) より

    p2 = (p1/a2 + omega* h ) * a1

    return p2
   }
