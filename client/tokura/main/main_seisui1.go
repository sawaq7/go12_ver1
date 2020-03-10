///
///  静水圧　U字管の計算 type1（key-in入力）
///

package main

import (
	"fmt"
	"os"
	"client/tokura/suiri"
//	"basic/maths"
	"basic/maths/sum"
	   )

func main() {

// 単位容積重量　（ω）をセット
//    const omega float64 = 1.0

    var a1 ,a2 ,a3, a4,p1 ,p2 ,h ,omega float64
    var cont string

   for{
      fmt.Printf ("単位容積重量 ω は（t/㎡）")
      fmt.Scanln(&omega)

      fmt.Printf ("U字管１の直径は（ｍ）")
      fmt.Scanln(&a1)

      fmt.Printf("U字管２の直径は（ｍ）")
      fmt.Scanln(&a2)

      fmt.Printf ("荷重の重さは（t）")
      fmt.Scanln (&p2 )

      fmt.Printf ("高度差は（ｍ）")
      fmt.Scanln (&h)

/// U字管の面積を計算する

      a3 = sum.Circle_Area(a1/2)
      a4 = sum.Circle_Area(a2/2)

      p1 =  suiri.Seisui1( a3 ,a4  ,p2  ,omega  ,h  )

/// 静水圧を計算する

     fmt.Println("静水圧は",p1,"ｔ")

     fmt.Printf ("next or end ")
     fmt.Scanln (&cont)

     if cont == "next" {
///   そのままスルー

      } else if cont == "end" {

         break

      } else {

      fmt.Println ("seisui1.go key-in error")
      os.Exit(1)
      }
   }
}