package pipe_line_st_wl_update

import (

	    "strconv"
	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/process2"
//	    "github.com/sawaq7/go12_ver1/client/tokura/storage3/trans4"
	    "github.com/sawaq7/go12_ver1/client/tokura/storage3"
//	    "github.com/sawaq7/go12_ver1/client/tokura/storage3/put1"

	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	    "github.com/sawaq7/go12_ver1/general/type5"

                                                   )

///
///     update water-line inf. in storage
///

func Pipe_line_st_wl_update(w http.ResponseWriter, r *http.Request) {

//	fmt.Fprintf( w, "sky_pipe_line_st_wl_update start %v\n" )

	var water_line  type4.Water_Line

	var idmy1 ,idmy2 int64

///
///      get current water-name
///

      water2_temp_minor , _ := storage3.Storage_tokura( "Water2_Temp" ,"trans" ,idmy1 , idmy2 , w , r  )

      water2_temp, _ := water2_temp_minor.([]type4.Water2_Temp)

//    water2_temp := trans4.Water2_temp( w , r  )

    updidw , err := strconv.Atoi(r.FormValue("id"))

	if err  != nil {

//	   fmt.Fprintf( w, "sky_pipe_line_st_wl_update :error updidw %v\n", updidw )

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}

    updid := int64(updidw)

    water_line.Id = updid

    water_line.Name = water2_temp[0].Name

	water_line.Section = r.FormValue("section")

	f_facter := r.FormValue("f_facter")
	water_line.Friction_Factor,_ =strconv.ParseFloat(f_facter,64)

	velocity := r.FormValue("velocity")
	water_line.Velocity,_ =strconv.ParseFloat(velocity,64)

	p_diameter := r.FormValue("p_diameter")
	water_line.Pipe_Diameter,_ =strconv.ParseFloat(p_diameter,64)

	p_length := r.FormValue("p_length")
	water_line.Pipe_Length,_ =strconv.ParseFloat(p_length,64)

//	fmt.Fprintf( w, "pipe_line_st_wl_update : water_line.Name %v\n", water_line.Name )
//	fmt.Fprintf( w, "pipe_line_st_wl_update : water_line.Section %v\n", water_line.Section )
//	fmt.Fprintf( w, "pipe_line_st_wl_update : water_line.Friction_Factor %v\n", water_line.Friction_Factor )
//	fmt.Fprintf( w, "pipe_line_st_wl_update : water_line.Velocity %v\n", water_line.Velocity )
//	fmt.Fprintf( w, "pipe_line_st_wl_update : water_line.Pipe_Diameter %v\n", water_line.Pipe_Diameter )
//	fmt.Fprintf( w, "pipe_line_st_wl_update : water_line.Pipe_Length %v\n", water_line.Pipe_Length )

///
///         put water-line inf. which was updated in storage
///

     general_work := make([]type5.General_Work, 1)
     general_work[0].Int64_Work  = updid
     general_work[0].String_Work = water2_temp[0].Name

     _ , _ = storage3.Storage_tokura( "Water_Line" ,"put2" ,general_work , water_line , w , r  )

//	put1.Water_line_update ( w , r ,updid ,water2_temp[0].Name ,water_line )

///
///         show water-line inf. on web
///

	process2.Pipe_line_st_wl_show ( water2_temp[0].Name ,w , r )

}
