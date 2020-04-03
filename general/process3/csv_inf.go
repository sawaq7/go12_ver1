package process3

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
	    "html/template"

	    "github.com/sawaq7/go12_ver1/general/html5"
	    "github.com/sawaq7/go12_ver1/general/datastore5/trans3"
//	    "time"
                                                )


func Csv_inf(w http.ResponseWriter, r *http.Request ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN  district_no  : 地域No

//    fmt.Fprintf( w, "csv_inf start \n" )  // チE��チE��}

// チE��プレート�EヘッダーをGET

     monitor := template.Must(template.New("html").Parse(html5.Csv_show))


// チE�Eタストアーから、csv惁E��をゲチE��

//    d_area_view := datastore2.Datastore_sgh( "D_Area","trans" ,district_no , w , r  )

    csv_inf := trans3.Csv_inf ( w ,r )

// モニターに表示
//   fmt.Fprintf( w, "csv_inf csv_inf %v\n" ,csv_inf)  // チE��チE��

	err := monitor.Execute( w, csv_inf )

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

