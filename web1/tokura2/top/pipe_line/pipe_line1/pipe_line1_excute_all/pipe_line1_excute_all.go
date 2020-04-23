////////////////////////////////////////////////////////////////////////
///                                                                  ///
///   make water-slope-line　　　　　　　　　　　　　　　　　　  ///
///                                                                  ///
////////////////////////////////////////////////////////////////////////

package pipe_line1_excute_all

import (
//	      "fmt"
	      "io"
	      "strings"
	      "net/http"
	      "bufio"
	      "github.com/sawaq7/go12_ver1/client/tokura/suiri"
	      "github.com/sawaq7/go12_ver1/storage2"
	      "cloud.google.com/go/storage"
                                         )



///  main process ///

func Pipe_line1_excute_all(w http.ResponseWriter, r *http.Request) {

    bucket := "sample-7777"
    filename1 := "water_inf.txt"
    filename2 := "grade_line.txt"

/// create water-slope-line

    writer_minor , _ := storage2.Storage_basic( "create" ,bucket ,filename2 , w , r  )

    writer, _ := writer_minor.(*storage.Writer)

//	writer := storage2.File_Create( w ,r ,bucket ,filename2 )

	defer writer.Close()

///    open water-inf-file

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename1 , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)

//    reader := storage2.File_Open(w ,r ,bucket ,filename1)

    defer reader.Close()

//    get file-reader

    sreader := bufio.NewReaderSize(reader, 4096)

    index := 0

    for {

      index ++

//      fmt.Fprintf(w, "pipe_line1_excute_all : lndex %v\n", index )

      //    read one-record
      line ,_  := sreader.ReadString('\n')

      // change water data from string-type to string-array-type by spliting brank
      str := strings.Fields(line)

      num := len(str)

//      fmt.Fprintf(w, "pipe_line1_excute_all : num %v\n", num )

      if num != 0 {
         if index == 1{

         //   skip header

//             fmt.Fprintf(w, "pipe_line1_excute_all (header) : line %s\n", line )

          }else{
//             fmt.Fprintf(w, "pipe_line1_excute_all (the other): line %s\n", line )

///    make water-slope-line

             ad_hp ,ad_hl ,ad_vhead ,ad_eneup ,ad_enedown ,ad_glineup ,ad_glinedown := suiri.Kansui1( line  )

//             fmt.Fprintf(w, "pipe_line1_excute_all : ad_hp %s\n", ad_hp )
//             fmt.Fprintf(w, "pipe_line1_excute_all : ad_hl %s\n", ad_hl )
//             fmt.Fprintf(w, "pipe_line1_excute_all : ad_vhead %s\n", ad_vhead )
//             fmt.Fprintf(w, "pipe_line1_excute_all : ad_eneup %s\n", ad_eneup )
//             fmt.Fprintf(w, "pipe_line1_excute_all : ad_enedown %s\n", ad_enedown )
//             fmt.Fprintf(w, "pipe_line1_excute_all : ad_glineup %s\n", ad_glineup )
//             fmt.Fprintf(w, "pipe_line1_excute_all : ad_glinedown %s\n", ad_glinedown )

///     write point-loss

//                 storage2.File_Write ( w ,bucket ,filename2 ,writer ,ad_hp )
                 storage2.File_write ( w ,writer ,ad_hp )

///    write line-loss

//                 storage2.File_Write ( w ,bucket ,filename2 ,writer ,ad_hl )
                 storage2.File_write ( w ,writer ,ad_hl )

///     write vilocity-head

//                 storage2.File_Write ( w ,bucket ,filename2 ,writer ,ad_vhead )
                 storage2.File_write ( w ,writer ,ad_vhead )

///     write enegy-line-up

//                 storage2.File_Write ( w ,bucket ,filename2 ,writer ,ad_eneup )
                 storage2.File_write ( w ,writer ,ad_eneup )

///     write enegy-line-down

//                 storage2.File_Write ( w ,bucket ,filename2 ,writer ,ad_enedown )
                 storage2.File_write ( w ,writer ,ad_enedown )

///     write water-slope-line-up

//                 storage2.File_Write ( w ,bucket ,filename2 ,writer ,ad_glineup )
                 storage2.File_write ( w ,writer ,ad_glineup )

///     write water-slope-line-down

//                 storage2.File_Write ( w ,bucket ,filename2 ,writer ,ad_glinedown )
                 storage2.File_write ( w ,writer ,ad_glinedown )

          }

      } else if num == 0 {

//          io.WriteString(w, "\n pipe_line1_excute_all : data end \n")

         break

      }
   }

//    fmt.Fprintf(w, "\n pipe_line1_excute_all : Calculate succeeded.\n" )

	return
}







