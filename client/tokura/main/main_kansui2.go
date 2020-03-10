///
/// 管水路　エネルギー線・動水勾配線を描く
///　　　　　データはfloat型　

package main

import (
	    "fmt"
	    "os"
	    "strings"
	    "bufio"
	    "client/tokura/suiri"
	    "basic/rw"
	                  )

func main() {

// 単位容積重量　（ω）をセット

   var fname ,fname2 ,line  string
   var index  ,num int

// ファイルオープン


   fname = "C:/Go_Original/src/client/tokura/file/water_inf.txt"
   fname2 = "C:/Go_Original/src/client/tokura/file/grade_line2.txt"

// 水路ファイル、オープン

//   freader , _ := os.OpenFile(fname, os.O_CREATE|os.O_RDWR, 0666)
   freader , _ := os.Open(fname)

// 導水勾配線ファイル、オープン

   writer , _ := os.OpenFile(fname2, os.O_CREATE|os.O_RDWR, 0666)

   index = 0        // レコードカウンターをinitialize

// ファイルリーダーをＧＥＴ
   reader := bufio.NewReaderSize(freader, 4096)

   for {

      index ++     // レコードカウンターをカウント

      fmt.Println ("main_kansui2 index " ,index)  // デバック

// ファイルを１行read

      line ,_  = reader.ReadString('\n')


//文字単位にスペースで分割

      str := strings.Fields(line)

      num = len(str)

      fmt.Println ("main_kansui2 num " ,num)  // デバック

      if num != 0 {
         if index == 1{

// ヘッダーを書く
             writer.WriteString(line)
             fmt.Println ("main_kansui2 ヘッダーwrite " ,line)  // デバック

          }else{

             fmt.Println ("main_kansui2 データwrite " ,line)  // デバック

/// 動水勾配線のテータを作成

             ad_hp ,ad_hl ,ad_vhead ,ad_eneup ,ad_enedown ,ad_glineup ,ad_glinedown := suiri.Kansui2( line  )
// ポイント損失情報をwrite

             rw.Wrline1( writer ,ad_hp)

// ライン損失情報をwrite

             rw.Wrline1( writer ,ad_hl)

// 速度水頭情報をwrite

             rw.Wrline1( writer ,ad_vhead)

// エネルギー線upをwrite

             rw.Wrline1( writer ,ad_eneup)

// エネルギー線downをwrite

             rw.Wrline1( writer ,ad_enedown)

// 動水勾配線upをwrite

             rw.Wrline1( writer ,ad_glineup)

// 動水勾配線downをwrite

             rw.Wrline1( writer ,ad_glinedown)

          }

      } else if num == 0 {

          fmt.Println ("main_kansui2 normal end" )   //デバック

        break

      }
   }

//  final process ,"example file delete ,rename ,close"

//   END :


//   os.Remove(fname) // 既存の静水ファイルを削除
//   os.Rename(fname2 ,fname) //ワークファイルを静水ファイルとして	再登録

   defer freader.Close()
   defer writer.Close()

   return
}
