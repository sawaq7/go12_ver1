package sky

import (

	    "net/http"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"
                                                            )

///
/// 　　   key-in car inf.
///

func init() {
	http.HandleFunc("/deliver_keyin_car", handler)
}



func handler(w http.ResponseWriter, r *http.Request) {

///
///     show car inf. which was keied-in on web
///

   process.Deliver_keyin_car(w , r )

}

