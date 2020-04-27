package equation


import (
//	    "fmt"
        "math"
	  		  )

///
///    Manning formula
///

func Suiri_Manningu2( n float64 ,D float64 ) float64 {

//     IN    n
//     IN    D

//    OUT           : ramuda

   var ramuda float64

//   fmt.Println ("Suiri_Manningu2 n　" ,n)
//   fmt.Println ("Suiri_Manningu2 D　" ,D)

   rwork := math.Pow( n ,2.0)
   rwork2 := math.Pow( D ,1.0/3.0)

   ramuda = 124.5 * rwork / rwork2

//   fmt.Println ("Suiri_Manningu2 rwork　" ,rwork)
//   fmt.Println ("Suiri_Manningu2 rwork2　" ,rwork2)
//   fmt.Println ("Suiri_Manningu2 ramuda " ,ramuda)

return ramuda
}
