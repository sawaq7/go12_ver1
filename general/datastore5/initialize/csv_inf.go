package initialize

import (

	    "net/http"
//	    "fmt"
        "cloud.google.com/go/datastore"


	    "general/type5"
	    "context"
	    "os"

                                                )

///                                   ///
/// delete all data of temporary file ///
///                                   ///

func Csv_inf(w http.ResponseWriter, r *http.Request )   {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ

//    fmt.Fprintf( w, "init/csv_inf start \n" )  // デバック

   projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)

	query := datastore.NewQuery("Csv_Inf")

    count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	csv_inf     := make([]type5.Csv_Inf, 0, count)

    keys, err := client.GetAll(ctx, query , &csv_inf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
    }

	for _, keysw := range keys {

      if err := client.Delete(ctx, datastore.IDKey("Csv_Inf", keysw.ID, nil)); err != nil {

		 http.Error(w, err.Error(), http.StatusInternalServerError)
		 return
	  }

    }
	return
}

