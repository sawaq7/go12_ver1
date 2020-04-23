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

//     IN    radius(m)

//    OUT    one        : circumference of circle

    pai := declare.Math_Const_Pai  ///   get pai

    fmt.Println("Circle_Circum",pai)

    return(2*pai*radius)
}
