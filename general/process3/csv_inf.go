package process3

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
	    "html/template"

	    "github.com/sawaq7/go12_ver1/general/html5"
	    "github.com/sawaq7/go12_ver1/general/datastore5/trans3"
//	    "time"
                                                )


func Csv_inf(w http.ResponseWriter, r *http.Request ) {

//     IN    w      縲縲: 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r      縲縲: 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ
//     IN  district_no  : 蝨ｰ蝓櫻o

//    fmt.Fprintf( w, "csv_inf start \n" )  // 繝・ヰ繝・け}

// 繝・Φ繝励Ξ繝ｼ繝医・繝倥ャ繝繝ｼ繧竪ET

     monitor := template.Must(template.New("html").Parse(html5.Csv_show))


// 繝・・繧ｿ繧ｹ繝医い繝ｼ縺九ｉ縲…sv諠・ｱ繧偵ご繝・ヨ

//    d_area_view := datastore2.Datastore_sgh( "D_Area","trans" ,district_no , w , r  )

    csv_inf := trans3.Csv_inf ( w ,r )

// 繝｢繝九ち繝ｼ縺ｫ陦ｨ遉ｺ
//   fmt.Fprintf( w, "csv_inf csv_inf %v\n" ,csv_inf)  // 繝・ヰ繝・け

	err := monitor.Execute( w, csv_inf )

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

