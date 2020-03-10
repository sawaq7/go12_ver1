///    　　　
///    動水勾配線(grade line)の作成
///    　　管水路の計算　　データはfloat型　

package suiri

import (
	    "fmt"
	    "client/tokura/equation"
	    "strings"
	    "strconv"
    	                 )

func Kansui2( wdeta string  ) ([]float64 ,[]float64 ,[]float64 ,[]float64,[]float64 ,[]float64 ,[]float64 ){

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
   var tflag ,wflag ,eflag ,index ,index2 int
   var char string


   fmt.Println ("func kansui2 水路データ　",wdeta )

// ラインデータを、ブランクで分割する

   str := strings.Fields(wdeta)

   fmt.Println ("kansui2 nummax" ,len(str))  // デバック

// 動水勾配線用データ・ワーク用のスライス・index・eflagを　initialize

   ad_hp := make([]float64 ,20 ,50)        // ①　hp　
   ad_hl := make([]float64 ,20 ,50)        // ②　hl　
   ad_vhead := make([]float64 ,20 ,50)     // ③　vhead
   ad_eneup := make([]float64 ,20 ,50)     // ④　eneup
   ad_enedown := make([]float64 ,20 ,50)   // ⑤　enedown
   ad_glineup := make([]float64 ,20 ,50)   // ⑥　glineup
   ad_glinedown := make([]float64 ,20 ,50) // ⑦　glinedown

   index = 0
   eflag = 0

// 1アイテムづつ、read

   for i := 0 ; i < len(str) ; i++ {

// コンマで分割する

      char = str[i]
      str2 := strings.Split(char, ","  )
//      fmt.Println("kansui2 str2" ,str2)  // デバック
// 水路データをsave

//      fmt.Println ("kansui2 num2 " ,len(str2))  // デバック

// 計算タイプフラグを、initialize
      tflag = 0

      for  j := 0 ; j < len(str2) ; j++ {

// ヘッダーをread

          char2 := str2[j]
//          fmt.Println ("kansui2 char2 " ,char2)  // デバック
          str3 := strings.Split(char2, ":"  )

//          fmt.Println ("kansui2 num3 " ,len(str3))  // デバック
//          fmt.Println("kansui2 str3" ,j ,str3)  // デバック

          switch str3[0] {

// 高さの場合
          case "H" :

             Hmax,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("kansui2 Hmax" ,Hmax)  // デバック

             break;

//　粗度係数の場合
          case "n" :

             s_coeff,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("kansui2 s_coeff" ,s_coeff)  // デバック

             break;

// ポイントの場合
          case "pt" :

             tflag = 1

          break;

//　＊＊係数の場合
          case "f" :

             f_coeff,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("kansui2 f_coeff" ,f_coeff)  // デバック
          break;

//　流速の場合
          case "v" :

             velocity,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("kansui2 velocity" ,velocity)  // デバック

          break;

// ラインの場合
          case "len":

             tflag = 2

          break;

// 内径の場合
          case "d" :

             diameter,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("kansui2 diameter" ,diameter)  // デバック

          break;

//長さの場合
          case "l" :

             length,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("kansui2 length" ,length)  // デバック

          break;

          }

/// 文字を分解したので、各種データを作成する

          if j == 2 {

             if tflag == 1 {    // ポイント損失を求める

                vhead = equation.Suiri_Vhead( velocity )  //速度水頭を求める
                hp = f_coeff * vhead
                fmt.Println("kansui2 hp" ,hp)  // デバック

             }else if tflag == 2 {   // ライン損失を求める

                ramuda := equation.Suiri_Manningu2( s_coeff ,diameter)  // 摩擦係数を求める
                vhead := equation.Suiri_Vhead( velocity )  //速度水頭を求める

                hl = ramuda * (length / diameter) * vhead
                fmt.Println("kansui2 hl" ,hl)  // デバック


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

         ad_hp[index] = hp
         fmt.Println("kansui2 hp(ad)" ,ad_hp)  // デバック

         if eflag == 1 {     // ラストデータの場合、速度水頭と摩擦損失は０

            hl    = 0.0
            vhead = 0.0
         }
         ad_hl[index] = hl
         fmt.Println("kansui2 hl(ad)" ,ad_hl)  // デバック

         ad_vhead[index] = vhead
         fmt.Println("kansui2 vhead(ad)" ,ad_vhead)  // デバック

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
         b_hl     = hl        //  摩擦損失をリセット

         if index == 0 {      // x,y座標用のindexをset

            index2 = 0

         } else {

          index2 = index * 2

         }

         ad_eneup[index2] =  x_eneup   //　x,y座標の作成
         ad_eneup[index2+1] =  y_eneup
         fmt.Println("kansui2 eneup(ad)" ,ad_eneup)  // デバック

//   エネルギー線を作成 (down)

         x_enedown = x_eneup
         y_enedown = y_eneup - hp

         ad_enedown[index2] =  x_enedown      //　x,y座標をセット
         ad_enedown[index2+1] =  y_enedown
         fmt.Println("kansui2 enedown(ad)" ,ad_enedown)  // デバック

//   動水勾配線を作成 (up)

         x_glineup = x_eneup
         y_glineup = y_eneup - vhead

         ad_glineup[index2] = x_glineup       //　x,y座標をセット
         ad_glineup[index2+1] = y_glineup
         fmt.Println("kansui2 glinedown(ad)" ,ad_glineup)  // デバック

//   動水勾配線を作成 (up)

         x_glinedown = x_eneup
         y_glinedown = y_glineup - hp

         ad_glinedown[index2] = x_glinedown    //　x,y座標をセット
         ad_glinedown[index2+1] = y_glinedown

         fmt.Println("kansui2 glinedown(ad)" ,ad_glinedown)  // デバック

         wflag = 0
         index ++

      }
   }

   return ad_hp ,ad_hl ,ad_vhead ,ad_eneup ,ad_enedown ,ad_glineup ,ad_glinedown
}