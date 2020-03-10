package pipe_line_ds_update

import (

	    "strconv"

	    "net/http"
//	    "fmt"
	    "client/tokura/suiri/process2"


                                                   )

func Pipe_line_ds_update(w http.ResponseWriter, r *http.Request) {

//	fmt.Fprintf( w, "sky_pipe_line_ds_update start %v\n" )  // デバック

///
///      指定したデータidをGET
///

    updidw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {

//	   fmt.Fprintf( w, "sky_pipe_line_ds_update :error updidw %v\n", updidw )  // デバック

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

///
///       配達情報の変更
///

	process2.Pipe_line_ds_update(w , r ,updid)

///
///        モニター　再表示
///

	process2.Pipe_line_ds_show(w , r )

}
