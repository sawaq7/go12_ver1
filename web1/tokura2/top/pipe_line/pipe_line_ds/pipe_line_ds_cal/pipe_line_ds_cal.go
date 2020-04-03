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

//   fmt.Fprintf( w, "sky/pipe_line_ds_cal start \n"  )  // 繝・ヰ繝・け

/// key-in 繝・・繧ｿ繧竪ET ///

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

//    fmt.Fprintf( w, "sky/pipe_line_ds_cal  \n" ,count )  // 繝・ヰ繝・け

    water_temp      := make([]type4.Water2, 0, count)

    keys, err := client.GetAll(ctx, query , &water_temp)
//	keys, err := q2.GetAll(c, &water_temp )
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
    }

//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : len(water) %v\n", len(water) )  // 繝・ヰ繝・け

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
///      豌ｴ霍ｯ繝ｩ繧､繝ｳ縺ｮ繝・・繧ｿ繧偵ご繝・ヨ
///

//    water_line := trans2.Water_line (1  ,water.Name , w ,r )

      water_line := datastore4.Datastore_tokura( "Water_Line"  ,"trans"  ,water.Name , w , r  )



     value, _ := water_line.([]type4.Water_Line)    // 遨ｺ繧､繝ｳ繧ｿ繝ｼ繝輔ぉ繧､繧ｹ螟画焚繧医ｊ繝舌Μ繝･繝ｼ蛟､繧偵ご繝・ヨ

///
///         蜍墓ｰｴ蜍ｾ驟咲ｷ壹・險育ｮ・///

    p_number ,ad_eneup ,ad_enedown ,ad_glineup ,ad_glinedown := cal.Pipe_line1( water  ,value  )

//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : p_number %v\n", p_number )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : ad_eneup %v\n", ad_eneup )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : ad_enedown %v\n", ad_enedown )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : ad_glineup %v\n", ad_glineup )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : ad_glinedown %v\n", ad_glinedown )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : ad_eneup len %v\n", len(ad_eneup) )  // 繝・ヰ繝・け

/// 繧ｰ繝ｩ繝輔・菴懈・ ///

    f_name := cal.Pipe_line1_make_graf( w ,r ,p_number ,ad_eneup ,ad_enedown ,ad_glineup ,ad_glinedown  )

//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : f_name %v\n", f_name )  // 繝・ヰ繝・け

/// 繧ｰ繝ｩ繝輔・陦ｨ遉ｺ ///

    cal.Pipe_line1_show_graf( w ,r ,f_name )

}

