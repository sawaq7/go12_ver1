package put1

import (

	    "net/http"
//	    "fmt"
	    "storage2"
	    "bufio"

	    "io"

	    "client/tokura/suiri/type4"
	    "cloud.google.com/go/storage"

                                                )

///                           　　　　　　　　　　　
///   ストレッジファイルに新しく水路情報を書く
///                          　　　　　　　　　　　

func Water_line( w http.ResponseWriter, r *http.Request ,water_inf type4.Water_Line ) {

//     IN     w         : レスポンスライター
//     IN     r         : リクエストパラメーター
//     IN   water_inf   : 水路ラインのスライス　　struct : Water_Line

//    fmt.Fprintf( w, "put1.water_line start \n" )  // デバック

    var lf_flag int64

    bucket := "sample-7777"
    filename1 := "Water_Line.txt"
    filename2 := "Water_Line_2.txt"

///
/// 　　　ファイルのリネーム
///

    storage2.File_Rename ( w ,r ,bucket ,filename1 ,filename2 )

///
///      差し替えた、水路ファイルを　（read file） オープン
///

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename2 , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)  // インターフェイス型を型変換

//    reader := storage2.File_Open(w ,r ,bucket ,filename2)

    sreader := bufio.NewReaderSize(reader, 4096)

///
///      新しく水路ファイルを作成
///

    writer_minor , _ := storage2.Storage_basic( "create" ,bucket ,filename1 , w , r  )

    writer, _ := writer_minor.(*storage.Writer)  // インターフェイス型を型変換

//    writer := storage2.File_Create( w ,r ,bucket ,filename1 )   // "Water_Line.txt"を再度作成

    defer writer.Close()

    index := 0

    lf_flag = 0

    for {

//      fmt.Fprintf(w, "put1.water_line : lndex %v\n", index )  // デバック

// ファイルを１行read

      line ,_  := sreader.ReadString('\n')

      num := len(line)

//      fmt.Fprintf(w, "put1.water_line : num %v\n", num )  // デバック

      if num > 1 {

         index ++     // レコードカウンターをカウント

//         fmt.Fprintf(w, "put1.water_line : line %s\n", line )  // デバック
         storage2.File_Write_Struct ( w ,writer ,lf_flag ,line )

      } else if num == 0 {

          index ++     // レコードカウンターをカウント

          lf_flag = 1

          storage2.File_Write_Struct ( w ,writer ,lf_flag ,water_inf )

//         io.WriteString(w, "\n put1.water_line : data end \n")   //デバック

         break

      }
   }

///
/// ワークファイルを削除
///

   storage2.File_Delete ( w , r ,bucket ,filename2 )

  return

}


