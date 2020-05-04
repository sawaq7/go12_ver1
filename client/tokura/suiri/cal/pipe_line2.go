package cal

import (
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/tokura/equation"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	    "github.com/sawaq7/go12_ver1/basic/type3"
//	    "strings"
//	    "strconv"
    	                 )

func Pipe_line2( water type4.Water2 ,water_line []type4.Water_Line  ) ( ) {



//     IN  wdeta :
//    OUT  one   :
//    OUT  two   :
//    OUT  three :
//    OUT  four  :
//    OUT  five  :
//    OUT  six   :
//    OUT  seven :


   var b_length float64
   var x_eneup ,y_eneup  ,y_enedown ,y_glineup float64

   var hp ,hl ,b_hl,vhead float64



//   fmt.Println ("cal.pipe_line2 water %v\n",water )
//   fmt.Println ("cal.pipe_line2 water_line %v\n",water_line )

//   allocate various work-area

   ad_hp := make([]float64 ,20 ,50)        // hp　
   ad_hl := make([]float64 ,20 ,50)        // 　hl　
   ad_vhead := make([]float64 ,20 ,50)     // vhead
   ad_eneup := make([]type3.Point ,20 ,50)     // 　eneup


   ad_enedown := make([]type3.Point ,20 ,50)   // 　enedown
   ad_glineup := make([]type3.Point ,20 ,50)   // 　glineup
   ad_glinedown := make([]type3.Point ,20 ,50) // 　glinedown

   eflag := 0

   line_num := len(water_line)

//   fmt.Println ("cal.pipe_line2 line_num　%v\n",line_num )

   Hmax := water.High

   s_coeff := water.Roughness_Factor

///
///      read one record and process it
///

   index := 0
   for pos, water_linew := range water_line {
     count := pos + 1
     if count == line_num {

        eflag = 1

     }


     f_coeff  := water_linew.Friction_Factor

     velocity := water_linew.Velocity

     diameter := water_linew.Pipe_Diameter

     length   := water_linew.Pipe_Length

     vhead = equation.Suiri_Vhead( velocity )
     hp = f_coeff * vhead

//     fmt.Println("cal.pipe_line2 hp" ,hp)

     ramuda := equation.Suiri_Manningu2( s_coeff ,diameter)
     vhead := equation.Suiri_Vhead( velocity )
     hl = ramuda * (length / diameter) * vhead

     ad_hp[index] = hp

//     fmt.Println("cal.pipe_line2 hp(ad)" ,ad_hp)

     if eflag == 1 {

        hl    = 0.0
        vhead = 0.0
     }

     ad_hl[index] = hl

//     fmt.Println("cal.pipe_line2 hl(ad)　%v\n" ,ad_hl)

     ad_vhead[index] = vhead

//     fmt.Println("cal.pipe_line2 vhead(ad)　%v\n" ,ad_vhead)

      if index == 0 {

         b_length = 0.0   //  initialize offset of horizontal-direction of x,y-coordinate
         x_eneup  = 0.0
         y_eneup = Hmax

      }else{

            y_eneup  = y_enedown - b_hl

      }

//　 make energy-line (up)

      x_eneup  = x_eneup + b_length

      b_length = length    //  reset offset of horizontal-direction
      b_hl     = hl

      ad_eneup[index].X = x_eneup //　make x,y-coordinate
      ad_eneup[index].Y = y_eneup
//         fmt.Println("cal.pipe_line2 eneup(ad)" ,ad_eneup)

//　 make energy-line (down)

      y_enedown = y_eneup - hp

      ad_enedown[index].X = x_eneup
      ad_enedown[index].Y = y_eneup - hp

//         fmt.Println("cal.pipe_line2 enedown(ad)" ,ad_enedown)

//　 make water-slope-line (up)


      y_glineup = y_eneup - vhead
      ad_glineup[index].X = x_eneup
      ad_glineup[index].Y = y_eneup - vhead

//         fmt.Println("cal.pipe_line2 glinedown(ad)" ,ad_glineup)

//　 make water-slope-line (down)

      ad_glinedown[index].X = x_eneup
      ad_glinedown[index].Y = y_glineup - hp

//         fmt.Println("cal.pipe_line2 glinedown(ad)" ,ad_glinedown)
      index ++

   }

///   set point-number

   p_number := index

   _ = p_number
   _ = ad_eneup
   _ = ad_enedown
   _ = ad_glineup
   _ = ad_glinedown

   return

}