package process

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "client/sgh/html2"
	    "client/sgh/datastore2"
	    "client/sgh/type2"
                                                )


func D_district_area_show(w http.ResponseWriter, r *http.Request ,district_no int64) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　district_no : 地区No

//    fmt.Fprintf( w, "d_district_area_show start \n" )  // デバック}

// テンプレートのヘッダーをGET

//     monitor := template.Must(template.New("html").Parse(html2.D_district_area_show))
     monitor := template.Must(template.New("html").Parse(html2.D_district_area))

// データストアーから、表示用データをGET

     d_area_view := datastore2.Datastore_sgh( "D_Area","trans" ,district_no , w , r  )

// 空インターフェイス変数よりバリュー値をゲット

    value, _ := d_area_view.([]type2.D_Area)

// モニターに表示

	err := monitor.Execute(w, value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

