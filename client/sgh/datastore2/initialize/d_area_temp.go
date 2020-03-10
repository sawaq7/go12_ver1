package initialize

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
//	    "html/template"
//	    "web/htmls/sgh"
	    "client/sgh/type2"
//	    "time"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                )

///                                   ///
/// delete all data of temporary file ///
///                                   ///

func D_area_temp(w http.ResponseWriter, r *http.Request )   {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ

//    fmt.Fprintf( w, "init/D_area_temp start \n" )  // デバック
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

    query := datastore.NewQuery("D_Area_Temp").Order("District_No")
//	q := datastore.NewQuery("D_Area_Temp").Order("District_No")

//	count, err := q.Count(c)
	count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

//    fmt.Fprintf( w, "init/D_area_temp count \n" ,count )  // デバック

	d_area_temp     := make([]type2.D_Area_Temp, 0, count)

	keys, err := client.GetAll(ctx, query , &d_area_temp)
//	keys, err := q.GetAll(c, &d_area_temp )
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
    }
//	for pos2, _ := range d_area_temp {
	for _, keysw := range keys {

/// 一時ファイルの削除

      if err := client.Delete(ctx, datastore.IDKey("D_Area_Temp", keysw.ID, nil)); err != nil {

//      key := datastore.NewKey(c, "D_Area_Temp"       , "", keys[pos2].IntID(), nil)
//	  if err := datastore.Delete(c, key  ); err != nil {

		  http.Error(w, err.Error(), http.StatusInternalServerError)
		  return
	  }
//      fmt.Fprintf( w, "init/d_area_temp pos2 %v   \n" , pos2  )  // デバック

    }
	return
}

