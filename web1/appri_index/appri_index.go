package appri_index

import (

	    "net/http"
	    "github.com/sawaq7/go12_ver1/general/html5"
	    "html/template"
                     )


func Appri_index(w http.ResponseWriter, r *http.Request) {

   var cdmy string

/// get template

    monitor := template.Must(template.New("html").Parse(html5.Appri_index))

// モニターに表示

	err := monitor.Execute(w, cdmy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
