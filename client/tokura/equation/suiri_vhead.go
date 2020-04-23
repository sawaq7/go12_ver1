package equation

import (
	    "fmt"
//	    "github.com/sawaq7/go12_ver1/basic/maths"
	    "github.com/sawaq7/go12_ver1/basic/declare"
	  		          )

///
///   calculate verocity head
///

func Suiri_Vhead( velocity float64 ) float64 {

//     IN    verocity
//    OUT    one        : verocity head　

   var hv float64

  ///　 get gravitational-acceleration
   gravi := declare.Math_Const_gravi

   hv = velocity * velocity / (2*gravi)
   fmt.Println ("Suiri_vhead hv  " ,hv)

return hv

}
