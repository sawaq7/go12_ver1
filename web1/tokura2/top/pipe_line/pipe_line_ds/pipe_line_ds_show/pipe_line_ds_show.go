package pipe_line_ds_show

import (

	    "net/http"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/process2"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	    "strconv"
//	    "fmt"

        "cloud.google.com/go/datastore"
        "context"
        "os"

                                                  )

func Pipe_line_ds_show(w http.ResponseWriter, r *http.Request) {

   var g type4.Water2

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

	g.Name = r.FormValue("water_name")

	water_high := r.FormValue("water_high")
	g.High,_ =strconv.ParseFloat(water_high,64)

	r_facter := r.FormValue("r_facter")
	g.Roughness_Factor,_ =strconv.ParseFloat(r_facter,64)

//	fmt.Fprintf( w, "pipe_line_ds_keyin : g.Name %v\n", g.Name )
//	fmt.Fprintf( w, "pipe_line_ds_keyin : g.High %v\n", g.High )

///    put new data in d.s.

    new_key := datastore.IncompleteKey("Water2", nil)

    if _, err = client.Put(ctx, new_key, &g ); err != nil {

		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

///     show water inf. on web

   process2.Pipe_line_ds_show(w , r )

}

