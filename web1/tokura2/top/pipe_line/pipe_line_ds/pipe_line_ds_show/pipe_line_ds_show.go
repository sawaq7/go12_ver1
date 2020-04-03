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

/// key-in 繝・・繧ｿ繧竪ET ///

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

	g.Name = r.FormValue("water_name")  // 豌ｴ霍ｯ蜷阪ｒ繧ｲ繝・ヨ

	water_high := r.FormValue("water_high")      // 豌ｴ霍ｯ鬮倥ｒ繧ｲ繝・ヨ
	g.High,_ =strconv.ParseFloat(water_high,64)  //縲float64縲縺ｫ螟画鋤

	r_facter := r.FormValue("r_facter")      // 邊礼ｲ剃ｿよ焚繧偵ご繝・ヨ
	g.Roughness_Factor,_ =strconv.ParseFloat(r_facter,64)  //縲float64縲縺ｫ螟画鋤

//	fmt.Fprintf( w, "pipe_line_ds_keyin : g.Name %v\n", g.Name )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "pipe_line_ds_keyin : g.High %v\n", g.High )  // 繝・ヰ繝・け

/// 繝・・繧ｿ繧ｹ繝医い繝ｼ縺ｫ繝・・繧ｿ繧偵そ繝・ヨ ///

    new_key := datastore.IncompleteKey("Water2", nil)

    if _, err = client.Put(ctx, new_key, &g ); err != nil {
//	if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Water2", nil), &g); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

/// 繝｢繝九ち繝ｼ陦ｨ遉ｺ ///

   process2.Pipe_line_ds_show(w , r )

}

