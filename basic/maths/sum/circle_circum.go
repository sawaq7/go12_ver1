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

//     IN    radius(m)  : 半征E//    OUT    one        : 冁E��

//    pai:= maths. Math_Pai_Get() //πをget
    pai := declare.Math_Const_Pai //冁E��玁E��GET
    fmt.Println("Circle_Circum",pai)

    return(2*pai*radius)
}
