package sky

import (

	    "net/http"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"
                                                  )
/// main 驟埼＃繝・・繧ｿ繧貞・蜉・///

func init() {
	http.HandleFunc("/deliver_keyin_all", handler)
}



func handler(w http.ResponseWriter, r *http.Request) {

/// 繝｢繝九ち繝ｼ陦ｨ遉ｺ ///

   process.Deliver_keyin_all(w , r )

}

