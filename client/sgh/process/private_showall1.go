package process

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "client/sgh/html2"
	    "client/sgh/type2"
//	    "client/sgh/datastore2/trans"
	    "client/sgh/datastore2"

                                                )


func Private_showall1(w http.ResponseWriter, r *http.Request) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ

//    fmt.Fprintf( w, "process.deliver1_show_all1 start \n" )  // デバック

// テンプレートのヘッダーをGET

    monitor := template.Must(template.New("html").Parse(html2.Private_showall1))

// データストアーから、表示用データをGET

    private_view := datastore2.Datastore_sgh( "Private" ,"trans" ,nil ,w , r  )

    // 空インターフェイス変数よりバリュー値をゲット

    value, _ := private_view.([]type2.Private)

// モニターに表示

	err := monitor.Execute(w, value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

