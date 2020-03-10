package initialize3

import (

//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
//	    "html/template"

	    "client/reserve/type6"

	    "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                )

///                                   ///
/// delete all data of temporary file ///
///                                   ///

func Guest_temp(w http.ResponseWriter, r *http.Request )   {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ

//    fmt.Fprintf( w, "init3/guest_temp start \n" )  // デバック

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return
    }

    query := datastore.NewQuery("Guest_Temp").Order("Guest_No")
//	q := datastore.NewQuery("Guest_Temp").Order("Guest_No")

    count, err := client.Count(ctx, query)
//	count, err := q.Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

//    fmt.Fprintf( w, "init3/guest_temp count \n" ,count )  // デバック

	guest_temp     := make([]type6.Guest_Temp, 0, count)

	keys, err := client.GetAll(ctx, query , &guest_temp)
//	keys, err := q.GetAll(c, &guest_temp )
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
    }

//	for pos2, _ := range guest_temp {
    for _, keysw := range keys {

/// 一時ファイルの削除

      if err := client.Delete(ctx, datastore.IDKey("Guest_Temp", keysw.ID, nil)); err != nil {

//      key := datastore.NewKey(c, "Guest_Temp"       , "", keys[pos2].IntID(), nil)
//	  if err := datastore.Delete(c, key  ); err != nil {

		  http.Error(w, err.Error(), http.StatusInternalServerError)
		  return
	  }
//      fmt.Fprintf( w, "init3/guest_temp pos2 %v   \n" , pos2  )  // デバック

    }
	return
}
