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
/// 逕ｻ蜒上ヵ繧｡繧､繝ｫ繧偵√え繧ｨ繝紋ｸ翫↓陦ｨ遉ｺ縺吶ｋ縲・///


func  Image_file_show( w http.ResponseWriter ,r *http.Request ,bucket string ,filename string) {

//     IN    w      : 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r      : 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ
//     IN  bucket     : 繝舌こ繝・ヨ蜷・//     IN  filename   : 繝輔ぃ繧､繝ｫ蜷・
	var g type5.Image_Show

//    fmt.Fprintf( w, "image_file_show start \n" )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "image_file_show : bucket %v\n", bucket )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "image_file_show : filename %v\n", filename )  // 繝・ヰ繝・け

    g.File_Name = filename

//	fmt.Fprintf( w, "image_file_show : g.File_Name %v\n", g.File_Name )  // 繝・ヰ繝・け

	const publicURL = "https://storage.googleapis.com/%s/%s"
	g.Url = fmt.Sprintf(publicURL, bucket, filename)



/// 繝｢繝九ち繝ｼ縲陦ｨ遉ｺ ///
    // 繝・Φ繝励Ξ繝ｼ繝医・繝倥ャ繝繝ｼ繧竪ET

     monitor := template.Must(template.New("html").Parse(html5.Image_file_show))
//     monitor := template.Must(template.New("html").Parse(html4.Pipe_line1_show_graf))

// 繝｢繝九ち繝ｼ縺ｫ陦ｨ遉ｺ

	err := monitor.Execute(w, g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
