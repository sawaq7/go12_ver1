package cal

import (
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/tokura/equation"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
//	    "github.com/sawaq7/go12_ver1/basic/type3"
//	    "strings"
//	    "strconv"
    	                 )

// func Pipe_line1( water type4.Water2 ,water_line []type4.Water_Line  ) (int ,[]type3.Point,[]type3.Point ,[]type3.Point ,[]type3.Point ) {

func Pipe_line1( water type4.Water2 ,water_line []type4.Water_Line  ) {

//     IN  wdeta : 水路データ
//    OUT  one   : ポイント損失のスライス
//    OUT  two   : ライン損失のスライス
//    OUT  three : 速度水頭のスライス
//    OUT  four  : エネルギー線（up）のスライス
//    OUT  five  : エネルギー線（down）のスライス
//    OUT  six   : 導水勾配線（up）のスライス
//    OUT  seven : 導水勾配線（down）のスライス


   var b_length float64
   var x_eneup ,y_eneup  ,y_enedown ,y_glineup float64

   var hp ,hl ,b_hl,vhead float64

   var ad_eneup_wk ,ad_enedown_wk  ,ad_glineup_wk ,ad_glinedown_wk  [3]float64

///
///     allocate various-work-area
///

   ad_hp := make([]float64 ,20 ,50)        // 1　hp　
   ad_hl := make([]float64 ,20 ,50)        // 2　hl　
   ad_vhead := make([]float64 ,20 ,50)         // 3 　vhead

//   ad_eneup := make([]type3.Point ,20 ,50)     // 4 　eneup
//   ad_enedown := make([]type3.Point ,20 ,50)   // 5　enedown
//   ad_glineup := make([]type3.Point ,20 ,50)   // 6　glineup
//   ad_glinedown := make([]type3.Point ,20 ,50) // 7　glinedown

//   ad_eneup := make([]type3.Point, 0)
//   ad_enedown := make([]type3.Point, 0)
//   ad_glineup := make([]type3.Point, 0)
//   ad_glinedown := make([]type3.Point, 0)


   eflag := 0

   line_num := len(water_line) // set line-no

//   fmt.Println ("cal.pipe_line1 line_num　%v\n",line_num )

   Hmax := water.High   //   set water-high

   s_coeff := water.Roughness_Factor   //　set roughness_factor

///
///  continue process while read records  until end-mark
///

   index := 0
   for pos, water_linew := range water_line {
     count := pos + 1
     if count == line_num {

        eflag = 1

     }

     f_coeff  := water_linew.Friction_Factor  //  set friction-factor

     velocity := water_linew.Velocity         // set velocity

     diameter := water_linew.Pipe_Diameter    //  set diameter

     length   := water_linew.Pipe_Length      // set pipe_length

///    cal. point-loss

     vhead = equation.Suiri_Vhead( velocity )  //    cal. velocity-head
     hp = f_coeff * vhead

///
///    cal. line-loss
///

     ramuda := equation.Suiri_Manningu2( s_coeff ,diameter)  // cal. friction-factor
     vhead := equation.Suiri_Vhead( velocity )               // cal. velocity-head
     hl = ramuda * (length / diameter) * vhead

///
///      make various data for water-slope-line
///

     ad_hp[index] = hp

     if eflag == 1 {     // if i'ts last process , velocity-head and friction-loss is zero

        hl    = 0.0
        vhead = 0.0
     }

     ad_hl[index] = hl

     ad_vhead[index] = vhead

      if index == 0 {

         b_length = 0.0   //  initialize offset of horizontal of x,y

         x_eneup  = 0.0
         y_eneup = Hmax

      }else{

            y_eneup  = y_enedown - b_hl

      }

///　  make energy-line(up)

      x_eneup  = x_eneup + b_length

      b_length = length    //  reset offset of horizontal
      b_hl     = hl

      ad_eneup_wk[0] = x_eneup //　make coordinate of x,y
      ad_eneup_wk[1] = y_eneup

///　  make energy-line(down)

      y_enedown = y_eneup - hp

      ad_enedown_wk[0] = x_eneup
      ad_enedown_wk[1] = y_eneup - hp

///　 make water-slope-line (up)


      y_glineup = y_eneup - vhead
      ad_glineup_wk[0] = x_eneup
      ad_glineup_wk[1] = y_eneup - vhead

///　 make water-slope-line (up)

      ad_glinedown_wk[0] = x_eneup
      ad_glinedown_wk[1] = y_glineup - hp

///
///　     set various inf.  in slice of struct
///

//      ad_eneup = append( ad_eneup , type3.Point {  ad_eneup_wk[0]  ,
//                                                   ad_eneup_wk[1]  ,
//                                                   ad_eneup_wk[2]   })

//      ad_enedown = append( ad_enedown , type3.Point {  ad_enedown_wk[0]  ,
//                                                       ad_enedown_wk[1] ,
//                                                       ad_enedown_wk[2]   })

//      ad_glineup = append( ad_glineup , type3.Point {  ad_glineup_wk[0]  ,
//                                                       ad_glineup_wk[1]  ,
//                                                       ad_glineup_wk[2]   })

//      ad_glinedown = append( ad_glinedown , type3.Point {  ad_glinedown_wk[0]  ,
//                                                           ad_glinedown_wk[1]  ,
//                                                           ad_glinedown_wk[2]  })

      index ++


   }

//   p_number := index   //  set point-no

//   return p_number ,ad_eneup ,ad_enedown ,ad_glineup ,ad_glinedown
     return

}