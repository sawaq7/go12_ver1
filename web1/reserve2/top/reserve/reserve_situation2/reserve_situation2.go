package reserve_situation2

import (

	    "net/http"
//	    "fmt"
//	    "github.com/sawaq7/go12_ver1/client/sgh"

        "github.com/sawaq7/go12_ver1/client/reserve"
//        "github.com/sawaq7/go12_ver1/client/reserve/type6"
//	    "strconv"
                                                  )

/// main 莠育ｴ・憾豕√ｒ陦ｨ遉ｺ縺吶ｋ ///

func Reserve_situation2(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "reserve_situation2 start \n" )  // 繝・ヰ繝・け

///
/// 繝｢繝九ち繝ｼ縲陦ｨ遉ｺ
///

   reserve.Reserve( w , r)

//	fmt.Fprintf( w, "reserve_situation2 : normal end \n" )  // 繝・ヰ繝・け




}
