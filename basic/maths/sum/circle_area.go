///
/// calculate area of circle
///

package sum

import (
        "fmt"
//        "basic/maths"
        "basic/declare"
                       )

func Circle_Area(radius float64 )float64 {

//     IN    radius(m)  : 半径
//    OUT    one        : 円の面積

//    pai:= maths. Math_Pai_Get() //円周率をGET
   pai := declare.Math_Const_Pai //円周率をGET
   fmt.Println("Circle_Area πは",pai)

   return(pai*radius*radius)
}