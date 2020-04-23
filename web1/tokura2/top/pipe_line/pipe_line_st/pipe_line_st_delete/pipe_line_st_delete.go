package pipe_line_st_delete

import (

	    "net/http"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/process2"
//	    "github.com/sawaq7/go12_ver1/client/tokura/storage3/delete1"
	    "github.com/sawaq7/go12_ver1/client/tokura/storage3"
//	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	    "strconv"
//	    "fmt"

                                                  )


///
///     delete water-line inf. in storage
///

func Pipe_line_st_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "sky.pipe_line_st_delete start \n" )

    var idmy int64

	id := r.FormValue("id")

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//	fmt.Fprintf( w, "pipe_line_ds_delete : delid %v\n", delid )

///
///       delete water which was selected
///

    _ , _ = storage3.Storage_tokura( "Water2" ,"delete" ,delid , idmy , w , r  )

//    delete1.Water2( w , r ,delid )

///
///      show water inf. on web
///

   process2.Pipe_line_st_show(w , r )

}

