package hydrostatic_pressure1_show

import (
	     "net/http"
	     "html/template"
	     "client/tokura/suiri/type4"

         "client/tokura/html4"

	                                  )

func Hydrostatic_pressure1_show(w http.ResponseWriter, r *http.Request) {

   var seisui type4.Seisui

// テンプレートのヘッダーをGET

	monitor := template.Must( template.New("html").Parse( html4.Hydrostatic_pressure1_show) )

// 各種入力データを表示

	err := monitor.Execute(w, seisui)

    if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
