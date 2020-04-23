package sky

import (

	    "net/http"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"
                                                            )

///
/// 　　   key-in delivery inf.
///

func init() {
	http.HandleFunc("/deliver_keyin_all", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {

///
///     show deliver inf. which was keied-in on web
///

   process.Deliver_keyin_all(w , r )

}

