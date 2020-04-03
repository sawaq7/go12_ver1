package process

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/sgh/html2"
	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2"

                                                )
///
/// 縲縲謖・ｮ壹＠縺溷慍蝓櫻O.縺ｮ蜿ｷ霆頑ュ蝣ｱ繧定｡ｨ遉ｺ縺吶ｋ
///

func Car_show(w http.ResponseWriter, r *http.Request ,district_no int64) {

//     IN    w      縲縲: 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r      縲縲: 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ
//     IN  district_no  : 蝨ｰ蝓櫻o

//    fmt.Fprintf( w, "car_show start \n" )  // 繝・ヰ繝・け}

// 繝・Φ繝励Ξ繝ｼ繝医・繝倥ャ繝繝ｼ繧竪ET

     monitor := template.Must(template.New("html").Parse(html2.Car_show))

// 繝・・繧ｿ繧ｹ繝医い繝ｼ縺九ｉ縲∬｡ｨ遉ｺ逕ｨ繝・・繧ｿ繧竪ET

//     car_view := datastore2.D_store( "Car" ,"trans"  ,district_no , w , r  )
     car_view := datastore2.Datastore_sgh( "Car" ,"trans"  ,district_no , w , r  )

     // 遨ｺ繧､繝ｳ繧ｿ繝ｼ繝輔ぉ繧､繧ｹ螟画焚繧医ｊ繝舌Μ繝･繝ｼ蛟､繧偵ご繝・ヨ

     value, _ := car_view.([]type2.Car)


// 繝｢繝九ち繝ｼ縺ｫ陦ｨ遉ｺ

	err := monitor.Execute(w, value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

