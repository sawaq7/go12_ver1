package pipe_line_st_keyin

import (

	    "net/http"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/process2"
//	    "fmt"

                                                  )

func Pipe_line_st_keyin(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "pipe_line_st_keyin start \n" )

///
///        show water-line inf. on web
///

   process2.Pipe_line_st_show(w , r )

}

