package pipe_line_st_cal

import (

	    "net/http"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/cal4"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
//	    "github.com/sawaq7/go12_ver1/client/tokura/storage3/trans4"
	    "github.com/sawaq7/go12_ver1/client/tokura/storage3"
//	    "strconv"
	    "fmt"

                                                  )

func Pipe_line_st_cal(w http.ResponseWriter, r *http.Request) {

//     IN     w         : response-writer
//     IN     r         : request-parameter

//   fmt.Fprintf( w, "sky/pipe_line_st_cal start \n"  )

   var water type4.Water2

   var idmy ,idmy2 int64

///
///        get current water-name
///

   water2_temp_minor , _ := storage3.Storage_tokura( "Water2_Temp" ,"trans" ,idmy , idmy2 , w , r  )

   //    water2_temp := trans4.Water2_temp( w , r  )

   water2_temp, _ := water2_temp_minor.([]type4.Water2_Temp)  // get value from interface data



    water.Name = water2_temp[0].Name
	water.High = water2_temp[0].High
	water.Roughness_Factor = water2_temp[0].Roughness_Factor

///
///    get water-linr inf.
///
      water_line_minor , _ := storage3.Storage_tokura( "Water_Line" ,"trans" ,water.Name , idmy , w , r  )
//    water_line := trans4.Water_line ( water.Name , w ,r )

      water_line, _ := water_line_minor.([]type4.Water_Line)

///
///         calculate water-slope-line
///

    p_number ,ad_eneup ,ad_enedown ,ad_glineup ,ad_glinedown := cal4.Pipe_line1( water  ,water_line  )

    fmt.Fprintf( w, "sky/pipe_line_st_cal : p_number %v\n", p_number )
    fmt.Fprintf( w, "sky/pipe_line_st_cal : ad_eneup %v\n", ad_eneup )
    fmt.Fprintf( w, "sky/pipe_line_st_cal : ad_enedown %v\n", ad_enedown )
   fmt.Fprintf( w, "sky/pipe_line_st_cal : ad_glineup %v\n", ad_glineup )
    fmt.Fprintf( w, "sky/pipe_line_st_cal : ad_glinedown %v\n", ad_glinedown )
    fmt.Fprintf( w, "sky/pipe_line_st_cal : ad_eneup len %v\n", len(ad_eneup) )

///
///           make graf on wab
///

//    f_name := cal.Pipe_line1_make_graf( w ,r ,p_number ,ad_eneup ,ad_enedown ,ad_glineup ,ad_glinedown  )

//    fmt.Fprintf( w, "sky/pipe_line_st_cal : f_name %v\n", f_name )

///
///           show graf on wab
///

//    cal.Pipe_line1_show_graf( w ,r ,f_name )

}
