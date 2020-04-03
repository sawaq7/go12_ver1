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

//	   fmt.Fprintf( w, "sky_csv_update start %v\n" )  // 繝・ヰ繝・け

/// 謖・ｮ壹＠縺溘ョ繝ｼ繧ｿid繧竪ET ///

    updidw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {

//	   fmt.Fprintf( w, "sky_csv_update :error updidw %v\n", updidw )  // 繝・ヰ繝・け

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

///
///    驟埼＃諠・ｱ縺ｮ螟画峩
///

	process3.Csv_update(w , r ,updid)

///
///    繝｢繝九ち繝ｼ縲蜀崎｡ｨ遉ｺ
///

	process3.Csv_inf ( w , r )

//	fmt.Fprintf( w, "csv_update normal end \n" )  // 繝・ヰ繝・け

}
