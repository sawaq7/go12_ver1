package pipe_line_ds_wl_show

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

func Pipe_line_ds_wl_show(w http.ResponseWriter, r *http.Request) {

//   fmt.Fprintf( w, "sky/pipe_line_ds_wl_show start \n"  )

   var g type4.Water_Line

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

/// チE��ポラリーファイル  get water-name from temp.-file

    query := datastore.NewQuery("Water2_Temp").Order("Name")

    count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

//    fmt.Fprintf( w, "sky/pipe_line_ds_wl_show  \n" ,count )

	water2_temp     := make([]type4.Water2_Temp, 0, count)

	keys, err := client.GetAll(ctx, query , &water2_temp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
    }

    keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

//    fmt.Fprintf( w, "sky/pipe_line_ds_wl_show : len(water2_temp) %v\n", len(water2_temp) )

	for pos2, water2_tempw := range water2_temp {

       g.Name = water2_tempw.Name
       g.Id   = keys_wk[pos2]

    }

//	g.Name = r.FormValue("water_name")

	g.Section = r.FormValue("section")

	f_facter := r.FormValue("f_facter")
	g.Friction_Factor,_ =strconv.ParseFloat(f_facter,64)

	velocity := r.FormValue("velocity")
	g.Velocity,_ =strconv.ParseFloat(velocity,64)

	p_diameter := r.FormValue("p_diameter")
	g.Pipe_Diameter,_ =strconv.ParseFloat(p_diameter,64)

	p_length := r.FormValue("p_length")
	g.Pipe_Length,_ =strconv.ParseFloat(p_length,64)

//	fmt.Fprintf( w, "sky/pipe_line_ds_wl_show : g.Name %v\n", g.Name )
//	fmt.Fprintf( w, "sky/pipe_line_ds_wl_show : g.Section %v\n", g.Section )
//	fmt.Fprintf( w, "sky/pipe_line_ds_wl_show : g.Friction_Factor %v\n", g.Friction_Factor )
//	fmt.Fprintf( w, "sky/pipe_line_ds_wl_show : g.Velocity %v\n", g.Velocity )
//	fmt.Fprintf( w, "sky/pipe_line_ds_wl_show : g.Pipe_Diameter %v\n", g.Pipe_Diameter )
//	fmt.Fprintf( w, "sky/pipe_line_ds_wl_show : g.Pipe_Length %v\n", g.Pipe_Length )

///    put water-line inf.

    new_key := datastore.IncompleteKey("Water_Line", nil)

    if _, err = client.Put(ctx, new_key, &g ); err != nil {

		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

///     show water-line inf.

   process2.Pipe_line_ds_wl_show(1 ,g.Name ,w , r )

}

