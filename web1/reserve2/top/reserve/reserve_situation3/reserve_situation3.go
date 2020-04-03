package reserve_situation3

import (

	    "net/http"
//	    "fmt"
//	    "github.com/sawaq7/go12_ver1/client/sgh"

        "github.com/sawaq7/go12_ver1/client/reserve"
//        "github.com/sawaq7/go12_ver1/client/reserve/type6"
//	    "strconv"

//	    "time"
                                                  )

/// main 莠育ｴ・憾豕√ｒ陦ｨ遉ｺ縺吶ｋ ///

func Reserve_situation3(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "reserve_situation3 start \n" )  // 繝・ヰ繝・け

    reserve_date  := r.FormValue("reserve_date")

///
///    繝｢繝九ち繝ｼ縲陦ｨ遉ｺ
///

    reserve.Reserve2( w , r ,reserve_date)

//	fmt.Fprintf( w, "reserve_situation3 : normal end \n" )  // 繝・ヰ繝・け

}
