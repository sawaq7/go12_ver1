package process2

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/tokura/html4"

	    "github.com/sawaq7/go12_ver1/client/tokura/datastore4"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
//	    "time"
                                                )
func Pipe_line_ds_show(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "process2.pipe_line_ds_show start \n" )  // 繝・ヰ繝・け

    var idmy int64

///
///      繝・Φ繝励Ξ繝ｼ繝医・繝倥ャ繝繝ｼ繧竪ET
///

     monitor := template.Must(template.New("html").Parse(html4.Pipe_line_ds_keyin))

///
///   繝・・繧ｿ繧ｹ繝医い繝ｼ縺九ｉ縲∬｡ｨ遉ｺ逕ｨ繝・・繧ｿ繧竪ET
///


//     water2_view := trans2.Water2 ( w ,r )

     water2_view := datastore4.Datastore_tokura( "Water2"  ,"trans"  ,idmy , w , r  )

     value, _ := water2_view.([]type4.Water2)    // 遨ｺ繧､繝ｳ繧ｿ繝ｼ繝輔ぉ繧､繧ｹ螟画焚繧医ｊ繝舌Μ繝･繝ｼ蛟､繧偵ご繝・ヨ


///
///             繝｢繝九ち繝ｼ縺ｫ陦ｨ遉ｺ
///

	err := monitor.Execute(w, value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

