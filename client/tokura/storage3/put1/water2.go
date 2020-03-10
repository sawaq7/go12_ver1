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
///   ストレッジファイルに水路ファイル情報を書く(水路ファイルがある場合）
///                          　　　　　　　　　　　

func Water2( w http.ResponseWriter, r *http.Request ,water_inf type4.Water2 ) {

//     IN     w         : レスポンスライター
//     IN     r         : リクエストパラメーター
//     IN   water_inf   : 水路情報のスライス　　struct : Water2

//    fmt.Fprintf( w, "put1.water2 start \n" )  // デバック

    var lf_flag int64

    bucket := "sample-7777"
    filename1 := "Water2.txt"
    filename2 := "Water2_2.txt"

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

//    writer := storage2.File_Create( w ,r ,bucket ,filename1 )   // "Water2.txt"を再度作成

    defer writer.Close()

    index := 0

    lf_flag = 0

    for {

//      fmt.Fprintf(w, "put1.water2 : lndex %v\n", index )  // デバック

// ファイルを１行read

      line ,_  := sreader.ReadString('\n')

      num := len(line)

//      fmt.Fprintf(w, "put1.water2 : num %v\n", num )  // デバック

      if num > 1 {

         index ++     // レコードカウンターをカウント

//         fmt.Fprintf(w, "put1.water2 : line %s\n", line )  // デバック
         storage2.File_Write_Struct ( w ,writer ,lf_flag ,line )

      } else if num == 0 {

          index ++     // レコードカウンターをカウント

          lf_flag = 1

          water_inf.Id = int64(index)

          storage2.File_Write_Struct ( w ,writer ,lf_flag ,water_inf )

//         io.WriteString(w, "\n put1.water2 : data end \n")   //デバック

         break

      }
   }

///
/// ワークファイルを削除
///

   storage2.File_Delete ( w , r ,bucket ,filename2 )

  return

}


