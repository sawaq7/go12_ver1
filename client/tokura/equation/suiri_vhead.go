///
/// 速度水頭を求める
///   hv = v2剰 / 2g

package equation

import (
	    "fmt"
//	    "basic/maths"
	    "basic/declare"
	  		          )

func Suiri_Vhead( velocity float64 ) float64 {

//     IN    verocity(m/s）  : 流速
//    OUT    one        : 速度水頭　＝　vの2剰　/2*g

   var hv float64

//　重力加速度をget

//   gravi := maths.Math_Gravi_Get()
   gravi := declare.Math_Const_gravi


   hv = velocity * velocity / (2*gravi)
   fmt.Println ("Suiri_vhead 速度水頭  " ,hv)


return hv
}