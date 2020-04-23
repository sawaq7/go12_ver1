package put1

import (

	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/storage2"
	    "bufio"

	    "io"

	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	    "cloud.google.com/go/storage"

                                                )

///                           　　　　　　　　　　　
///   ストレチE��ファイルに新しく水路惁E��を書ぁE///                          　　　　　　　　　　　

func Water_line( w http.ResponseWriter, r *http.Request ,water_inf type4.Water_Line ) {

//     IN     w         : レスポンスライター
//     IN     r         : リクエストパラメーター
//     IN   water_inf   : 水路ラインのスライス　　struct : Water_Line

//    fmt.Fprintf( w, "put1.water_line start \n" )  // チE��チE��

    var lf_flag int64

    bucket := "sample-7777"
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

    index := 0

    lf_flag = 0

    for {

//      fmt.Fprintf(w, "put1.water_line : lndex %v\n", index )  // チE��チE��

// ファイルを１行read

      line ,_  := sreader.ReadString('\n')

      num := len(line)

//      fmt.Fprintf(w, "put1.water_line : num %v\n", num )  // チE��チE��

      if num > 1 {

         index ++     // レコードカウンターをカウンチE
//         fmt.Fprintf(w, "put1.water_line : line %s\n", line )  // チE��チE��
         storage2.File_Write_Struct ( w ,writer ,lf_flag ,line )

      } else if num == 0 {

          index ++     // レコードカウンターをカウンチE
          lf_flag = 1

          storage2.File_Write_Struct ( w ,writer ,lf_flag ,water_inf )

//         io.WriteString(w, "\n put1.water_line : data end \n")   //チE��チE��

         break

      }
   }

///
/// ワークファイルを削除
///

   storage2.File_Delete ( w , r ,bucket ,filename2 )

  return

}


