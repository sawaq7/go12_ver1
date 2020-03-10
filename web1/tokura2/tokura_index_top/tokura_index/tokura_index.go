package tokura_index

import (

	    "net/http"
	    "client/tokura/html4"
	    "html/template"
                                                  )

///
/// show the menu of tokuras' applications
///

func Tokura_index(w http.ResponseWriter, r *http.Request) {

   var cdmy string

/// テンプレートのヘッダーをGET

//    monitor := template.Must(template.New("html").Parse(html4.Tokura_index))
    monitor := template.Must(template.New("html").Parse(html4.Tokura_index2))

// モニターに表示

	err := monitor.Execute(w, cdmy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

