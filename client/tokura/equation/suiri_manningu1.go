package equation


import (
	    "fmt"
        "math"
	  		  )
///
///    Manning formula
///

func Suiri_Manningu1( n float64 ,R float64 ,I float64 ) float64 {

//     IN    n
//     IN    R
//     IN    I

//    OUT    one  :   velocity

   var cons ,cons2 ,V float64

       cons = 2.0/3.0
       cons2 = 1.0/2.0

       fmt.Println ("Suiri_Manningu1 cons　" ,cons)
       fmt.Println ("Suiri_Manningu1 cons2　" ,cons2)

       rwork := math.Pow( R ,cons)
       rwork2 := math.Pow( I ,cons2)

       V = 1/n * rwork * rwork2

       fmt.Println ("Suiri_Manningu1 征E��部　" ,rwork)
       fmt.Println ("Suiri_Manningu1 動水勾配部　" ,rwork2)

return V
}
