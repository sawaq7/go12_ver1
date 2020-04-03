package tokura_index_pl

import (

	    "net/http"
	    "github.com/sawaq7/go12_ver1/client/tokura/html4"
	    "html/template"
                                                  )

///
/// show the menu of tokuras' applications
///


func Tokura_index_pl(w http.ResponseWriter, r *http.Request) {

   var cdmy string

/// チE��プレート�EヘッダーをGET

    monitor := template.Must(template.New("html").Parse(html4.Tokura_index_pl))

// モニターに表示

	err := monitor.Execute(w, cdmy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

