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


func Pipe_line_ds_wl_update(w http.ResponseWriter, r *http.Request ,updid int64) {

//    fmt.Fprintf( w, "pipe_line_ds_wl_update start \n" )  // デバック
//    fmt.Fprintf( w, "pipe_line_ds_wl_update : updid %v\n", updid )  // デバック


    var g type4.Water_Line

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

    key := datastore.IDKey("Water_Line", updid, nil)

    if err := client.Get(ctx, key , &g ) ; err != nil {
//	key := datastore.NewKey(c, "Water_Line", "", updid, nil)
//	if err := datastore.Get(c, key, &g); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	g.Name = r.FormValue("water_name")  // 水路名をゲット
	g.Section = r.FormValue("section")  // 区間名をゲット

	f_facter := r.FormValue("f_facter")                   // 摩擦係数をゲット
	g.Friction_Factor,_ =strconv.ParseFloat(f_facter,64)  //　float64　に変換

	velocity := r.FormValue("velocity")                   // 速度をゲット
	g.Velocity,_ =strconv.ParseFloat(velocity,64)         //　float64　に変換

	p_diameter := r.FormValue("p_diameter")      // 摩擦係数をゲット
	g.Pipe_Diameter,_ =strconv.ParseFloat(p_diameter,64)  //　float64　に変換

	p_length := r.FormValue("p_length")      // 摩擦係数をゲット
	g.Pipe_Length,_ =strconv.ParseFloat(p_length,64)  //　float64　に変換

//	fmt.Fprintf( w, "pipe_line_ds_wl_update : g.Name %v\n", g.Name )  // デバック
//	fmt.Fprintf( w, "pipe_line_ds_wl_update : g.Section %v\n", g.Section )  // デバック
//	fmt.Fprintf( w, "pipe_line_ds_wl_update : g.Friction_Factor %v\n", g.Friction_Factor )  // デバック
//	fmt.Fprintf( w, "pipe_line_ds_wl_update : g.Velocity %v\n", g.Velocity )  // デバック
//	fmt.Fprintf( w, "pipe_line_ds_wl_update : g.Pipe_Diameter %v\n", g.Pipe_Diameter )  // デバック
//	fmt.Fprintf( w, "pipe_line_ds_wl_update : g.Pipe_Length %v\n", g.Pipe_Length )  // デバック

    if _, err := client.Put(ctx, key, &g ); err != nil {
//	if _, err := datastore.Put(c, key, &g); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

}
