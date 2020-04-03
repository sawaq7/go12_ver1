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

//     IN    radius(m)  : 蜊雁ｾ・//    OUT    one        : 蜀・・髱｢遨・
//    pai:= maths. Math_Pai_Get() //蜀・捉邇・ｒGET
   pai := declare.Math_Const_Pai //蜀・捉邇・ｒGET
   fmt.Println("Circle_Area ﾏ縺ｯ",pai)

   return(pai*radius*radius)
}
