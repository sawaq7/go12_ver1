    ///////////////////////////////////////////////////////
   ///    main_kansui1_2                               ///
  ///     水路情報を　GET                              ///
 ///      データはstring型                            ///
////////////////////////////////////////////////////////

package main

import (
	     "fmt"
	     "os"
	     "strings"
	     "bufio"
	     "client/tokura/suiri"
	                  )

func main() {

// ファイルオープン

   fname  := "C:/Go_Original/src/client/tokura/file/water_inf.txt"

// 水路情報ファイル、オープン

   freader , _ := os.OpenFile(fname, os.O_CREATE|os.O_RDWR, 0666)

// ファイルリーダーをＧＥＴ

   reader := bufio.NewReaderSize(freader, 4096)

// 導水勾配線ファイル、オープン

   index := 0        // レコードカウンターをinitialize


   for {

      index ++     // レコードカウンターをカウント

      fmt.Println ("main_kansui1_2 index " ,index)  // デバック

// ファイルを１行read

      line ,_  := reader.ReadString('\n')


//文字単位にスペースで分割

      str := strings.Fields(line)

      num := len(str)

      fmt.Println ("main_kansui1_2 num " ,num)  // デバック

      if num != 0 {
         if index == 1 {  // ヘッダーを表示

             fmt.Println ("main_kansui1_2 ヘッダーwrite " ,line)  // デバック

          }else{         // 読み飛ばす

             suiro_name ,water_high ,roughness_factor := suiri.Kansui1_2( line  )
             fmt.Println ("main_kansui1_2 水路名 " ,suiro_name)  // デバック
             fmt.Println ("main_kansui1_2 水路 高 " ,water_high)  // デバック
             fmt.Println ("main_kansui1_2 粗度係数 " ,roughness_factor)  // デバック

          }

      } else if num == 0 { // End check

          fmt.Println ("main_kansui1_2 normal end" )   //デバック
         break

      }
   }

//  final process ,"example file delete ,rename ,close"

//   END :


//   os.Remove(fname) // 既存の静水ファイルを削除

   defer freader.Close()

   return
}






