package d_district_showall2_sample

import (

	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/temp/process1000"

//	    "strconv"
//	    "time"
                             )

func D_district_showall2_sample(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "d_district_showall2_sample start \n" )  // γEγE―

/// γ’γγΏγΌγθ‘¨η€Ί ///

   process1000.D_district_showall1_sample(w , r )

//	fmt.Fprintf( w, "d_district_showall2_sample : normal end \n" )  // γEγE―

}
