package cal

import (

//	    "strconv"
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	    "net/http"
	    "fmt"
//	    "github.com/sawaq7/go12_ver1/storage2/get"

	    "html/template"
        "github.com/sawaq7/go12_ver1/client/tokura/html4"

//        "github.com/sawaq7/go12_ver1/general/type5"
//        "github.com/sawaq7/go12_ver1/general/html5"

        "os"
        "cloud.google.com/go/datastore"
        "context"

                                                   )

///
/// 繧ｰ繝ｩ繝・繝輔ぃ繧､繝ｫ・亥ｰ取ｰｴ蜍ｾ驟咲ｷ夂ｾ､・峨ｒ縲√え繧ｨ繝紋ｸ翫↓陦ｨ遉ｺ縺吶ｋ縺ｨ縺ｨ繧ゅ↓
/// 繝・・繧ｿ繧ｹ繝医い縺ｫ繧ｰ繝ｩ繝・繝輔ぃ繧､繝ｫ諠・ｱ繧堤匳骭ｲ縺吶ｋ縲・///


func  Pipe_line1_show_graf( w http.ResponseWriter ,r *http.Request ,f_name string) {

//     IN     w         : 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN     r         : 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ繝ｼ
//     IN  f_name 縲縲  : 繝輔ぃ繧､繝ｫ蜷・
	var g type4.Water_Slope // 逕ｻ蜒上ヵ繧｡繧､繝ｫ陦ｨ遉ｺ逕ｨ讒矩菴薙窶掖ype5.Image_Show窶昴→蜷後ヵ繧ｩ繝ｼ繝槭ャ繝・
//    fmt.Fprintf( w, "pipe_line1_show_graf start \n" )  // 繝・ヰ繝・け

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

    ctx := context.Background()
//	c := appengine.NewContext(r)

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return
    }

    g.File_Name = f_name

	bucket := "sample-7777"

	const publicURL = "https://storage.googleapis.com/%s/%s"
	g.Url = fmt.Sprintf(publicURL, bucket, g.File_Name)

//	fmt.Fprintf( w, "pipe_line1_show_graf : g.File_Name %v\n", g.File_Name )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "pipe_line1_show_graf : g.Url %v\n", g.Url )  // 繝・ヰ繝・け

///
/// 繝・・繧ｿ繧ｹ繝医い繝ｼ縺ｫ繝九Η繝ｼ繝・・繧ｿ繧偵そ繝・ヨ
///

    new_key := datastore.IncompleteKey("Water_Slope", nil)

    if _, err = client.Put(ctx, new_key, &g ); err != nil {
//	if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Water_Slope", nil), &g); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

///
/// 繝・Φ繝励Ξ繝ｼ繝医・繝倥ャ繝繝ｼ繧竪ET
///

//     monitor := template.Must(template.New("html").Parse(html5.Image_file_show))
     monitor := template.Must(template.New("html").Parse(html4.Pipe_line1_show_graf))

///
/// 繝｢繝九ち繝ｼ縺ｫ陦ｨ遉ｺ
///

	err = monitor.Execute(w, g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

//    get.Image_file_show( w ,r  ,bucket ,f_name)   test excute


}
