package trans2

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"

	    "client/tokura/suiri/type4"

	    "cloud.google.com/go/datastore"
        "context"
        "os"
                                                )

///                           　　　　　　　　　　　
/// データストアーから導水勾配線の情報をGETする（導水勾配線ファイル）
///                          　　　　　　　　　　　

func Water_slope( w http.ResponseWriter, r *http.Request )  ([]type4.Water_Slope ) {

//     IN     w         : レスポンスライター
//     IN     r         : リクエストパラメーター
//     OUT        　　  : 導水勾配線のスライス

//    fmt.Fprintf( w, "trans2.Water_slope start \n" )  // デバック

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)

    query := datastore.NewQuery("Water_Slope").Order("File_Name")
//	q := datastore.NewQuery("Water_Slope").Order("File_Name")

    count, err := client.Count(ctx, query)
//	count, err := q.Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

	water_slope      := make([]type4.Water_Slope, 0, count)

	water_slope_view := make([]type4.Water_Slope, 0)

    keys, err := client.GetAll(ctx, query , &water_slope)
//	keys, err := q.GetAll(c, &water_slope)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
//		fmt.Fprintf( w, "water_slope err \n" ,err)  // デバック
		return	nil
	}

	keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	for pos, water_slopew := range water_slope {

///  機能によりチェック項目をセット

         water_slope_view = append(water_slope_view, type4.Water_Slope { keys_wk[pos]       ,
                                                                         water_slopew.File_Name  ,
                                                                         water_slopew.Url          })


	}

    return	water_slope_view
}

