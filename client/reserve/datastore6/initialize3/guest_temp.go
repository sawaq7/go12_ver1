package initialize3

import (

	    "net/http"
//	    "fmt"
//	    "html/template"

	    "github.com/sawaq7/go12_ver1/client/reserve/type6"

	    "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                )

///                                   ///
/// delete all data of temporary file ///
///                                   ///

func Guest_temp(w http.ResponseWriter, r *http.Request )   {

//     IN    w      　　: response-writer
//     IN    r      　　: request-parameter

//    fmt.Fprintf( w, "init3/guest_temp start \n" )

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

    query := datastore.NewQuery("Guest_Temp").Order("Guest_No")

    count, err := client.Count(ctx, query)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

//    fmt.Fprintf( w, "init3/guest_temp count \n" ,count )

	guest_temp     := make([]type6.Guest_Temp, 0, count)

	keys, err := client.GetAll(ctx, query , &guest_temp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
    }

    for _, keysw := range keys {

/// 一時ファイルの削除

      if err := client.Delete(ctx, datastore.IDKey("Guest_Temp", keysw.ID, nil)); err != nil {

		  http.Error(w, err.Error(), http.StatusInternalServerError)
		  return
	  }
//      fmt.Fprintf( w, "init3/guest_temp pos2 %v   \n" , pos2  )

    }
	return
}
