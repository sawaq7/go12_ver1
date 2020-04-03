package process2

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/tokura/html4"

	    "github.com/sawaq7/go12_ver1/client/tokura/datastore4"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
                                                )

///
///       蟆取ｰｴ蜍ｾ驟咲ｷ壹Μ繧ｹ繝医ｒ陦ｨ遉ｺ縺吶ｋ縲・///

func Water_slope_show(w http.ResponseWriter, r *http.Request) {

//     IN     w         : 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN     r         : 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ繝ｼ

//    fmt.Fprintf( w, "process2.water_slope_show start \n" )  // 繝・ヰ繝・け

    var idmy int64

///
///     繝・Φ繝励Ξ繝ｼ繝医・繝倥ャ繝繝ｼ繧竪ET
///

     monitor := template.Must(template.New("html").Parse(html4.Water_slope_show))

///
///      繝・・繧ｿ繧ｹ繝医い繝ｼ縺九ｉ縲∬｡ｨ遉ｺ逕ｨ繝・・繧ｿ繧竪ET
///

//     water_slope_view := trans2.Water_slope ( w ,r )

     water_slope_view := datastore4.Datastore_tokura( "Water_Slope"  ,"trans"  ,idmy , w , r  )

     value, _ := water_slope_view.([]type4.Water_Slope)    // 遨ｺ繧､繝ｳ繧ｿ繝ｼ繝輔ぉ繧､繧ｹ螟画焚繧医ｊ繝舌Μ繝･繝ｼ蛟､繧偵ご繝・ヨ

///
///       繝｢繝九ち繝ｼ縺ｫ陦ｨ遉ｺ
///

	err := monitor.Execute( w, value )
	if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
}

