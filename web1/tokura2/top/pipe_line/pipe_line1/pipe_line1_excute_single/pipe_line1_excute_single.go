////////////////////////////////////////////////////////////////////////
///                                                                  ///
///   make water-slope-line　　　　　　　　　　　　　　　　　　　  ///
///                                                                  ///
////////////////////////////////////////////////////////////////////////

package pipe_line1_excute_single

import (
//	      "fmt"
	      "io"
	      "strings"
	      "net/http"
	      "bufio"
	      "github.com/sawaq7/go12_ver1/client/tokura/suiri"
	      "github.com/sawaq7/go12_ver1/storage2"
	      "strconv"
	      "cloud.google.com/go/storage"

                                         )



///  main process ///

func Pipe_line1_excute_single(w http.ResponseWriter, r *http.Request) {

    water_id , err := strconv.Atoi(r.FormValue("water_id"))
//    fmt.Fprintf( w, "pipe_line1_excute_single  water_id %v\n", water_id )

	if err  != nil {

//	   fmt.Fprintf( w, "pipe_line1_excute_single :error water_id"  )

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}

    bucket := "sample-7777"
    filename1 := "water_inf.txt"
    filename2 := "grade_line.txt"

///     create water-slope-line

    writer_minor , _ := storage2.Storage_basic( "create" ,bucket ,filename2 , w , r  )

    writer, _ := writer_minor.(*storage.Writer)

//    writer := storage2.File_Create( w ,r ,bucket ,filename2 )

	defer writer.Close()

///    open water-inf-file

//    reader := storage2.File_Open(w ,r ,bucket ,filename1)

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename1 , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)

    defer reader.Close()

//    get file-reader

    sreader := bufio.NewReaderSize(reader, 4096)

    index := 0

    for {

      index ++

//      fmt.Fprintf(w, "pipe_line1_excute_single : lndex %v\n", index )

///   read one-record

      line ,_  := sreader.ReadString('\n')

      // change water data from string-type to string-array-type by spliting brank
      str := strings.Fields(line)

      num := len(str)

//      fmt.Fprintf(w, "pipe_line1_excute_single : num %v\n", num )

      if num != 0 {
         if index == 1 || index != water_id {

         //   skip header

//             fmt.Fprintf(w, "pipe_line1_excute_single (header) : line %s\n", line )

          }else{
//             fmt.Fprintf(w, "pipe_line1_excute_single (the other): line %s\n", line )

///    make water-slope-line

             ad_hp ,ad_hl ,ad_vhead ,ad_eneup ,ad_enedown ,ad_glineup ,ad_glinedown := suiri.Kansui1( line  )

//             fmt.Fprintf(w, "pipe_line1_excute_single : ad_hp %s\n", ad_hp )
//             fmt.Fprintf(w, "pipe_line1_excute_single : ad_hl %s\n", ad_hl )
//             fmt.Fprintf(w, "pipe_line1_excute_single : ad_vhead %s\n", ad_vhead )
//             fmt.Fprintf(w, "pipe_line1_excute_single : ad_eneup %s\n", ad_eneup )
//             fmt.Fprintf(w, "pipe_line1_excute_single : ad_enedown %s\n", ad_enedown )
//             fmt.Fprintf(w, "pipe_line1_excute_single : ad_glineup %s\n", ad_glineup )
//             fmt.Fprintf(w, "pipe_line1_excute_single : ad_glinedown %s\n", ad_glinedown )

///     write point-loss

//               storage2.File_Write ( w ,bucket ,filename2 ,writer ,ad_hp )
               storage2.File_write ( w ,writer ,ad_hp )

///    write line-loss

//               storage2.File_Write ( w ,bucket ,filename2 ,writer ,ad_hl )
               storage2.File_write ( w ,writer ,ad_hl )

///     write vilocity-head

//               storage2.File_Write ( w ,bucket ,filename2 ,writer ,ad_vhead )
               storage2.File_write ( w ,writer ,ad_vhead )

///     write enegy-line-up

//               storage2.File_Write ( w ,bucket ,filename2 ,writer ,ad_eneup )
               storage2.File_write ( w ,writer ,ad_eneup )

///     write enegy-line-down

//               storage2.File_Write ( w ,bucket ,filename2 ,writer ,ad_enedown )
               storage2.File_write ( w ,writer ,ad_enedown )

///     write water-slope-line-up

//               storage2.File_Write ( w ,bucket ,filename2 ,writer ,ad_glineup )
               storage2.File_write ( w ,writer ,ad_glineup )

///     write water-slope-line-down

//               storage2.File_Write ( w ,bucket ,filename2 ,writer ,ad_glinedown )
               storage2.File_write ( w ,writer ,ad_glinedown )

          }

      } else if num == 0 {

//          io.WriteString(w, "\n pipe_line1_excute_single : data end \n")

         break

      }
   }

// end process

//	fmt.Fprintf(w, "\n pipe_line1_excute_single : Calculate succeeded.\n" )

	return
}







