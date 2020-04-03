package tokura_index

import (

	    "net/http"
	    "github.com/sawaq7/go12_ver1/client/tokura/html4"
	    "html/template"
                                                  )

///
/// show the menu of tokuras' applications
///

func Tokura_index(w http.ResponseWriter, r *http.Request) {

   var cdmy string

/// チE��プレート�EヘッダーをGET

//    monitor := template.Must(template.New("html").Parse(html4.Tokura_index))
    monitor := template.Must(template.New("html").Parse(html4.Tokura_index2))

// モニターに表示

	err := monitor.Execute(w, cdmy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

