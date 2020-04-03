package sky

import (

//	    "strconv"
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"

//	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
                                                   )

func init() {
	http.HandleFunc("/deliver_update_all", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {

//	   fmt.Fprintf( w, "sky_deliver_update_all start %v\n" )  // チE��チE��

/// 配達惁E��の変更 ///

	process.Deliver_update_all(w , r )

/// モニター　再表示 ///

	process.Deliver_showall1(w , r )

//	http.Redirect(w, r, "/", http.StatusFound)


}
