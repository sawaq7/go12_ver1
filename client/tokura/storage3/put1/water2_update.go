package put1

import (

	    "net/http"
//	    "fmt"
	    "storage2"
	    "bufio"

	    "io"
        "client/tokura/storage3/struct_set"
	    "client/tokura/suiri/type4"
	    "cloud.google.com/go/storage"
	    "general/type5"

                                                )

///                           　　　　　　　　　　　
///   ストレッジファイルに水路ファイル情報を書き換える
///                          　　　　　　　　　　　

func Water2_update( w http.ResponseWriter, r *http.Request ,updid int64 ,water_inf type4.Water2 ) {

//     IN     w         : レスポンスライター
//     IN     r         : リクエストパラメーター
//     IN     updid     : アップデートするレコードid
//     IN   water_inf   : 水路情報のスライス　　struct : Water2

//    fmt.Fprintf( w, "put1.water2_update start \n" )  // デバック

    var lf_flag int64

    bucket := "sample-7777"
    filename1 := "Water2.txt"
    filename2 := "Water2_2.txt"

    lf_flag   = 1

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

    for {

//      fmt.Fprintf(w, "put1.water2_update : lndex %v\n", index )  // デバック

// ファイルを１行read

      line ,_  := sreader.ReadString('\n')

      num := len(line)

//      fmt.Fprintf(w, "put1.water2_update : num %v\n", num )  // デバック

      if num > 1 {

         index ++     // レコードカウンターをカウント

//         fmt.Fprintf(w, "put1.water2_update : line %s\n", line )  // デバック

         water2_struct := struct_set.Water2( w , line )

         general_work := make([]type5.General_Work, 1 )    // ワークエリア確保

         general_work[0].Sw_Work    = writer     //　ストレッジライターをセット
         general_work[0].Int64_Work = lf_flag    //　改行フラグをセット

         if  water2_struct.Id == updid {

           storage2.Storage_basic( "write2" ,general_work ,water_inf , w , r  )

//         storage2.File_Write_Struct ( w ,writer ,lf_flag ,water_inf )

         } else {

           storage2.Storage_basic( "write2" ,general_work ,water2_struct , w , r  )

//         storage2.File_Write_Struct ( w ,writer ,lf_flag ,water2_struct )

         }

      } else if num == 0 {

//          io.WriteString(w, "\n put1.water2_update : data end \n")   //デバック

         break

      }
   }

///
/// ワークファイルを削除
///

   storage2.File_Delete ( w , r ,bucket ,filename2 )

  return

}


