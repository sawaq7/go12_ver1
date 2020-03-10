package process4

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "client/reserve/html6"
//	    "client/reserve/type6"
	    "client/reserve/datastore6/trans5"

                                                )


func Medical_record_show(w http.ResponseWriter, r *http.Request ,guest_no int64) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN  guest_no     : ゲストNo

//    fmt.Fprintf( w, "medical_record_show start \n" )  // デバック}

// テンプレートのヘッダーをGET

    monitor := template.Must(template.New("html").Parse(html6.Medical_record_show))

    guest_medical_record_slice := trans5.Guest_medical_record( guest_no , w , r  )

	err := monitor.Execute(w, guest_medical_record_slice)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
