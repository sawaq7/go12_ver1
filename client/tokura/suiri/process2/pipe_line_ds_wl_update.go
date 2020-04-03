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


func Pipe_line_ds_wl_update(w http.ResponseWriter, r *http.Request ,updid int64) {

//    fmt.Fprintf( w, "pipe_line_ds_wl_update start \n" )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "pipe_line_ds_wl_update : updid %v\n", updid )  // 繝・ヰ繝・け


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

	g.Name = r.FormValue("water_name")  // 豌ｴ霍ｯ蜷阪ｒ繧ｲ繝・ヨ
	g.Section = r.FormValue("section")  // 蛹ｺ髢灘錐繧偵ご繝・ヨ

	f_facter := r.FormValue("f_facter")                   // 鞫ｩ謫ｦ菫よ焚繧偵ご繝・ヨ
	g.Friction_Factor,_ =strconv.ParseFloat(f_facter,64)  //縲float64縲縺ｫ螟画鋤

	velocity := r.FormValue("velocity")                   // 騾溷ｺｦ繧偵ご繝・ヨ
	g.Velocity,_ =strconv.ParseFloat(velocity,64)         //縲float64縲縺ｫ螟画鋤

	p_diameter := r.FormValue("p_diameter")      // 鞫ｩ謫ｦ菫よ焚繧偵ご繝・ヨ
	g.Pipe_Diameter,_ =strconv.ParseFloat(p_diameter,64)  //縲float64縲縺ｫ螟画鋤

	p_length := r.FormValue("p_length")      // 鞫ｩ謫ｦ菫よ焚繧偵ご繝・ヨ
	g.Pipe_Length,_ =strconv.ParseFloat(p_length,64)  //縲float64縲縺ｫ螟画鋤

//	fmt.Fprintf( w, "pipe_line_ds_wl_update : g.Name %v\n", g.Name )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "pipe_line_ds_wl_update : g.Section %v\n", g.Section )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "pipe_line_ds_wl_update : g.Friction_Factor %v\n", g.Friction_Factor )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "pipe_line_ds_wl_update : g.Velocity %v\n", g.Velocity )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "pipe_line_ds_wl_update : g.Pipe_Diameter %v\n", g.Pipe_Diameter )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "pipe_line_ds_wl_update : g.Pipe_Length %v\n", g.Pipe_Length )  // 繝・ヰ繝・け

    if _, err := client.Put(ctx, key, &g ); err != nil {
//	if _, err := datastore.Put(c, key, &g); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

}
