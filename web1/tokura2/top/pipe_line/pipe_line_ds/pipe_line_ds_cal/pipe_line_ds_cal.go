package pipe_line_ds_cal

import (


	    "net/http"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/cal"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
//	    "github.com/sawaq7/go12_ver1/client/tokura/datastore4"
	    "github.com/sawaq7/go12_ver1/client/tokura/datastore4/trans2"
//	    "github.com/sawaq7/go12_ver1/basic/type3"
//	    "strconv"
	    "fmt"

        "cloud.google.com/go/datastore"
        "context"
        "os"

                                                  )

func Pipe_line_ds_cal(w http.ResponseWriter, r *http.Request) {

//   fmt.Fprintf( w, "sky/pipe_line_ds_cal start \n"  )

   var water type4.Water2

   project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)

    query := datastore.NewQuery("Water2_Temp").Order("Name")

    count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

//    fmt.Fprintf( w, "sky/pipe_line_ds_cal  \n" ,count )

    water_temp      := make([]type4.Water2, 0, count)

    keys, err := client.GetAll(ctx, query , &water_temp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
    }

//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : len(water) %v\n", len(water) )

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
///      get water-line inf.
///

    water_line := trans2.Water_line (1  ,water.Name , w ,r )

//      water_line := datastore4.Datastore_tokura( "Water_Line"  ,"trans"  ,water.Name , w , r  )

//      _ = datastore4.Datastore_tokura( "Water_Line"  ,"trans"  ,water.Name , w , r  )



//     value, _ := water_line.([]type4.Water_Line)

         fmt.Fprintf( w, "sky/pipe_line_ds_cal : water %v\n", water )
         fmt.Fprintf( w, "sky/pipe_line_ds_cal : value %v\n", water_line )

///
///         動水勾配線の計算
///

//    p_number ,ad_eneup ,ad_enedown ,ad_glineup ,ad_glinedown := cal.Pipe_line1( water  ,value  )
//      cal.Pipe_line1( water  ,value  )
      cal.Pipe_line1( water  ,water_line  )

//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : p_number %v\n", p_number )  // デバック
//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : ad_eneup %v\n", ad_eneup )  // デバック/
//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : ad_enedown %v\n", ad_enedown )  // デバック
//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : ad_glineup %v\n", ad_glineup )  // デバック
//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : ad_glinedown %v\n", ad_glinedown )  // デバック
//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : ad_eneup len %v\n", len(ad_eneup) )  // デバック

/// グラフの作成 ///

//    p_number := 1
//    ad_eneup := make([]type3.Point ,20 ,50)
//   ad_enedown := make([]type3.Point ,20 ,50)
//   ad_glineup := make([]type3.Point ,20 ,50)
//   ad_glinedown := make([]type3.Point ,20 ,50)

//   ad_eneup[0].X = 1.
//      ad_enedown[0].Y = 1.
//      ad_glineup[0].X = 1.
//      ad_glinedown[0].Y = 1.

//    f_name := cal.Pipe_line1_make_graf( w ,r ,p_number ,ad_eneup ,ad_enedown ,ad_glineup ,ad_glinedown  )

//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : f_name %v\n", f_name )  // デバック

/// グラフの表示 ///
//    f_name := "1"
//    cal.Pipe_line1_show_graf( w ,r ,f_name )

}

