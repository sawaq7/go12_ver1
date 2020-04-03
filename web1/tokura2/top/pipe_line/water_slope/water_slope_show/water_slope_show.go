package water_slope_show

import (

	    "net/http"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/process2"

                                                  )

///
/// 蟆取ｰｴ蜍ｾ驟咲ｷ壹Μ繧ｹ繝医ｒ陦ｨ遉ｺ縺吶ｋ縲・///

func Water_slope_show(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "water_slope_show start \n" )  // 繝・ヰ繝・け

/// 繝｢繝九ち繝ｼ縲陦ｨ遉ｺ ///

    process2.Water_slope_show(w , r )

//	fmt.Fprintf( w, "water_slope_show : normal end \n" )  // 繝・ヰ繝・け

}

