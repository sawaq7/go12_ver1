package initialize

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
        "cloud.google.com/go/datastore"

	    "net/http"
//	    "fmt"

	    "general/type5"
	    "context"
	    "os"

                                                )

///                                   ///
/// delete all data of temporary file ///
///                                   ///

func Storage_b_o_temp(w http.ResponseWriter, r *http.Request )   {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ

//    fmt.Fprintf( w, "init/storage_b_o_temp start \n" )  // デバック

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)

	query := datastore.NewQuery("Storage_B_O_Temp")

//	count, err := q.Count(c)
    count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

//    fmt.Fprintf( w, "init/storage_b_o_temp count \n" ,count )  // デバック

	storage_b_o_temp     := make([]type5.Storage_B_O_Temp, 0, count)

	keys, err := client.GetAll(ctx, query , &storage_b_o_temp)

//	keys, err := q.GetAll(ctx, &storage_b_o_temp )
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
    }

	for _, keysw := range keys {

      if err := client.Delete(ctx, datastore.IDKey("Storage_B_O_Temp", keysw.ID, nil)); err != nil {

		 http.Error(w, err.Error(), http.StatusInternalServerError)
		 return
	  }

    }

	return

}

