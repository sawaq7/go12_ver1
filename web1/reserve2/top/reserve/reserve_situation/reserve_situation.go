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

/// main δΊη΄EΆζ³γθ‘¨η€Ίγγ ///

func Reserve_situation(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "reserve_situation start \n" )  // γEγE―

    reserve_date  := r.FormValue("reserve_date")

/// γ’γγΏγΌγθ‘¨η€Ί ///

	process4.Reserve_situation(w , r ,reserve_date)

//	fmt.Fprintf( w, "reserve_situation : normal end \n" )  // γEγE―




}
