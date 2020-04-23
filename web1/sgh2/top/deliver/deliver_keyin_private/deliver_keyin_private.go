package sky

import (

	    "net/http"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"
                                                           )

///
/// 　　   key-in private inf.
///

func init() {
	http.HandleFunc("/deliver_keyin_private", handler)
}



func handler(w http.ResponseWriter, r *http.Request) {

///
///     show private inf. which was keied-in on web
///

   process.Deliver_keyin_private(w , r )

}