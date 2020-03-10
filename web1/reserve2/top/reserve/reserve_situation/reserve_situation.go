package reserve_situation

import (
//
	    "net/http"
//	    "fmt"
//	    "client/sgh"

        "client/reserve/process4"
//        "client/reserve/type6"
//	    "strconv"

                                                  )

/// main 予約状況を表示する ///

func Reserve_situation(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "reserve_situation start \n" )  // デバック

    reserve_date  := r.FormValue("reserve_date")

/// モニター　表示 ///

	process4.Reserve_situation(w , r ,reserve_date)

//	fmt.Fprintf( w, "reserve_situation : normal end \n" )  // デバック




}
