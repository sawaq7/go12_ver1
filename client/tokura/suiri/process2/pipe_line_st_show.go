package process2

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/tokura/html4"
        "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	    "github.com/sawaq7/go12_ver1/client/tokura/storage3"
	    "github.com/sawaq7/go12_ver1/storage2"
                                                )
func Pipe_line_st_show(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "process2.pipe_line_st_show start \n" )

    var idmy1 ,idmy2 int64

    skip_flag := 1

    bucket := "sample-7777"

///
///            check whether or not exist Water2-file
///

    objects :=  storage2.Object_List ( w  ,r , bucket )  //   get object list

    for _ , objectsw := range objects {

      if objectsw == "Water2.txt" {

	    skip_flag = 0

	  }

    }

//    fmt.Fprintf(w, "process2.pipe_line_st_show : skip_flag %v\n", skip_flag )

///
///          set template and show water-inf.  on web
///

     monitor := template.Must(template.New("html").Parse(html4.Pipe_line_st_keyin))

     if skip_flag == 0 {

     water2_view_minor , _ := storage3.Storage_tokura( "Water2" ,"trans" ,idmy1 , idmy2 , w , r  )

//       water2_view := trans4.Water2 ( w ,r )

       water2_view, _ := water2_view_minor.([]type4.Water2)

       err := monitor.Execute(w, water2_view)

       if err != nil {

	    http.Error(w, err.Error(), http.StatusInternalServerError)

	   }

     } else {

       water2_view := make([]type4.Water2, 0)
       err := monitor.Execute(w, water2_view)

       if err != nil {

	    http.Error(w, err.Error(), http.StatusInternalServerError)

	   }

     }

}

