///
/// 征E��を求めめE///

package equation

import (
	    "fmt"
	  		  )

func Suiri_Keisin( area float64 ,S float64  ) float64 {

//     IN    area(㎡�E�E : 流積（水路の断面積！E//     IN    �E�　       : 潤辺�E�水路が壁と接してぁE��長さ！E//    OUT    one        : 征E��　�E�　流積　/ 潤辺

   var R float64

   R = area/S
   fmt.Println ("Suiri_Keisin R  " ,R)


return R
}
