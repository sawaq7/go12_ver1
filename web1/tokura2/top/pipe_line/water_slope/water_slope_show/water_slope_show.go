package water_slope_show

import (

	    "net/http"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/process2"

                                                  )

///
/// 導水勾配線リストを表示する、E///

func Water_slope_show(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "water_slope_show start \n" )  // チE��チE��

/// モニター　表示 ///

    process2.Water_slope_show(w , r )

//	fmt.Fprintf( w, "water_slope_show : normal end \n" )  // チE��チE��

}

