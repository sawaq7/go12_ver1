package pipe_line_ds_wl_show

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

func Pipe_line_ds_wl_show(w http.ResponseWriter, r *http.Request) {

//   fmt.Fprintf( w, "sky/pipe_line_ds_wl_show start \n"  )  // 繝・ヰ繝・け

/// key-in 繝・・繧ｿ繧竪ET ///

   var g type4.Water_Line

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

/// 繝・Φ繝昴Λ繝ｪ繝ｼ繝輔ぃ繧､繝ｫ繧医ｊ縲∵ｰｴ霍ｯ蜷阪ｒ繧ｲ繝・ヨ

    query := datastore.NewQuery("Water2_Temp").Order("Name")
//    q2 := datastore.NewQuery("Water2_Temp").Order("Name")

//    count, err := q2.Count(c)
    count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

//    fmt.Fprintf( w, "sky/pipe_line_ds_wl_show  \n" ,count )  // 繝・ヰ繝・け

	water2_temp     := make([]type4.Water2_Temp, 0, count)


//	keys, err := q2.GetAll(c, &water2_temp )
	keys, err := client.GetAll(ctx, query , &water2_temp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
    }

    keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

//    fmt.Fprintf( w, "sky/pipe_line_ds_wl_show : len(water2_temp) %v\n", len(water2_temp) )  // 繝・ヰ繝・け

	for pos2, water2_tempw := range water2_temp {

       g.Name = water2_tempw.Name  /// 豌ｴ霍ｯ蜷阪・繧ｻ繝・ヨ
//       g.Id   = keys[pos2].IntID()
       g.Id   = keys_wk[pos2]

    }

//	g.Name = r.FormValue("water_name")  // 豌ｴ霍ｯ蜷阪ｒ繧ｲ繝・ヨ

	g.Section = r.FormValue("section")  // 蛹ｺ髢灘錐繧偵ご繝・ヨ

	f_facter := r.FormValue("f_facter")                   // 鞫ｩ謫ｦ菫よ焚繧偵ご繝・ヨ
	g.Friction_Factor,_ =strconv.ParseFloat(f_facter,64)  //縲float64縲縺ｫ螟画鋤

	velocity := r.FormValue("velocity")                   // 騾溷ｺｦ繧偵ご繝・ヨ
	g.Velocity,_ =strconv.ParseFloat(velocity,64)         //縲float64縲縺ｫ螟画鋤

	p_diameter := r.FormValue("p_diameter")      // 鞫ｩ謫ｦ菫よ焚繧偵ご繝・ヨ
	g.Pipe_Diameter,_ =strconv.ParseFloat(p_diameter,64)  //縲float64縲縺ｫ螟画鋤

	p_length := r.FormValue("p_length")      // 鞫ｩ謫ｦ菫よ焚繧偵ご繝・ヨ
	g.Pipe_Length,_ =strconv.ParseFloat(p_length,64)  //縲float64縲縺ｫ螟画鋤

//	fmt.Fprintf( w, "sky/pipe_line_ds_wl_show : g.Name %v\n", g.Name )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "sky/pipe_line_ds_wl_show : g.Section %v\n", g.Section )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "sky/pipe_line_ds_wl_show : g.Friction_Factor %v\n", g.Friction_Factor )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "sky/pipe_line_ds_wl_show : g.Velocity %v\n", g.Velocity )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "sky/pipe_line_ds_wl_show : g.Pipe_Diameter %v\n", g.Pipe_Diameter )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "sky/pipe_line_ds_wl_show : g.Pipe_Length %v\n", g.Pipe_Length )  // 繝・ヰ繝・け

/// 繝・・繧ｿ繧ｹ繝医い繝ｼ縺ｫ繝・・繧ｿ繧偵そ繝・ヨ ///

    new_key := datastore.IncompleteKey("Water_Line", nil)

    if _, err = client.Put(ctx, new_key, &g ); err != nil {
//	if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Water_Line", nil), &g); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

/// 繝｢繝九ち繝ｼ陦ｨ遉ｺ ///

   process2.Pipe_line_ds_wl_show(1 ,g.Name ,w , r )


}

