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

//    fmt.Fprintf( w, "sky.pipe_line_st_wl_keyin start \n" )  // チE��チE��

    var water2_temp type4.Water2_Temp

    var idmy int64

    bucket := "sample-7777"

    filename1 := "Water2.txt"

///
///       セレクトしたレコードidをゲチE��
///

    select_idw , err := strconv.Atoi(r.FormValue("id"))

	if err  != nil {

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    select_id := int64(select_idw)

//    fmt.Fprintf( w, "sky.pipe_line_st_wl_keyin : select_idw %v\n", select_idw )  // チE��チE��
//    fmt.Fprintf( w, "sky.pipe_line_st_wl_keyin : select_id %v\n", select_id )  // チE��チE��

///
///     水路ファイルめE オープン
///

//    reader := storage2.File_Open(w ,r ,bucket ,filename1)

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename1 , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)  // インターフェイス型を型変換

    defer reader.Close()

    sreader := bufio.NewReaderSize(reader, 4096)

    index := 0

    for {

//      fmt.Fprintf(w, "sky.pipe_line_st_wl_keyin : lndex %v\n", index )  // チE��チE��

      line ,_  := sreader.ReadString('\n')   // ファイルを１行read

      num := len(line)

//      fmt.Fprintf(w, "sky.pipe_line_st_wl_keyin : num %v\n", num )  // チE��チE��

      if num > 1 {

         index ++     // レコードカウンターをカウンチE
//        fmt.Fprintf(w, "sky.pipe_line_st_wl_keyin : line %s\n", line )  // チE��チE��

         water2_struct_minor , _ := storage3.Storage_tokura( "Water2" ,"struct_set" ,line , idmy , w , r  )

//         water2_struct := struct_set.Water2( w , line )   //  string型�EチE�Eタを構造体型に変換

         water2_struct, _ := water2_struct_minor.(type4.Water2)  // インターフェイス型を型変換

         if  water2_struct.Id == select_id {

           water2_temp.Id   = int64( 1 )
           water2_temp.Name = water2_struct.Name
           water2_temp.High = water2_struct.High
           water2_temp.Roughness_Factor = water2_struct.Roughness_Factor

//           fmt.Fprintf( w, "sky.pipe_line_st_wl_keyin : water2_temp %v\n", water2_temp )  // チE��チE��

           break

         }
      } else if num == 0 {

//          io.WriteString(w, "\n sky.pipe_line_st_wl_keyin : can't find data \n")   //チE��チE��

         break

      }
   }

///
///         temporary-fileをイニシャライズ  & ニューチE�EタをセチE��
///

      _ , _ = storage3.Storage_tokura( "Water2_Temp" ,"initialize" ,water2_temp , idmy , w , r  )

//    initialize3.Water2_temp (w , r ,water2_temp) // temporary-fileをイニシャライズ

///
/// 　　　　　　モニター表示
///

   process2.Pipe_line_st_wl_show ( water2_temp.Name ,w , r )

}

