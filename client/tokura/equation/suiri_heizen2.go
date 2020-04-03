///
/// 水琁E��　管の長さを求める（�Eーゼン・ウイリアムスの公式！E///
///


package equation


import (
	    "fmt"
        "math"
        "github.com/sawaq7/go12_ver1/basic/maths/sum"
	  		              )

func Suiri_Heizen2( ch float64 ,D float64 ,velocity float64  ,high float64 ) float64 {

//     IN    ch         : 流E��係数
//     IN    D(�E�！E    : 管の冁E��E//     IN velocity(m/s) : 流E��E//     IN    high�E�E)   : 損失水頭
//    OUT    one�E�E)    : 管の長ぁE
   var length ,Q float64

// 流E��を求めめE
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
