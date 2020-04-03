    ///////////////////////////////////////////////////////
   ///                main_kansui1                     ///
  ///     邂｡豌ｴ霍ｯ縲繧ｨ繝阪Ν繧ｮ繝ｼ邱壹・蜍墓ｰｴ蜍ｾ驟咲ｷ壹ｒ謠上￥        ///
 ///          繝・・繧ｿ縺ｯstring蝙・                       ///
////////////////////////////////////////////////////////

package main

import (
	     "fmt"
	     "os"
	     "strings"
	     "bufio"
	     "github.com/sawaq7/go12_ver1/client/tokura/suiri"
	     "github.com/sawaq7/go12_ver1/basic/rw"
	                  )

func main() {

// 蜊倅ｽ榊ｮｹ遨埼㍾驥上・委会ｼ峨ｒ繧ｻ繝・ヨ

   var fname ,fname2 ,line  string
   var index  ,num int

// 繝輔ぃ繧､繝ｫ繧ｪ繝ｼ繝励Φ


   fname  = "C:/Go_Original/src/github.com/sawaq7/go12_ver1/client/tokura/file/water_inf.txt"
   fname2  = "C:/Go_Original/src/github.com/sawaq7/go12_ver1/client/tokura/file/grade_line.txt"

// 豌ｴ霍ｯ繝輔ぃ繧､繝ｫ縲√が繝ｼ繝励Φ

   freader , _ := os.OpenFile(fname, os.O_CREATE|os.O_RDWR, 0666)
//   freader , _ := os.Open(fname)

// 蟆取ｰｴ蜍ｾ驟咲ｷ壹ヵ繧｡繧､繝ｫ縲√が繝ｼ繝励Φ

   writer , _ := os.OpenFile(fname2, os.O_CREATE|os.O_RDWR, 0666)
//   writer , _ := os.Open(fname2)
   index = 0        // 繝ｬ繧ｳ繝ｼ繝峨き繧ｦ繝ｳ繧ｿ繝ｼ繧段nitialize

// 繝輔ぃ繧､繝ｫ繝ｪ繝ｼ繝繝ｼ繧抵ｼｧ・･・ｴ
   reader := bufio.NewReaderSize(freader, 4096)

   for {

      index ++     // 繝ｬ繧ｳ繝ｼ繝峨き繧ｦ繝ｳ繧ｿ繝ｼ繧偵き繧ｦ繝ｳ繝・
      fmt.Println ("main_kansui1 index " ,index)  // 繝・ヰ繝・け

// 繝輔ぃ繧､繝ｫ繧抵ｼ題｡罫ead

      line ,_  = reader.ReadString('\n')


//譁・ｭ怜腰菴阪↓繧ｹ繝壹・繧ｹ縺ｧ蛻・牡

      str := strings.Fields(line)

      num = len(str)

      fmt.Println ("main_kansui1 num " ,num)  // 繝・ヰ繝・け

      if num != 0 {
         if index == 1{

// 繝倥ャ繝繝ｼ繧呈嶌縺・繝・・繧ｿ髢薙・繝悶Λ繝ｳ繧ｯ縺ｯ谺｡縺ｮ繧ｳ繝ｼ繝峨〒隱ｿ謨ｴ・・             writer.WriteString(line)
             fmt.Println ("main_kansui1 繝倥ャ繝繝ｼwrite " ,line)  // 繝・ヰ繝・け

          }else{

             fmt.Println ("main_kansui1 繝・・繧ｿwrite " ,line)  // 繝・ヰ繝・け

/// 蜍墓ｰｴ蜍ｾ驟咲ｷ壹ｂ繝・・繧ｿ繧剃ｽ懈・

             ad_hp ,ad_hl ,ad_vhead ,ad_eneup ,ad_enedown ,ad_glineup ,ad_glinedown := suiri.Kansui1( line  )

// 繝昴う繝ｳ繝域錐螟ｱ諠・ｱ繧蜘rite

             rw.Wrline2(  writer , ad_hp )

// 繝ｩ繧､繝ｳ謳榊､ｱ諠・ｱ繧蜘rite

             rw.Wrline2(  writer , ad_hl )

// 騾溷ｺｦ豌ｴ鬆ｭ諠・ｱ繧蜘rite

             rw.Wrline2(  writer , ad_vhead )

// 繧ｨ繝阪Ν繧ｮ繝ｼ邱嗽p繧蜘rite

             rw.Wrline2(  writer , ad_eneup )

// 繧ｨ繝阪Ν繧ｮ繝ｼ邱單own繧蜘rite

             rw.Wrline2(  writer , ad_enedown )

// 蜍墓ｰｴ蜍ｾ驟咲ｷ嗽p繧蜘rite

             rw.Wrline2(  writer , ad_glineup )

// 蜍墓ｰｴ蜍ｾ驟咲ｷ單own繧蜘rite

             rw.Wrline2(  writer , ad_glinedown )

          }

      } else if num == 0 {
          fmt.Println ("main_kansui1 normal end" )   //繝・ヰ繝・け
         break

      }
   }

//  final process ,"example file delete ,rename ,close"

//   END :


//   os.Remove(fname) // 譌｢蟄倥・髱呎ｰｴ繝輔ぃ繧､繝ｫ繧貞炎髯､
//   os.Rename(fname2 ,fname) //繝ｯ繝ｼ繧ｯ繝輔ぃ繧､繝ｫ繧帝撕豌ｴ繝輔ぃ繧､繝ｫ縺ｨ縺励※	蜀咲匳骭ｲ

   defer freader.Close()
   defer writer.Close()

   return
}






