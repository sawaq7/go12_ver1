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

/// main δΊη΄EΆζ³γθ‘¨η€Ίγγ ///

func Reserve_situation3(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "reserve_situation3 start \n" )  // γEγE―

    reserve_date  := r.FormValue("reserve_date")

///
///    γ’γγΏγΌγθ‘¨η€Ί
///

    reserve.Reserve2( w , r ,reserve_date)

//	fmt.Fprintf( w, "reserve_situation3 : normal end \n" )  // γEγE―

}
