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

//	   fmt.Fprintf( w, "sky_csv_update start %v\n" )  // ãEãE¯

/// æE®ãããã¼ã¿idãGET ///

    updidw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {

//	   fmt.Fprintf( w, "sky_csv_update :error updidw %v\n", updidw )  // ãEãE¯

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

///
///    ééæE ±ã®å¤æ´
///

	process3.Csv_update(w , r ,updid)

///
///    ã¢ãã¿ã¼ãåè¡¨ç¤º
///

	process3.Csv_inf ( w , r )

//	fmt.Fprintf( w, "csv_update normal end \n" )  // ãEãE¯

}
