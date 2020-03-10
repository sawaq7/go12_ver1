///
/// 水理学　管の長さを求める（ヘーゼン・ウイリアムスの公式）
///
///


package equation


import (
	    "fmt"
        "math"
        "basic/maths/sum"
	  		              )

func Suiri_Heizen2( ch float64 ,D float64 ,velocity float64  ,high float64 ) float64 {

//     IN    ch         : 流速係数
//     IN    D(ｍ）     : 管の内径
//     IN velocity(m/s) : 流速
//     IN    high（m)   : 損失水頭
//    OUT    one（m)    : 管の長さ

   var length ,Q float64

// 流量を求める

   Q = sum.Circle_Area(D/2.0 ) * velocity

   rwork := math.Pow( ch ,-1.85)
   rwork2 := math.Pow( D,-4.87)
   rwork3 := math.Pow( Q,1.85)

   length = high / (10.666 * rwork * rwork2 * rwork3 )

   fmt.Println ("Suiri_He-zen2 rwork　" ,rwork)
   fmt.Println ("Suiri_He-zen2 rwork2　" ,rwork2)
   fmt.Println ("Suiri_He-zen2 rwork3　" ,rwork3)
   fmt.Println ("Suiri_He-zen2 length " ,length)

return length

}