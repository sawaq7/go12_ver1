package private_showall1

import (

	    "net/http"
//	    "fmt"
	    "client/sgh/process"

//	    "time"
                                                  )

func Private_showall1(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "private_showall1 start \n" )  // デバック

/// モニター　再表示 ///

	process.Private_showall1(w , r )

//	fmt.Fprintf( w, "private_showall1 : normal end \n" )  // デバック




}
