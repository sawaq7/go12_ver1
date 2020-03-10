package pipe_line_st_delete

import (

	    "net/http"
	    "client/tokura/suiri/process2"
//	    "client/tokura/storage3/delete1"
	    "client/tokura/storage3"
//	    "client/tokura/suiri/type4"
	    "strconv"
//	    "fmt"

                                                  )

func Pipe_line_st_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "sky.pipe_line_st_delete start \n" )  // デバック

    var idmy int64

///
///       指示したidをGET
///

	id := r.FormValue("id")

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//	fmt.Fprintf( w, "pipe_line_ds_delete : delid %v\n", delid )  // デバック

///
///     指示した水路データを削除
///

    _ , _ = storage3.Storage_tokura( "Water2" ,"delete" ,delid , idmy , w , r  )

//    delete1.Water2( w , r ,delid )

///
///      モニター表示
///

   process2.Pipe_line_st_show(w , r )

}

