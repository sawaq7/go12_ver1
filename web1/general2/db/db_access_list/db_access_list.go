package db_access_list

import (

	    "net/http"
	    "github.com/sawaq7/go12_ver1/general/process3"
                                                  )

///
///        show d.b. access list inf. on web
///

func Db_access_list(w http.ResponseWriter, r *http.Request) {

///
///     show d.b. access list inf. on web
///

    process3.Db_access_list(w , r )

}
