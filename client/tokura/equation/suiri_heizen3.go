package equation


import (
	    "fmt"
        "math"
              )

///
///      formura of Hezen Williams
///         genre : hydraulics

func Suiri_Heizen3( ch float64 ,I float64 ,Q float64  ) float64 {

//     IN    ch
//     IN    I
//     IN    Q
//     IN    high
//    OUT    one      : tube length

   var D float64

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
