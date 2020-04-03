package process4

import (
//
	    "net/http"
//	    "fmt"
	    "html/template"
//	    "github.com/sawaq7/go12_ver1/client/reserve/datastore6"
        "github.com/sawaq7/go12_ver1/client/reserve/datastore6/trans5"
	    "github.com/sawaq7/go12_ver1/client/reserve/html6"
//	    "github.com/sawaq7/go12_ver1/client/reserve/type6"
//	    "github.com/sawaq7/go12_ver1/general/type5"
//	    "time"
                                                )


func Guest_show(w http.ResponseWriter, r *http.Request) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ

//    fmt.Fprintf( w, "guest_show start \n" )  // チE��チE��

// チE��プレート�EヘッダーをGET

     monitor := template.Must(template.New("html").Parse(html6.Guest_show))

// チE�Eタストアーから、表示用チE�EタをGET

     reserve_minor := trans5.Reserve (w ,r )

//     general_work := make([]type5.General_Work, 2)
//     general_work[0].Int64_Work = 0          // 地区惁E��
//     general_work[1].Int64_Work = 0          //　コースNO

//     deliver_view := datastore2.Datastore_sgh( "D_District" ,"trans" ,general_work , w , r  )

     // 空インターフェイス変数よりバリュー値をゲチE��

//     value, _ := deliver_view.([]type2.D_District)


// モニターに表示

//    fmt.Fprintf( w, "trans.d_district reserve_minor %v\n" ,reserve_minor )  // チE��チE��

    err := monitor.Execute(w, reserve_minor)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

