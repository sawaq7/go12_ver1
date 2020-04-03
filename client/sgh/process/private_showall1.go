package process

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/sgh/html2"
	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
//	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2/trans"
	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2"

                                                )


func Private_showall1(w http.ResponseWriter, r *http.Request) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ

//    fmt.Fprintf( w, "process.deliver1_show_all1 start \n" )  // チE��チE��

// チE��プレート�EヘッダーをGET

    monitor := template.Must(template.New("html").Parse(html2.Private_showall1))

// チE�Eタストアーから、表示用チE�EタをGET

    private_view := datastore2.Datastore_sgh( "Private" ,"trans" ,nil ,w , r  )

    // 空インターフェイス変数よりバリュー値をゲチE��

    value, _ := private_view.([]type2.Private)

// モニターに表示

	err := monitor.Execute(w, value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

