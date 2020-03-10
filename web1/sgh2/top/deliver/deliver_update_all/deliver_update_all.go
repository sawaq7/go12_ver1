package sky

import (

//	    "strconv"
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
	    "client/sgh/process"

//	    "client/sgh/type2"
                                                   )

func init() {
	http.HandleFunc("/deliver_update_all", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {

//	   fmt.Fprintf( w, "sky_deliver_update_all start %v\n" )  // デバック

/// 配達情報の変更 ///

	process.Deliver_update_all(w , r )

/// モニター　再表示 ///

	process.Deliver_showall1(w , r )

//	http.Redirect(w, r, "/", http.StatusFound)


}
