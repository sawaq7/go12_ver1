package process4

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/reserve/html6"
//	    "github.com/sawaq7/go12_ver1/client/reserve/type6"
	    "github.com/sawaq7/go12_ver1/client/reserve/datastore6/trans5"

                                                )


func Medical_record_show(w http.ResponseWriter, r *http.Request ,guest_no int64) {

//     IN    w      縲縲: 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r      縲縲: 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ
//     IN  guest_no     : 繧ｲ繧ｹ繝・o

//    fmt.Fprintf( w, "medical_record_show start \n" )  // 繝・ヰ繝・け}

// 繝・Φ繝励Ξ繝ｼ繝医・繝倥ャ繝繝ｼ繧竪ET

    monitor := template.Must(template.New("html").Parse(html6.Medical_record_show))

    guest_medical_record_slice := trans5.Guest_medical_record( guest_no , w , r  )

	err := monitor.Execute(w, guest_medical_record_slice)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
