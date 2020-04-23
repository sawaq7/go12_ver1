package pipe_line_ds_wl_keyin

import (

	    "net/http"
	    "strconv"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/process2"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	    "github.com/sawaq7/go12_ver1/client/tokura/datastore4"

//	    "fmt"

        "cloud.google.com/go/datastore"
        "context"
        "os"

                                                 )

func Pipe_line_ds_wl_keyin(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "pipe_line_ds_wl_keyin start \n" )

    var g  type4.Water2

    var g2 type4.Water2_Temp

    var idmy int64

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

    updidw , err := strconv.Atoi(r.FormValue("id"))

	if err  != nil {
//	   fmt.Fprintf( w, "pipe_line_ds_wl_keyin :error updidw %v\n", updidw )
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}

    updid := int64(updidw)

//    fmt.Fprintf( w, "pipe_line_ds_wl_keyin : updidw %v\n", updidw )
//    fmt.Fprintf( w, "pipe_line_ds_wl_keyin : updid %v\n", updid )

     key := datastore.IDKey("Water2", updid, nil)

    if err := client.Get(ctx, key , &g ) ; err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

///
///          renewal temp.-file
///

///    initialize2.Water2_temp (w , r )

     // ini. temporary-file
     _ = datastore4.Datastore_tokura( "Water2_Temp"  ,"initialize"  ,idmy , w , r  )

    g2.Name = g.Name
    g2.High = g.High
    g2.Roughness_Factor = g.Roughness_Factor

//    fmt.Fprintf( w, "pipe_line_ds_wl_keyin : g2.Name %v\n", g2.Name )

    // put new-temporary-file
    new_key := datastore.IncompleteKey("Water2_Temp", nil)

    if _, err = client.Put(ctx, new_key, &g2 ); err != nil {

		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

///
///     show water-line inf.
///

   process2.Pipe_line_ds_wl_show(1 ,g.Name ,w , r )

}

