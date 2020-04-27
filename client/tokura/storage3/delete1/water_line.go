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
///   delete one-record in Water_Line file
///                          　　　　　　　　　　　

func Water_line( w http.ResponseWriter, r *http.Request ,delid int64 ,wname string  ) {

//     IN     w         : レスポンスライター
//     IN     r         : リクエストパラメーター
//     IN   delid       :
//     IN    wname      : water-name

//    fmt.Fprintf( w, "delete1.water_line start \n" )

    var lf_flag int64

    bucket    := "sample-7777"
    filename1 := "Water_Line.txt"
    filename2 := "Water_Line_2.txt"

///
/// 　　　rename file-name
///

    storage2.File_Rename ( w ,r ,bucket ,filename1 ,filename2 )

///
///      open water-file which was renamed
///

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename2 , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)

//    reader := storage2.File_Open(w ,r ,bucket ,filename2)

    sreader := bufio.NewReaderSize(reader, 4096)

///
///        make new water-file
///

    writer_minor , _ := storage2.Storage_basic( "create" ,bucket ,filename1 , w , r  )

    writer, _ := writer_minor.(*storage.Writer)

//    writer := storage2.File_Create( w ,r ,bucket ,filename1 )

    defer writer.Close()

    id_count := 0
    lf_flag   = 1

///
///   read Water_Line　file
///

    for {

      line ,_  := sreader.ReadString('\n')    // read one-record

      num := len(line)

//      fmt.Fprintf(w, "delete1.water_line : line %s\n", line )
//      fmt.Fprintf(w, "delete1.water_line : num %v\n", num )

      if num  > 1 {

         id_count ++

         water_line_struct := struct_set.Water_line( w , line )  //　change format which is struct

         if delid != water_line_struct.Id   {

           if delid <  water_line_struct.Id   &&
              wname == water_line_struct.Name    {

             water_line_struct.Id --

///             fmt.Fprintf(w, "delete1.water_line : water_line_struct.Id 1 %v\n", water_line_struct.Id )
           }

           storage2.File_Write_Struct ( w ,writer ,lf_flag ,water_line_struct )

         }else if water_line_struct.Name != wname {    //　when water-name is difficult

           storage2.File_Write_Struct ( w ,writer ,lf_flag ,water_line_struct )

         }

      } else if num == 0 {    // check whether or not end

          io.WriteString(w, "\n delete1.water_line : data end \n")

         break

      }

   }

///
///       delete work-file
///

   storage2.File_Delete ( w , r ,bucket ,filename2 )

  return

}
