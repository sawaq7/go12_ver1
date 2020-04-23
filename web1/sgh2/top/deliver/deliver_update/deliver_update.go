package deliver_update

import (

	    "strconv"

	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"

//	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
                                                            )

///
/// 　　   upfdate delivery inf.
///

func Deliver_update(w http.ResponseWriter, r *http.Request) {

//	   fmt.Fprintf( w, "sky_deliver_update start %v\n" )

    updidw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {

//	   fmt.Fprintf( w, "sky_deliver_update :error updidw %v\n", updidw )

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

    //    update delivery inf. on web
	process.Deliver_update_single(w , r ,updid)

    //    show delivery inf. on web
	process.Deliver_showall1(w , r )

}
