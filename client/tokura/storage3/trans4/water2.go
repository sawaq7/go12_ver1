package trans4

import (

	    "net/http"
//	    "fmt"

	    "bufio"

	    "github.com/sawaq7/go12_ver1/storage2"
	    "io"

	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	    "github.com/sawaq7/go12_ver1/client/tokura/storage3/struct_set"
                                                )

///                           　　　　　　　　　　　
///   ストレチE��から水路ファイル惁E��をGETする
///                          　　　　　　　　　　　

func Water2( w http.ResponseWriter, r *http.Request )  ([]type4.Water2 ) {

//     IN     w         : レスポンスライター
//     IN     r         : リクエストパラメーター
//     OUT        　　  : 水路ファイル・チE�Eタストアのスライス

//   fmt.Fprintf( w, "trans4.water2 start \n" )  // チE��チE��

    bucket := "sample-7777"
    filename1 := "Water2.txt"

    water2_view := make([]type4.Water2, 0)   //   Water2　の表示エリアを確俁E
///
///     Water2 ファイル�E�ストレチE���E�オープン
///

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename1 , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)  // インターフェイス型を型変換

//    reader := storage2.File_Open(w ,r ,bucket ,filename1)

    defer reader.Close()

// ファイルリーダー(string用�E�を�E��E��E�

    sreader := bufio.NewReaderSize(reader, 4096)

    index := 0

    for {

      index ++     // レコードカウンターをカウンチE
//      fmt.Fprintf(w, "trans4.water2 : lndex %v\n", index )  // チE��チE��

// ファイルを１行read

      line ,_  := sreader.ReadString('\n')

      num := len(line)

//      fmt.Fprintf(w, "trans4.water2 : num %v\n", num )  // チE��チE��

      if num > 1 {

//         fmt.Fprintf(w, "trans4.water2 : line %s\n", line )  // チE��チE��

///
///   ラインチE�Eタを、構造体にセチE��
///

         water2_struct := struct_set.Water2( w , line )

         water2_view = append( water2_view ,water2_struct )   // ラインチE�Eタを追加

      } else if num == 0 {

//          io.WriteString(w, "\n trans4.water2 : data end \n")   //チE��チE��

         break

      }
   }

   return	water2_view

}

