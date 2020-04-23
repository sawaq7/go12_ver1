///    　　　
///      make water-slope-line
///    　

package suiri

import (

//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/tokura/equation"
	    "strings"
	    "strconv"
    	                 )

func Kansui2( wdeta string  ) ([]float64 ,[]float64 ,[]float64 ,[]float64,[]float64 ,[]float64 ,[]float64 ){

//     IN  wdeta : water data
//    OUT  one   : point loss
//    OUT  two   : line loss
//    OUT  three : velocity head
//    OUT  four  : energy line up
//    OUT  five  : energy line down
//    OUT  six   : water-slope-line up
//    OUT  seven : water-slope-line down

   var f_coeff ,velocity ,s_coeff ,diameter ,length ,b_length float64
   var x_eneup ,y_eneup ,x_enedown ,y_enedown float64
   var x_glineup ,y_glineup ,x_glinedown ,y_glinedown float64
   var Hmax ,hp ,hl ,b_hl,vhead float64
   var tflag ,wflag ,eflag ,index ,index2 int
   var char string

//   fmt.Println ("func kansui2 wdeta　",wdeta )

   // change water data from string-type to string-array-type by spliting brank
   str := strings.Fields(wdeta)

//   fmt.Println ("kansui2 nummax" ,len(str))

///
///   make work area with string-slice
///

   ad_hp := make([]float64 ,20 ,50)        // ①　hp　
   ad_hl := make([]float64 ,20 ,50)        // ②　hl　
   ad_vhead := make([]float64 ,20 ,50)     // ③　vhead
   ad_eneup := make([]float64 ,20 ,50)     // ④　eneup
   ad_enedown := make([]float64 ,20 ,50)   // ⑤　enedown
   ad_glineup := make([]float64 ,20 ,50)   // ⑥　glineup
   ad_glinedown := make([]float64 ,20 ,50) // ⑦　glinedown

   index = 0
   eflag = 0

///
///     make various data from one-string-record
///

   for i := 0 ; i < len(str) ; i++ {

    // separate string data by spliting comma
      char = str[i]
      str2 := strings.Split(char, ","  )
//      fmt.Println("kansui2 str2" ,str2)

//      fmt.Println ("kansui2 num2 " ,len(str2))

// 計算タイプフラグを、initialize
      tflag = 0

      for  j := 0 ; j < len(str2) ; j++ {

// ヘッダーをread

          char2 := str2[j]
//          fmt.Println ("kansui2 char2 " ,char2)
          str3 := strings.Split(char2, ":"  )

//          fmt.Println ("kansui2 num3 " ,len(str3))
//          fmt.Println("kansui2 str3" ,j ,str3)

          switch str3[0] {

           //    high
           case "H" :

             Hmax,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("kansui2 Hmax" ,Hmax)

            break;

            //　coefficient
            case "n" :

             s_coeff,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("kansui2 s_coeff" ,s_coeff)

             break;

            // point
            case "pt" :

             tflag = 1

          break;

          // f_coeff
          case "f" :

             f_coeff,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("kansui2 f_coeff" ,f_coeff)
          break;

          //　velocity
          case "v" :

             velocity,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("kansui2 velocity" ,velocity)

          break;

          //  line
          case "len":

             tflag = 2

          break;

          //    diameter
          case "d" :

             diameter,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("kansui2 diameter" ,diameter)

          break;

          //    length
          case "l" :

             length,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("kansui2 length" ,length)

          break;

          }

          ///    since adjust various data , do various calculations

          if j == 2 {

             if tflag == 1 {    // make point loss
                vhead = equation.Suiri_Vhead( velocity )   //  calculate  velocity head
                hp = f_coeff * vhead
//                fmt.Println("kansui2 hp" ,hp)

             }else if tflag == 2 {   // make line loss

                ramuda := equation.Suiri_Manningu2( s_coeff ,diameter)  // frictional coefficient
                vhead := equation.Suiri_Vhead( velocity )     //   calculate  velocity head
                hl = ramuda * (length / diameter) * vhead
//                fmt.Println("kansui2 hl" ,hl)


             }
          }
      }

      //  whether or not various data write

      if tflag == 2 {

         wflag = 1

      } else if i == len(str)-1 {

         wflag = 1
         eflag = 1
      }

      // if wflag equal one ,write various-work-area
      if wflag == 1 {

         ad_hp[index] = hp
//         fmt.Println("kansui2 hp(ad)" ,ad_hp)

         if eflag == 1 {     // when the data is last , the data is irregular process
            hl    = 0.0
            vhead = 0.0
         }
         ad_hl[index] = hl
//         fmt.Println("kansui2 hl(ad)" ,ad_hl)

         ad_vhead[index] = vhead
//         fmt.Println("kansui2 vhead(ad)" ,ad_vhead)

         if index == 0 {

            b_length = 0.0
            x_eneup  = 0.0
            y_eneup = Hmax

         }else{

            y_eneup  = y_enedown - b_hl
         }

         x_eneup  = x_eneup + b_length

         b_length = length
         b_hl     = hl

         if index == 0 {

            index2 = 0

         } else {

          index2 = index * 2

         }

///　     make energy line up

         ad_eneup[index2] =  x_eneup   //　make x,y coordinate
         ad_eneup[index2+1] =  y_eneup
//         fmt.Println("kansui2 eneup(ad)" ,ad_eneup)  //　make x,y coordinate

///　     make energy-line-down

         x_enedown = x_eneup
         y_enedown = y_eneup - hp

         ad_enedown[index2] =  x_enedown     //　make x,y coordinate
         ad_enedown[index2+1] =  y_enedown
//         fmt.Println("kansui2 enedown(ad)" ,ad_enedown)  //　make x,y coordinate

///　     make water-slope-line-up

         x_glineup = x_eneup
         y_glineup = y_eneup - vhead

         ad_glineup[index2] = x_glineup       //　make x,y coordinate
         ad_glineup[index2+1] = y_glineup
//         fmt.Println("kansui2 glinedown(ad)" ,ad_glineup)

///　     make water-slope-line-down

         x_glinedown = x_eneup
         y_glinedown = y_glineup - hp

         ad_glinedown[index2] = x_glinedown    //　make x,y coordinate
         ad_glinedown[index2+1] = y_glinedown

//         fmt.Println("kansui2 glinedown(ad)" ,ad_glinedown)

         wflag = 0
         index ++

      }
   }

   return ad_hp ,ad_hl ,ad_vhead ,ad_eneup ,ad_enedown ,ad_glineup ,ad_glinedown
}
