package appri_index

import (

	    "net/http"
	    "general/html5"
	    "html/template"
                     )


func Appri_index(w http.ResponseWriter, r *http.Request) {

   var cdmy string

/// テンプレートのヘッダーをGET

    monitor := template.Must(template.New("html").Parse(html5.Appri_index))

// モニターに表示

	err := monitor.Execute(w, cdmy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
