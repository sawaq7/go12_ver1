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

/// main 予紁E��況を表示する ///

func Reserve_situation(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "reserve_situation start \n" )  // チE��チE��

    reserve_date  := r.FormValue("reserve_date")

/// モニター　表示 ///

	process4.Reserve_situation(w , r ,reserve_date)

//	fmt.Fprintf( w, "reserve_situation : normal end \n" )  // チE��チE��




}
