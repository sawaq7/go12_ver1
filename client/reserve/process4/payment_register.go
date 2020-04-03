package process4

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/reserve/html6"
//	    "github.com/sawaq7/go12_ver1/client/reserve/type6"
	    "github.com/sawaq7/go12_ver1/client/reserve/datastore6/trans5"
                                                )


func Payment_register(w http.ResponseWriter, r *http.Request ,guest_no int64) {

//     IN    w      縲縲: 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r      縲縲: 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ
//     IN  guest_no     : 繧ｲ繧ｹ繝・o

//    fmt.Fprintf( w, "payment_register start \n" )  // 繝・ヰ繝・け}

// 繝・Φ繝励Ξ繝ｼ繝医・繝倥ャ繝繝ｼ繧竪ET

    monitor := template.Must(template.New("html").Parse(html6.Payment_register))

// 繝・・繧ｿ繧ｹ繝医い繝ｼ縺九ｉ縲∬｡ｨ遉ｺ逕ｨ繝・・繧ｿ繧竪ET

//    d_area_view := datastore2.Datastore_sgh( "D_Area","trans" ,district_no , w , r  )

    guest_payment_slice := trans5.Guest_payment ( guest_no , w , r  )

// 遨ｺ繧､繝ｳ繧ｿ繝ｼ繝輔ぉ繧､繧ｹ螟画焚繧医ｊ繝舌Μ繝･繝ｼ蛟､繧偵ご繝・ヨ

//    value, _ := d_area_view.([]type2.D_Area)

// 繝｢繝九ち繝ｼ縺ｫ陦ｨ遉ｺ
//   fmt.Fprintf( w, "payment_register d_area_view %v\n" ,d_area_view)  // 繝・ヰ繝・け

//	err := monitor.Execute(w, value)

	err := monitor.Execute(w, guest_payment_slice)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

