package process2

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/tokura/html4"

	    "github.com/sawaq7/go12_ver1/client/tokura/datastore4"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
//	    "time"
                                                )
func Pipe_line_ds_show(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "process2.pipe_line_ds_show start \n" )  // ãEãE¯

    var idmy int64

///
///      ãE³ãã¬ã¼ããEãããã¼ãGET
///

     monitor := template.Must(template.New("html").Parse(html4.Pipe_line_ds_keyin))

///
///   ãEEã¿ã¹ãã¢ã¼ãããè¡¨ç¤ºç¨ãEEã¿ãGET
///


//     water2_view := trans2.Water2 ( w ,r )

     water2_view := datastore4.Datastore_tokura( "Water2"  ,"trans"  ,idmy , w , r  )

     value, _ := water2_view.([]type4.Water2)    // ç©ºã¤ã³ã¿ã¼ãã§ã¤ã¹å¤æ°ããããªã¥ã¼å¤ãã²ãE


///
///             ã¢ãã¿ã¼ã«è¡¨ç¤º
///

	err := monitor.Execute(w, value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

