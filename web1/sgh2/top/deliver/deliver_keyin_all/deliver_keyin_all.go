package sky

import (

	    "net/http"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"
                                                  )
/// main 配達チE�Eタを�E劁E///

func init() {
	http.HandleFunc("/deliver_keyin_all", handler)
}



func handler(w http.ResponseWriter, r *http.Request) {

/// モニター表示 ///

   process.Deliver_keyin_all(w , r )

}

