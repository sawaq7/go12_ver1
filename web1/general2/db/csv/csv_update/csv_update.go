package csv_update

import (

	    "strconv"

	    "net/http"
//	    "fmt"
//        "github.com/sawaq7/go12_ver1/general/type5"
        "github.com/sawaq7/go12_ver1/general/process3"

                                                   )

func Csv_update(w http.ResponseWriter, r *http.Request) {

//	   fmt.Fprintf( w, "sky_csv_update start %v\n" )

                              ///    get the id which is selected for updating
    updidw , err := strconv.Atoi(r.FormValue("id"))

	if err  != nil {
//	   fmt.Fprintf( w, "sky_csv_update :error updidw %v\n", updidw )

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

///
///       change csv inf.
///

	process3.Csv_update(w , r ,updid)

///
///      show csv inf. on web
///

	process3.Csv_inf ( w , r )

//	fmt.Fprintf( w, "csv_update normal end \n" )

}
