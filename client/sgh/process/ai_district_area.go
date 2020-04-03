package process

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/sgh/html2"
	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2"
//	    "time"
                                                )


func Ai_district_area(w http.ResponseWriter, r *http.Request ,district_no int64) {

//     IN  district_no  : 地域No
//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ

    fmt.Fprintf( w, "ai_district_area start \n" )  // チE��チE��}

// チE��プレート�EヘッダーをGET

     monitor := template.Must(template.New("html").Parse(html2.Ai_district_area))

// チE�Eタストアーから、表示用チE�EタをGET

//     d_area_view := datastore2.D_store( "D_Area","trans" ,district_no , w , r  )
     d_area_view := datastore2.Datastore_sgh( "D_Area","trans" ,district_no , w , r  )

// 空インターフェイス変数よりバリュー値をゲチE��

    value, _ := d_area_view.([]type2.D_Area)

// モニターに表示

	err := monitor.Execute(w, value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

