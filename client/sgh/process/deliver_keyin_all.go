package process

import (

	    "net/http"
	    "fmt"
	    "html/template"
	    "client/sgh/html2"
	    "client/sgh/datastore2"
	    "client/sgh/type2"
	    "general/type5"

                                                )
func Deliver_keyin_all(w http.ResponseWriter, r *http.Request) {

//     IN    w      　: レスポンスライター
//     IN    r      　: リクエストパラメータ

    fmt.Fprintf( w, "process.deliver_keyin_all start \n" )  // デバック

// テンプレートのヘッダーをGET

     monitor := template.Must(template.New("html").Parse(html2.Deliver_keyin_all))

// データストアーから、表示用データをGET

     general_work := make([]type5.General_Work, 2)
     general_work[0].Int64_Work = 0          // 地区情報
     general_work[1].Int64_Work = 0          //　コースNO

     deliver_view := datastore2.Datastore_sgh( "Deliver" ,"trans" ,general_work , w , r  )

     // 空インターフェイス変数よりバリュー値をゲット

     value, _ := deliver_view.([]type2.Deliver)

// モニターに表示

	err := monitor.Execute(w, value)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

