package equation


import (
	    "fmt"
        "math"
        "github.com/sawaq7/go12_ver1/basic/maths/sum"
	  		                                            )

///
///      formura of Hezen Williams
///

func Suiri_Heizen2( ch float64 ,D float64 ,velocity float64  ,high float64 ) float64 {

//     IN    ch         :
//     IN    D
//     IN velocity
//     IN    high

//    OUT    one    : tube length

   var length ,Q float64

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
