package reserve_situation2

import (

	    "net/http"
//	    "fmt"
//	    "github.com/sawaq7/go12_ver1/client/sgh"

        "github.com/sawaq7/go12_ver1/client/reserve"
//        "github.com/sawaq7/go12_ver1/client/reserve/type6"
//	    "strconv"
                                                  )

/// main δΊη΄EΆζ³γθ‘¨η€Ίγγ ///

func Reserve_situation2(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "reserve_situation2 start \n" )  // γEγE―

///
/// γ’γγΏγΌγθ‘¨η€Ί
///

   reserve.Reserve( w , r)

//	fmt.Fprintf( w, "reserve_situation2 : normal end \n" )  // γEγE―




}
