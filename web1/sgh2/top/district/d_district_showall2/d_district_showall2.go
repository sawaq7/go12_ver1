package d_district_showall2

import (

	    "net/http"
//	    "fmt"
	    "client/sgh/process"

//	    "strconv"
//	    "time"
                                                  )

/// main 配達地域のデータを表示する ///

func D_district_showall2(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "d_district_showall2 start \n" )  // デバック

/// モニター　表示 ///

   process.D_district_showall1(w , r )

//	fmt.Fprintf( w, "d_district_showall2 : normal end \n" )  // デバック




}
