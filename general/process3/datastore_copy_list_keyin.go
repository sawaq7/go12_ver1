package process3

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/general/html5"

	    "github.com/sawaq7/go12_ver1/general/datastore5/trans3"
                                                )

func Datastore_copy_list_keyin(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "process.datastore_copy_list_keyin start \n" )  // 繝・ヰ繝・け

// 繝・Φ繝励Ξ繝ｼ繝医・繝倥ャ繝繝ｼ繧竪ET

     monitor := template.Must(template.New("html").Parse(html5.Datastore_copy_list_keyin))

// 繝・・繧ｿ繧ｹ繝医い繝ｼ縺九ｉ縲∬｡ｨ遉ｺ逕ｨ繝・・繧ｿ繧竪ET


     ds_copy_list_view := trans3.Ds_copy_list ( w ,r )

// 繝｢繝九ち繝ｼ縺ｫ陦ｨ遉ｺ

	err := monitor.Execute(w, ds_copy_list_view)
	if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
}

