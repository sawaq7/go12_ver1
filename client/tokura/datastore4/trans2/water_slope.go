package trans2

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"

	    "cloud.google.com/go/datastore"
        "context"
        "os"
                                                )

///                           　　　　　　　　　　　
/// チE�Eタストアーから導水勾配線�E惁E��をGETする�E�導水勾配線ファイル�E�E///                          　　　　　　　　　　　

func Water_slope( w http.ResponseWriter, r *http.Request )  ([]type4.Water_Slope ) {

//     IN     w         : レスポンスライター
//     IN     r         : リクエストパラメーター
//     OUT        　　  : 導水勾配線�Eスライス

//    fmt.Fprintf( w, "trans2.Water_slope start \n" )  // チE��チE��

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
//		fmt.Fprintf( w, "water_slope err \n" ,err)  // チE��チE��
		return	nil
	}

	keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	for pos, water_slopew := range water_slope {

///  機�EによりチェチE��頁E��をセチE��

         water_slope_view = append(water_slope_view, type4.Water_Slope { keys_wk[pos]       ,
                                                                         water_slopew.File_Name  ,
                                                                         water_slopew.Url          })


	}

    return	water_slope_view
}

