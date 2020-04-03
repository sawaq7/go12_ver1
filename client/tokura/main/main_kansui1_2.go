    ///////////////////////////////////////////////////////
   ///    main_kansui1_2                               ///
  ///     豌ｴ霍ｯ諠・ｱ繧偵GET                              ///
 ///      繝・・繧ｿ縺ｯstring蝙・                           ///
////////////////////////////////////////////////////////

package main

import (
	     "fmt"
	     "os"
	     "strings"
	     "bufio"
	     "github.com/sawaq7/go12_ver1/client/tokura/suiri"
	                  )

func main() {

// 繝輔ぃ繧､繝ｫ繧ｪ繝ｼ繝励Φ

   fname  := "C:/Go_Original/src/github.com/sawaq7/go12_ver1/client/tokura/file/water_inf.txt"

// 豌ｴ霍ｯ諠・ｱ繝輔ぃ繧､繝ｫ縲√が繝ｼ繝励Φ

   freader , _ := os.OpenFile(fname, os.O_CREATE|os.O_RDWR, 0666)

// 繝輔ぃ繧､繝ｫ繝ｪ繝ｼ繝繝ｼ繧抵ｼｧ・･・ｴ

   reader := bufio.NewReaderSize(freader, 4096)

// 蟆取ｰｴ蜍ｾ驟咲ｷ壹ヵ繧｡繧､繝ｫ縲√が繝ｼ繝励Φ

   index := 0        // 繝ｬ繧ｳ繝ｼ繝峨き繧ｦ繝ｳ繧ｿ繝ｼ繧段nitialize


   for {

      index ++     // 繝ｬ繧ｳ繝ｼ繝峨き繧ｦ繝ｳ繧ｿ繝ｼ繧偵き繧ｦ繝ｳ繝・
      fmt.Println ("main_kansui1_2 index " ,index)  // 繝・ヰ繝・け

// 繝輔ぃ繧､繝ｫ繧抵ｼ題｡罫ead

      line ,_  := reader.ReadString('\n')


//譁・ｭ怜腰菴阪↓繧ｹ繝壹・繧ｹ縺ｧ蛻・牡

      str := strings.Fields(line)

      num := len(str)

      fmt.Println ("main_kansui1_2 num " ,num)  // 繝・ヰ繝・け

      if num != 0 {
         if index == 1 {  // 繝倥ャ繝繝ｼ繧定｡ｨ遉ｺ

             fmt.Println ("main_kansui1_2 繝倥ャ繝繝ｼwrite " ,line)  // 繝・ヰ繝・け

          }else{         // 隱ｭ縺ｿ鬟帙・縺・
             suiro_name ,water_high ,roughness_factor := suiri.Kansui1_2( line  )
             fmt.Println ("main_kansui1_2 豌ｴ霍ｯ蜷・" ,suiro_name)  // 繝・ヰ繝・け
             fmt.Println ("main_kansui1_2 豌ｴ霍ｯ 鬮・" ,water_high)  // 繝・ヰ繝・け
             fmt.Println ("main_kansui1_2 邊怜ｺｦ菫よ焚 " ,roughness_factor)  // 繝・ヰ繝・け

          }

      } else if num == 0 { // End check

          fmt.Println ("main_kansui1_2 normal end" )   //繝・ヰ繝・け
         break

      }
   }

//  final process ,"example file delete ,rename ,close"

//   END :


//   os.Remove(fname) // 譌｢蟄倥・髱呎ｰｴ繝輔ぃ繧､繝ｫ繧貞炎髯､

   defer freader.Close()

   return
}






