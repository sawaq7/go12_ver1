package cal

import (

//	    "strconv"
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "client/tokura/suiri/type4"
	    "net/http"
	    "fmt"
//	    "storage2/get"

	    "html/template"
        "client/tokura/html4"

//        "general/type5"
//        "general/html5"

        "os"
        "cloud.google.com/go/datastore"
        "context"

                                                   )

///
/// グラフ-ファイル（導水勾配線群）を、ウエブ上に表示するとともに
/// データストアにグラフ-ファイル情報を登録する。
///


func  Pipe_line1_show_graf( w http.ResponseWriter ,r *http.Request ,f_name string) {

//     IN     w         : レスポンスライター
//     IN     r         : リクエストパラメーター
//     IN  f_name 　　  : ファイル名

	var g type4.Water_Slope // 画像ファイル表示用構造体　”type5.Image_Show”と同フォーマット

//    fmt.Fprintf( w, "pipe_line1_show_graf start \n" )  // デバック

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

//	fmt.Fprintf( w, "pipe_line1_show_graf : g.File_Name %v\n", g.File_Name )  // デバック
//	fmt.Fprintf( w, "pipe_line1_show_graf : g.Url %v\n", g.Url )  // デバック

///
/// データストアーにニューデータをセット
///

    new_key := datastore.IncompleteKey("Water_Slope", nil)

    if _, err = client.Put(ctx, new_key, &g ); err != nil {
//	if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Water_Slope", nil), &g); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

///
/// テンプレートのヘッダーをGET
///

//     monitor := template.Must(template.New("html").Parse(html5.Image_file_show))
     monitor := template.Must(template.New("html").Parse(html4.Pipe_line1_show_graf))

///
/// モニターに表示
///

	err = monitor.Execute(w, g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

//    get.Image_file_show( w ,r  ,bucket ,f_name)   test excute


}
