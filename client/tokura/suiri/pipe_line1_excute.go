///    　　　
///    動水勾配緁Egrade line)の作�E
///    　　　　チE�Eタはstring型　

package suiri

import (
	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/tokura/equation"
	    "strings"
	    "strconv"
    	                 )

func Pipe_line1_excute( wdeta string  ) ([]string ,[]string ,[]string ,[]string,[]string ,[]string ,[]string ) {

//     IN  wdeta : 水路チE�Eタ
//    OUT  one   : ポイント損失のスライス
//    OUT  two   : ライン損失のスライス
//    OUT  three : 速度水頭のスライス
//    OUT  four  : エネルギー線！Ep�E��Eスライス
//    OUT  five  : エネルギー線！Eown�E��Eスライス
//    OUT  six   : 導水勾配線！Ep�E��Eスライス
//    OUT  seven : 導水勾配線！Eown�E��Eスライス

   var f_coeff ,velocity ,s_coeff ,diameter ,length ,b_length float64
   var x_eneup ,y_eneup ,x_enedown ,y_enedown float64
   var x_glineup ,y_glineup ,x_glinedown ,y_glinedown float64
   var Hmax ,hp ,hl ,b_hl,vhead float64
   var tflag ,wflag ,eflag ,index int
   var char string


   fmt.Println ("func Pipe_line1_excute 水路チE�Eタ　",wdeta )

// ラインチE�Eタを、ブランクで刁E��する

   str := strings.Fields(wdeta)

   fmt.Println ("Pipe_line1_excute nummax" ,len(str))  // チE��チE��

// 動水勾配線用チE�Eタ・ワーク用のスライス・index・eflagを　initialize

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

// 1アイチE��づつ、read

   for i := 0 ; i < len(str) ; i++ {

// コンマで刁E��する

      char = str[i]
      str2 := strings.Split(char, ","  )
//      fmt.Println("Pipe_line1_excute str2" ,str2)  // チE��チE��
// 水路チE�Eタをsave

//      fmt.Println ("Pipe_line1_excute num2 " ,len(str2))  // チE��チE��

// 計算タイプフラグを、initialize
      tflag = 0

      for  j := 0 ; j < len(str2) ; j++ {

// ヘッダーをread

          char2 := str2[j]
//          fmt.Println ("Pipe_line1_excute char2 " ,char2)  // チE��チE��
          str3 := strings.Split(char2, ":"  )

//          fmt.Println ("Pipe_line1_excute num3 " ,len(str3))  // チE��チE��
//          fmt.Println("Pipe_line1_excute str3" ,j ,str3)  // チE��チE��

          switch str3[0] {

// 高さの場吁E          case "H" :

             Hmax,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("Pipe_line1_excute Hmax" ,Hmax)  // チE��チE��

             break;

//　粗度係数の場吁E          case "n" :

             s_coeff,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("Pipe_line1_excute s_coeff" ,s_coeff)  // チE��チE��

             break;

// ポイント�E場吁E          case "pt" :

             tflag = 1

          break;

//　�E�＊係数の場吁E          case "f" :

             f_coeff,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("Pipe_line1_excute f_coeff" ,f_coeff)  // チE��チE��
          break;

//　流E���E場吁E          case "v" :

             velocity,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("Pipe_line1_excute velocity" ,velocity)  // チE��チE��

          break;

// ラインの場吁E          case "len":

             tflag = 2

          break;

// 冁E��E�E場吁E          case "d" :

             diameter,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("Pipe_line1_excute diameter" ,diameter)  // チE��チE��

          break;

//長さ�E場吁E          case "l" :

             length,_ =strconv.ParseFloat(str3[1],64)
//             fmt.Println("Pipe_line1_excute length" ,length)  // チE��チE��

          break;

          }

/// 斁E��を刁E��したので、各種チE�Eタを作�Eする

          if j == 2 {

             if tflag == 1 {    // ポイント損失を求めめE
                vhead = equation.Suiri_Vhead( velocity )  //速度水頭を求めめE                hp = f_coeff * vhead
                fmt.Println("Pipe_line1_excute hp" ,hp)  // チE��チE��

             }else if tflag == 2 {   // ライン損失を求めめE
                ramuda := equation.Suiri_Manningu2( s_coeff ,diameter)  // 摩擦係数を求めめE                vhead := equation.Suiri_Vhead( velocity )  //速度水頭を求めめE
                hl = ramuda * (length / diameter) * vhead
                fmt.Println("Pipe_line1_excute hl" ,hl)  // チE��チE��


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

         ad_hp[index] = strconv.FormatFloat( hp, 'f' ,8 ,64 )
         fmt.Println("Pipe_line1_excute hp(ad)" ,ad_hp)  // チE��チE��

         if eflag == 1 {     // ラストデータの場合、E��度水頭と摩擦損失は�E�E
            hl    = 0.0
            vhead = 0.0
         }
         ad_hl[index] = strconv.FormatFloat( hl, 'f' ,8 ,64 )
         fmt.Println("Pipe_line1_excute hl(ad)" ,ad_hl)  // チE��チE��

         ad_vhead[index] = strconv.FormatFloat( vhead, 'f' ,8 ,64 )
         fmt.Println("Pipe_line1_excute vhead(ad)" ,ad_vhead)  // チE��チE��

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

         ad_eneup[index] = strings.Join( ad_wk2, "," )   //　x,y座標�E作�E
         fmt.Println("Pipe_line1_excute eneup(ad)" ,ad_eneup)  // チE��チE��

//　 エネルギー線を作�E (down)

         x_enedown = x_eneup
         y_enedown = y_eneup - hp
         ad_wk[0] = x_enedown
         ad_wk[1] = y_enedown

         ad_wk2[0] = strconv.FormatFloat( ad_wk[0], 'f' ,8 ,64 )  // 斁E���Eに変換
         ad_wk2[1] = strconv.FormatFloat( ad_wk[1], 'f' ,8 ,64 )  // 斁E���Eに変換

         ad_enedown[index] = strings.Join( ad_wk2, "," )   //　x,y座標�E作�E

         fmt.Println("Pipe_line1_excute enedown(ad)" ,ad_enedown)  // チE��チE��

//　 動水勾配線を作�E (up)

         x_glineup = x_eneup
         y_glineup = y_eneup - vhead
         ad_wk[0] = x_glineup
         ad_wk[1] = y_glineup

         ad_wk2[0] = strconv.FormatFloat( ad_wk[0], 'f' ,8 ,64 )  // 斁E���Eに変換
         ad_wk2[1] = strconv.FormatFloat( ad_wk[1], 'f' ,8 ,64 )  // 斁E���Eに変換

         ad_glineup[index] = strings.Join( ad_wk2, "," )   //　x,y座標�E作�E

         fmt.Println("Pipe_line1_excute glinedown(ad)" ,ad_glineup)  // チE��チE��

//　 動水勾配線を作�E (up)

         x_glinedown = x_eneup
         y_glinedown = y_glineup - hp

         ad_wk[0] = x_glinedown
         ad_wk[1] = y_glinedown

         ad_wk2[0] = strconv.FormatFloat( ad_wk[0], 'f' ,8 ,64 )  // 斁E���Eに変換
         ad_wk2[1] = strconv.FormatFloat( ad_wk[1], 'f' ,8 ,64 )  // 斁E���Eに変換


         ad_glinedown[index] = strings.Join( ad_wk2, "," )   //　x,y座標�E作�E

         fmt.Println("Pipe_line1_excute glinedown(ad)" ,ad_glinedown)  // チE��チE��

         wflag = 0
         index ++

      }
   }

   return ad_hp ,ad_hl ,ad_vhead ,ad_eneup ,ad_enedown ,ad_glineup ,ad_glinedown
}
