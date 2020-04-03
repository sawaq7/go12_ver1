package sky

import (

	    "net/http"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"
                                                  )

/// main 車両番号を�E力すめE///


func init() {
	http.HandleFunc("/deliver_keyin_car", handler)
}



func handler(w http.ResponseWriter, r *http.Request) {

/// モニター表示 ///

   process.Deliver_keyin_car(w , r )

}

