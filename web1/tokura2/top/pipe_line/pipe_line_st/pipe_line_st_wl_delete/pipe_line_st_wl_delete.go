package pipe_line_st_wl_delete

import (

	"net/http"
	"strconv"
//	"fmt"
	"github.com/sawaq7/go12_ver1/client/tokura/suiri/process2"
	"github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
//	"github.com/sawaq7/go12_ver1/client/tokura/storage3/trans4"
	"github.com/sawaq7/go12_ver1/client/tokura/storage3"

                                            )

func Pipe_line_st_wl_delete(w http.ResponseWriter, r *http.Request) {

//   fmt.Fprintf( w, "pipe_line_st_wl_delete start \n" )  // γEγE―

    var idmy1 ,idmy2 int64

///
///       γ«γ¬γ³γγEζ°΄θ·―ζE ±γγ²γE
///

      water2_temp_minor , _ := storage3.Storage_tokura( "Water2_Temp" ,"trans" ,idmy1 , idmy2 , w , r  )

      water2_temp, _ := water2_temp_minor.([]type4.Water2_Temp)  // γ€γ³γΏγΌγγ§γ€γΉεγεε€ζ

//    water2_temp := trans4.Water2_temp( w , r  )

///
///       ζE€ΊγγidγGET
///

    id := r.FormValue("id")
//    fmt.Fprintf( w, "pipe_line_st_wl_delete : id %v\n", id )  // γEγE―

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//    fmt.Fprintf( w, "pipe_line_st_wl_delete : delidw %v\n", delidw )  // γEγE―
//    fmt.Fprintf( w, "pipe_line_st_wl_delete : delid %v\n", delid )  // γEγE―

///
///     ζE€Ίγγζ°΄θ·―γ©γ€γ³γEEγΏγει€
///

      _ , _ = storage3.Storage_tokura( "Water_Line" ,"delete" ,delid , water2_temp[0].Name , w , r  )

//    delete1.Water_line( w , r ,delid , water2_temp[0].Name )

///
///        γ’γγΏγΌγεθ‘¨η€Ί
///

	process2.Pipe_line_st_wl_show ( water2_temp[0].Name ,w , r )

}
