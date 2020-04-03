package pipe_line_ds_wl_update

import (

	    "strconv"
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/process2"

	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"

	    "cloud.google.com/go/datastore"
         "context"
         "os"


                                                   )

func Pipe_line_ds_wl_update(w http.ResponseWriter, r *http.Request) {

//	fmt.Fprintf( w, "sky_pipe_line_ds_wl_update start %v\n" )  // 繝・ヰ繝・け

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

/// 謖・ｮ壹＠縺溘ョ繝ｼ繧ｿid繧竪ET ///

    updidw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {

//	   fmt.Fprintf( w, "sky_pipe_line_ds_wl_update :error updidw %v\n", updidw )  // 繝・ヰ繝・け

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

/// 驟埼＃諠・ｱ縺ｮ螟画峩 ///

	process2.Pipe_line_ds_wl_update(w , r ,updid)

/// 繝｢繝九ち繝ｼ縲蜀崎｡ｨ遉ｺ ///

    key := datastore.IDKey("Water_Line", updid, nil)

    if err := client.Get(ctx, key , &g ) ; err != nil {
//    key := datastore.NewKey(c, "Water_Line", "", updid, nil)
//    if err := datastore.Get(c, key, &g); err != nil {              //縲縲豌ｴ霍ｯ蜷阪ｒGET
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	process2.Pipe_line_ds_wl_show(1 ,g.Name ,w , r )

//	http.Redirect(w, r, "/", http.StatusFound)


}
