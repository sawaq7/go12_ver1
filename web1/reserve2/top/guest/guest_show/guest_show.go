package guest_show

import (

	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/reserve/process4"

//	    "strconv"
//	    "time"
                                                  )

/// main 顧客惁E��を表示する ///

func Guest_show(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "guest_show start \n" )  // チE��チE��

/// モニター　表示 ///

   process4.Guest_show(w , r )

//	fmt.Fprintf( w, "guest_show : normal end \n" )  // チE��チE��




}
