package process3

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
	    "html/template"

	    "general/html5"
	    "general/datastore5/trans3"
//	    "time"
                                                )


func Csv_inf(w http.ResponseWriter, r *http.Request ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN  district_no  : 地域No

//    fmt.Fprintf( w, "csv_inf start \n" )  // デバック}

// テンプレートのヘッダーをGET

     monitor := template.Must(template.New("html").Parse(html5.Csv_show))


// データストアーから、csv情報をゲット

//    d_area_view := datastore2.Datastore_sgh( "D_Area","trans" ,district_no , w , r  )

    csv_inf := trans3.Csv_inf ( w ,r )

// モニターに表示
//   fmt.Fprintf( w, "csv_inf csv_inf %v\n" ,csv_inf)  // デバック

	err := monitor.Execute( w, csv_inf )

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

