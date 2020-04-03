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

func Pipe_line_st_update(w http.ResponseWriter, r *http.Request) {

    var water2 type4.Water2

///
///          key-in γEEγΏγGET
///

    updidw , err := strconv.Atoi(r.FormValue("id"))     // idγγ²γE
	if err  != nil {

//	   fmt.Fprintf( w, "sky_pipe_line_ds_update :error updidw %v\n", updidw )  // γEγE―

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

    water2.Id = updid                      // γ¬γ³γΌγidγ»γE

	water2.Name = r.FormValue("water_name")  // ζ°΄θ·―εγγ²γE

	water_high := r.FormValue("water_high")      // ζ°΄θ·―ι«γγ²γE
	water2.High,_ =strconv.ParseFloat(water_high,64)  //γfloat64γγ«ε€ζ

	r_facter := r.FormValue("r_facter")      // η²η²δΏζ°γγ²γE
	water2.Roughness_Factor,_ =strconv.ParseFloat(r_facter,64)  //γfloat64γγ«ε€ζ

//	fmt.Fprintf( w, "sky.pipe_line_st_update : water2.Name %v\n", water2.Name )  // γEγE―
//	fmt.Fprintf( w, "sky.pipe_line_st_update : water2.High %v\n", water2.High )  // γE
///
///         γΉγγ¬γEΈγ«γ’γEEγEEγγγΌγΏγγ»γE
///

    _ , _ = storage3.Storage_tokura( "Water2" ,"put3" ,updid , water2 , w , r  )

//	put1.Water2_update ( w , r ,updid ,water2 )

///
///           γ’γγΏγΌθ‘¨η€Ί
///

   process2.Pipe_line_st_show(w , r )

}

