package d_district_showall2

import (

	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"

//	    "strconv"
//	    "time"
                                                  )

///
/// 　　   show district inf.
///

func D_district_showall2(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "d_district_showall2 start \n" )

   //    show district inf.  on web
   process.D_district_showall1(w , r )

//	fmt.Fprintf( w, "d_district_showall2 : normal end \n" )




}
