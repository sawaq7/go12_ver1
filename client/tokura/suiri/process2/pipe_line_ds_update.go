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


func Pipe_line_ds_update(w http.ResponseWriter, r *http.Request ,updid int64) {

//    fmt.Fprintf( w, "pipe_line_ds_update start \n" )
//    fmt.Fprintf( w, "pipe_line_ds_update : updid %v\n", updid )


    var g type4.Water2

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return
    }

    key := datastore.IDKey("Water2", updid, nil)

    if err := client.Get(ctx, key , &g ) ; err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	g.Name = r.FormValue("water_name")

	water_high := r.FormValue("water_high")
	g.High,_ =strconv.ParseFloat(water_high,64)

	r_facter := r.FormValue("r_facter")
	g.Roughness_Factor,_ =strconv.ParseFloat(r_facter,64)

//	fmt.Fprintf( w, "pipe_line_ds_update : g.Name %v\n", g.Name )
//	fmt.Fprintf( w, "pipe_line_ds_update : g.High %v\n", g.High )
//	fmt.Fprintf( w, "pipe_line_ds_update : g.Roughness_Factor %v\n", g.Roughness_Factor )

    if _, err := client.Put(ctx, key, &g ); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

}
