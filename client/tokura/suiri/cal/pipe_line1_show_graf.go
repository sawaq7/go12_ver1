package cal

import (

//	    "strconv"

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
/// グラチEファイル�E�導水勾配線群�E�を、ウエブ上に表示するとともに
/// チE�EタストアにグラチEファイル惁E��を登録する、E///


func  Pipe_line1_show_graf( w http.ResponseWriter ,r *http.Request ,f_name string) {

//     IN     w         : レスポンスライター
//     IN     r         : リクエストパラメーター
//     IN  f_name 　　  : file-name

	var g type4.Water_Slope // ”Water_Slope" and type5.Image_Show”is same format

//    fmt.Fprintf( w, "pipe_line1_show_graf start \n" )

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return
    }

    g.File_Name = f_name

	bucket := "sample-7777"

	const publicURL = "https://storage.googleapis.com/%s/%s"
	g.Url = fmt.Sprintf(publicURL, bucket, g.File_Name)

//	fmt.Fprintf( w, "pipe_line1_show_graf : g.File_Name %v\n", g.File_Name )
//	fmt.Fprintf( w, "pipe_line1_show_graf : g.Url %v\n", g.Url )

///
///     put new data in d.s.
///

    new_key := datastore.IncompleteKey("Water_Slope", nil)

    if _, err = client.Put(ctx, new_key, &g ); err != nil {

		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

///
///     set template
///

     monitor := template.Must(template.New("html").Parse(html4.Pipe_line1_show_graf))

///
///     show water-slope inf. on web
///

	err = monitor.Execute(w, g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
