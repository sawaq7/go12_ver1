package process

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/sgh/html2"
	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
	    "github.com/sawaq7/go12_ver1/general/type5"
	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2"
//	    "time"
                                                )


func D_district_keyin_all(w http.ResponseWriter, r *http.Request) {

//     IN    w      縲縲: 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r      縲縲: 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ

    fmt.Fprintf( w, "d_district_keyin_all start \n" )  // 繝・ヰ繝・け

// 繝・Φ繝励Ξ繝ｼ繝医・繝倥ャ繝繝ｼ繧竪ET

     monitor := template.Must(template.New("html").Parse(html2.D_district_keyin_all))

// 繝・・繧ｿ繧ｹ繝医い繝ｼ縺九ｉ縲∬｡ｨ遉ｺ逕ｨ繝・・繧ｿ繧竪ET

//     d_district_view := trans.D_district ( 0 ,0 ,w ,r )

     general_work := make([]type5.General_Work, 2)
     general_work[0].Int64_Work = 0          // 蝨ｰ蛹ｺ諠・ｱ
     general_work[1].Int64_Work = 0          //縲繧ｳ繝ｼ繧ｹNO

//     deliver_view := datastore2.D_store( "D_District" ,"trans" ,general_work , w , r  )
     deliver_view := datastore2.Datastore_sgh( "D_District" ,"trans" ,general_work , w , r  )

     // 遨ｺ繧､繝ｳ繧ｿ繝ｼ繝輔ぉ繧､繧ｹ螟画焚繧医ｊ繝舌Μ繝･繝ｼ蛟､繧偵ご繝・ヨ

     value, _ := deliver_view.([]type2.D_District)


// 繝｢繝九ち繝ｼ縺ｫ陦ｨ遉ｺ

    err := monitor.Execute(w, value)
//	err := monitor.Execute(w, d_district_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

