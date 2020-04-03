package process

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/sgh/html2"
	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2"
//	    "time"
                                                )


func D_district_area(w http.ResponseWriter, r *http.Request ,district_no int64) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN  district_no  : 地域No

//    fmt.Fprintf( w, "d_district_area start \n" )  // チE��チE��}

// チE��プレート�EヘッダーをGET

     monitor := template.Must(template.New("html").Parse(html2.D_district_area))
//     monitor := template.Must(template.New("html").Parse(html2.D_district_area_type2))

// チE�Eタストアーから、表示用チE�EタをGET

    d_area_view := datastore2.Datastore_sgh( "D_Area","trans" ,district_no , w , r  )


// 空インターフェイス変数よりバリュー値をゲチE��

    value, _ := d_area_view.([]type2.D_Area)

// モニターに表示
//   fmt.Fprintf( w, "d_district_area d_area_view %v\n" ,d_area_view)  // チE��チE��

	err := monitor.Execute(w, value)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

