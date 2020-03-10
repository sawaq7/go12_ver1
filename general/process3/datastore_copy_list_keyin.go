package process3

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "general/html5"

	    "general/datastore5/trans3"
                                                )

func Datastore_copy_list_keyin(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "process.datastore_copy_list_keyin start \n" )  // デバック

// テンプレートのヘッダーをGET

     monitor := template.Must(template.New("html").Parse(html5.Datastore_copy_list_keyin))

// データストアーから、表示用データをGET


     ds_copy_list_view := trans3.Ds_copy_list ( w ,r )

// モニターに表示

	err := monitor.Execute(w, ds_copy_list_view)
	if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
}

