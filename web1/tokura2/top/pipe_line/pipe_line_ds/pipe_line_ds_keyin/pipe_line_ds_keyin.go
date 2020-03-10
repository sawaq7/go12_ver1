package pipe_line_ds_keyin

import (

	    "net/http"
	    "client/tokura/suiri/process2"
//	    "fmt"

                                                  )

func Pipe_line_ds_keyin(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "pipe_line_ds_keyin start \n" )  // デバック


/// モニター表示 ///

   process2.Pipe_line_ds_show(w , r )

}

