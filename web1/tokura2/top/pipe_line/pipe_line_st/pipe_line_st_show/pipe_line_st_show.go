package pipe_line_st_show

import (

	    "net/http"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/process2"
//	    "github.com/sawaq7/go12_ver1/client/tokura/storage3/put1"
	    "github.com/sawaq7/go12_ver1/client/tokura/storage3"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	    "strconv"
	    "storage2"
//	    "fmt"

                                                  )

///
///     sky ζ°΄θ·―γEEγΏγγγ‘γ€γ«γ«θΏ½ε  (γΉγγ¬γEΈEE///

func Pipe_line_st_show(w http.ResponseWriter, r *http.Request) {

    var water2 type4.Water2

    var idmy int64

    new_flag := 1

    bucket := "sample-7777"

///
/// key-in γEEγΏγGET
///

	water2.Name = r.FormValue("water_name")  // ζ°΄θ·―εγγ²γE

	water_high := r.FormValue("water_high")      // ζ°΄θ·―ι«γγ²γE
	water2.High,_ =strconv.ParseFloat(water_high,64)  //γfloat64γγ«ε€ζ

	r_facter := r.FormValue("r_facter")      // η²η²δΏζ°γγ²γE
	water2.Roughness_Factor,_ =strconv.ParseFloat(r_facter,64)  //γfloat64γγ«ε€ζ

//	fmt.Fprintf( w, "pipe_line_ds_keyin : water2.Name %v\n", water2.Name )  // γEγE―
//	fmt.Fprintf( w, "pipe_line_ds_keyin : water2.High %v\n", water2.High )  // γEγE―

///
///             Water2γγγ‘γ€γ«γγγγγγ§γE―
///

    objects_minor , _ := storage2.Storage_basic( "list2" ,bucket ,idmy, w , r  )

    objects, _ := objects_minor.([]string)  // γ€γ³γΏγΌγγ§γ€γΉεγεε€ζ

//    objects :=  storage2.Object_List ( w  ,r , bucket )  // γγ±γEεEEγͺγγΈγ§γ―γγγ²γEγγ

    for _ , objectsw := range objects {

      if objectsw == "Water2.txt" {

	    new_flag = 0

	  }

    }

//    fmt.Fprintf(w, "process2.pipe_line_ds_keyin : new_flag %v\n", new_flag )  // γEγE―

///
///         γΉγγ¬γEΈγ«γEEγΏγγ»γE
///

    if new_flag == 0 {

      _ , _ = storage3.Storage_tokura( "Water2" ,"put" ,water2 , idmy , w , r  )

//      put1.Water2 ( w , r ,water2 )   //  γγ‘γ€γ«γ«θΏ½ε 

    } else {                          //  γγ‘γ€γ«γζ°θ¦δ½ζE

      _ , _ = storage3.Storage_tokura( "Water2" ,"put2" ,water2 , idmy , w , r  )

//	  put1.Water2_new ( w , r ,water2 )

	}

///
///           γ’γγΏγΌθ‘¨η€Ί
///

   process2.Pipe_line_st_show(w , r )

}

