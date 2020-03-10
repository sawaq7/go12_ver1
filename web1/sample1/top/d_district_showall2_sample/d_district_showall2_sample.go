package d_district_showall2_sample

import (

	    "net/http"
//	    "fmt"
	    "temp/process1000"

//	    "strconv"
//	    "time"
                             )

func D_district_showall2_sample(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "d_district_showall2_sample start \n" )  // デバック

/// モニター　表示 ///

   process1000.D_district_showall1_sample(w , r )

//	fmt.Fprintf( w, "d_district_showall2_sample : normal end \n" )  // デバック

}
