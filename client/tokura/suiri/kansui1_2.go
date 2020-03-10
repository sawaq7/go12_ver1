///    　　　
///    水路情報より水路名をget
///    　　　　

package suiri

import (
	    "fmt"
	    "client/tokura/equation"
	    "strings"
	    "strconv"
    	                 )

func Kansui1_2( wdeta string  ) (string ,string ,string) {

//    OUT  one   : 水路ナンバー
//    OUT  two   : 水路名
//    OUT  three : 水路高
//    OUT  four  : 粗度係数


   var f_coeff ,velocity ,s_coeff ,diameter ,length ,b_length float64
   var x_eneup ,y_eneup ,x_enedown ,y_enedown float64
   var x_glineup ,y_glineup ,x_glinedown ,y_glinedown float64
   var Hmax ,hp ,hl ,b_hl,vhead float64
   var tflag ,wflag ,eflag ,index int
   var char ,water_name ,water_high ,roughness_factor  string


   fmt.Println ("func kansui1_2 水路データ　",wdeta )

// ラインデータを、ブランクで分割する

   str := strings.Fields(wdeta)

   fmt.Println ("kansui1_2 nummax" ,len(str))  // デバック

// 動水勾配線用データ・ワーク用のスライス・index・eflagを　initialize

   ad_wk := make([]float64 ,2 ,5 )
   ad_wk2 := make([]string ,2 ,5 )

   index = 0
   eflag = 0

// 1アイテムづつ、read

   for i := 0 ; i < len(str) ; i++ {

// コンマで分割する

      char = str[i]
      str2 := strings.Split(char, ","  )
//      fmt.Println("kansui1_2 str2" ,str2)  // デバック
// 水路データをsave

//      fmt.Println ("kansui1_2 num2 " ,len(str2))  // デバック

// 計算タイプフラグを、initialize
      tflag = 0

      for  j := 0 ; j < len(str2) ; j++ {

// ヘッダーをread


          char2 := str2[j]
          if i == 0 && j == 0 {
             fmt.Println ("suiro name " ,char2)  // デバック
          }
//          fmt.Println ("kansui1_2 char2 " ,char2)  // デバック
          str3 := strings.Split(char2, ":"  )

//          fmt.Println ("kansui1_2 num3 " ,len(str3))  // デバック
//          fmt.Println("kansui1_2 str3" ,j ,str3)  // デバック

          switch str3[0] {

// 高さの場合
          case "H" :

             Hmax,_ =strconv.ParseFloat(str3[1],64)
             water_high = str3[1]
//             fmt.Println("kansui1_2 Hmax" ,Hmax)  // デバック

             break;

//　粗度係数の場合
          case "n" :

             s_coeff,_ =strconv.ParseFloat(str3[1],64)
             roughness_factor = str3[1]
//             fmt.Println("kansui1_2 s_coeff" ,s_coeff)  // デバック

             break;

// ポイントの場合
          case "pt" :

             tflag = 1

          break;

//　＊＊係数の場合
          case "f" :

             f_coeff,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("kansui1_2 f_coeff" ,f_coeff)  // デバック
          break;

//　流速の場合
          case "v" :

             velocity,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("kansui1_2 velocity" ,velocity)  // デバック

          break;

// ラインの場合
          case "len":

             tflag = 2

          break;

// 内径の場合
          case "d" :

             diameter,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("kansui1_2 diameter" ,diameter)  // デバック

          break;

//長さの場合
          case "l" :

             length,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("kansui1_2 length" ,length)  // デバック

          break;
//水路名の場合
          case "na" :

             water_name = str3[1]
             fmt.Println("水路名" ,str3[1])

          break;

          }

/// 文字を分解したので、各種データを作成する

          if j == 2 {

             if tflag == 1 {    // ポイント損失を求める

                vhead = equation.Suiri_Vhead( velocity )  //速度水頭を求める
                hp = f_coeff * vhead
                fmt.Println("kansui1_2 hp" ,hp)  // デバック

             }else if tflag == 2 {   // ライン損失を求める

                ramuda := equation.Suiri_Manningu2( s_coeff ,diameter)  // 摩擦係数を求める
                vhead := equation.Suiri_Vhead( velocity )  //速度水頭を求める

                hl = ramuda * (length / diameter) * vhead
                fmt.Println("kansui1_2 hl" ,hl)  // デバック


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

// データがそろったので、書き込み指示し動水勾配線用データを作成する
      if wflag == 1 {

         if eflag == 1 {     // ラストデータの場合、速度水頭と摩擦損失は０

            hl    = 0.0
            vhead = 0.0
         }


//　 エネルギー線を作成 (up)



         if index == 0 {
            b_length = 0.0   //  x,y座標 水平方向のオフセットをinitialize
            x_eneup  = 0.0
            y_eneup = Hmax
         }else{
            y_eneup  = y_enedown - b_hl
         }
         x_eneup  = x_eneup + b_length

         b_length = length    //  水平方向のオフセットをリセット
         b_hl     = hl

         ad_wk[0] = x_eneup
         ad_wk[1] = y_eneup

         ad_wk2[0] = strconv.FormatFloat( ad_wk[0], 'f' ,8 ,64 )  // 文字列に変換
         ad_wk2[1] = strconv.FormatFloat( ad_wk[1], 'f' ,8 ,64 )  // 文字列に変換


//　 エネルギー線を作成 (down)

         x_enedown = x_eneup
         y_enedown = y_eneup - hp
         ad_wk[0] = x_enedown
         ad_wk[1] = y_enedown

         ad_wk2[0] = strconv.FormatFloat( ad_wk[0], 'f' ,8 ,64 )  // 文字列に変換
         ad_wk2[1] = strconv.FormatFloat( ad_wk[1], 'f' ,8 ,64 )  // 文字列に変換

//　 動水勾配線を作成 (up)

         x_glineup = x_eneup
         y_glineup = y_eneup - vhead
         ad_wk[0] = x_glineup
         ad_wk[1] = y_glineup

         ad_wk2[0] = strconv.FormatFloat( ad_wk[0], 'f' ,8 ,64 )  // 文字列に変換
         ad_wk2[1] = strconv.FormatFloat( ad_wk[1], 'f' ,8 ,64 )  // 文字列に変換

//　 動水勾配線を作成 (up)

         x_glinedown = x_eneup
         y_glinedown = y_glineup - hp

         ad_wk[0] = x_glinedown
         ad_wk[1] = y_glinedown

         ad_wk2[0] = strconv.FormatFloat( ad_wk[0], 'f' ,8 ,64 )  // 文字列に変換
         ad_wk2[1] = strconv.FormatFloat( ad_wk[1], 'f' ,8 ,64 )  // 文字列に変換

         wflag = 0
         index ++

      }
   }

   return water_name ,water_high ,roughness_factor
}