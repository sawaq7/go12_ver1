package put1

import (

	    "net/http"
//	    "fmt"
	    "storage2"
	    "bufio"

	    "io"
        "github.com/sawaq7/go12_ver1/client/tokura/storage3/struct_set"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	    "cloud.google.com/go/storage"

                                                )

///                           　　　　　　　　　　　
///   持E��した水路惁E��を変更する
///                          　　　　　　　　　　　

func Water_line_update( w http.ResponseWriter, r *http.Request ,updid int64 ,wname string ,water_inf type4.Water_Line ) {

//     IN     w         : レスポンスライター
//     IN     r         : リクエストパラメーター
//     IN     updid     : アチE�EチE�Eトするレコードid
//     IN    wname      : 水路吁E//     IN   water_inf   : 水路惁E��のスライス　　struct : Water_Line

//   fmt.Fprintf( w, "put1.water_line_update start \n" )  // チE��チE��

    var lf_flag int64

    bucket := "sample-7777"
    filename1 := "Water_Line.txt"
    filename2 := "Water_Line_2.txt"

    lf_flag   = 1

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

    index := 0

    for {

//      fmt.Fprintf(w, "put1.water_line_update : lndex %v\n", index )  // チE��チE��

// ファイルを１行read

      line ,_  := sreader.ReadString('\n')

      num := len(line)

//     fmt.Fprintf(w, "put1.water_line_update : num %v\n", num )  // チE��チE��

      if num > 1 {

         index ++     // レコードカウンターをカウンチE
//         fmt.Fprintf(w, "put1.water_line_update : line %s\n", line )  // チE��チE��

         water_line_struct := struct_set.Water_line( w , line )

       if  water_line_struct.Id   == updid &&
           water_line_struct.Name == wname      {

         storage2.File_Write_Struct ( w ,writer ,lf_flag ,water_inf )

       } else {

         storage2.File_Write_Struct ( w ,writer ,lf_flag ,water_line_struct )

       }

      } else if num == 0 {

//          io.WriteString(w, "\n put1.water_line_update : data end \n")   //チE��チE��

         break

      }
   }

///
/// ワークファイルを削除
///

   storage2.File_Delete ( w , r ,bucket ,filename2 )

  return

}


