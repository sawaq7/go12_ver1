package initialize

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
//	    "html/template"
//	    "web/htmls/sgh"
	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
//	    "time"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                )

///                                   ///
/// delete all data of temporary file ///
///                                   ///

func D_area_temp(w http.ResponseWriter, r *http.Request )   {

//     IN    w      縲縲: 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r      縲縲: 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ

//    fmt.Fprintf( w, "init/D_area_temp start \n" )  // 繝・ヰ繝・け
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

//    fmt.Fprintf( w, "init/D_area_temp count \n" ,count )  // 繝・ヰ繝・け

	d_area_temp     := make([]type2.D_Area_Temp, 0, count)

	keys, err := client.GetAll(ctx, query , &d_area_temp)
//	keys, err := q.GetAll(c, &d_area_temp )
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
    }
//	for pos2, _ := range d_area_temp {
	for _, keysw := range keys {

/// 荳譎ゅヵ繧｡繧､繝ｫ縺ｮ蜑企勁

      if err := client.Delete(ctx, datastore.IDKey("D_Area_Temp", keysw.ID, nil)); err != nil {

//      key := datastore.NewKey(c, "D_Area_Temp"       , "", keys[pos2].IntID(), nil)
//	  if err := datastore.Delete(c, key  ); err != nil {

		  http.Error(w, err.Error(), http.StatusInternalServerError)
		  return
	  }
//      fmt.Fprintf( w, "init/d_area_temp pos2 %v   \n" , pos2  )  // 繝・ヰ繝・け

    }
	return
}

