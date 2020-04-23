package guest_show

import (

	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/reserve/process4"

//	    "strconv"
//	    "time"
                                                  )
///
///     show guest inf. on web
///

func Guest_show(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "guest_show start \n" )

///
///     show guest inf. on web
///

   process4.Guest_show(w , r )

//	fmt.Fprintf( w, "guest_show : normal end \n" )

}
