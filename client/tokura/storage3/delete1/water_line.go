package delete1

import (

	    "net/http"
//	    "fmt"
	    "storage2"
	    "bufio"
	    "io"

	    "client/tokura/storage3/struct_set"
	    "cloud.google.com/go/storage"

                                                )

///                           　　　　　　　　　　　
///   水路ラインファイルから指示したライン情報を削除する
///                          　　　　　　　　　　　

func Water_line( w http.ResponseWriter, r *http.Request ,delid int64 ,wname string  ) {

//     IN     w         : レスポンスライター
//     IN     r         : リクエストパラメーター
//     IN   delid       : 削除するラインNO　　struct : Water_Line
//     IN    wname      : 水路名

//    fmt.Fprintf( w, "delete1.water_line start \n" )  // デバック

    var lf_flag int64

    bucket    := "sample-7777"
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

    id_count := 0
    lf_flag   = 1

///
///   Water_Line　ファイルを読む
///

    for {

      line ,_  := sreader.ReadString('\n')   // ファイルを１行read

      num := len(line)

//      fmt.Fprintf(w, "delete1.water_line : line %s\n", line )  // デバック
//      fmt.Fprintf(w, "delete1.water_line : num %v\n", num )  // デバック

      if num  > 1 {

         id_count ++

         water_line_struct := struct_set.Water_line( w , line )  //　ラインデータをWater_Lineのフォーマットに変換

         if delid != water_line_struct.Id   {     // 削除レコードをスキップ

           if delid <  water_line_struct.Id   &&
              wname == water_line_struct.Name    {    //レコードNOの調整

             water_line_struct.Id --

//             fmt.Fprintf(w, "delete1.water_line : water_line_struct.Id 1 %v\n", water_line_struct.Id )  // デバック
           }

           storage2.File_Write_Struct ( w ,writer ,lf_flag ,water_line_struct )

         }else if water_line_struct.Name != wname { //　水路名が違う場合ライト

           storage2.File_Write_Struct ( w ,writer ,lf_flag ,water_line_struct )

         }

      } else if num == 0 {    // リード終了かチェック

          io.WriteString(w, "\n delete1.water_line : data end \n")   //デバック

         break

      }

   }

///
/// ワークファイルを削除
///

   storage2.File_Delete ( w , r ,bucket ,filename2 )

  return

}


