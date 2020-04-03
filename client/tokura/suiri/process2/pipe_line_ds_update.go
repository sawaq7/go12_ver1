package process2

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
	    "strconv"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
//	    "time"

        "os"
        "cloud.google.com/go/datastore"
        "context"
                                                )


func Pipe_line_ds_update(w http.ResponseWriter, r *http.Request ,updid int64) {

//    fmt.Fprintf( w, "pipe_line_ds_update start \n" )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "pipe_line_ds_update : updid %v\n", updid )  // 繝・ヰ繝・け


    var g type4.Water2

///
///   繝励Ο繧ｸ繧ｧ繧ｯ繝亥錐繧偵ご繝・ヨ
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

	g.Name = r.FormValue("water_name")  // 豌ｴ霍ｯ蜷阪ｒ繧ｲ繝・ヨ

	water_high := r.FormValue("water_high")      // 豌ｴ霍ｯ鬮倥ｒ繧ｲ繝・ヨ
	g.High,_ =strconv.ParseFloat(water_high,64)  //縲float64縲縺ｫ螟画鋤

	r_facter := r.FormValue("r_facter")      // 邊礼ｲ剃ｿよ焚繧偵ご繝・ヨ
	g.Roughness_Factor,_ =strconv.ParseFloat(r_facter,64)  //縲float64縲縺ｫ螟画鋤

//	fmt.Fprintf( w, "pipe_line_ds_update : g.Name %v\n", g.Name )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "pipe_line_ds_update : g.High %v\n", g.High )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "pipe_line_ds_update : g.Roughness_Factor %v\n", g.Roughness_Factor )  // 繝・ヰ繝・け

//	if _, err := datastore.Put(c, key, &g); err != nil {
    if _, err := client.Put(ctx, key, &g ); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

}
