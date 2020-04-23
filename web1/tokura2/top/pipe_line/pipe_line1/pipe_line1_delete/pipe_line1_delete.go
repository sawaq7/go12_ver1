    ///////////////////////////////////////////////////////
   ///    pipe_line_show                               ///
  ///     show water inf.                      ///
////////////////////////////////////////////////////////

package pipe_line1_delete

import (
//	     "fmt"
	     "strings"
	     "bufio"
	     "io"
	     "net/http"
	     "strconv"
	     "github.com/sawaq7/go12_ver1/client/tokura/suiri"
	     "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	     "github.com/sawaq7/go12_ver1/storage2"
	     "cloud.google.com/go/storage"

	                                     )

///  main process ///

func Pipe_line1_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "pipe_line1_show delete \n" )

    water := make([]type4.Water,100 )

    pos := 0

    bucket := "sample-7777"
    filename1 := "water_inf.txt"
    filename2 := "work.txt"

    water_id , err := strconv.Atoi(r.FormValue("water_id"))
//    fmt.Fprintf( w, "pipe_line1_excute_delete water_id %v\n", water_id )

	if err  != nil {

//	   fmt.Fprintf( w, "pipe_line1_excute_delete :error water_id"  )

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}

//  rename file name

    storage2.File_Rename ( w ,r ,bucket ,filename1 ,filename2 )

//    open new file

//    reader := storage2.File_Open(w ,r ,bucket ,filename2)

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename2 , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)

//    defer reader.Close()

//    get file-reader

    sreader := bufio.NewReaderSize(reader, 4096)

//    create file

    writer_minor , _ := storage2.Storage_basic( "create" ,bucket ,filename1 , w , r  )

    writer, _ := writer_minor.(*storage.Writer)

//    writer := storage2.File_Create( w ,r ,bucket ,filename1 )

    defer writer.Close()


///、open water-slope-line-file

   index := 0

   for {
      index ++
//      fmt.Fprintf(w, "pipe_line1_delete : lndex %v\n", index )

//    read one-record

      line ,_  := sreader.ReadString('\n')

      // change water data from string-type to string-array-type by spliting brank
      str := strings.Fields(line)

      num := len(str)

//      fmt.Fprintf(w, "pipe_line1_delete : num %v\n", num )

      if num != 0 && index != water_id {

          storage2.File_write ( w ,writer ,str )

         if index == 1 {
         //   skip header

//             fmt.Fprintf(w, "pipe_line1_delete (header) : line %s\n", line )

          }else{

         ///     get water's header
             pos ++
//             fmt.Fprintf(w, "pipe_line1_delete (the other): line %s\n", line )

             water[pos-1].No = strconv.Itoa(index)
             water[pos-1].Name ,water[pos-1].High ,water[pos-1].Roughness_factor = suiri.Kansui1_2( line  )

//             fmt.Fprintf(w, "pipe_line1_delete : 水路ナンバ�E %v\n", water[pos-1].No )
//             fmt.Fprintf(w, "pipe_line1_delete : 水路吁E%s\n", water[pos-1] .Name)
//             fmt.Fprintf(w, "pipe_line1_delete : 水路髁E%s\n", water[pos-1].High )
//             fmt.Fprintf(w, "pipe_line1_delete : 粗度係数 %s\n", water[pos-1].Roughness_factor )

          }

      } else if num == 0 {                                         // End check

//          io.WriteString(w, "\n pipe_line1_delete : data end \n")

         break

      }
   }

//    delete work-file

   storage2.Storage_basic( "delete" ,bucket ,filename2 , w , r  )

//   storage2.File_Delete ( w , r ,bucket ,filename2 )

/// compress slice

//   fmt.Fprintf(w, "pipe_line1_delete : len(water) cap(water) %v\n", len(water)  ,cap(water))

   water2 := make([]type4.Water, pos )
   copy ( water2 ,water[0:pos] )

///  show water-line on web

   suiri.Pipe_line1_show( w ,pos , water2 )

//	fmt.Fprintf(w, "\n pipe_line1_delete : Calculate succeeded.\n" )

    return

}
