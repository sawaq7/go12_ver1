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
///     get water inf.
///                          　　　　　　　　　　　

func Water2_temp( w http.ResponseWriter, r *http.Request )  ([]type4.Water2_Temp ) {

//     IN     w         : レスポンスライター
//     IN     r         : リクエストパラメーター

//     OUT        　　  : water inf.

//    fmt.Fprintf( w, "trans4.water2_temp start \n" )

    bucket := "sample-7777"
    filename1 := "Water2_Temp.txt"

    water2_temp_view := make([]type4.Water2_Temp, 0)
///
///     open Water2_Temp in storage
///

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename1 , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)

//    reader := storage2.File_Open(w ,r ,bucket ,filename1)

    defer reader.Close()

//    get file reader

    sreader := bufio.NewReaderSize(reader, 4096)

    index := 0

    for {

      index ++
//      fmt.Fprintf(w, "trans4.water2_temp : lndex %v\n", index )

//    read one-record

      line ,_  := sreader.ReadString('\n')

      num := len(line)

//      fmt.Fprintf(w, "trans4.water2_temp : num %v\n", num )

      if num > 1 {

//        fmt.Fprintf(w, "trans4.water2_temp : line %s\n", line )

///
///   set water-temp.-inf.
///

         water2_temp_struct := struct_set.Water2_temp( w , line )

         water2_temp_view = append( water2_temp_view ,water2_temp_struct )   // add one-record

      } else if num == 0 {

//         io.WriteString(w, "\n trans4.water2_temp : data end \n")

         break

      }
   }

   return	water2_temp_view

}

