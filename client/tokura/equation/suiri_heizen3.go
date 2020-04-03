///
/// 水琁E��　管の冁E��E��求める（�Eーゼン・ウイリアムスの公式！E///
///


package equation


import (
	    "fmt"
        "math"
              )

func Suiri_Heizen3( ch float64 ,I float64 ,Q float64  ) float64 {

//     IN    ch         : 流E��係数
//     IN    I          : 動水勾酁E//     IN    Q(m/s) : 流E��E//     IN    high�E�E)   : 損失水頭
//    OUT    one�E�E)    : 管の長ぁE
   var D float64

// 流E��を求めめE
   rwork := math.Pow( ch ,-0.38)
   rwork2 := math.Pow( I,-0.205)
   rwork3 := math.Pow( Q,0.38)

   D = 1.6258 * rwork * rwork3 * rwork2

   fmt.Println ("Suiri_He-zen3 rwork　" ,rwork)
   fmt.Println ("Suiri_He-zen3 rwork2　" ,rwork2)
   fmt.Println ("Suiri_He-zen3 rwork3　" ,rwork3)
   fmt.Println ("Suiri_He-zen3 length " ,D)

return D

}
