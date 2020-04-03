package process2

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/tokura/html4"
        "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	    "github.com/sawaq7/go12_ver1/client/tokura/storage3"
	    "storage2"
                                                )
func Pipe_line_st_show(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "process2.pipe_line_st_show start \n" )  // ãEãE¯

    var idmy1 ,idmy2 int64

    skip_flag := 1

    bucket := "sample-7777"

///
///             Water2ããã¡ã¤ã«ãããããã§ãE¯
///

    objects :=  storage2.Object_List ( w  ,r , bucket )  // ãã±ãEåEEãªãã¸ã§ã¯ããã²ãEãã

    for _ , objectsw := range objects {

      if objectsw == "Water2.txt" {

	    skip_flag = 0

	  }

    }

//    fmt.Fprintf(w, "process2.pipe_line_st_show : skip_flag %v\n", skip_flag )  // ãEãE¯

///
///            ãè¡¨ç¤ºç¨ãEEã¿ã»ãE³ãã¬ã¼ããEãããã¼ãGETãã¦è¡¨ç¤º
///

     monitor := template.Must(template.New("html").Parse(html4.Pipe_line_st_keyin))

     if skip_flag == 0 {

     water2_view_minor , _ := storage3.Storage_tokura( "Water2" ,"trans" ,idmy1 , idmy2 , w , r  )

//       water2_view := trans4.Water2 ( w ,r )

       water2_view, _ := water2_view_minor.([]type4.Water2)  // ã¤ã³ã¿ã¼ãã§ã¤ã¹åãåå¤æ

       err := monitor.Execute(w, water2_view)

       if err != nil {

	    http.Error(w, err.Error(), http.StatusInternalServerError)

	   }

     } else {

       water2_view := make([]type4.Water2, 0)   //   Water2ãã®è¡¨ç¤ºã¨ãªã¢ãç¢ºä¿E
       err := monitor.Execute(w, water2_view)

       if err != nil {

	    http.Error(w, err.Error(), http.StatusInternalServerError)

	   }

     }

}

