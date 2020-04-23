package suiri

import (
	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"

        "github.com/sawaq7/go12_ver1/client/tokura/html4"
                                      )


func Pipe_line1_show( w http.ResponseWriter , wnumber int , water []type4.Water  ) {

//     IN    w     : レスポンスライター
//     IN  wnumbwr : water no
//     IN  water   :



//    fmt.Fprintf( w, "suiri/Pipe_line1_show start \n" )
//    fmt.Fprintf( w, "suiri/Pipe_line1_show wnumber  %v \n" , wnumber)

    water2 := make([]type4.Water, wnumber)

    for pos, waterwk := range water {

			water2[pos].No  = waterwk.No
			water2[pos].Name = waterwk.Name
			water2[pos].High = waterwk.High
            water2[pos].Roughness_factor = waterwk.Roughness_factor
//			fmt.Fprintf( w, "suiri/Pipe_line1_show pos ,NO  ,Name %v %v %v \n" , pos ,water2[pos].No ,water2[pos].Name)
//			fmt.Fprintf( w, "suiri/Pipe_line1_show name High ,Roughness_factor %v %v  \n" , water2[pos].High ,water2[pos].Roughness_factor)
	}

///    set template
	monitor := template.Must(template.New("html").Parse( html4.Pipe_line1_show))
//	monitor := template.Must(template.New("html").Parse( html4.Pipe_line1_show2))



//    show water-inf. on web

	err := monitor.Execute(w, water2)
    if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

