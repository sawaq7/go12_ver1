///
/// 静水圧　U字管の計算    type2　（ファイル入力）
///

package main

import (
	    "fmt"
	    "os"
	    "strconv"
	    "strings"
	    "bufio"
	    "client/tokura/suiri"
	    "basic/rw"
	    "basic/maths/sum"
    	                 )

func main() {

// 単位容積重量　（ω）をセット

   var omega ,drad1 ,drad2 ,press1 ,press2,high ,area1 ,area2 float64
   var fname ,fname2 ,cdmy string
   var index ,num int

// ファイルオープン

   fname = "C:/Go_Original/src/client/tokura/file/seisui_inf.txt"
   fname2 = "C:/Go_Original/src/client/tokura/file/seisui.txt"

   ad_fdata := make([]float64 ,6)        // keep work data for etc float data

// 静水情報ファイル、オープン

   file , err := os.Open(fname)

// ファイルリーダーをＧＥＴ
   reader := bufio.NewReaderSize(file, 4096)

// 静水ファイル、オープン

   writer , err := os.OpenFile(fname2, os.O_CREATE|os.O_RDWR, 0666)
   if err !=nil {
      fmt.Println ("seisui2 open error" ,err)
      os.Exit(1)
   }

   index = 0        // レコードカウンターをinitialize

   for {

      index ++     // レコードカウンターをカウント

      fmt.Println ("main_seisui2　index " ,index)  // デバック

// ファイルを１行read

      line ,_  := reader.ReadString('\n')

//文字単位にスペースで分割

      str := strings.Fields(line)

      num = len(str)

      fmt.Println ("main_seisui2　num " ,num)  // デバック

      if num == 0 {  //　END　チェック

         fmt.Println ("main_seisui2 normal end")  // デバック
         goto END
      }

      if index != 1{   // 見出し以外をmake

         omega ,_ =strconv.ParseFloat(str[0],64)
         drad1 ,_ =strconv.ParseFloat(str[1],64)
         drad2 ,_ =strconv.ParseFloat(str[2],64)
         press1 ,_ =strconv.ParseFloat(str[3],64)
         cdmy = str[4]
         high ,_ =strconv.ParseFloat(str[5],64)

         fmt.Println ( "main_seisui2 file data " ,omega, drad1 , drad2 ,press1 ,cdmy ,high ) // デバック

// U字管の面積を計算する

         area1 = sum.Circle_Area(drad1/2)
         area2 = sum.Circle_Area(drad2/2)

         press2 =  suiri.Seisui1( area1 ,area2  ,press1  ,omega  ,high  )

         fmt.Println("main_seisui2 圧力は",press2,"ｔ")   //デバック

// U字管の面積をrewrite
         ad_fdata[0] = omega
         ad_fdata[1] = drad1
         ad_fdata[2] = drad2
         ad_fdata[3] = press1
         ad_fdata[4] = press2
         ad_fdata[5] = high
         rw.Wrline1(  writer , ad_fdata )

      }
   }
   //  final process ,"example file delete ,rename ,close"

   END :

   defer file.Close()
   defer writer.Close()

   return
}






