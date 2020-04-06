package get

import (

	    "github.com/sawaq7/go12_ver1/general/type5"
	    "net/http"
//	    "fmt"

	    "html/template"
        "github.com/sawaq7/go12_ver1/general/html5"
                         )

///
/// 画像ファイルを、ウエブ上に表示する、Epart2
///


func  Image_file_show2( w http.ResponseWriter ,r *http.Request ,image_show type5.Image_Show ) {

//     IN    w       : レスポンスライター
//     IN    r       : リクエストパラメータ
//     IN image_show : 構造体　”画像データ E
//    fmt.Fprintf( w, "image_file_show2 start \n" )

//	fmt.Fprintf( w, "image_file_show2 : image_show.File_Name %v\n", image_show.File_Name )

//   template get

     monitor := template.Must(template.New("html").Parse(html5.Image_file_show))

// モニターに表示

	err := monitor.Execute(w, image_show)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
