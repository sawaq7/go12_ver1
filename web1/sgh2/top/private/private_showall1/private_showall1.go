package private_showall1

import (

	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"

//	    "time"
                                                  )

///
/// 　　   show private inf. on web
///

func Private_showall1(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "private_showall1 start \n" )

    //    show private inf.  on web
	process.Private_showall1(w , r )

//	fmt.Fprintf( w, "private_showall1 : normal end \n" )

}