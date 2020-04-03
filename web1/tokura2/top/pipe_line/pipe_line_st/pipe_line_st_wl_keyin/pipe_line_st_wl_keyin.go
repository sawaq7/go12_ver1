package pipe_line_st_wl_keyin

import (

	    "net/http"
	    "strconv"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/process2"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	    "storage2"
	    "bufio"

	    "io"
	    "github.com/sawaq7/go12_ver1/client/tokura/storage3"

//	    "fmt"

                                                  )

func Pipe_line_st_wl_keyin(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "sky.pipe_line_st_wl_keyin start \n" )  // γEγE―

    var water2_temp type4.Water2_Temp

    var idmy int64

    bucket := "sample-7777"

    filename1 := "Water2.txt"

///
///       γ»γ¬γ―γγγγ¬γ³γΌγidγγ²γE
///

    select_idw , err := strconv.Atoi(r.FormValue("id"))

	if err  != nil {

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    select_id := int64(select_idw)

//    fmt.Fprintf( w, "sky.pipe_line_st_wl_keyin : select_idw %v\n", select_idw )  // γEγE―
//    fmt.Fprintf( w, "sky.pipe_line_st_wl_keyin : select_id %v\n", select_id )  // γEγE―

///
///     ζ°΄θ·―γγ‘γ€γ«γE γͺγΌγγ³
///

//    reader := storage2.File_Open(w ,r ,bucket ,filename1)

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename1 , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)  // γ€γ³γΏγΌγγ§γ€γΉεγεε€ζ

    defer reader.Close()

    sreader := bufio.NewReaderSize(reader, 4096)

    index := 0

    for {

//      fmt.Fprintf(w, "sky.pipe_line_st_wl_keyin : lndex %v\n", index )  // γEγE―

      line ,_  := sreader.ReadString('\n')   // γγ‘γ€γ«γοΌθ‘read

      num := len(line)

//      fmt.Fprintf(w, "sky.pipe_line_st_wl_keyin : num %v\n", num )  // γEγE―

      if num > 1 {

         index ++     // γ¬γ³γΌγγ«γ¦γ³γΏγΌγγ«γ¦γ³γE
//        fmt.Fprintf(w, "sky.pipe_line_st_wl_keyin : line %s\n", line )  // γEγE―

         water2_struct_minor , _ := storage3.Storage_tokura( "Water2" ,"struct_set" ,line , idmy , w , r  )

//         water2_struct := struct_set.Water2( w , line )   //  stringεγEγEEγΏγζ§ι δ½εγ«ε€ζ

         water2_struct, _ := water2_struct_minor.(type4.Water2)  // γ€γ³γΏγΌγγ§γ€γΉεγεε€ζ

         if  water2_struct.Id == select_id {

           water2_temp.Id   = int64( 1 )
           water2_temp.Name = water2_struct.Name
           water2_temp.High = water2_struct.High
           water2_temp.Roughness_Factor = water2_struct.Roughness_Factor

//           fmt.Fprintf( w, "sky.pipe_line_st_wl_keyin : water2_temp %v\n", water2_temp )  // γEγE―

           break

         }
      } else if num == 0 {

//          io.WriteString(w, "\n sky.pipe_line_st_wl_keyin : can't find data \n")   //γEγE―

         break

      }
   }

///
///         temporary-fileγγ€γγ·γ£γ©γ€γΊ  & γγ₯γΌγEEγΏγγ»γE
///

      _ , _ = storage3.Storage_tokura( "Water2_Temp" ,"initialize" ,water2_temp , idmy , w , r  )

//    initialize3.Water2_temp (w , r ,water2_temp) // temporary-fileγγ€γγ·γ£γ©γ€γΊ

///
/// γγγγγγγ’γγΏγΌθ‘¨η€Ί
///

   process2.Pipe_line_st_wl_show ( water2_temp.Name ,w , r )

}

