///    　　　
///    動水勾配線(grade line)の作成
///    　　　　データはstring型　

package suiri

import (
	    "fmt"
	    "client/tokura/equation"
	    "strings"
	    "strconv"
    	                 )

func Kansui1( wdeta string  ) ([]string ,[]string ,[]string ,[]string,[]string ,[]string ,[]string ) {

//     IN  wdeta : 水路データ
//    OUT  one   : ポイント損失のスライス
//    OUT  two   : ライン損失のスライス
//    OUT  three : 速度水頭のスライス
//    OUT  four  : エネルギー線（up）のスライス
//    OUT  five  : エネルギー線（down）のスライス
//    OUT  six   : 導水勾配線（up）のスライス
//    OUT  seven : 導水勾配線（down）のスライス

   var f_coeff ,velocity ,s_coeff ,diameter ,length ,b_length float64
   var x_eneup ,y_eneup ,x_enedown ,y_enedown float64
   var x_glineup ,y_glineup ,x_glinedown ,y_glinedown float64
   var Hmax ,hp ,hl ,b_hl,vhead float64
   var tflag ,wflag ,eflag ,index int
   var char string


   fmt.Println ("func kansui1 水路データ　",wdeta )

// ラインデータを、ブランクで分割する

   str := strings.Fields(wdeta)

   fmt.Println ("kansui1 nummax" ,len(str))  // デバック

// 動水勾配線用データ・ワーク用のスライス・index・eflagを　initialize

   ad_hp := make([]string ,20 ,50)        // ①　hp　
   ad_hl := make([]string ,20 ,50)        // ②　hl　
   ad_vhead := make([]string ,20 ,50)     // ③　vhead
   ad_eneup := make([]string ,20 ,50)     // ④　eneup
   ad_enedown := make([]string ,20 ,50)   // ⑤　enedown
   ad_glineup := make([]string ,20 ,50)   // ⑥　glineup
   ad_glinedown := make([]string ,20 ,50) // ⑦　glinedown
   ad_wk := make([]float64 ,2 ,5 )
   ad_wk2 := make([]string ,2 ,5 )

   index = 0
   eflag = 0

// 1アイテムづつ、read

   for i := 0 ; i < len(str) ; i++ {

// コンマで分割する

      char = str[i]
      str2 := strings.Split(char, ","  )
//      fmt.Println("kansui1 str2" ,str2)  // デバック
// 水路データをsave

//      fmt.Println ("kansui1 num2 " ,len(str2))  // デバック

// 計算タイプフラグを、initialize
      tflag = 0

      for  j := 0 ; j < len(str2) ; j++ {

// ヘッダーをread

          char2 := str2[j]
//          fmt.Println ("kansui1 char2 " ,char2)  // デバック
          str3 := strings.Split(char2, ":"  )

//          fmt.Println ("kansui1 num3 " ,len(str3))  // デバック
//          fmt.Println("kansui1 str3" ,j ,str3)  // デバック

          switch str3[0] {

// 高さの場合
          case "H" :

             Hmax,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("kansui1 Hmax" ,Hmax)  // デバック

             break;

//　粗度係数の場合
          case "n" :

             s_coeff,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("kansui1 s_coeff" ,s_coeff)  // デバック

             break;

// ポイントの場合
          case "pt" :

             tflag = 1

          break;

//　＊＊係数の場合
          case "f" :

             f_coeff,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("kansui1 f_coeff" ,f_coeff)  // デバック
          break;

//　流速の場合
          case "v" :

             velocity,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("kansui1 velocity" ,velocity)  // デバック

          break;

// ラインの場合
          case "len":

             tflag = 2

          break;

// 内径の場合
          case "d" :

             diameter,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("kansui1 diameter" ,diameter)  // デバック

          break;

//長さの場合
          case "l" :

             length,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("kansui1 length" ,length)  // デバック

          break;

          }

/// 文字を分解したので、各種データを作成する

          if j == 2 {

             if tflag == 1 {    // ポイント損失を求める

                vhead = equation.Suiri_Vhead( velocity )  //速度水頭を求める
                hp = f_coeff * vhead
                fmt.Println("kansui1 hp" ,hp)  // デバック

             }else if tflag == 2 {   // ライン損失を求める

                ramuda := equation.Suiri_Manningu2( s_coeff ,diameter)  // 摩擦係数を求める
                vhead := equation.Suiri_Vhead( velocity )  //速度水頭を求める

                hl = ramuda * (length / diameter) * vhead
                fmt.Println("kansui1 hl" ,hl)  // デバック


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

         ad_hp[index] = strconv.FormatFloat( hp, 'f' ,8 ,64 )
         fmt.Println("kansui1 hp(ad)" ,ad_hp)  // デバック

         if eflag == 1 {     // ラストデータの場合、速度水頭と摩擦損失は０

            hl    = 0.0
            vhead = 0.0
         }
         ad_hl[index] = strconv.FormatFloat( hl, 'f' ,8 ,64 )
         fmt.Println("kansui1 hl(ad)" ,ad_hl)  // デバック

         ad_vhead[index] = strconv.FormatFloat( vhead, 'f' ,8 ,64 )
         fmt.Println("kansui1 vhead(ad)" ,ad_vhead)  // デバック

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

         ad_eneup[index] = strings.Join( ad_wk2, "," )   //　x,y座標の作成
         fmt.Println("kansui1 eneup(ad)" ,ad_eneup)  // デバック

//　 エネルギー線を作成 (down)

         x_enedown = x_eneup
         y_enedown = y_eneup - hp
         ad_wk[0] = x_enedown
         ad_wk[1] = y_enedown

         ad_wk2[0] = strconv.FormatFloat( ad_wk[0], 'f' ,8 ,64 )  // 文字列に変換
         ad_wk2[1] = strconv.FormatFloat( ad_wk[1], 'f' ,8 ,64 )  // 文字列に変換

         ad_enedown[index] = strings.Join( ad_wk2, "," )   //　x,y座標の作成

         fmt.Println("kansui1 enedown(ad)" ,ad_enedown)  // デバック

//　 動水勾配線を作成 (up)

         x_glineup = x_eneup
         y_glineup = y_eneup - vhead
         ad_wk[0] = x_glineup
         ad_wk[1] = y_glineup

         ad_wk2[0] = strconv.FormatFloat( ad_wk[0], 'f' ,8 ,64 )  // 文字列に変換
         ad_wk2[1] = strconv.FormatFloat( ad_wk[1], 'f' ,8 ,64 )  // 文字列に変換

         ad_glineup[index] = strings.Join( ad_wk2, "," )   //　x,y座標の作成

         fmt.Println("kansui1 glinedown(ad)" ,ad_glineup)  // デバック

//　 動水勾配線を作成 (up)

         x_glinedown = x_eneup
         y_glinedown = y_glineup - hp

         ad_wk[0] = x_glinedown
         ad_wk[1] = y_glinedown

         ad_wk2[0] = strconv.FormatFloat( ad_wk[0], 'f' ,8 ,64 )  // 文字列に変換
         ad_wk2[1] = strconv.FormatFloat( ad_wk[1], 'f' ,8 ,64 )  // 文字列に変換


         ad_glinedown[index] = strings.Join( ad_wk2, "," )   //　x,y座標の作成

         fmt.Println("kansui1 glinedown(ad)" ,ad_glinedown)  // デバック

         wflag = 0
         index ++

      }
   }

   return ad_hp ,ad_hl ,ad_vhead ,ad_eneup ,ad_enedown ,ad_glineup ,ad_glinedown
}