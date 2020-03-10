package tokura_index_hp

import (

	    "net/http"
	    "client/tokura/html4"
	    "html/template"
                                                  )

///
/// show the menu of tokuras' applications
///



func Tokura_index_hp(w http.ResponseWriter, r *http.Request) {

   var cdmy string

/// テンプレートのヘッダーをGET

    monitor := template.Must(template.New("html").Parse(html4.Tokura_index_hp))

// モニターに表示

	err := monitor.Execute(w, cdmy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

