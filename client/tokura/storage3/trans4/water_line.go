package trans4

import (

	    "net/http"
//	    "fmt"

	    "bufio"

	    "storage2"
	    "io"

	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	    "github.com/sawaq7/go12_ver1/client/tokura/storage3/struct_set"
                                                )

///                           　　　　　　　　　　　
///   持E��した水路の水路ライン惁E��をGETする
///                          　　　　　　　　　　　

func Water_line( wname string ,w http.ResponseWriter, r *http.Request )  ([]type4.Water_Line ) {

//     IN   wname       : 水路吁E　　　　
//     IN     w         : レスポンスライター
//     IN     r         : リクエストパラメーター
//     OUT        　　  : 水路ファイル・チE�Eタストアのスライス

//    fmt.Fprintf( w, "trans4.water_line start \n" )  // チE��チE��

    bucket := "sample-7777"
    filename1 := "Water_Line.txt"

    water_line_view := make([]type4.Water_Line, 0)   //   Water_Line　の表示エリアを確俁E
///
///     Water_Line ファイル�E�ストレチE���E�オープン
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
//      fmt.Fprintf(w, "trans4.water_line : lndex %v\n", index )  // チE��チE��

// ファイルを１行read

      line ,_  := sreader.ReadString('\n')

      num := len(line)

//      fmt.Fprintf(w, "trans4.water_line : num %v\n", num )  // チE��チE��

      if num > 1 {

//        fmt.Fprintf(w, "trans4.water_line : line %s\n", line )  // チE��チE��

///
///   ラインチE�Eタを、構造体にセチE��
///

         water_line_struct := struct_set.Water_line( w , line )

         if water_line_struct.Name == wname {                         // 水路名が同じかチェチE��

           water_line_view = append( water_line_view ,water_line_struct )   // ラインチE�Eタを追加

         }

      } else if num == 0 {

//          io.WriteString(w, "\n trans4.water_line : data end \n")   //チE��チE��

         break

      }
   }

   return	water_line_view

}

