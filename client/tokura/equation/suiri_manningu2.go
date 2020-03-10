///
/// 水理学　（マンニングの公式：摩擦係数を求める）
///　　　　λ = 124.5 * nの2乗 / Dの1/3剰

package equation


import (
	    "fmt"
        "math"
	  		  )

func Suiri_Manningu2( n float64 ,D float64 ) float64 {

//     IN    n          : マンニング粗度係数
//     IN    Ｄ(ｍ）    : 管の内径
//    OUT    out1（m/s) : 摩擦係数

   var ramuda float64

   fmt.Println ("Suiri_Manningu2 n　" ,n)
   fmt.Println ("Suiri_Manningu2 D　" ,D)

   rwork := math.Pow( n ,2.0)
   rwork2 := math.Pow( D ,1.0/3.0)

   ramuda = 124.5 * rwork / rwork2

   fmt.Println ("Suiri_Manningu2 rwork　" ,rwork)
   fmt.Println ("Suiri_Manningu2 rwork2　" ,rwork2)
   fmt.Println ("Suiri_Manningu2 ramuda (マーニング）" ,ramuda)



return ramuda
}