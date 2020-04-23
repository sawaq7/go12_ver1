package water_slope_show

import (

	    "net/http"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/process2"

                                                  )

///
///     show water-slope-line list on web
///

func Water_slope_show(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "water_slope_show start \n" )

///
/// 　　　　　show water-slope-line inf. on web
///

    process2.Water_slope_show(w , r )

//	fmt.Fprintf( w, "water_slope_show : normal end \n" )

}

