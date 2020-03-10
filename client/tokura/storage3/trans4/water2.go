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
///   ストレッジから水路ファイル情報をGETする
///                          　　　　　　　　　　　

func Water2( w http.ResponseWriter, r *http.Request )  ([]type4.Water2 ) {

//     IN     w         : レスポンスライター
//     IN     r         : リクエストパラメーター
//     OUT        　　  : 水路ファイル・データストアのスライス

//   fmt.Fprintf( w, "trans4.water2 start \n" )  // デバック

    bucket := "sample-7777"
    filename1 := "Water2.txt"

    water2_view := make([]type4.Water2, 0)   //   Water2　の表示エリアを確保

///
///     Water2 ファイル（ストレッジ）オープン
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

//      fmt.Fprintf(w, "trans4.water2 : lndex %v\n", index )  // デバック

// ファイルを１行read

      line ,_  := sreader.ReadString('\n')

      num := len(line)

//      fmt.Fprintf(w, "trans4.water2 : num %v\n", num )  // デバック

      if num > 1 {

//         fmt.Fprintf(w, "trans4.water2 : line %s\n", line )  // デバック

///
///   ラインデータを、構造体にセット
///

         water2_struct := struct_set.Water2( w , line )

         water2_view = append( water2_view ,water2_struct )   // ラインデータを追加

      } else if num == 0 {

//          io.WriteString(w, "\n trans4.water2 : data end \n")   //デバック

         break

      }
   }

   return	water2_view

}

