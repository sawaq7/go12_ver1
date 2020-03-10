package sky

import (

	    "net/http"
	    "client/sgh/process"
                                                  )
/// main 配達データを入力 ///

func init() {
	http.HandleFunc("/deliver_keyin_all", handler)
}



func handler(w http.ResponseWriter, r *http.Request) {

/// モニター表示 ///

   process.Deliver_keyin_all(w , r )

}

