///
///  髱呎ｰｴ蝨ｧ縲U蟄礼ｮ｡縺ｮ險育ｮ・type1・・ey-in蜈･蜉幢ｼ・///

package main

import (
	"fmt"
	"os"
	"github.com/sawaq7/go12_ver1/client/tokura/suiri"
//	"github.com/sawaq7/go12_ver1/basic/maths"
	"github.com/sawaq7/go12_ver1/basic/maths/sum"
	   )

func main() {

// 蜊倅ｽ榊ｮｹ遨埼㍾驥上・委会ｼ峨ｒ繧ｻ繝・ヨ
//    const omega float64 = 1.0

    var a1 ,a2 ,a3, a4,p1 ,p2 ,h ,omega float64
    var cont string

   for{
      fmt.Printf ("蜊倅ｽ榊ｮｹ遨埼㍾驥・ﾏ・縺ｯ・・/緕｡・・)
      fmt.Scanln(&omega)

      fmt.Printf ("U蟄礼ｮ｡・代・逶ｴ蠕・・・茨ｽ搾ｼ・)
      fmt.Scanln(&a1)

      fmt.Printf("U蟄礼ｮ｡・偵・逶ｴ蠕・・・茨ｽ搾ｼ・)
      fmt.Scanln(&a2)

      fmt.Printf ("闕ｷ驥阪・驥阪＆縺ｯ・・・・)
      fmt.Scanln (&p2 )

      fmt.Printf ("鬮伜ｺｦ蟾ｮ縺ｯ・茨ｽ搾ｼ・)
      fmt.Scanln (&h)

/// U蟄礼ｮ｡縺ｮ髱｢遨阪ｒ險育ｮ励☆繧・
      a3 = sum.Circle_Area(a1/2)
      a4 = sum.Circle_Area(a2/2)

      p1 =  suiri.Seisui1( a3 ,a4  ,p2  ,omega  ,h  )

/// 髱呎ｰｴ蝨ｧ繧定ｨ育ｮ励☆繧・
     fmt.Println("髱呎ｰｴ蝨ｧ縺ｯ",p1,"・・)

     fmt.Printf ("next or end ")
     fmt.Scanln (&cont)

     if cont == "next" {
///   縺昴・縺ｾ縺ｾ繧ｹ繝ｫ繝ｼ

      } else if cont == "end" {

         break

      } else {

      fmt.Println ("seisui1.go key-in error")
      os.Exit(1)
      }
   }
}
