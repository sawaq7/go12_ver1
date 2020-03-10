package process

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "client/sgh/html2"
	    "client/sgh/type2"
	    "client/sgh/datastore2"

                                                )
///
/// 　　指定した地域NO.の号車情報を表示する
///

func Car_show(w http.ResponseWriter, r *http.Request ,district_no int64) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN  district_no  : 地域No

//    fmt.Fprintf( w, "car_show start \n" )  // デバック}

// テンプレートのヘッダーをGET

     monitor := template.Must(template.New("html").Parse(html2.Car_show))

// データストアーから、表示用データをGET

//     car_view := datastore2.D_store( "Car" ,"trans"  ,district_no , w , r  )
     car_view := datastore2.Datastore_sgh( "Car" ,"trans"  ,district_no , w , r  )

     // 空インターフェイス変数よりバリュー値をゲット

     value, _ := car_view.([]type2.Car)


// モニターに表示

	err := monitor.Execute(w, value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

