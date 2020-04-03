///
/// 水琁E��　流E��を求める（�Eーゼン・ウイリアムスの公式！E///

package equation


import (
	    "fmt"
        "math"
	  		  )

func Suiri_Heizen1( ch float64 ,D float64 ,I float64 ) float64 {

//     IN    ch         : 流E��係数
//     IN    D(�E�！E    : 管の冁E��E//     IN    I          : 動水勾酁E//    OUT    one�E�E/s)  : 速度

   var V float64

   rwork := math.Pow( D ,0.63)
   rwork2 := math.Pow( I ,0.54)

   V = 0.3564 * ch * rwork * rwork2

   fmt.Println ("Suiri_He-zen1 征E��部　" ,rwork)
   fmt.Println ("Suiri_He-zen1 動水勾配部　" ,rwork2)
   fmt.Println ("Suiri_He-zen1 V " ,V)

return V
}
