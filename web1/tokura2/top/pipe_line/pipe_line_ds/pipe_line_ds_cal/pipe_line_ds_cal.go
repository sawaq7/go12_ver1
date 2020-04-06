package pipe_line_ds_cal

import (

//        "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/cal"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	    "github.com/sawaq7/go12_ver1/client/tokura/datastore4"
//	    "strconv"
//	    "fmt"

        "cloud.google.com/go/datastore"
        "context"
        "os"

                                                  )

func Pipe_line_ds_cal(w http.ResponseWriter, r *http.Request) {

//   fmt.Fprintf( w, "sky/pipe_line_ds_cal start \n"  )  // チE��チE��

/// key-in チE�EタをGET ///

   var water type4.Water2

   project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)

    query := datastore.NewQuery("Water2_Temp").Order("Name")
//    q2 := datastore.NewQuery("Water2_Temp").Order("Name")

//	count, err := q2.Count(c)
    count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

//    fmt.Fprintf( w, "sky/pipe_line_ds_cal  \n" ,count )  // チE��チE��

    water_temp      := make([]type4.Water2, 0, count)

    keys, err := client.GetAll(ctx, query , &water_temp)
//	keys, err := q2.GetAll(c, &water_temp )
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
    }

//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : len(water) %v\n", len(water) )  // チE��チE��

    keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	for pos2, waterw := range water_temp {

//       water.Id   = keys[pos2].IntID()
       water.Id   = keys_wk[pos2]

	   water.Name = waterw.Name
	   water.High =  waterw.High
	   water.Roughness_Factor = waterw.Roughness_Factor

    }

///
///      水路ラインのチE�EタをゲチE��
///

//    water_line := trans2.Water_line (1  ,water.Name , w ,r )

      water_line := datastore4.Datastore_tokura( "Water_Line"  ,"trans"  ,water.Name , w , r  )



     value, _ := water_line.([]type4.Water_Line)    // 空インターフェイス変数よりバリュー値をゲチE��

///
///         動水勾配線�E計箁E///

    p_number ,ad_eneup ,ad_enedown ,ad_glineup ,ad_glinedown := cal.Pipe_line1( water  ,value  )

//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : p_number %v\n", p_number )  // チE��チE��
//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : ad_eneup %v\n", ad_eneup )  // チE��チE��
//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : ad_enedown %v\n", ad_enedown )  // チE��チE��
//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : ad_glineup %v\n", ad_glineup )  // チE��チE��
//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : ad_glinedown %v\n", ad_glinedown )  // チE��チE��
//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : ad_eneup len %v\n", len(ad_eneup) )  // チE��チE��

/// グラフ�E作�E ///

    f_name := cal.Pipe_line1_make_graf( w ,r ,p_number ,ad_eneup ,ad_enedown ,ad_glineup ,ad_glinedown  )

//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : f_name %v\n", f_name )  // チE��チE��

/// グラフ�E表示 ///

    cal.Pipe_line1_show_graf( w ,r ,f_name )

}
