package reserve_situation3

import (

	    "net/http"
//	    "fmt"
//	    "client/sgh"

        "client/reserve"
//        "client/reserve/type6"
//	    "strconv"

//	    "time"
                                                  )

/// main 予約状況を表示する ///

func Reserve_situation3(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "reserve_situation3 start \n" )  // デバック

    reserve_date  := r.FormValue("reserve_date")

///
///    モニター　表示
///

    reserve.Reserve2( w , r ,reserve_date)

//	fmt.Fprintf( w, "reserve_situation3 : normal end \n" )  // デバック

}
