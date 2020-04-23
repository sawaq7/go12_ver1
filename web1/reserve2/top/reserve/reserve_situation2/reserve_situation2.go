package reserve_situation2

import (

	    "net/http"
//	    "fmt"
//	    "github.com/sawaq7/go12_ver1/client/sgh"

        "github.com/sawaq7/go12_ver1/client/reserve"
//        "github.com/sawaq7/go12_ver1/client/reserve/type6"
//	    "strconv"
                                                  )

///
///     show reserve situation2 inf.
///

func Reserve_situation2(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "reserve_situation2 start \n" )

///
///     show reserve situation2 inf.
///

   reserve.Reserve( w , r)

//	fmt.Fprintf( w, "reserve_situation2 : normal end \n" )

}
