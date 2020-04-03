///
/// calculate  circumference of circle
///

package sum

import (
         "fmt"
//         "github.com/sawaq7/go12_ver1/basic/maths"
         "github.com/sawaq7/go12_ver1/basic/declare"
                        )

func Circle_Circum(radius float64 )float64 {

//     IN    radius(m)  : 蜊雁ｾ・//    OUT    one        : 蜀・捉

//    pai:= maths. Math_Pai_Get() //ﾏ繧暖et
    pai := declare.Math_Const_Pai //蜀・捉邇・ｒGET
    fmt.Println("Circle_Circum",pai)

    return(2*pai*radius)
}
