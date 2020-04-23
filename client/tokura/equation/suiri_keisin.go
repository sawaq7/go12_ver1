package equation

import (
	    "fmt"
	  		  )

///
///      formura ( wetted perimeter )
///         genre : hydraulics

func Suiri_Keisin( area float64 ,S float64  ) float64 {

//     IN    area
//     IN    S　

//    OUT    one        : wetted perimeter

   var R float64

   R = area/S
   fmt.Println ("Suiri_Keisin R  " ,R)


return R
}
