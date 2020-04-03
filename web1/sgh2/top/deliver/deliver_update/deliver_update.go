package deliver_update

import (

	    "strconv"

	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"

//	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
                                                   )

func Deliver_update(w http.ResponseWriter, r *http.Request) {

//	   fmt.Fprintf( w, "sky_deliver_update start %v\n" )  // 繝・ヰ繝・け

/// 謖・ｮ壹＠縺溘ョ繝ｼ繧ｿid繧竪ET ///

    updidw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {

//	   fmt.Fprintf( w, "sky_deliver_update :error updidw %v\n", updidw )  // 繝・ヰ繝・け

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

/// 驟埼＃諠・ｱ縺ｮ螟画峩 ///

	process.Deliver_update_single(w , r ,updid)

/// 繝｢繝九ち繝ｼ縲蜀崎｡ｨ遉ｺ ///

	process.Deliver_showall1(w , r )

}
