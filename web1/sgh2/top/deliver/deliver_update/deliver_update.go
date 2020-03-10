package deliver_update

import (

	    "strconv"

	    "net/http"
//	    "fmt"
	    "client/sgh/process"

//	    "client/sgh/type2"
                                                   )

func Deliver_update(w http.ResponseWriter, r *http.Request) {

//	   fmt.Fprintf( w, "sky_deliver_update start %v\n" )  // デバック

/// 指定したデータidをGET ///

    updidw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {

//	   fmt.Fprintf( w, "sky_deliver_update :error updidw %v\n", updidw )  // デバック

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

/// 配達情報の変更 ///

	process.Deliver_update_single(w , r ,updid)

/// モニター　再表示 ///

	process.Deliver_showall1(w , r )

}
