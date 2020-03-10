package process

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "client/sgh/datastore2/trans"
	    "client/sgh/html2"
//	    "client/sgh/type2"
//	    "general/type5"
//	    "time"
                                                )


func D_district_showall1(w http.ResponseWriter, r *http.Request) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ

//    fmt.Fprintf( w, "d_district_show_all1 start \n" )  // デバック

// テンプレートのヘッダーをGET

     monitor := template.Must(template.New("html").Parse(html2.D_district_showall1))

// データストアーから、表示用データをGET

     d_district_view := trans.D_district2 ( w ,r )

//     general_work := make([]type5.General_Work, 2)
//     general_work[0].Int64_Work = 0          // 地区情報
//     general_work[1].Int64_Work = 0          //　コースNO

//     deliver_view := datastore2.Datastore_sgh( "D_District" ,"trans" ,general_work , w , r  )

     // 空インターフェイス変数よりバリュー値をゲット

//     value, _ := deliver_view.([]type2.D_District)


// モニターに表示

	err := monitor.Execute(w, d_district_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

