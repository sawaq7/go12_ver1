///
/// get several constant number
///

package maths

// the structure for constant number

var Math_Const struct{
     pai float64        // ①　円周率
     gravi float64      // ②　重力加速度
}




func Math_Pai_Get()float64 {

//    OUT    one        : 円周率

     Math_Const.pai = 3.1416 // 円周率をセット

     return(Math_Const.pai)
  }

func Math_Gravi_Get()float64 {

//    OUT    one        : 重力加速度

     Math_Const.gravi = 9.8066 // 重力加速度をセット

     return(Math_Const.gravi)
  }



