package pipe_line_st_wl_show

import (

	    "net/http"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/process2"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	    "strconv"
//	    "fmt"

//	    "github.com/sawaq7/go12_ver1/client/tokura/storage3/trans4"

//	    "github.com/sawaq7/go12_ver1/client/tokura/storage3/check4"
	    "github.com/sawaq7/go12_ver1/client/tokura/storage3"
                                                  )

///
///         show new water-line on web
///

func Pipe_line_st_wl_show(w http.ResponseWriter, r *http.Request) {

//   fmt.Fprintf( w, "sky/pipe_line_st_wl_show start \n"  )

   var water_line type4.Water_Line

   var idmy ,idmy2 int64

///
///        get water-name in temp.-file
///
     water2_temp_minor , _ := storage3.Storage_tokura( "Water2_Temp" ,"trans" ,idmy , idmy2 , w , r  )

     water2_temp, _ := water2_temp_minor.([]type4.Water2_Temp)

//   water2_temp := trans4.Water2_temp ( w ,r )

///
///        get an existing water-line's number which was selected
///

    record_number_temp , _ := storage3.Storage_tokura( "Water_Line" ,"check" ,water2_temp[0].Name , idmy , w , r  )

    record_number, _ := record_number_temp.(int64)

//    record_number := check4.Water_line_re_num( water2_temp[0].Name  ,w , r  )

	for _, water2_tempw := range water2_temp {

       water_line.Name = water2_tempw.Name
       water_line.Id   = record_number + int64( 1 )

    }

	water_line.Section = r.FormValue("section")

	f_facter := r.FormValue("f_facter")
	water_line.Friction_Factor,_ =strconv.ParseFloat(f_facter,64)

	velocity := r.FormValue("velocity")
	water_line.Velocity,_ =strconv.ParseFloat(velocity,64)

	p_diameter := r.FormValue("p_diameter")
	water_line.Pipe_Diameter,_ =strconv.ParseFloat(p_diameter,64)

	p_length := r.FormValue("p_length")
	water_line.Pipe_Length,_ =strconv.ParseFloat(p_length,64)  //　float64　に変換

//	fmt.Fprintf( w, "sky/pipe_line_st_wl_show : water_line.Name %v\n", water_line.Name )
//	fmt.Fprintf( w, "sky/pipe_line_st_wl_show : water_line.Section %v\n", water_line.Section )
//	fmt.Fprintf( w, "sky/pipe_line_st_wl_show : water_line.Friction_Factor %v\n", water_line.Friction_Factor )
//	fmt.Fprintf( w, "sky/pipe_line_st_wl_show : water_line.Velocity %v\n", water_line.Velocity )
//	fmt.Fprintf( w, "sky/pipe_line_st_wl_show : water_line.Pipe_Diameter %v\n", water_line.Pipe_Diameter )
//	fmt.Fprintf( w, "sky/pipe_line_st_wl_show : water_line.Pipe_Length %v\n", water_line.Pipe_Length )

///                           　　　　　　　　　　　
///       put water-line inf. in storage
///                          　　　　　　　　　　　

    _ , _ = storage3.Storage_tokura( "Water_Line" ,"put" ,water_line , idmy , w , r  )

//   put1.Water_line( w , r ,water_line )

///
///           show water-line inf. on web
///

   process2.Pipe_line_st_wl_show ( water_line.Name ,w , r )

}
