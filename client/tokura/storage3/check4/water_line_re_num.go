package check4

import (

	    "net/http"
//	    "fmt"

	    "bufio"

	    "storage2"
	    "io"

	    "client/tokura/storage3/struct_set"
                                                )

///                           　　　　　　　　　　　
///   指定した水路名の水路ライン数をゲット
///                          　　　　　　　　　　　

func Water_line_re_num( wname string ,w http.ResponseWriter, r *http.Request )  (record_number int64) {

//     IN   wname       : 水路名 　　　　
//     IN     w         : レスポンスライター
//     IN     r         : リクエストパラメーター

//     OUT        　　  : 水路ライン数

//    fmt.Fprintf( w, "trans4.water_line_re_num start \n" )  // デバック

    bucket := "sample-7777"
    filename1 := "Water_Line.txt"

///
///     Water_Line ファイル（ストレッジ）オープン
///

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename1 , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)  // インターフェイス型を型変換

//    reader := storage2.File_Open(w ,r ,bucket ,filename1)

    defer reader.Close()

// ファイルリーダー(string用）をＧＥＴ

    sreader := bufio.NewReaderSize(reader, 4096)

    index := 0

    record_number = 0

    for {

      index ++     // レコードカウンターをカウント

//      fmt.Fprintf(w, "trans4.water_line_re_num : lndex %v\n", index )  // デバック

// ファイルを１行read

      line ,_  := sreader.ReadString('\n')

      num := len(line)

//      fmt.Fprintf(w, "trans4.water_line_re_num : num %v\n", num )  // デバック

      if num > 1 {

//         fmt.Fprintf(w, "trans4.water_line_re_num : line %s\n", line )  // デバック

///
///   ラインデータを、構造体にセット
///

         water_line_struct := struct_set.Water_line( w , line )

         if water_line_struct.Name == wname {

          record_number ++  // 水路ライン数追加

         }

      } else if num == 0 {

//          io.WriteString(w, "\n trans4.water_line_re_num : data end \n")   //デバック

         break

      }
   }

   return	record_number

}

