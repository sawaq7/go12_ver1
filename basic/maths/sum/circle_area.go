///
/// calculate area of circle
///

package sum

import (
        "fmt"
//        "github.com/sawaq7/go12_ver1/basic/maths"
        "github.com/sawaq7/go12_ver1/basic/declare"
                       )

func Circle_Area(radius float64 )float64 {

//     IN    radius(m)  : 半征E//    OUT    one        : 冁E�E面穁E
//    pai:= maths. Math_Pai_Get() //冁E��玁E��GET
   pai := declare.Math_Const_Pai //冁E��玁E��GET
   fmt.Println("Circle_Area πは",pai)

   return(pai*radius*radius)
}
