///
/// 水理学　流速を求める（ヘーゼン・ウイリアムスの公式）
///

package equation


import (
	    "fmt"
        "math"
	  		  )

func Suiri_Heizen1( ch float64 ,D float64 ,I float64 ) float64 {

//     IN    ch         : 流速係数
//     IN    D(ｍ）     : 管の内径
//     IN    I          : 動水勾配
//    OUT    one（m/s)  : 速度

   var V float64

   rwork := math.Pow( D ,0.63)
   rwork2 := math.Pow( I ,0.54)

   V = 0.3564 * ch * rwork * rwork2

   fmt.Println ("Suiri_He-zen1 径深部　" ,rwork)
   fmt.Println ("Suiri_He-zen1 動水勾配部　" ,rwork2)
   fmt.Println ("Suiri_He-zen1 V " ,V)

return V
}