package delete1

import (

	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/storage2"
	    "bufio"
	    "io"

	    "github.com/sawaq7/go12_ver1/client/tokura/storage3/struct_set"
	    "cloud.google.com/go/storage"
                                                )

///                           　　　　　　　　　　　
///   ストレチE��ファイルから持E��した水路惁E��を削除する
///                          　　　　　　　　　　　

func Water2( w http.ResponseWriter, r *http.Request ,delid int64 ) {

//     IN     w         : レスポンスライター
//     IN     r         : リクエストパラメーター
//     IN   delid       : 削除するラインNO　　struct : Water2

//    fmt.Fprintf( w, "delete1.water2 start \n" )  // チE��チE��

    var lf_flag int64

    bucket    := "sample-7777"
    filename1 := "Water2.txt"
    filename2 := "Water2_2.txt"

///
/// 　　　ファイルのリネ�Eム
///

    storage2.File_Rename ( w ,r ,bucket ,filename1 ,filename2 )

///
///      差し替えた、水路ファイルを　�E�Eead file�E�Eオープン
///

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename2 , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)  // インターフェイス型を型変換

//    reader := storage2.File_Open(w ,r ,bucket ,filename2)

    sreader := bufio.NewReaderSize(reader, 4096)

///
///      新しく水路ファイルを作�E
///

    writer_minor , _ := storage2.Storage_basic( "create" ,bucket ,filename1 , w , r  )

    writer, _ := writer_minor.(*storage.Writer)  // インターフェイス型を型変換

//    writer := storage2.File_Create( w ,r ,bucket ,filename1 )   // "Water2.txt"を�E度作�E

    defer writer.Close()

    index    := 0
    id_count := 0
    lf_flag   = 1

///
///   Water2　ファイルを読む
///

    for {

      line ,_  := sreader.ReadString('\n')   // ファイルを１行read

      num := len(line)

//      fmt.Fprintf(w, "delete1.water2 : line %s\n", line )  // チE��チE��
//      fmt.Fprintf(w, "delete1.water2 : num %v\n", num )  // チE��チE��

      if num  > 1 {

         id_count ++

         if delid != int64(id_count)     {  // 削除レコードをスキチE�E

           index ++     // レコードカウンターをカウンチE
//           fmt.Fprintf(w, "delete1.water2 : lndex %v\n", index )  // チE��チE��

           water2_struct := struct_set.Water2( w , line )  //　ラインチE�EタをWater2のフォーマットに変換

           water2_struct.Id = int64(index)

           storage2.File_Write_Struct ( w ,writer ,lf_flag ,water2_struct )

         }

      } else if num == 0 {    // リード終亁E��チェチE��

//         io.WriteString(w, "\n delete1.water2 : data end \n")   //チE��チE��

         break

      }

   }

///
/// ワークファイルを削除
///

   storage2.File_Delete ( w , r ,bucket ,filename2 )

  return

}


