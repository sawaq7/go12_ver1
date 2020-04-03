package process3

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/general/html5"

	    "github.com/sawaq7/go12_ver1/general/datastore5/trans3"
                                                )

func Db_access_list(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "process3.db_access_list start \n" )  // 繝・ヰ繝・け

///
///    繝・Φ繝励Ξ繝ｼ繝医・繝倥ャ繝繝ｼ繧竪ET
///

     monitor := template.Must(template.New("html").Parse(html5.Db_access_list))

///
///     繝・・繧ｿ繧ｹ繝医い繝ｼ縺九ｉ縲∬｡ｨ遉ｺ逕ｨ繝・・繧ｿ繧竪ET
///

     db_access_list := trans3.Db_access_list ( w ,r )

///
///       繝｢繝九ち繝ｼ縺ｫ陦ｨ遉ｺ
///

	err := monitor.Execute(w, db_access_list)
	if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
}

