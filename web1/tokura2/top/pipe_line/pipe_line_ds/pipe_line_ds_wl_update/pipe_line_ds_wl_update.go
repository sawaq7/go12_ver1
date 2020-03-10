package pipe_line_ds_wl_update

import (

	    "strconv"
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
	    "client/tokura/suiri/process2"

	    "client/tokura/suiri/type4"

	    "cloud.google.com/go/datastore"
         "context"
         "os"


                                                   )

func Pipe_line_ds_wl_update(w http.ResponseWriter, r *http.Request) {

//	fmt.Fprintf( w, "sky_pipe_line_ds_wl_update start %v\n" )  // デバック

	var g  type4.Water_Line

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

/// 指定したデータidをGET ///

    updidw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {

//	   fmt.Fprintf( w, "sky_pipe_line_ds_wl_update :error updidw %v\n", updidw )  // デバック

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

/// 配達情報の変更 ///

	process2.Pipe_line_ds_wl_update(w , r ,updid)

/// モニター　再表示 ///

    key := datastore.IDKey("Water_Line", updid, nil)

    if err := client.Get(ctx, key , &g ) ; err != nil {
//    key := datastore.NewKey(c, "Water_Line", "", updid, nil)
//    if err := datastore.Get(c, key, &g); err != nil {              //　　水路名をGET
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	process2.Pipe_line_ds_wl_show(1 ,g.Name ,w , r )

//	http.Redirect(w, r, "/", http.StatusFound)


}
