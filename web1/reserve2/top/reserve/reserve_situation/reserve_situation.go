package reserve_situation

import (
//
	    "net/http"
//	    "fmt"
//	    "github.com/sawaq7/go12_ver1/client/sgh"

        "github.com/sawaq7/go12_ver1/client/reserve/process4"
//        "github.com/sawaq7/go12_ver1/client/reserve/type6"
//	    "strconv"

                                                  )

///
///     show reserve situation inf.
///

func Reserve_situation(w http.ResponseWriter, r *http.Request) {

//     IN    w      : response-writer
//     IN    r      : request- paramete

//    fmt.Fprintf( w, "reserve_situation start \n" )

    reserve_date  := r.FormValue("reserve_date")

///
///     show reserve situation inf.
///

	process4.Reserve_situation(w , r ,reserve_date)

//	fmt.Fprintf( w, "reserve_situation : normal end \n" )




}
