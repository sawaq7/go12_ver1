///
/// 水理学　（マンニングの公式：流速を求める）
///     v = 1/n * Rの2/3剰 * Iの1/2剰

package equation


import (
	    "fmt"
        "math"
	  		  )

func Suiri_Manningu1( n float64 ,R float64 ,I float64 ) float64 {

//     IN    n          : マンニング粗度係数
//     IN    R(ｍ）     : 径深
//     IN    I          : 動水勾配
//    OUT    one（m/s)  : 速度

   var cons ,cons2 ,V float64

       cons = 2.0/3.0
       cons2 = 1.0/2.0

       fmt.Println ("Suiri_Manningu1 cons　" ,cons)
       fmt.Println ("Suiri_Manningu1 cons2　" ,cons2)

       rwork := math.Pow( R ,cons)
       rwork2 := math.Pow( I ,cons2)

       V = 1/n * rwork * rwork2

       fmt.Println ("Suiri_Manningu1 径深部　" ,rwork)
       fmt.Println ("Suiri_Manningu1 動水勾配部　" ,rwork2)
       fmt.Println ("Suiri_Manningu1 V (マーニング）" ,V)



return V
}