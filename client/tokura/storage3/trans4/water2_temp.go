package trans4

import (

	    "net/http"
//	    "fmt"

	    "bufio"

	    "storage2"
	    "io"

	    "client/tokura/suiri/type4"
	    "client/tokura/storage3/struct_set"
                                                )

///                           　　　　　　　　　　　
///   カレントの水路情報をGETする
///                          　　　　　　　　　　　

func Water2_temp( w http.ResponseWriter, r *http.Request )  ([]type4.Water2_Temp ) {

//     IN     w         : レスポンスライター
//     IN     r         : リクエストパラメーター
//     OUT        　　  : 水路ファイル・データストアのスライス

//    fmt.Fprintf( w, "trans4.water2_temp start \n" )  // デバック

    bucket := "sample-7777"
    filename1 := "Water2_Temp.txt"

    water2_temp_view := make([]type4.Water2_Temp, 0)   //   Water2_Temp　の表示エリアを確保

///
///     Water2_Temp ファイル（ストレッジ）オープン
///

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename1 , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)  // インターフェイス型を型変換

//    reader := storage2.File_Open(w ,r ,bucket ,filename1)

    defer reader.Close()

// ファイルリーダー(string用）をＧＥＴ

    sreader := bufio.NewReaderSize(reader, 4096)

    index := 0

    for {

      index ++     // レコードカウンターをカウント

//      fmt.Fprintf(w, "trans4.water2_temp : lndex %v\n", index )  // デバック

// ファイルを１行read

      line ,_  := sreader.ReadString('\n')

      num := len(line)

//      fmt.Fprintf(w, "trans4.water2_temp : num %v\n", num )  // デバック

      if num > 1 {

//        fmt.Fprintf(w, "trans4.water2_temp : line %s\n", line )  // デバック

///
///   ラインデータを、構造体にセット
///

         water2_temp_struct := struct_set.Water2_temp( w , line )

         water2_temp_view = append( water2_temp_view ,water2_temp_struct )   // ラインデータを追加

      } else if num == 0 {

//         io.WriteString(w, "\n trans4.water2_temp : data end \n")   //デバック

         break

      }
   }

   return	water2_temp_view

}

