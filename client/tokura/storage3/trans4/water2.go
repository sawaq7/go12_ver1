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
///    get water inf.
///                          　　　　　　　　　　　

func Water2( w http.ResponseWriter, r *http.Request )  ([]type4.Water2 ) {

//     IN     w         : レスポンスライター
//     IN     r         : リクエストパラメーター
//     OUT        　　  :  slice of struct's format which is type4.Water2

//   fmt.Fprintf( w, "trans4.water2 start \n" )

    bucket := "sample-7777"
    filename1 := "Water2.txt"

    water2_view := make([]type4.Water2, 0)
///
///     open Water2_Temp in storage
///

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename1 , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)

//    reader := storage2.File_Open(w ,r ,bucket ,filename1)

    defer reader.Close()

//    open Water2_Temp in storage

    sreader := bufio.NewReaderSize(reader, 4096)

    index := 0

    for {

      index ++
//      fmt.Fprintf(w, "trans4.water2 : lndex %v\n", index )

//       read one-record

      line ,_  := sreader.ReadString('\n')

      num := len(line)

//      fmt.Fprintf(w, "trans4.water2 : num %v\n", num )

      if num > 1 {

//         fmt.Fprintf(w, "trans4.water2 : line %s\n", line )

///
///      change format which is struct
///

         water2_struct := struct_set.Water2( w , line )

         water2_view = append( water2_view ,water2_struct )   //   add water data

      } else if num == 0 {

//          io.WriteString(w, "\n trans4.water2 : data end \n")

         break

      }
   }

   return	water2_view

}

