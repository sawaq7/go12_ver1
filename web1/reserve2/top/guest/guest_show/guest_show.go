package guest_show

import (

	    "net/http"
//	    "fmt"
	    "client/reserve/process4"

//	    "strconv"
//	    "time"
                                                  )

/// main 顧客情報を表示する ///

func Guest_show(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "guest_show start \n" )  // デバック

/// モニター　表示 ///

   process4.Guest_show(w , r )

//	fmt.Fprintf( w, "guest_show : normal end \n" )  // デバック




}
