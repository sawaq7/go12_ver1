package reserve_situation2

import (

	    "net/http"
//	    "fmt"
//	    "client/sgh"

        "client/reserve"
//        "client/reserve/type6"
//	    "strconv"
                                                  )

/// main 予約状況を表示する ///

func Reserve_situation2(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "reserve_situation2 start \n" )  // デバック

///
/// モニター　表示
///

   reserve.Reserve( w , r)

//	fmt.Fprintf( w, "reserve_situation2 : normal end \n" )  // デバック




}
