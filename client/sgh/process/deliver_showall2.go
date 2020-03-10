package process

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "client/sgh/html2"
	    "client/sgh/type2"
	    "general/type5"
//	    "client/sgh/datastore2/trans"

        "client/sgh/datastore2"
//	    "time"
                                                )


func Deliver_showall2( course_no int64 ,w http.ResponseWriter, r *http.Request ) {

//     IN  course_no  : コースNo
//     IN    w      　: レスポンスライター
//     IN    r      　: リクエストパラメータ

//   fmt.Fprintf( w, "process.deliver_show_all2 start \n" )  // デバック

///
///     テンプレートのヘッダーをGET
///

     monitor := template.Must(template.New("html").Parse(html2.Deliver_showall1))

///
///    データストアーから、表示用データをGET
///


     general_work := make([]type5.General_Work, 2)
     general_work[0].Int64_Work = 0          // 地区情報
     general_work[1].Int64_Work = course_no  //　コースNO

//     deliver_view := datastore2.D_store( "Deliver"  ,"trans"  ,general_work , w , r  )
     deliver_view := datastore2.Datastore_sgh( "Deliver"  ,"trans"  ,general_work , w , r  )

     // 空インターフェイス変数よりバリュー値をゲット

     value, _ := deliver_view.([]type2.Deliver)

///
///      モニターに表示
///

    err := monitor.Execute(w, value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

