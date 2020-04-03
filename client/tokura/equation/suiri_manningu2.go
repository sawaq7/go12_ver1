///
/// æ°´çE­¦ãEãEã³ãã³ã°ã®å¬å¼ï¼æ©æ¦ä¿æ°ãæ±ããï¼E///ããããÎ» = 124.5 * nã®2ä¹E/ Dã®1/3å°

package equation


import (
	    "fmt"
        "math"
	  		  )

func Suiri_Manningu2( n float64 ,D float64 ) float64 {

//     IN    n          : ãã³ãã³ã°ç²åº¦ä¿æ°
//     IN    E¤(Eï¼E   : ç®¡ã®åE¾E//    OUT    out1EE/s) : æ©æ¦ä¿æ°

   var ramuda float64

   fmt.Println ("Suiri_Manningu2 nã" ,n)
   fmt.Println ("Suiri_Manningu2 Dã" ,D)

   rwork := math.Pow( n ,2.0)
   rwork2 := math.Pow( D ,1.0/3.0)

   ramuda = 124.5 * rwork / rwork2

   fmt.Println ("Suiri_Manningu2 rworkã" ,rwork)
   fmt.Println ("Suiri_Manningu2 rwork2ã" ,rwork2)
   fmt.Println ("Suiri_Manningu2 ramuda (ããEãã³ã°EE ,ramuda)



return ramuda
}
