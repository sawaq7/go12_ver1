package reserve_situation3

import (

	    "net/http"
//	    "fmt"
//	    "github.com/sawaq7/go12_ver1/client/sgh"

        "github.com/sawaq7/go12_ver1/client/reserve"
//        "github.com/sawaq7/go12_ver1/client/reserve/type6"
//	    "strconv"

//	    "time"
                                                  )

///
///     show reserve situation3 inf.
///

func Reserve_situation3(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "reserve_situation3 start \n" )

    reserve_date  := r.FormValue("reserve_date")

///
///     show reserve situation3 inf.
///

    reserve.Reserve2( w , r ,reserve_date)

//	fmt.Fprintf( w, "reserve_situation3 : normal end \n" )

}
