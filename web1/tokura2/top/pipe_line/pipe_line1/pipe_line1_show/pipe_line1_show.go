package pipe_line1_show

import (
//	     "fmt"
	     "strings"
	     "bufio"
	     "io"
	     "net/http"
	     "strconv"
	     "github.com/sawaq7/go12_ver1/client/tokura/suiri"
	     "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	     "storage2"
	     "html/template"
	     "github.com/sawaq7/go12_ver1/client/tokura/html4"
	                               )

///  main process ///

func Pipe_line1_show(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "Pipe_line1_show start \n" )  // 繝・ヰ繝・け

    water := make([]type4.Water,100 )
//    water := make([]type4.Water,4 )

// 豌ｴ霍ｯ諠・ｱ繧ｫ繧ｦ繝ｳ繧ｿ繝ｼ繧段nitialize

    pos := 0

// 繝舌こ繝・ヨ蜷阪・繝輔ぃ繧､繝ｫ蜷阪繧ｻ繝・ヨ

    bucket := "sample-7777"
    filename1 := "water_inf.txt"

// 豌ｴ霍ｯ諠・ｱ繝輔ぃ繧､繝ｫ縲・・ead file・・繧ｪ繝ｼ繝励Φ

//	reader  := storage2.File_Open(w ,r ,bucket ,filename1)

	reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename1 , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)  // 繧､繝ｳ繧ｿ繝ｼ繝輔ぉ繧､繧ｹ蝙九ｒ蝙句､画鋤

    defer reader.Close()

// 繝輔ぃ繧､繝ｫ繝ｪ繝ｼ繝繝ｼ(string逕ｨ・峨ｒ・ｧ・･・ｴ

    sreader := bufio.NewReaderSize(reader, 4096)

// 蟆取ｰｴ蜍ｾ驟咲ｷ壹ヵ繧｡繧､繝ｫ縲√が繝ｼ繝励Φ

   index := 0        // 繝ｬ繧ｳ繝ｼ繝峨き繧ｦ繝ｳ繧ｿ繝ｼ繧段nitialize


   for {
      index ++     // 繝ｬ繧ｳ繝ｼ繝峨き繧ｦ繝ｳ繧ｿ繝ｼ繧偵き繧ｦ繝ｳ繝・
//      fmt.Fprintf(w, "pipe_line1_show : lndex %v\n", index )  // 繝・ヰ繝・け

// 繝輔ぃ繧､繝ｫ繧抵ｼ題｡罫ead

      line ,_  := sreader.ReadString('\n')

      str := strings.Fields(line)

      num := len(str)

//      fmt.Fprintf(w, "pipe_line1_show : num %v\n", num )  // 繝・ヰ繝・け

      if num != 0 {
         if index == 1 {

// 繝倥ャ繝繝ｼ縺ｯ縲√せ繝ｫ繝ｼ縺吶ｋ

//             fmt.Fprintf(w, "pipe_line1_show (header) : line %s\n", line )  // 繝・ヰ繝・け

          }else{

/// 豌ｴ霍ｯ繝倥ャ繝繝ｼ諠・ｱ繧偵GET
             pos ++     // 豌ｴ霍ｯ諠・ｱ繧ｫ繧ｦ繝ｳ繧ｿ繝ｼ繧偵き繧ｦ繝ｳ繝・
//             fmt.Fprintf(w, "pipe_line1_show (the other): line %s\n", line )  // 繝・ヰ繝・け

             water[pos-1].No = strconv.Itoa(index) //縲謨ｴ謨ｰ繧呈枚蟄励↓螟画鋤
             water[pos-1].Name ,water[pos-1].High ,water[pos-1].Roughness_factor = suiri.Kansui1_2( line  )

//             fmt.Fprintf(w, "pipe_line1_show : 豌ｴ霍ｯ繝翫Φ繝舌・ %v\n", water[pos-1].No )  // 繝・ヰ繝・け
//             fmt.Fprintf(w, "pipe_line1_show : 豌ｴ霍ｯ蜷・%s\n", water[pos-1] .Name)  // 繝・ヰ繝・け
//             fmt.Fprintf(w, "pipe_line1_show : 豌ｴ霍ｯ鬮・%s\n", water[pos-1].High )  // 繝・ヰ繝・け
//             fmt.Fprintf(w, "pipe_line1_show : 邊怜ｺｦ菫よ焚 %s\n", water[pos-1].Roughness_factor )  // 繝・ヰ繝・け

          }

      } else if num == 0 {                                         // End check

//          io.WriteString(w, "\n pipe_line1_show : data end \n")   //繝・ヰ繝・け

         break

      }
   }
//   fmt.Fprintf(w, "pipe_line1_show : len(water) cap(water) %v\n", len(water)  ,cap(water))  // 繝・ヰ繝・け

// 繧ｹ繝ｩ繧､繧ｹ繧貞悸邵ｮ

   water2 := make([]type4.Water, pos )
   copy ( water2 ,water[0:pos] ) // 豕ｨ・壹ョ繝ｼ繧ｿ縺ｯ縲・縲縺九ｉ縲pos・・縲縺ｾ縺ｧ

/// 豌ｴ霍ｯ諠・ｱ縲陦ｨ遉ｺ

   monitor := template.Must(template.New("html").Parse( html4.Pipe_line1_show))
    err := monitor.Execute(w, water2)
    if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	}

//   suiri.Pipe_line1_show( w ,pos , water2 )

// end process

//    fmt.Fprintf(w, "\n pipe_line1_show : Calculate succeeded.\n" )

    return

}
