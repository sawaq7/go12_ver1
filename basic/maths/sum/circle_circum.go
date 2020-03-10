///
/// calculate  circumference of circle
///

package sum

import (
         "fmt"
//         "basic/maths"
         "basic/declare"
                        )

func Circle_Circum(radius float64 )float64 {

//     IN    radius(m)  : 半径
//    OUT    one        : 円周

//    pai:= maths. Math_Pai_Get() //πをget
    pai := declare.Math_Const_Pai //円周率をGET
    fmt.Println("Circle_Circum",pai)

    return(2*pai*radius)
}