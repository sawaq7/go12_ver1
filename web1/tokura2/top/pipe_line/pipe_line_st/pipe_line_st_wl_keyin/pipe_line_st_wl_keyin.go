package pipe_line_st_wl_keyin

import (

	    "net/http"
	    "strconv"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/process2"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	    "storage2"
	    "bufio"

	    "io"
	    "github.com/sawaq7/go12_ver1/client/tokura/storage3"

//	    "fmt"

                                                  )

func Pipe_line_st_wl_keyin(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "sky.pipe_line_st_wl_keyin start \n" )  // 繝・ヰ繝・け

    var water2_temp type4.Water2_Temp

    var idmy int64

    bucket := "sample-7777"

    filename1 := "Water2.txt"

///
///       繧ｻ繝ｬ繧ｯ繝医＠縺溘Ξ繧ｳ繝ｼ繝永d繧偵ご繝・ヨ
///

    select_idw , err := strconv.Atoi(r.FormValue("id"))

	if err  != nil {

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    select_id := int64(select_idw)

//    fmt.Fprintf( w, "sky.pipe_line_st_wl_keyin : select_idw %v\n", select_idw )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "sky.pipe_line_st_wl_keyin : select_id %v\n", select_id )  // 繝・ヰ繝・け

///
///     豌ｴ霍ｯ繝輔ぃ繧､繝ｫ繧・ 繧ｪ繝ｼ繝励Φ
///

//    reader := storage2.File_Open(w ,r ,bucket ,filename1)

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename1 , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)  // 繧､繝ｳ繧ｿ繝ｼ繝輔ぉ繧､繧ｹ蝙九ｒ蝙句､画鋤

    defer reader.Close()

    sreader := bufio.NewReaderSize(reader, 4096)

    index := 0

    for {

//      fmt.Fprintf(w, "sky.pipe_line_st_wl_keyin : lndex %v\n", index )  // 繝・ヰ繝・け

      line ,_  := sreader.ReadString('\n')   // 繝輔ぃ繧､繝ｫ繧抵ｼ題｡罫ead

      num := len(line)

//      fmt.Fprintf(w, "sky.pipe_line_st_wl_keyin : num %v\n", num )  // 繝・ヰ繝・け

      if num > 1 {

         index ++     // 繝ｬ繧ｳ繝ｼ繝峨き繧ｦ繝ｳ繧ｿ繝ｼ繧偵き繧ｦ繝ｳ繝・
//        fmt.Fprintf(w, "sky.pipe_line_st_wl_keyin : line %s\n", line )  // 繝・ヰ繝・け

         water2_struct_minor , _ := storage3.Storage_tokura( "Water2" ,"struct_set" ,line , idmy , w , r  )

//         water2_struct := struct_set.Water2( w , line )   //  string蝙九・繝・・繧ｿ繧呈ｧ矩菴灘梛縺ｫ螟画鋤

         water2_struct, _ := water2_struct_minor.(type4.Water2)  // 繧､繝ｳ繧ｿ繝ｼ繝輔ぉ繧､繧ｹ蝙九ｒ蝙句､画鋤

         if  water2_struct.Id == select_id {

           water2_temp.Id   = int64( 1 )
           water2_temp.Name = water2_struct.Name
           water2_temp.High = water2_struct.High
           water2_temp.Roughness_Factor = water2_struct.Roughness_Factor

//           fmt.Fprintf( w, "sky.pipe_line_st_wl_keyin : water2_temp %v\n", water2_temp )  // 繝・ヰ繝・け

           break

         }
      } else if num == 0 {

//          io.WriteString(w, "\n sky.pipe_line_st_wl_keyin : can't find data \n")   //繝・ヰ繝・け

         break

      }
   }

///
///         temporary-file繧偵う繝九す繝｣繝ｩ繧､繧ｺ  & 繝九Η繝ｼ繝・・繧ｿ繧偵そ繝・ヨ
///

      _ , _ = storage3.Storage_tokura( "Water2_Temp" ,"initialize" ,water2_temp , idmy , w , r  )

//    initialize3.Water2_temp (w , r ,water2_temp) // temporary-file繧偵う繝九す繝｣繝ｩ繧､繧ｺ

///
/// 縲縲縲縲縲縲繝｢繝九ち繝ｼ陦ｨ遉ｺ
///

   process2.Pipe_line_st_wl_show ( water2_temp.Name ,w , r )

}

