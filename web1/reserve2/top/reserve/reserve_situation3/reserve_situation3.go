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

/// main 予紁E��況を表示する ///

func Reserve_situation3(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "reserve_situation3 start \n" )  // チE��チE��

    reserve_date  := r.FormValue("reserve_date")

///
///    モニター　表示
///

    reserve.Reserve2( w , r ,reserve_date)

//	fmt.Fprintf( w, "reserve_situation3 : normal end \n" )  // チE��チE��

}
