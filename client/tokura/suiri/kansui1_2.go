///    　　　
///    get various water-data from water inf.
///    　　　　

package suiri

import (
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/tokura/equation"
	    "strings"
	    "strconv"
    	                 )

func Kansui1_2( wdeta string  ) (string ,string ,string) {

//    OUT  one   : water-no
//    OUT  two   : water-high
//    OUT  three : roughness_factor

   var f_coeff ,velocity ,s_coeff ,diameter ,length ,b_length float64
   var x_eneup ,y_eneup ,x_enedown ,y_enedown float64
   var x_glineup ,y_glineup ,x_glinedown ,y_glinedown float64
   var Hmax ,hp ,hl ,b_hl,vhead float64
   var tflag ,wflag ,eflag ,index int
   var char ,water_name ,water_high ,roughness_factor  string

//   fmt.Println ("func kansui1_2 水路チE�Eタ　",wdeta )

   // change water data from string-type to string-array-type by spliting brank
   str := strings.Fields(wdeta)

//   fmt.Println ("kansui1_2 nummax" ,len(str))

///    allocate work-area

   ad_wk := make([]float64 ,2 ,5 )
   ad_wk2 := make([]string ,2 ,5 )

   index = 0
   eflag = 0

///
///     make various data from one-string-record
///

   for i := 0 ; i < len(str) ; i++ {

      // separate string data by spliting comma
      char = str[i]
      str2 := strings.Split(char, ","  )
//      fmt.Println("kansui1_2 str2" ,str2)

//      fmt.Println ("kansui1_2 num2 " ,len(str2))

      tflag = 0

      for  j := 0 ; j < len(str2) ; j++ {

          char2 := str2[j]
          if i == 0 && j == 0 {
//             fmt.Println ("suiro name " ,char2)
          }
//          fmt.Println ("kansui1_2 char2 " ,char2)
          str3 := strings.Split(char2, ":"  )

//          fmt.Println ("kansui1_2 num3 " ,len(str3))
//          fmt.Println("kansui1_2 str3" ,j ,str3)

          switch str3[0] {

           //   water high
           case "H" :

             Hmax,_ =strconv.ParseFloat(str3[1],64)
             water_high = str3[1]
//             fmt.Println("kansui1_2 Hmax" ,Hmax)

             break;

           //　 s_coeff
           case "n" :

             s_coeff,_ =strconv.ParseFloat(str3[1],64)
             roughness_factor = str3[1]
//             fmt.Println("kansui1_2 s_coeff" ,s_coeff)

             break;

           //  point
           case "pt" :

             tflag = 1

          break;

           //　friction-coeff
           case "f" :

             f_coeff,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("kansui1_2 f_coeff" ,f_coeff)
          break;

           //　  velocity
           case "v" :

             velocity,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("kansui1_2 velocity" ,velocity)

          break;

          // line
          case "len":

             tflag = 2

          break;

          // diameter
          case "d" :

             diameter,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("kansui1_2 diameter" ,diameter)

          break;

          //  length
          case "l" :

             length,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("kansui1_2 length" ,length)

          break;

          //    water name
          case "na" :

             water_name = str3[1]
//             fmt.Println("water name ,str3[1])

          break;

          }

/// 斁E��を刁E��したので、各種チE�Eタを作�Eする

          if j == 2 {

             if tflag == 1 {    // ポイント損失を求めめE
                vhead = equation.Suiri_Vhead( velocity )
                hp = f_coeff * vhead
//                fmt.Println("kansui1_2 hp" ,hp)

             }else if tflag == 2 {   // ライン損失を求めめE
                ramuda := equation.Suiri_Manningu2( s_coeff ,diameter)
                vhead := equation.Suiri_Vhead( velocity )
                hl = ramuda * (length / diameter) * vhead
//                fmt.Println("kansui1_2 hl" ,hl)


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

         if eflag == 1 {    // when the data is last , the data is irregular process
            hl    = 0.0
            vhead = 0.0
         }

///　     make energy line up

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

         ad_wk[0] = x_eneup
         ad_wk[1] = y_eneup

         ad_wk2[0] = strconv.FormatFloat( ad_wk[0], 'f' ,8 ,64 )
         ad_wk2[1] = strconv.FormatFloat( ad_wk[1], 'f' ,8 ,64 )

///　     make energy line up

         x_enedown = x_eneup
         y_enedown = y_eneup - hp
         ad_wk[0] = x_enedown
         ad_wk[1] = y_enedown

         ad_wk2[0] = strconv.FormatFloat( ad_wk[0], 'f' ,8 ,64 )
         ad_wk2[1] = strconv.FormatFloat( ad_wk[1], 'f' ,8 ,64 )

///　     make water-slope-line-up

         x_glineup = x_eneup
         y_glineup = y_eneup - vhead
         ad_wk[0] = x_glineup
         ad_wk[1] = y_glineup

         ad_wk2[0] = strconv.FormatFloat( ad_wk[0], 'f' ,8 ,64 )
         ad_wk2[1] = strconv.FormatFloat( ad_wk[1], 'f' ,8 ,64 )

///　     make water-slope-line-down

         x_glinedown = x_eneup
         y_glinedown = y_glineup - hp

         ad_wk[0] = x_glinedown
         ad_wk[1] = y_glinedown

         ad_wk2[0] = strconv.FormatFloat( ad_wk[0], 'f' ,8 ,64 )
         ad_wk2[1] = strconv.FormatFloat( ad_wk[1], 'f' ,8 ,64 )

         wflag = 0
         index ++

      }
   }

   return water_name ,water_high ,roughness_factor
}
