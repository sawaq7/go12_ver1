package process3

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/general/html5"

	    "github.com/sawaq7/go12_ver1/general/datastore5/trans3"
                                                )

func Datastore_copy_list_keyin(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "process.datastore_copy_list_keyin start \n" )

//    set template

     monitor := template.Must(template.New("html").Parse(html5.Datastore_copy_list_keyin))

//    get copy list

     ds_copy_list_view := trans3.Ds_copy_list ( w ,r )

//    set web

	err := monitor.Execute(w, ds_copy_list_view)
	if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
}

