package csv_update

import (

	    "strconv"
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
//        "general/type5"
        "general/process3"

                                                   )

func Csv_update(w http.ResponseWriter, r *http.Request) {

//	   fmt.Fprintf( w, "sky_csv_update start %v\n" )  // デバック

/// 指定したデータidをGET ///

    updidw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {

//	   fmt.Fprintf( w, "sky_csv_update :error updidw %v\n", updidw )  // デバック

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

///
///    配達情報の変更
///

	process3.Csv_update(w , r ,updid)

///
///    モニター　再表示
///

	process3.Csv_inf ( w , r )

//	fmt.Fprintf( w, "csv_update normal end \n" )  // デバック

}
