///
/// 速度水頭を求めめE///   hv = v2剰 / 2g

package equation

import (
	    "fmt"
//	    "github.com/sawaq7/go12_ver1/basic/maths"
	    "github.com/sawaq7/go12_ver1/basic/declare"
	  		          )

func Suiri_Vhead( velocity float64 ) float64 {

//     IN    verocity(m/s�E�E : 流E��E//    OUT    one        : 速度水頭　�E�　vの2剰　/2*g

   var hv float64

//　重力加速度をget

//   gravi := maths.Math_Gravi_Get()
   gravi := declare.Math_Const_gravi


   hv = velocity * velocity / (2*gravi)
   fmt.Println ("Suiri_vhead 速度水頭  " ,hv)


return hv
}
