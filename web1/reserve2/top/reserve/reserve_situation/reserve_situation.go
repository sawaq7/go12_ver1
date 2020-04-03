package reserve_situation

import (
//
	    "net/http"
//	    "fmt"
//	    "github.com/sawaq7/go12_ver1/client/sgh"

        "github.com/sawaq7/go12_ver1/client/reserve/process4"
//        "github.com/sawaq7/go12_ver1/client/reserve/type6"
//	    "strconv"

                                                  )

/// main 莠育ｴ・憾豕√ｒ陦ｨ遉ｺ縺吶ｋ ///

func Reserve_situation(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "reserve_situation start \n" )  // 繝・ヰ繝・け

    reserve_date  := r.FormValue("reserve_date")

/// 繝｢繝九ち繝ｼ縲陦ｨ遉ｺ ///

	process4.Reserve_situation(w , r ,reserve_date)

//	fmt.Fprintf( w, "reserve_situation : normal end \n" )  // 繝・ヰ繝・け




}
