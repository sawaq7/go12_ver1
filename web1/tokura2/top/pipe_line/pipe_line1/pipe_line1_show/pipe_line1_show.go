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
	     "storage2"
	     "html/template"
	     "github.com/sawaq7/go12_ver1/client/tokura/html4"
	                               )

///  main process ///

func Pipe_line1_show(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "Pipe_line1_show start \n" )  // ãEãE¯

    water := make([]type4.Water,100 )
//    water := make([]type4.Water,4 )

// æ°´è·¯æE ±ã«ã¦ã³ã¿ã¼ãinitialize

    pos := 0

// ãã±ãEåãEãã¡ã¤ã«åãã»ãE

    bucket := "sample-7777"
    filename1 := "water_inf.txt"

// æ°´è·¯æE ±ãã¡ã¤ã«ãEEead fileEEãªã¼ãã³

//	reader  := storage2.File_Open(w ,r ,bucket ,filename1)

	reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename1 , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)  // ã¤ã³ã¿ã¼ãã§ã¤ã¹åãåå¤æ

    defer reader.Close()

// ãã¡ã¤ã«ãªã¼ãã¼(stringç¨EãE§E¥E´

    sreader := bufio.NewReaderSize(reader, 4096)

// å°æ°´å¾éç·ãã¡ã¤ã«ããªã¼ãã³

   index := 0        // ã¬ã³ã¼ãã«ã¦ã³ã¿ã¼ãinitialize


   for {
      index ++     // ã¬ã³ã¼ãã«ã¦ã³ã¿ã¼ãã«ã¦ã³ãE
//      fmt.Fprintf(w, "pipe_line1_show : lndex %v\n", index )  // ãEãE¯

// ãã¡ã¤ã«ãï¼è¡read

      line ,_  := sreader.ReadString('\n')

      str := strings.Fields(line)

      num := len(str)

//      fmt.Fprintf(w, "pipe_line1_show : num %v\n", num )  // ãEãE¯

      if num != 0 {
         if index == 1 {

// ãããã¼ã¯ãã¹ã«ã¼ãã

//             fmt.Fprintf(w, "pipe_line1_show (header) : line %s\n", line )  // ãEãE¯

          }else{

/// æ°´è·¯ãããã¼æE ±ããGET
             pos ++     // æ°´è·¯æE ±ã«ã¦ã³ã¿ã¼ãã«ã¦ã³ãE
//             fmt.Fprintf(w, "pipe_line1_show (the other): line %s\n", line )  // ãEãE¯

             water[pos-1].No = strconv.Itoa(index) //ãæ´æ°ãæå­ã«å¤æ
             water[pos-1].Name ,water[pos-1].High ,water[pos-1].Roughness_factor = suiri.Kansui1_2( line  )

//             fmt.Fprintf(w, "pipe_line1_show : æ°´è·¯ãã³ããE %v\n", water[pos-1].No )  // ãEãE¯
//             fmt.Fprintf(w, "pipe_line1_show : æ°´è·¯åE%s\n", water[pos-1] .Name)  // ãEãE¯
//             fmt.Fprintf(w, "pipe_line1_show : æ°´è·¯é«E%s\n", water[pos-1].High )  // ãEãE¯
//             fmt.Fprintf(w, "pipe_line1_show : ç²åº¦ä¿æ° %s\n", water[pos-1].Roughness_factor )  // ãEãE¯

          }

      } else if num == 0 {                                         // End check

//          io.WriteString(w, "\n pipe_line1_show : data end \n")   //ãEãE¯

         break

      }
   }
//   fmt.Fprintf(w, "pipe_line1_show : len(water) cap(water) %v\n", len(water)  ,cap(water))  // ãEãE¯

// ã¹ã©ã¤ã¹ãå§ç¸®

   water2 := make([]type4.Water, pos )
   copy ( water2 ,water[0:pos] ) // æ³¨Eãã¼ã¿ã¯ãEããããposEEãã¾ã§

/// æ°´è·¯æE ±ãè¡¨ç¤º

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
