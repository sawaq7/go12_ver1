package pipe_line_ds_update

import (

	    "strconv"

	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/process2"


                                                   )

func Pipe_line_ds_update(w http.ResponseWriter, r *http.Request) {

//	fmt.Fprintf( w, "sky_pipe_line_ds_update start %v\n" )  // チE��チE��

///
///      持E��したデータidをGET
///

    updidw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {

//	   fmt.Fprintf( w, "sky_pipe_line_ds_update :error updidw %v\n", updidw )  // チE��チE��

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

///
///       配達惁E��の変更
///

	process2.Pipe_line_ds_update(w , r ,updid)

///
///        モニター　再表示
///

	process2.Pipe_line_ds_show(w , r )

}
