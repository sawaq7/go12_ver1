package sum

import (
        "fmt"

        "github.com/sawaq7/go12_ver1/basic/declare"
                       )

///
/// calculate area of circle
///

func Circle_Area(radius float64 )float64 {

//     IN    radius(m)  :

//    OUT    one        : area of circle

///
/// calculate area of circle
///

   pai := declare.Math_Const_Pai    ///   get pai

   fmt.Println("Circle_Area πは",pai)

   return(pai*radius*radius)
}
