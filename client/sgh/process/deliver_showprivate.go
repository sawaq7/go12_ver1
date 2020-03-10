package process

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
	    "fmt"
	    "html/template"
	    "client/sgh/html2"
	    "client/sgh/datastore2"

	    "client/sgh/type2"
	    "general/type5"

//	    "time"
                                                )

/// sub  car no ごとの　配達データを表示する　 ///

func Deliver_showprivate(w http.ResponseWriter, r *http.Request ,private_no int64) {


//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN  private_no   : プライベートNo

    fmt.Fprintf( w, "deliver_showprivate start \n" )  // デバック
    fmt.Fprintf( w, "deliver_showprivate : private_no %v\n", private_no )  // デバック

// テンプレートのヘッダーをGET

     monitor := template.Must(template.New("html").Parse(html2.Deliver_showprivate))

// データストアーから、表示用データをGET

     general_work := make([]type5.General_Work, 2)
     general_work[0].Int64_Work = 2          // 個人情報
     general_work[1].Int64_Work = private_no

//     deliver_view := datastore2.D_store( "Deliver" ,"trans" ,general_work , w , r  )
     deliver_view := datastore2.Datastore_sgh( "Deliver" ,"trans" ,general_work , w , r  )

     // 空インターフェイス変数よりバリュー値をゲット

     value, _ := deliver_view.([]type2.Deliver)

// モニターに表示

	err := monitor.Execute(w, value)

//	err := monitor.Execute(w, deliver_view)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

