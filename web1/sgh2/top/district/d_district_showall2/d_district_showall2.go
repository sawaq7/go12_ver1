package d_district_showall2

import (

	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"

//	    "strconv"
//	    "time"
                                                  )

/// main ιιε°εγEγEEγΏγθ‘¨η€Ίγγ ///

func D_district_showall2(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "d_district_showall2 start \n" )  // γEγE―

/// γ’γγΏγΌγθ‘¨η€Ί ///

   process.D_district_showall1(w , r )

//	fmt.Fprintf( w, "d_district_showall2 : normal end \n" )  // γEγE―




}
