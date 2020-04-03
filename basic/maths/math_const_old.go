///
/// get several constant number
///

package maths

// the structure for constant number

var Math_Const struct{
     pai float64        // 竭縲蜀・捉邇・     gravi float64      // 竭｡縲驥榊鴨蜉騾溷ｺｦ
}




func Math_Pai_Get()float64 {

//    OUT    one        : 蜀・捉邇・
     Math_Const.pai = 3.1416 // 蜀・捉邇・ｒ繧ｻ繝・ヨ

     return(Math_Const.pai)
  }

func Math_Gravi_Get()float64 {

//    OUT    one        : 驥榊鴨蜉騾溷ｺｦ

     Math_Const.gravi = 9.8066 // 驥榊鴨蜉騾溷ｺｦ繧偵そ繝・ヨ

     return(Math_Const.gravi)
  }



