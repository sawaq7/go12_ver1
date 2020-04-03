package get

import (

	    "github.com/sawaq7/go12_ver1/general/type5"
	    "net/http"
	    "fmt"

	    "html/template"
        "github.com/sawaq7/go12_ver1/general/html5"
//        "github.com/sawaq7/go12_ver1/client/tokura/html4"

                             )

///
/// 画像ファイルを、ウエブ上に表示する、E///


func  Image_file_show( w http.ResponseWriter ,r *http.Request ,bucket string ,filename string) {

//     IN    w      : レスポンスライター
//     IN    r      : リクエストパラメータ
//     IN  bucket     : バケチE��吁E//     IN  filename   : ファイル吁E
	var g type5.Image_Show

//    fmt.Fprintf( w, "image_file_show start \n" )  // チE��チE��
//    fmt.Fprintf( w, "image_file_show : bucket %v\n", bucket )  // チE��チE��
//    fmt.Fprintf( w, "image_file_show : filename %v\n", filename )  // チE��チE��

    g.File_Name = filename

//	fmt.Fprintf( w, "image_file_show : g.File_Name %v\n", g.File_Name )  // チE��チE��

	const publicURL = "https://storage.googleapis.com/%s/%s"
	g.Url = fmt.Sprintf(publicURL, bucket, filename)



/// モニター　表示 ///
    // チE��プレート�EヘッダーをGET

     monitor := template.Must(template.New("html").Parse(html5.Image_file_show))
//     monitor := template.Must(template.New("html").Parse(html4.Pipe_line1_show_graf))

// モニターに表示

	err := monitor.Execute(w, g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
