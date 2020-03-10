package process4

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "client/reserve/html6"
//	    "client/reserve/type6"
	    "client/reserve/datastore6/trans5"
//	    "strconv"
                                                )


func Medical_xray_show(w http.ResponseWriter, r *http.Request ,guest_no int64) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN  guest_no     : ゲストNo

//    fmt.Fprintf( w, "medical_xray_show start \n" )  // デバック


// テンプレートのヘッダーをGET

    monitor := template.Must(template.New("html").Parse(html6.Medical_xray_show))

// データストアーから、表示用データをGET

//    d_area_view := datastore2.D_Area( "D_Area","trans" ,district_no , w , r  )

    guest_medical_xray_slice := trans5.Guest_medical_xray( guest_no , w , r  )

// 空インターフェイス変数よりバリュー値をゲット

//    value, _ := d_area_view.([]type2.D_Area)

// モニターに表示
//   fmt.Fprintf( w, "medical_xray_show d_area_view %v\n" ,d_area_view)  // デバック

	err := monitor.Execute(w, guest_medical_xray_slice)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

