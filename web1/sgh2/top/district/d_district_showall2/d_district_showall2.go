package d_district_showall2

import (

	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"

//	    "strconv"
//	    "time"
                                                  )

/// main 驟埼＃蝨ｰ蝓溘・繝・・繧ｿ繧定｡ｨ遉ｺ縺吶ｋ ///

func D_district_showall2(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "d_district_showall2 start \n" )  // 繝・ヰ繝・け

/// 繝｢繝九ち繝ｼ縲陦ｨ遉ｺ ///

   process.D_district_showall1(w , r )

//	fmt.Fprintf( w, "d_district_showall2 : normal end \n" )  // 繝・ヰ繝・け




}
