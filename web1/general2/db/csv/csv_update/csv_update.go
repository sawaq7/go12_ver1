package csv_update

import (

	    "strconv"
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
//        "github.com/sawaq7/go12_ver1/general/type5"
        "github.com/sawaq7/go12_ver1/general/process3"

                                                   )

func Csv_update(w http.ResponseWriter, r *http.Request) {

//	   fmt.Fprintf( w, "sky_csv_update start %v\n" )  // チE��チE��

/// 持E��したデータidをGET ///

    updidw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {

//	   fmt.Fprintf( w, "sky_csv_update :error updidw %v\n", updidw )  // チE��チE��

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

///
///    配達惁E��の変更
///

	process3.Csv_update(w , r ,updid)

///
///    モニター　再表示
///

	process3.Csv_inf ( w , r )

//	fmt.Fprintf( w, "csv_update normal end \n" )  // チE��チE��

}
