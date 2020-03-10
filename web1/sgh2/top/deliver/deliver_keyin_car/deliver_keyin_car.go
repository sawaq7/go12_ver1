package sky

import (

	    "net/http"
	    "client/sgh/process"
                                                  )

/// main 車両番号を入力する ///


func init() {
	http.HandleFunc("/deliver_keyin_car", handler)
}



func handler(w http.ResponseWriter, r *http.Request) {

/// モニター表示 ///

   process.Deliver_keyin_car(w , r )

}

