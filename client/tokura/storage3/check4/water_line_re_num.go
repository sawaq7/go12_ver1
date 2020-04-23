package check4

import (

	    "net/http"
//	    "fmt"

	    "bufio"

	    "github.com/sawaq7/go12_ver1/storage2"
	    "io"

	    "github.com/sawaq7/go12_ver1/client/tokura/storage3/struct_set"
                                                )

///                           　　　　　　　　　　　
///   持E��した水路名�E水路ライン数をゲチE��
///                          　　　　　　　　　　　

func Water_line_re_num( wname string ,w http.ResponseWriter, r *http.Request )  (record_number int64) {

//     IN   wname       : 水路吁E　　　　
//     IN     w         : レスポンスライター
//     IN     r         : リクエストパラメーター

//     OUT        　　  : 水路ライン数

//    fmt.Fprintf( w, "trans4.water_line_re_num start \n" )  // チE��チE��

    bucket := "sample-7777"
    filename1 := "Water_Line.txt"

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

    record_number = 0

    for {

      index ++     // レコードカウンターをカウンチE
//      fmt.Fprintf(w, "trans4.water_line_re_num : lndex %v\n", index )  // チE��チE��

// ファイルを１行read

      line ,_  := sreader.ReadString('\n')

      num := len(line)

//      fmt.Fprintf(w, "trans4.water_line_re_num : num %v\n", num )  // チE��チE��

      if num > 1 {

//         fmt.Fprintf(w, "trans4.water_line_re_num : line %s\n", line )  // チE��チE��

///
///   ラインチE�Eタを、構造体にセチE��
///

         water_line_struct := struct_set.Water_line( w , line )

         if water_line_struct.Name == wname {

          record_number ++  // 水路ライン数追加

         }

      } else if num == 0 {

//          io.WriteString(w, "\n trans4.water_line_re_num : data end \n")   //チE��チE��

         break

      }
   }

   return	record_number

}

