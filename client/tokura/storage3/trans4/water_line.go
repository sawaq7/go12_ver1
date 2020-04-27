package trans4

import (

	    "net/http"
//	    "fmt"

	    "bufio"

	    "github.com/sawaq7/go12_ver1/storage2"
	    "io"

	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	    "github.com/sawaq7/go12_ver1/client/tokura/storage3/struct_set"
                                                )

///                           　　　　　　　　　　　
///      get water inf.
///                          　　　　　　　　　　　

func Water_line( wname string ,w http.ResponseWriter, r *http.Request )  ([]type4.Water_Line ) {

//     IN   wname       : water-name　　
//     IN     w         : レスポンスライター
//     IN     r         : リクエストパラメーター
//     OUT        　　  : water-line inf.

//    fmt.Fprintf( w, "trans4.water_line start \n" )

    bucket := "sample-7777"
    filename1 := "Water_Line.txt"

    water_line_view := make([]type4.Water_Line, 0)

///
///     open Water_Line file
///

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename1 , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)

//    reader := storage2.File_Open(w ,r ,bucket ,filename1)

    defer reader.Close()

//   get file reader

    sreader := bufio.NewReaderSize(reader, 4096)

    index := 0

    for {

      index ++
//      fmt.Fprintf(w, "trans4.water_line : lndex %v\n", index )

//     get file-reader

      line ,_  := sreader.ReadString('\n')

      num := len(line)

//      fmt.Fprintf(w, "trans4.water_line : num %v\n", num )

      if num > 1 {

//        fmt.Fprintf(w, "trans4.water_line : line %s\n", line )

///
///   set water-temp.-inf.
///

         water_line_struct := struct_set.Water_line( w , line )

         if water_line_struct.Name == wname {               //   whether or not water name is same with it

           water_line_view = append( water_line_view ,water_line_struct )   // add one-record

         }

      } else if num == 0 {

//          io.WriteString(w, "\n trans4.water_line : data end \n")

         break

      }
   }

   return	water_line_view

}
