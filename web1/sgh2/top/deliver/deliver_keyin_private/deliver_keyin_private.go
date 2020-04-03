package sky

import (

	    "net/http"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"
                                                  )

/// main 個人番号を�E力すめE///


func init() {
	http.HandleFunc("/deliver_keyin_private", handler)
}



func handler(w http.ResponseWriter, r *http.Request) {

/// モニター表示 ///

//   deliver.Deliver1_show(w , r )
   process.Deliver_keyin_private(w , r )

}

