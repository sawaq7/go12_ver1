package d_district_showall2

import (

	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"

//	    "strconv"
//	    "time"
                                                  )

/// main 配達地域�EチE�Eタを表示する ///

func D_district_showall2(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "d_district_showall2 start \n" )  // チE��チE��

/// モニター　表示 ///

   process.D_district_showall1(w , r )

//	fmt.Fprintf( w, "d_district_showall2 : normal end \n" )  // チE��チE��




}
