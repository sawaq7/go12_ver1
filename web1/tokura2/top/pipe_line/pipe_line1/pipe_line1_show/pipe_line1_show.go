package pipe_line1_show

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
	     "html/template"
	     "github.com/sawaq7/go12_ver1/client/tokura/html4"
	                               )

///  main process ///

func Pipe_line1_show(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "Pipe_line1_show start \n" )

    water := make([]type4.Water,100 )

    pos := 0

    bucket := "sample-7777"
    filename1 := "water_inf.txt"

///    open water-inf-file

//	reader  := storage2.File_Open(w ,r ,bucket ,filename1)

	reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename1 , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)

    defer reader.Close()

//    get file-reader

    sreader := bufio.NewReaderSize(reader, 4096)

   index := 0


   for {
      index ++
//      fmt.Fprintf(w, "pipe_line1_show : lndex %v\n", index )

///   read one-record

      line ,_  := sreader.ReadString('\n')

      str := strings.Fields(line)

      num := len(str)

//      fmt.Fprintf(w, "pipe_line1_show : num %v\n", num )

      if num != 0 {
         if index == 1 {

         //   skip header

//             fmt.Fprintf(w, "pipe_line1_show (header) : line %s\n", line )

          }else{

          //   get water-header
             pos ++
//             fmt.Fprintf(w, "pipe_line1_show (the other): line %s\n", line )

             water[pos-1].No = strconv.Itoa(index)
             water[pos-1].Name ,water[pos-1].High ,water[pos-1].Roughness_factor = suiri.Kansui1_2( line  )

//             fmt.Fprintf(w, "pipe_line1_show : 水路ナンバ�E %v\n", water[pos-1].No )
//             fmt.Fprintf(w, "pipe_line1_show : 水路吁E%s\n", water[pos-1] .Name)
//             fmt.Fprintf(w, "pipe_line1_show : 水路髁E%s\n", water[pos-1].High )
//             fmt.Fprintf(w, "pipe_line1_show : 粗度係数 %s\n", water[pos-1].Roughness_factor )

          }

      } else if num == 0 {                                         // End check

//          io.WriteString(w, "\n pipe_line1_show : data end \n")

         break

      }
   }
//   fmt.Fprintf(w, "pipe_line1_show : len(water) cap(water) %v\n", len(water)  ,cap(water))

//    compress slice

   water2 := make([]type4.Water, pos )
   copy ( water2 ,water[0:pos] ) // 注�E�データは、E　から　pos�E�E　まで

///   show water-line inf.

   monitor := template.Must(template.New("html").Parse( html4.Pipe_line1_show))
    err := monitor.Execute(w, water2)
    if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	}

//   suiri.Pipe_line1_show( w ,pos , water2 )

// end process

//    fmt.Fprintf(w, "\n pipe_line1_show : Calculate succeeded.\n" )

    return

}
