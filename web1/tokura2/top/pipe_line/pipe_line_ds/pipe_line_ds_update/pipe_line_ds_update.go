package pipe_line_ds_update

import (

	    "strconv"

	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/process2"


                                                   )

func Pipe_line_ds_update(w http.ResponseWriter, r *http.Request) {

//	fmt.Fprintf( w, "sky_pipe_line_ds_update start %v\n" )  // ãEãE¯

///
///      æE®ãããã¼ã¿idãGET
///

    updidw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {

//	   fmt.Fprintf( w, "sky_pipe_line_ds_update :error updidw %v\n", updidw )  // ãEãE¯

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

///
///       ééæE ±ã®å¤æ´
///

	process2.Pipe_line_ds_update(w , r ,updid)

///
///        ã¢ãã¿ã¼ãåè¡¨ç¤º
///

	process2.Pipe_line_ds_show(w , r )

}
