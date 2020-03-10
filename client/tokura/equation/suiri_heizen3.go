///
/// 水理学　管の内径を求める（ヘーゼン・ウイリアムスの公式）
///
///


package equation


import (
	    "fmt"
        "math"
              )

func Suiri_Heizen3( ch float64 ,I float64 ,Q float64  ) float64 {

//     IN    ch         : 流速係数
//     IN    I          : 動水勾配
//     IN    Q(m/s) : 流速
//     IN    high（m)   : 損失水頭
//    OUT    one（m)    : 管の長さ

   var D float64

// 流量を求める

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