///
/// 髱呎ｰｴ蝨ｧ縲U蟄礼ｮ｡縺ｮ險育ｮ・   type2縲・医ヵ繧｡繧､繝ｫ蜈･蜉幢ｼ・///

package main

import (
	    "fmt"
	    "os"
	    "strconv"
	    "strings"
	    "bufio"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri"
	    "github.com/sawaq7/go12_ver1/basic/rw"
	    "github.com/sawaq7/go12_ver1/basic/maths/sum"
    	                 )

func main() {

// 蜊倅ｽ榊ｮｹ遨埼㍾驥上・委会ｼ峨ｒ繧ｻ繝・ヨ

   var omega ,drad1 ,drad2 ,press1 ,press2,high ,area1 ,area2 float64
   var fname ,fname2 ,cdmy string
   var index ,num int

// 繝輔ぃ繧､繝ｫ繧ｪ繝ｼ繝励Φ

   fname = "C:/Go_Original/src/github.com/sawaq7/go12_ver1/client/tokura/file/seisui_inf.txt"
   fname2 = "C:/Go_Original/src/github.com/sawaq7/go12_ver1/client/tokura/file/seisui.txt"

   ad_fdata := make([]float64 ,6)        // keep work data for etc float data

// 髱呎ｰｴ諠・ｱ繝輔ぃ繧､繝ｫ縲√が繝ｼ繝励Φ

   file , err := os.Open(fname)

// 繝輔ぃ繧､繝ｫ繝ｪ繝ｼ繝繝ｼ繧抵ｼｧ・･・ｴ
   reader := bufio.NewReaderSize(file, 4096)

// 髱呎ｰｴ繝輔ぃ繧､繝ｫ縲√が繝ｼ繝励Φ

   writer , err := os.OpenFile(fname2, os.O_CREATE|os.O_RDWR, 0666)
   if err !=nil {
      fmt.Println ("seisui2 open error" ,err)
      os.Exit(1)
   }

   index = 0        // 繝ｬ繧ｳ繝ｼ繝峨き繧ｦ繝ｳ繧ｿ繝ｼ繧段nitialize

   for {

      index ++     // 繝ｬ繧ｳ繝ｼ繝峨き繧ｦ繝ｳ繧ｿ繝ｼ繧偵き繧ｦ繝ｳ繝・
      fmt.Println ("main_seisui2縲index " ,index)  // 繝・ヰ繝・け

// 繝輔ぃ繧､繝ｫ繧抵ｼ題｡罫ead

      line ,_  := reader.ReadString('\n')

//譁・ｭ怜腰菴阪↓繧ｹ繝壹・繧ｹ縺ｧ蛻・牡

      str := strings.Fields(line)

      num = len(str)

      fmt.Println ("main_seisui2縲num " ,num)  // 繝・ヰ繝・け

      if num == 0 {  //縲END縲繝√ぉ繝・け

         fmt.Println ("main_seisui2 normal end")  // 繝・ヰ繝・け
         goto END
      }

      if index != 1{   // 隕句・縺嶺ｻ･螟悶ｒmake

         omega ,_ =strconv.ParseFloat(str[0],64)
         drad1 ,_ =strconv.ParseFloat(str[1],64)
         drad2 ,_ =strconv.ParseFloat(str[2],64)
         press1 ,_ =strconv.ParseFloat(str[3],64)
         cdmy = str[4]
         high ,_ =strconv.ParseFloat(str[5],64)

         fmt.Println ( "main_seisui2 file data " ,omega, drad1 , drad2 ,press1 ,cdmy ,high ) // 繝・ヰ繝・け

// U蟄礼ｮ｡縺ｮ髱｢遨阪ｒ險育ｮ励☆繧・
         area1 = sum.Circle_Area(drad1/2)
         area2 = sum.Circle_Area(drad2/2)

         press2 =  suiri.Seisui1( area1 ,area2  ,press1  ,omega  ,high  )

         fmt.Println("main_seisui2 蝨ｧ蜉帙・",press2,"・・)   //繝・ヰ繝・け

// U蟄礼ｮ｡縺ｮ髱｢遨阪ｒrewrite
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






