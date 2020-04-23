package pipe_line_st_update

import (

	    "net/http"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/process2"
//	    "github.com/sawaq7/go12_ver1/client/tokura/storage3/put1"
	    "github.com/sawaq7/go12_ver1/client/tokura/storage3"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	    "strconv"
//	    "storage2"
//	    "fmt"

                                                  )

///
///     update water inf. in storage
///

func Pipe_line_st_update(w http.ResponseWriter, r *http.Request) {

    var water2 type4.Water2

///
///          get key-in data
///

    updidw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {

//	   fmt.Fprintf( w, "sky_pipe_line_ds_update :error updidw %v\n", updidw )

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

    water2.Id = updid

	water2.Name = r.FormValue("water_name")

	water_high := r.FormValue("water_high")
	water2.High,_ =strconv.ParseFloat(water_high,64)

	r_facter := r.FormValue("r_facter")
	water2.Roughness_Factor,_ =strconv.ParseFloat(r_facter,64)

//	fmt.Fprintf( w, "sky.pipe_line_st_update : water2.Name %v\n", water2.Name )
//	fmt.Fprintf( w, "sky.pipe_line_st_update : water2.High %v\n", water2.High )

///
///        put Water2 inf.  on storage
///

    _ , _ = storage3.Storage_tokura( "Water2" ,"put3" ,updid , water2 , w , r  )

//	put1.Water2_update ( w , r ,updid ,water2 )

///
///          show water inf. on web
///

   process2.Pipe_line_st_show(w , r )

}

