package process4

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
	    "html/template"
	    "client/reserve/html6"
//	    "client/reserve/type6"
	    "client/reserve/datastore6/trans5"
//	    "time"
                                                )


func Reserve_situation(w http.ResponseWriter, r *http.Request ,reserve_date string) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN  reserve_date : 予約日

//    fmt.Fprintf( w, "reserve_situation start \n" )  // デバック}

///
///       テンプレートのヘッダーをSET
///

    monitor := template.Must(template.New("html").Parse(html6.Reserve_situation))

///
///      データストアーから、表示用データをGET
///

//    d_area_view := datastore2.Datastore_sgh( "D_Area","trans" ,district_no , w , r  )
    guest_reserve_minor_slice := trans5.Guest_reserve_minor2( reserve_date , w , r  )

// 空インターフェイス変数よりバリュー値をゲット

//    value, _ := d_area_view.([]type2.D_Area)

///
///         モニターに表示
///

//   fmt.Fprintf( w, "reserve_situation d_area_view %v\n" ,d_area_view)  // デバック

	err := monitor.Execute(w, guest_reserve_minor_slice)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

