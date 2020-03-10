///    the calculation of hydraulics　　　静水圧の計算
///    hydrostatic pressure
///    U-tube　　　　　　　　　　　　　　Ｕ字管の計算

package suiri

import "fmt"

func Seisui1( a1 float64 ,a2 float64 ,p1 float64 ,omega float64 ,h float64 ) float64 {

//     IN    a1(ｍ）    : 円形のピストンの断面積１
//     IN    a2(ｍ）    : 円形のピストンの断面積２
//     IN    p(t)       : 圧力１
//     IN omega（t/㎥） : 単位容積重量
//     IN    h(m)       : 高度差
//    OUT    one        : 圧力２

   var p2 float64

    fmt.Println ("func Seisui1 U字管１断面積　",a1,"㎡" )
    fmt.Println ("func Seisui1 U字管２断面積　",a2,"㎡" )
    fmt.Println ("func Seisui1 荷重　",p1,"t" )
    fmt.Println ("func Seisui1 単位容積重量　",omega,"t")
    fmt.Println ("func Seisui1 高度差　",h,"ｍ")

// p1/a1 = ( p2/a2 + omega * h ) より

    p2 = (p1/a2 + omega* h ) * a1

    return p2
   }