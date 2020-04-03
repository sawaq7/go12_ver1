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
/// チE�Eタストアーから水路ファイル惁E��をGETする�E�水路ファイル・チE�Eタストア�E�E///                          　　　　　　　　　　　

func Water2( w http.ResponseWriter, r *http.Request )  ([]type4.Water2 ) {

//     IN     w         : レスポンスライター
//     IN     r         : リクエストパラメーター
//     OUT        　　  : 水路ファイル・チE�Eタストアのスライス

//    fmt.Fprintf( w, "trans2.water2 start \n" )  // チE��チE��

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )  // チE��チE��

      project_name = "sample-7777"

	}
    ctx := context.Background()
//    c := appengine.NewContext(r)

	query := datastore.NewQuery("Water2").Order("Name")
//	q := datastore.NewQuery("Water2").Order("Name")

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return  nil
    }

//	count, err := q.Count(c)
    count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

	water2      := make([]type4.Water2, 0, count)

	water2_view := make([]type4.Water2, 0)

//	keys, err := q.GetAll(c, &water2)
	keys, err := client.GetAll(ctx, query , &water2)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
//		fmt.Fprintf( w, "water2 err \n" ,err)  // チE��チE��
		return	nil
	}

	keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	for pos, water2w := range water2 {

///  機�EによりチェチE��頁E��をセチE��

         water2_view = append(water2_view, type4.Water2 {           keys_wk[pos]      ,
                                                                    water2w.Name   ,
                                                                    water2w.High   ,
                                                                    water2w.Roughness_Factor   })

	}

    return	water2_view
}

