package pipe_line_ds_wl_show

import (

//        "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
	    "client/tokura/suiri/process2"
	    "client/tokura/suiri/type4"
	    "strconv"
//	    "fmt"

	    "cloud.google.com/go/datastore"
        "context"
        "os"

                                                  )

func Pipe_line_ds_wl_show(w http.ResponseWriter, r *http.Request) {

//   fmt.Fprintf( w, "sky/pipe_line_ds_wl_show start \n"  )  // デバック

/// key-in データをGET ///

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

/// テンポラリーファイルより、水路名をゲット

    query := datastore.NewQuery("Water2_Temp").Order("Name")
//    q2 := datastore.NewQuery("Water2_Temp").Order("Name")

//    count, err := q2.Count(c)
    count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

//    fmt.Fprintf( w, "sky/pipe_line_ds_wl_show  \n" ,count )  // デバック

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

//    fmt.Fprintf( w, "sky/pipe_line_ds_wl_show : len(water2_temp) %v\n", len(water2_temp) )  // デバック

	for pos2, water2_tempw := range water2_temp {

       g.Name = water2_tempw.Name  /// 水路名のセット
//       g.Id   = keys[pos2].IntID()
       g.Id   = keys_wk[pos2]

    }

//	g.Name = r.FormValue("water_name")  // 水路名をゲット

	g.Section = r.FormValue("section")  // 区間名をゲット

	f_facter := r.FormValue("f_facter")                   // 摩擦係数をゲット
	g.Friction_Factor,_ =strconv.ParseFloat(f_facter,64)  //　float64　に変換

	velocity := r.FormValue("velocity")                   // 速度をゲット
	g.Velocity,_ =strconv.ParseFloat(velocity,64)         //　float64　に変換

	p_diameter := r.FormValue("p_diameter")      // 摩擦係数をゲット
	g.Pipe_Diameter,_ =strconv.ParseFloat(p_diameter,64)  //　float64　に変換

	p_length := r.FormValue("p_length")      // 摩擦係数をゲット
	g.Pipe_Length,_ =strconv.ParseFloat(p_length,64)  //　float64　に変換

//	fmt.Fprintf( w, "sky/pipe_line_ds_wl_show : g.Name %v\n", g.Name )  // デバック
//	fmt.Fprintf( w, "sky/pipe_line_ds_wl_show : g.Section %v\n", g.Section )  // デバック
//	fmt.Fprintf( w, "sky/pipe_line_ds_wl_show : g.Friction_Factor %v\n", g.Friction_Factor )  // デバック
//	fmt.Fprintf( w, "sky/pipe_line_ds_wl_show : g.Velocity %v\n", g.Velocity )  // デバック
//	fmt.Fprintf( w, "sky/pipe_line_ds_wl_show : g.Pipe_Diameter %v\n", g.Pipe_Diameter )  // デバック
//	fmt.Fprintf( w, "sky/pipe_line_ds_wl_show : g.Pipe_Length %v\n", g.Pipe_Length )  // デバック

/// データストアーにデータをセット ///

    new_key := datastore.IncompleteKey("Water_Line", nil)

    if _, err = client.Put(ctx, new_key, &g ); err != nil {
//	if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Water_Line", nil), &g); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

/// モニター表示 ///

   process2.Pipe_line_ds_wl_show(1 ,g.Name ,w , r )


}

