package equation


import (
	    "fmt"
        "math"
	  		  )

///
///      formura of Hezen Williams
///          genre : hydraulics

func Suiri_Heizen1( ch float64 ,D float64 ,I float64 ) float64 {

//     IN    ch
//     IN    D

//     IN    I
//    OUT               :velocity

   var V float64

   rwork := math.Pow( D ,0.63)
   rwork2 := math.Pow( I ,0.54)

   V = 0.3564 * ch * rwork * rwork2

   fmt.Println ("Suiri_He-zen1 rwork　" ,rwork)
   fmt.Println ("Suiri_He-zen1 動水勾配部　" ,rwork2)
   fmt.Println ("Suiri_He-zen1 V " ,V)

return V
}
