///
/// 径深を求める
///

package equation

import (
	    "fmt"
	  		  )

func Suiri_Keisin( area float64 ,S float64  ) float64 {

//     IN    area(㎡）  : 流積（水路の断面積）
//     IN    Ｓ　       : 潤辺（水路が壁と接している長さ）
//    OUT    one        : 径深　＝　流積　/ 潤辺

   var R float64

   R = area/S
   fmt.Println ("Suiri_Keisin R  " ,R)


return R
}