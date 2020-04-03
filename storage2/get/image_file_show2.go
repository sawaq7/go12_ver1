package get

import (

	    "github.com/sawaq7/go12_ver1/general/type5"
	    "net/http"
//	    "fmt"

	    "html/template"
        "github.com/sawaq7/go12_ver1/general/html5"
                         )

///
/// 逕ｻ蜒上ヵ繧｡繧､繝ｫ繧偵√え繧ｨ繝紋ｸ翫↓陦ｨ遉ｺ縺吶ｋ縲・part2
///


func  Image_file_show2( w http.ResponseWriter ,r *http.Request ,image_show type5.Image_Show ) {

//     IN    w       : 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r       : 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ
//     IN image_show : 讒矩菴薙窶晉判蜒上ョ繝ｼ繧ｿ窶・
//    fmt.Fprintf( w, "image_file_show2 start \n" )  // 繝・ヰ繝・け

//	fmt.Fprintf( w, "image_file_show2 : image_show.File_Name %v\n", image_show.File_Name )  // 繝・ヰ繝・け

/// 繝｢繝九ち繝ｼ縲陦ｨ遉ｺ ///
    // 繝・Φ繝励Ξ繝ｼ繝医・繝倥ャ繝繝ｼ繧竪ET

     monitor := template.Must(template.New("html").Parse(html5.Image_file_show))

// 繝｢繝九ち繝ｼ縺ｫ陦ｨ遉ｺ

	err := monitor.Execute(w, image_show)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
