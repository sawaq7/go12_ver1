package process2

import (

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

//    fmt.Fprintf( w, "pipe_line_ds_wl_update start \n" )
//    fmt.Fprintf( w, "pipe_line_ds_wl_update : updid %v\n", updid )


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

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	g.Name = r.FormValue("water_name")
	g.Section = r.FormValue("section")

	f_facter := r.FormValue("f_facter")
	g.Friction_Factor,_ =strconv.ParseFloat(f_facter,64)

	velocity := r.FormValue("velocity")
	g.Velocity,_ =strconv.ParseFloat(velocity,64)

	p_diameter := r.FormValue("p_diameter")
	g.Pipe_Diameter,_ =strconv.ParseFloat(p_diameter,64)

	p_length := r.FormValue("p_length")      // 摩擦係数をゲチE��
	g.Pipe_Length,_ =strconv.ParseFloat(p_length,64)

//	fmt.Fprintf( w, "pipe_line_ds_wl_update : g.Name %v\n", g.Name )
//	fmt.Fprintf( w, "pipe_line_ds_wl_update : g.Section %v\n", g.Section )
//	fmt.Fprintf( w, "pipe_line_ds_wl_update : g.Friction_Factor %v\n", g.Friction_Factor )
//	fmt.Fprintf( w, "pipe_line_ds_wl_update : g.Velocity %v\n", g.Velocity )
//	fmt.Fprintf( w, "pipe_line_ds_wl_update : g.Pipe_Diameter %v\n", g.Pipe_Diameter )
//	fmt.Fprintf( w, "pipe_line_ds_wl_update : g.Pipe_Length %v\n", g.Pipe_Length )

    if _, err := client.Put(ctx, key, &g ); err != nil {

		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

}
