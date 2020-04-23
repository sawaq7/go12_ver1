package pipe_line_ds_update

import (

	    "strconv"

	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/process2"


                                                   )

func Pipe_line_ds_update(w http.ResponseWriter, r *http.Request) {

//	fmt.Fprintf( w, "sky_pipe_line_ds_update start %v\n" )

    updidw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {

//	   fmt.Fprintf( w, "sky_pipe_line_ds_update :error updidw %v\n", updidw )

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

///
///       update water-line inf.
///

	process2.Pipe_line_ds_update(w , r ,updid)

///
///        show water-line inf.
///

	process2.Pipe_line_ds_show(w , r )

}
