package get

import (

	    "github.com/sawaq7/go12_ver1/general/type5"
	    "net/http"
//	    "fmt"

	    "html/template"
        "github.com/sawaq7/go12_ver1/general/html5"
                         )

///
///         show image-inf. on web ( part2 )
///


func  Image_file_show2( w http.ResponseWriter ,r *http.Request ,image_show type5.Image_Show ) {

//     IN    w      : response-writer
//     IN    r      : request- paramete
//     IN image_show : struct ( type5.Image_Show )

//    fmt.Fprintf( w, "image_file_show2 start \n" )

//	fmt.Fprintf( w, "image_file_show2 : image_show.File_Name %v\n", image_show.File_Name )

///
///   template get
///

     monitor := template.Must(template.New("html").Parse(html5.Image_file_show))

///
///   show image-inf. on web
///

	err := monitor.Execute(w, image_show)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
