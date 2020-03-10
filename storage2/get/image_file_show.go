package get

import (

	    "general/type5"
	    "net/http"
	    "fmt"

	    "html/template"
        "general/html5"
//        "client/tokura/html4"

                             )

///
/// 画像ファイルを、ウエブ上に表示する。
///


func  Image_file_show( w http.ResponseWriter ,r *http.Request ,bucket string ,filename string) {

//     IN    w      : レスポンスライター
//     IN    r      : リクエストパラメータ
//     IN  bucket     : バケット名
//     IN  filename   : ファイル名

	var g type5.Image_Show

//    fmt.Fprintf( w, "image_file_show start \n" )  // デバック
//    fmt.Fprintf( w, "image_file_show : bucket %v\n", bucket )  // デバック
//    fmt.Fprintf( w, "image_file_show : filename %v\n", filename )  // デバック

    g.File_Name = filename

//	fmt.Fprintf( w, "image_file_show : g.File_Name %v\n", g.File_Name )  // デバック

	const publicURL = "https://storage.googleapis.com/%s/%s"
	g.Url = fmt.Sprintf(publicURL, bucket, filename)



/// モニター　表示 ///
    // テンプレートのヘッダーをGET

     monitor := template.Must(template.New("html").Parse(html5.Image_file_show))
//     monitor := template.Must(template.New("html").Parse(html4.Pipe_line1_show_graf))

// モニターに表示

	err := monitor.Execute(w, g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
