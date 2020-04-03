package initialize2

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
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

//     IN    w      : 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r      : 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ

//    fmt.Fprintf( w, "init/water2_temp start \n" )  // 繝・ヰ繝・け

///
///   繝励Ο繧ｸ繧ｧ繧ｯ繝亥錐繧偵ご繝・ヨ
///
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

    query := datastore.NewQuery("Water2_Temp").Order("Name")
//    q := datastore.NewQuery("Water2_Temp").Order("Name")

    count, err := client.Count(ctx, query)
//	count, err := q.Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

//    fmt.Fprintf( w, "init/water2_temp count \n" ,count )  // 繝・ヰ繝・け

	water2_temp     := make([]type4.Water2_Temp, 0, count)

	keys, err := client.GetAll(ctx, query , &water2_temp)
//	keys, err := q.GetAll(c, &water2_temp )
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
    }

//	for pos2, _ := range water2_temp {
	for _, keysw := range keys {

/// 荳譎ゅヵ繧｡繧､繝ｫ縺ｮ蜑企勁

//      key := datastore.NewKey(c, "Water2_Temp"       , "", keys[pos2].IntID(), nil)

      if err := client.Delete(ctx, datastore.IDKey("Water2_Temp", keysw.ID, nil)); err != nil {
//	  if err := datastore.Delete(c, key  ); err != nil {

		  http.Error(w, err.Error(), http.StatusInternalServerError)
		  return
	  }
//      fmt.Fprintf( w, "init/water2_temp pos2 %v   \n" , pos2  )  // 繝・ヰ繝・け

    }
	return
}

