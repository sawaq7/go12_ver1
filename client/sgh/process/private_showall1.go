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

//     IN    w      縲縲: 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r      縲縲: 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ

//    fmt.Fprintf( w, "process.deliver1_show_all1 start \n" )  // 繝・ヰ繝・け

// 繝・Φ繝励Ξ繝ｼ繝医・繝倥ャ繝繝ｼ繧竪ET

    monitor := template.Must(template.New("html").Parse(html2.Private_showall1))

// 繝・・繧ｿ繧ｹ繝医い繝ｼ縺九ｉ縲∬｡ｨ遉ｺ逕ｨ繝・・繧ｿ繧竪ET

    private_view := datastore2.Datastore_sgh( "Private" ,"trans" ,nil ,w , r  )

    // 遨ｺ繧､繝ｳ繧ｿ繝ｼ繝輔ぉ繧､繧ｹ螟画焚繧医ｊ繝舌Μ繝･繝ｼ蛟､繧偵ご繝・ヨ

    value, _ := private_view.([]type2.Private)

// 繝｢繝九ち繝ｼ縺ｫ陦ｨ遉ｺ

	err := monitor.Execute(w, value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

