///
/// 水琁E��　�E�ガンギレー・クチE��ーの公式！E///

package equation


import (
	    "fmt"
        "math"
	  		  )

func Suiri_Gangirei1( n float64 ,R float64 ,I float64 ) float64 {

//     IN    n          : 粗度係数
//     IN    R(�E�！E    : 征E��
//     IN    I          : 動水勾酁E//    OUT    one�E�E/s)  : 速度

   var C1 ,C2 ,C3,V float64

   rwork := math.Pow( R ,0.5)

   fmt.Println ("Suiri_Gangirei1 rwork　" ,rwork)

   C1 = 23 + 1/n +0.00155 / I

   C2 = 1 + ( 23 + 0.00155 / I ) * n / rwork

   C3 = C1 / C2

   fmt.Println ("Suiri_Gangirei1 C3　" ,C3)

   V = C3 * math.Pow( R*I ,0.5)

   fmt.Println ("Suiri_Gangirei1 V " ,V)

return V
}
