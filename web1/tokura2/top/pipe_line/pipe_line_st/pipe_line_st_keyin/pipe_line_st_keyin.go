package pipe_line_st_keyin

import (

	    "net/http"
	    "client/tokura/suiri/process2"
//	    "fmt"

                                                  )

func Pipe_line_st_keyin(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "pipe_line_st_keyin start \n" )  // デバック

///
///        モニター表示
///

   process2.Pipe_line_st_show(w , r )

}

