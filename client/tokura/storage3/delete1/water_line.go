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
///   水路ラインファイルから持E��したライン惁E��を削除する
///                          　　　　　　　　　　　

func Water_line( w http.ResponseWriter, r *http.Request ,delid int64 ,wname string  ) {

//     IN     w         : レスポンスライター
//     IN     r         : リクエストパラメーター
//     IN   delid       : 削除するラインNO　　struct : Water_Line
//     IN    wname      : 水路吁E
//    fmt.Fprintf( w, "delete1.water_line start \n" )  // チE��チE��

    var lf_flag int64

    bucket    := "sample-7777"
    filename1 := "Water_Line.txt"
    filename2 := "Water_Line_2.txt"

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

//    writer := storage2.File_Create( w ,r ,bucket ,filename1 )   // "Water_Line.txt"を�E度作�E

    defer writer.Close()

    id_count := 0
    lf_flag   = 1

///
///   Water_Line　ファイルを読む
///

    for {

      line ,_  := sreader.ReadString('\n')   // ファイルを１行read

      num := len(line)

//      fmt.Fprintf(w, "delete1.water_line : line %s\n", line )  // チE��チE��
//      fmt.Fprintf(w, "delete1.water_line : num %v\n", num )  // チE��チE��

      if num  > 1 {

         id_count ++

         water_line_struct := struct_set.Water_line( w , line )  //　ラインチE�EタをWater_Lineのフォーマットに変換

         if delid != water_line_struct.Id   {     // 削除レコードをスキチE�E

           if delid <  water_line_struct.Id   &&
              wname == water_line_struct.Name    {    //レコードNOの調整

             water_line_struct.Id --

//             fmt.Fprintf(w, "delete1.water_line : water_line_struct.Id 1 %v\n", water_line_struct.Id )  // チE��チE��
           }

           storage2.File_Write_Struct ( w ,writer ,lf_flag ,water_line_struct )

         }else if water_line_struct.Name != wname { //　水路名が違う場合ライチE
           storage2.File_Write_Struct ( w ,writer ,lf_flag ,water_line_struct )

         }

      } else if num == 0 {    // リード終亁E��チェチE��

          io.WriteString(w, "\n delete1.water_line : data end \n")   //チE��チE��

         break

      }

   }

///
/// ワークファイルを削除
///

   storage2.File_Delete ( w , r ,bucket ,filename2 )

  return

}


