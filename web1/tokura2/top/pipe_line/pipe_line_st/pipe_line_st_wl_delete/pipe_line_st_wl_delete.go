package pipe_line_st_wl_delete

import (

	"net/http"
	"strconv"
//	"fmt"
	"client/tokura/suiri/process2"
	"client/tokura/suiri/type4"
//	"client/tokura/storage3/trans4"
	"client/tokura/storage3"

                                            )

func Pipe_line_st_wl_delete(w http.ResponseWriter, r *http.Request) {

//   fmt.Fprintf( w, "pipe_line_st_wl_delete start \n" )  // デバック

    var idmy1 ,idmy2 int64

///
///       カレントの水路情報をゲット
///

      water2_temp_minor , _ := storage3.Storage_tokura( "Water2_Temp" ,"trans" ,idmy1 , idmy2 , w , r  )

      water2_temp, _ := water2_temp_minor.([]type4.Water2_Temp)  // インターフェイス型を型変換

//    water2_temp := trans4.Water2_temp( w , r  )

///
///       指示したidをGET
///

    id := r.FormValue("id")
//    fmt.Fprintf( w, "pipe_line_st_wl_delete : id %v\n", id )  // デバック

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//    fmt.Fprintf( w, "pipe_line_st_wl_delete : delidw %v\n", delidw )  // デバック
//    fmt.Fprintf( w, "pipe_line_st_wl_delete : delid %v\n", delid )  // デバック

///
///     指示した水路ラインデータを削除
///

      _ , _ = storage3.Storage_tokura( "Water_Line" ,"delete" ,delid , water2_temp[0].Name , w , r  )

//    delete1.Water_line( w , r ,delid , water2_temp[0].Name )

///
///        モニター　再表示
///

	process2.Pipe_line_st_wl_show ( water2_temp[0].Name ,w , r )

}
