///    　　　
///    水路惁E��より水路名をget
///    　　　　

package suiri

import (
	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/tokura/equation"
	    "strings"
	    "strconv"
    	                 )

func Kansui1_2( wdeta string  ) (string ,string ,string) {

//    OUT  one   : 水路ナンバ�E
//    OUT  two   : 水路吁E//    OUT  three : 水路髁E//    OUT  four  : 粗度係数


   var f_coeff ,velocity ,s_coeff ,diameter ,length ,b_length float64
   var x_eneup ,y_eneup ,x_enedown ,y_enedown float64
   var x_glineup ,y_glineup ,x_glinedown ,y_glinedown float64
   var Hmax ,hp ,hl ,b_hl,vhead float64
   var tflag ,wflag ,eflag ,index int
   var char ,water_name ,water_high ,roughness_factor  string


   fmt.Println ("func kansui1_2 水路チE�Eタ　",wdeta )

// ラインチE�Eタを、ブランクで刁E��する

   str := strings.Fields(wdeta)

   fmt.Println ("kansui1_2 nummax" ,len(str))  // チE��チE��

// 動水勾配線用チE�Eタ・ワーク用のスライス・index・eflagを　initialize

   ad_wk := make([]float64 ,2 ,5 )
   ad_wk2 := make([]string ,2 ,5 )

   index = 0
   eflag = 0

// 1アイチE��づつ、read

   for i := 0 ; i < len(str) ; i++ {

// コンマで刁E��する

      char = str[i]
      str2 := strings.Split(char, ","  )
//      fmt.Println("kansui1_2 str2" ,str2)  // チE��チE��
// 水路チE�Eタをsave

//      fmt.Println ("kansui1_2 num2 " ,len(str2))  // チE��チE��

// 計算タイプフラグを、initialize
      tflag = 0

      for  j := 0 ; j < len(str2) ; j++ {

// ヘッダーをread


          char2 := str2[j]
          if i == 0 && j == 0 {
             fmt.Println ("suiro name " ,char2)  // チE��チE��
          }
//          fmt.Println ("kansui1_2 char2 " ,char2)  // チE��チE��
          str3 := strings.Split(char2, ":"  )

//          fmt.Println ("kansui1_2 num3 " ,len(str3))  // チE��チE��
//          fmt.Println("kansui1_2 str3" ,j ,str3)  // チE��チE��

          switch str3[0] {

// 高さの場吁E          case "H" :

             Hmax,_ =strconv.ParseFloat(str3[1],64)
             water_high = str3[1]
//             fmt.Println("kansui1_2 Hmax" ,Hmax)  // チE��チE��

             break;

//　粗度係数の場吁E          case "n" :

             s_coeff,_ =strconv.ParseFloat(str3[1],64)
             roughness_factor = str3[1]
//             fmt.Println("kansui1_2 s_coeff" ,s_coeff)  // チE��チE��

             break;

// ポイント�E場吁E          case "pt" :

             tflag = 1

          break;

//　�E�＊係数の場吁E          case "f" :

             f_coeff,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("kansui1_2 f_coeff" ,f_coeff)  // チE��チE��
          break;

//　流E���E場吁E          case "v" :

             velocity,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("kansui1_2 velocity" ,velocity)  // チE��チE��

          break;

// ラインの場吁E          case "len":

             tflag = 2

          break;

// 冁E��E�E場吁E          case "d" :

             diameter,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("kansui1_2 diameter" ,diameter)  // チE��チE��

          break;

//長さ�E場吁E          case "l" :

             length,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("kansui1_2 length" ,length)  // チE��チE��

          break;
//水路名�E場吁E          case "na" :

             water_name = str3[1]
             fmt.Println("水路吁E ,str3[1])

          break;

          }

/// 斁E��を刁E��したので、各種チE�Eタを作�Eする

          if j == 2 {

             if tflag == 1 {    // ポイント損失を求めめE
                vhead = equation.Suiri_Vhead( velocity )  //速度水頭を求めめE                hp = f_coeff * vhead
                fmt.Println("kansui1_2 hp" ,hp)  // チE��チE��

             }else if tflag == 2 {   // ライン損失を求めめE
                ramuda := equation.Suiri_Manningu2( s_coeff ,diameter)  // 摩擦係数を求めめE                vhead := equation.Suiri_Vhead( velocity )  //速度水頭を求めめE
                hl = ramuda * (length / diameter) * vhead
                fmt.Println("kansui1_2 hl" ,hl)  // チE��チE��


             }
          }
      }

//  書き込み可能か判断する

      if tflag == 2 {

         wflag = 1

      } else if i == len(str)-1 {

         wflag = 1
         eflag = 1
      }

// チE�Eタがそろった�Eで、書き込み持E��し動水勾配線用チE�Eタを作�Eする
      if wflag == 1 {

         if eflag == 1 {     // ラストデータの場合、E��度水頭と摩擦損失は�E�E
            hl    = 0.0
            vhead = 0.0
         }


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

         ad_wk[0] = x_eneup
         ad_wk[1] = y_eneup

         ad_wk2[0] = strconv.FormatFloat( ad_wk[0], 'f' ,8 ,64 )  // 斁E���Eに変換
         ad_wk2[1] = strconv.FormatFloat( ad_wk[1], 'f' ,8 ,64 )  // 斁E���Eに変換


//　 エネルギー線を作�E (down)

         x_enedown = x_eneup
         y_enedown = y_eneup - hp
         ad_wk[0] = x_enedown
         ad_wk[1] = y_enedown

         ad_wk2[0] = strconv.FormatFloat( ad_wk[0], 'f' ,8 ,64 )  // 斁E���Eに変換
         ad_wk2[1] = strconv.FormatFloat( ad_wk[1], 'f' ,8 ,64 )  // 斁E���Eに変換

//　 動水勾配線を作�E (up)

         x_glineup = x_eneup
         y_glineup = y_eneup - vhead
         ad_wk[0] = x_glineup
         ad_wk[1] = y_glineup

         ad_wk2[0] = strconv.FormatFloat( ad_wk[0], 'f' ,8 ,64 )  // 斁E���Eに変換
         ad_wk2[1] = strconv.FormatFloat( ad_wk[1], 'f' ,8 ,64 )  // 斁E���Eに変換

//　 動水勾配線を作�E (up)

         x_glinedown = x_eneup
         y_glinedown = y_glineup - hp

         ad_wk[0] = x_glinedown
         ad_wk[1] = y_glinedown

         ad_wk2[0] = strconv.FormatFloat( ad_wk[0], 'f' ,8 ,64 )  // 斁E���Eに変換
         ad_wk2[1] = strconv.FormatFloat( ad_wk[1], 'f' ,8 ,64 )  // 斁E���Eに変換

         wflag = 0
         index ++

      }
   }

   return water_name ,water_high ,roughness_factor
}
