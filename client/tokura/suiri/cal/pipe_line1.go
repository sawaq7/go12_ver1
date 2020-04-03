package cal

import (
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/tokura/equation"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	    "github.com/sawaq7/go12_ver1/basic/type3"
//	    "strings"
//	    "strconv"
    	                 )

// func Pipe_line1( wdeta string  ) ([]string ,[]string ,[]string ,[]string,[]string ,[]string ,[]string ) {

func Pipe_line1( water type4.Water2 ,water_line []type4.Water_Line  ) (int ,[]type3.Point,[]type3.Point ,[]type3.Point ,[]type3.Point ) {



//     IN  wdeta : 水路チE�Eタ
//    OUT  one   : ポイント損失のスライス
//    OUT  two   : ライン損失のスライス
//    OUT  three : 速度水頭のスライス
//    OUT  four  : エネルギー線！Ep�E��Eスライス
//    OUT  five  : エネルギー線！Eown�E��Eスライス
//    OUT  six   : 導水勾配線！Ep�E��Eスライス
//    OUT  seven : 導水勾配線！Eown�E��Eスライス


   var b_length float64
   var x_eneup ,y_eneup  ,y_enedown ,y_glineup float64

   var hp ,hl ,b_hl,vhead float64



//   fmt.Println ("cal.pipe_line1 water %v\n",water )
//   fmt.Println ("cal.pipe_line1 water_line %v\n",water_line )

// 動水勾配線用チE�Eタ・ワーク用のスライス・index・eflagを　initialize

   ad_hp := make([]float64 ,20 ,50)        // ①　hp　
   ad_hl := make([]float64 ,20 ,50)        // ②　hl　
   ad_vhead := make([]float64 ,20 ,50)     // ③　vhead
   ad_eneup := make([]type3.Point ,20 ,50)     // ④　eneup


   ad_enedown := make([]type3.Point ,20 ,50)   // ⑤　enedown
   ad_glineup := make([]type3.Point ,20 ,50)   // ⑥　glineup
   ad_glinedown := make([]type3.Point ,20 ,50) // ⑦　glinedown




   eflag := 0

   line_num := len(water_line) // 水路ライン数セチE��

//   fmt.Println ("cal.pipe_line1 line_num　%v\n",line_num )

   Hmax := water.High   // 水路H-MAXをセチE��

   s_coeff := water.Roughness_Factor   //　粗度係数をセチE��

///
///   1水路ライン、read
///
   index := 0
   for pos, water_linew := range water_line {
     count := pos + 1
     if count == line_num {

        eflag = 1

     }


     f_coeff  := water_linew.Friction_Factor  // 摩擦係数をセチE��

     velocity := water_linew.Velocity         // 速度をセチE��

     diameter := water_linew.Pipe_Diameter    // 管征E
     length   := water_linew.Pipe_Length      // 管長

/// ポイント損失を求めめE
     vhead = equation.Suiri_Vhead( velocity )  //速度水頭を求めめE     hp = f_coeff * vhead

//     fmt.Println("cal.pipe_line1 hp" ,hp)  // チE��チE��

/// ライン損失を求めめE
     ramuda := equation.Suiri_Manningu2( s_coeff ,diameter)  // 摩擦係数を求めめE     vhead := equation.Suiri_Vhead( velocity )               //速度水頭を求めめE     hl = ramuda * (length / diameter) * vhead

//     fmt.Println("cal.pipe_line1 hl" ,hl)  // チE��チE��

// 動水勾配線用チE�Eタを作�Eする

     ad_hp[index] = hp

//     fmt.Println("cal.pipe_line1 hp(ad)" ,ad_hp)  // チE��チE��

     if eflag == 1 {     // ラストデータの場合、E��度水頭と摩擦損失は�E�E
        hl    = 0.0
        vhead = 0.0
     }

     ad_hl[index] = hl

//     fmt.Println("cal.pipe_line1 hl(ad)　%v\n" ,ad_hl)  // チE��チE��

     ad_vhead[index] = vhead

//     fmt.Println("cal.pipe_line1 vhead(ad)　%v\n" ,ad_vhead)  // チE��チE��

//　 エネルギー線を作�E (up)



      if index == 0 {

         b_length = 0.0   //  x,y座樁E水平方向�EオフセチE��をinitialize
         x_eneup  = 0.0
         y_eneup = Hmax

      }else{

            y_eneup  = y_enedown - b_hl

      }

      x_eneup  = x_eneup + b_length

      b_length = length    //  水平方向�EオフセチE��をリセチE��
      b_hl     = hl

      ad_eneup[index].X = x_eneup //　x,y座標�E作�E
      ad_eneup[index].Y = y_eneup
//         fmt.Println("cal.pipe_line1 eneup(ad)" ,ad_eneup)  // チE��チE��

//　 エネルギー線を作�E (down)

      y_enedown = y_eneup - hp

      ad_enedown[index].X = x_eneup
      ad_enedown[index].Y = y_eneup - hp

//         fmt.Println("cal.pipe_line1 enedown(ad)" ,ad_enedown)  // チE��チE��

//　 動水勾配線を作�E (up)


      y_glineup = y_eneup - vhead
      ad_glineup[index].X = x_eneup
      ad_glineup[index].Y = y_eneup - vhead

//         fmt.Println("cal.pipe_line1 glinedown(ad)" ,ad_glineup)  // チE��チE��

//　 動水勾配線を作�E (up)

      ad_glinedown[index].X = x_eneup
      ad_glinedown[index].Y = y_glineup - hp

//         fmt.Println("cal.pipe_line1 glinedown(ad)" ,ad_glinedown)  // チE��チE��
      index ++


   }
/// ポイント数セチE��　///
   p_number := index

   return p_number ,ad_eneup ,ad_enedown ,ad_glineup ,ad_glinedown

}
