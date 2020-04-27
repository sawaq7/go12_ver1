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
///     get water-line inf.
///                          　　　　　　　　　　　

func Water_line_re_num( wname string ,w http.ResponseWriter, r *http.Request )  (record_number int64) {

//     IN   wname       : water-name　　　　
//     IN     w         : レスポンスライター
//     IN     r         : リクエストパラメーター

//     OUT        　　  : water-number

//    fmt.Fprintf( w, "trans4.water_line_re_num start \n" )

    bucket := "sample-7777"
    filename1 := "Water_Line.txt"

///
///     open Water_Line file
///

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename1 , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)

//    reader := storage2.File_Open(w ,r ,bucket ,filename1)

    defer reader.Close()

//      get file-reader

    sreader := bufio.NewReaderSize(reader, 4096)

    index := 0

    record_number = 0

    for {

      index ++
//      fmt.Fprintf(w, "trans4.water_line_re_num : lndex %v\n", index )

//     read one-record

      line ,_  := sreader.ReadString('\n')

      num := len(line)

//      fmt.Fprintf(w, "trans4.water_line_re_num : num %v\n", num )

      if num > 1 {

//         fmt.Fprintf(w, "trans4.water_line_re_num : line %s\n", line )

///
///   change format which is struct
///

         water_line_struct := struct_set.Water_line( w , line )

         if water_line_struct.Name == wname {

          record_number ++  //   add water-line number

         }

      } else if num == 0 {

//          io.WriteString(w, "\n trans4.water_line_re_num : data end \n")

         break

      }
   }

   return	record_number

}
