package deliver_update

import (

	    "strconv"

	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"

//	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
                                                   )

func Deliver_update(w http.ResponseWriter, r *http.Request) {

//	   fmt.Fprintf( w, "sky_deliver_update start %v\n" )  // ãEãE¯

/// æE®ãããã¼ã¿idãGET ///

    updidw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {

//	   fmt.Fprintf( w, "sky_deliver_update :error updidw %v\n", updidw )  // ãEãE¯

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

/// ééæE ±ã®å¤æ´ ///

	process.Deliver_update_single(w , r ,updid)

/// ã¢ãã¿ã¼ãåè¡¨ç¤º ///

	process.Deliver_showall1(w , r )

}
