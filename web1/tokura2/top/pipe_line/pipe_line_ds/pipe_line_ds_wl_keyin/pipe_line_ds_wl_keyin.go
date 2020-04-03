package pipe_line_ds_wl_keyin

import (

//        "google.golang.org/appengine"
//        "google.golang.org/appengine/datastore"
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

//    fmt.Fprintf( w, "pipe_line_ds_wl_keyin start \n" )  // 繝・ヰ繝・け

    var g  type4.Water2

    var g2 type4.Water2_Temp

    var idmy int64

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

    updidw , err := strconv.Atoi(r.FormValue("id"))

	if err  != nil {
//	   fmt.Fprintf( w, "pipe_line_ds_wl_keyin :error updidw %v\n", updidw )  // 繝・ヰ繝・け
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}

    updid := int64(updidw)

//    fmt.Fprintf( w, "pipe_line_ds_wl_keyin : updidw %v\n", updidw )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "pipe_line_ds_wl_keyin : updid %v\n", updid )  // 繝・ヰ繝・け

     key := datastore.IDKey("Water2", updid, nil)
//	key := datastore.NewKey(c, "Water2", "", updid, nil)

    if err := client.Get(ctx, key , &g ) ; err != nil {
//	if err := datastore.Get(c, key, &g); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

///
///          豌ｴ霍ｯ諠・ｱ繧偵う繝九す繝｣繝ｩ繧､繧ｺ  & 繧ｻ繝・ヨ
///

///    initialize2.Water2_temp (w , r ) // temporary-file繧偵う繝九す繝｣繝ｩ繧､繧ｺ

     _ = datastore4.Datastore_tokura( "Water2_Temp"  ,"initialize"  ,idmy , w , r  )

    g2.Name = g.Name
    g2.High = g.High
    g2.Roughness_Factor = g.Roughness_Factor

//    fmt.Fprintf( w, "pipe_line_ds_wl_keyin : g2.Name %v\n", g2.Name )  // 繝・ヰ繝・け
    new_key := datastore.IncompleteKey("Water2_Temp", nil)

    if _, err = client.Put(ctx, new_key, &g2 ); err != nil {
//	if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Water2_Temp", nil), &g2); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

///
///     繝｢繝九ち繝ｼ陦ｨ遉ｺ
///

   process2.Pipe_line_ds_wl_show(1 ,g.Name ,w , r )

}

