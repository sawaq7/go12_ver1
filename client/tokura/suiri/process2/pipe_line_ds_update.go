package process2

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
	    "strconv"
//	    "fmt"
	    "client/tokura/suiri/type4"
//	    "time"

        "os"
        "cloud.google.com/go/datastore"
        "context"
                                                )


func Pipe_line_ds_update(w http.ResponseWriter, r *http.Request ,updid int64) {

//    fmt.Fprintf( w, "pipe_line_ds_update start \n" )  // デバック
//    fmt.Fprintf( w, "pipe_line_ds_update : updid %v\n", updid )  // デバック


    var g type4.Water2

///
///   プロジェクト名をゲット
///
    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return
    }

    key := datastore.IDKey("Water2", updid, nil)

    if err := client.Get(ctx, key , &g ) ; err != nil {

//	key := datastore.NewKey(c, "Water2", "", updid, nil)
//	if err := datastore.Get(c, key, &g); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	g.Name = r.FormValue("water_name")  // 水路名をゲット

	water_high := r.FormValue("water_high")      // 水路高をゲット
	g.High,_ =strconv.ParseFloat(water_high,64)  //　float64　に変換

	r_facter := r.FormValue("r_facter")      // 粗粒係数をゲット
	g.Roughness_Factor,_ =strconv.ParseFloat(r_facter,64)  //　float64　に変換

//	fmt.Fprintf( w, "pipe_line_ds_update : g.Name %v\n", g.Name )  // デバック
//	fmt.Fprintf( w, "pipe_line_ds_update : g.High %v\n", g.High )  // デバック
//	fmt.Fprintf( w, "pipe_line_ds_update : g.Roughness_Factor %v\n", g.Roughness_Factor )  // デバック

//	if _, err := datastore.Put(c, key, &g); err != nil {
    if _, err := client.Put(ctx, key, &g ); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

}
