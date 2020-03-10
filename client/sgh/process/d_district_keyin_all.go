package process

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
	    "fmt"
	    "html/template"
	    "client/sgh/html2"
	    "client/sgh/type2"
	    "general/type5"
	    "client/sgh/datastore2"
//	    "time"
                                                )


func D_district_keyin_all(w http.ResponseWriter, r *http.Request) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ

    fmt.Fprintf( w, "d_district_keyin_all start \n" )  // デバック

// テンプレートのヘッダーをGET

     monitor := template.Must(template.New("html").Parse(html2.D_district_keyin_all))

// データストアーから、表示用データをGET

//     d_district_view := trans.D_district ( 0 ,0 ,w ,r )

     general_work := make([]type5.General_Work, 2)
     general_work[0].Int64_Work = 0          // 地区情報
     general_work[1].Int64_Work = 0          //　コースNO

//     deliver_view := datastore2.D_store( "D_District" ,"trans" ,general_work , w , r  )
     deliver_view := datastore2.Datastore_sgh( "D_District" ,"trans" ,general_work , w , r  )

     // 空インターフェイス変数よりバリュー値をゲット

     value, _ := deliver_view.([]type2.D_District)


// モニターに表示

    err := monitor.Execute(w, value)
//	err := monitor.Execute(w, d_district_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

