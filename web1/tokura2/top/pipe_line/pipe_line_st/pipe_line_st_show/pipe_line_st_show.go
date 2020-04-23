package pipe_line_st_show

import (

	    "net/http"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/process2"
//	    "github.com/sawaq7/go12_ver1/client/tokura/storage3/put1"
	    "github.com/sawaq7/go12_ver1/client/tokura/storage3"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	    "strconv"
	    "github.com/sawaq7/go12_ver1/storage2"
//	    "fmt"

                                                  )

///
///     put new water inf. in storage
///

func Pipe_line_st_show(w http.ResponseWriter, r *http.Request) {

    var water2 type4.Water2

    var idmy int64

    new_flag := 1

    bucket := "sample-7777"

///
/// key-in チE�EタをGET
///

	water2.Name = r.FormValue("water_name")

	water_high := r.FormValue("water_high")
	water2.High,_ =strconv.ParseFloat(water_high,64)

	r_facter := r.FormValue("r_facter")
	water2.Roughness_Factor,_ =strconv.ParseFloat(r_facter,64)

//	fmt.Fprintf( w, "pipe_line_ds_keyin : water2.Name %v\n", water2.Name )
//	fmt.Fprintf( w, "pipe_line_ds_keyin : water2.High %v\n", water2.High )

///
///      check whether or not exist file "Water2"
///

    objects_minor , _ := storage2.Storage_basic( "list2" ,bucket ,idmy, w , r  )

    objects, _ := objects_minor.([]string)

//    objects :=  storage2.Object_List ( w  ,r , bucket )

    for _ , objectsw := range objects {

      if objectsw == "Water2.txt" {

	    new_flag = 0

	  }

    }

//    fmt.Fprintf(w, "process2.pipe_line_ds_keyin : new_flag %v\n", new_flag )

///
///         put Water2 inf. in storage
///

    if new_flag == 0 {

      _ , _ = storage3.Storage_tokura( "Water2" ,"put" ,water2 , idmy , w , r  )

//      put1.Water2 ( w , r ,water2 )   //  add an existing file

    } else {                          //  make a new file

      _ , _ = storage3.Storage_tokura( "Water2" ,"put2" ,water2 , idmy , w , r  )

//	  put1.Water2_new ( w , r ,water2 )

	}

///
///           show water inf. on web
///

   process2.Pipe_line_st_show(w , r )

}

