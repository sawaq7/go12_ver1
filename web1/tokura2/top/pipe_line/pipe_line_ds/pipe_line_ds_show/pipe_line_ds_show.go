package pipe_line_ds_show

import (

//        "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
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

/// key-in ãEEã¿ãGET ///

   var g type4.Water2

   project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

    ctx := context.Background()
//    c := appengine.NewContext(r)

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return
    }

	g.Name = r.FormValue("water_name")  // æ°´è·¯åãã²ãE

	water_high := r.FormValue("water_high")      // æ°´è·¯é«ãã²ãE
	g.High,_ =strconv.ParseFloat(water_high,64)  //ãfloat64ãã«å¤æ

	r_facter := r.FormValue("r_facter")      // ç²ç²ä¿æ°ãã²ãE
	g.Roughness_Factor,_ =strconv.ParseFloat(r_facter,64)  //ãfloat64ãã«å¤æ

//	fmt.Fprintf( w, "pipe_line_ds_keyin : g.Name %v\n", g.Name )  // ãEãE¯
//	fmt.Fprintf( w, "pipe_line_ds_keyin : g.High %v\n", g.High )  // ãEãE¯

/// ãEEã¿ã¹ãã¢ã¼ã«ãEEã¿ãã»ãE ///

    new_key := datastore.IncompleteKey("Water2", nil)

    if _, err = client.Put(ctx, new_key, &g ); err != nil {
//	if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Water2", nil), &g); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

/// ã¢ãã¿ã¼è¡¨ç¤º ///

   process2.Pipe_line_ds_show(w , r )

}

