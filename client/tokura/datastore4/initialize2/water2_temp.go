package initialize2

import (

	    "net/http"

//	    "fmt"

	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
//	    "time"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                )

///
///            delete all data of temporary file
///

func Water2_temp(w http.ResponseWriter, r *http.Request )   {

//     IN    w      : レスポンスライター
//     IN    r      : リクエストパラメータ

//    fmt.Fprintf( w, "init/water2_temp start \n" )

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return
    }

    query := datastore.NewQuery("Water2_Temp").Order("Name")

    count, err := client.Count(ctx, query)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

//    fmt.Fprintf( w, "init/water2_temp count \n" ,count )

	water2_temp     := make([]type4.Water2_Temp, 0, count)

	keys, err := client.GetAll(ctx, query , &water2_temp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
    }

	for _, keysw := range keys {

/// 一時ファイルの削除

      if err := client.Delete(ctx, datastore.IDKey("Water2_Temp", keysw.ID, nil)); err != nil {

		  http.Error(w, err.Error(), http.StatusInternalServerError)
		  return
	  }
//      fmt.Fprintf( w, "init/water2_temp pos2 %v   \n" , pos2  )

    }
	return
}

