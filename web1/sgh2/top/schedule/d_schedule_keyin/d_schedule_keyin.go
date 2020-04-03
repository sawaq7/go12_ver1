package d_schedule_keyin

import (

	    "net/http"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2"
  	    "strconv"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/client/sgh/type2"

	    "cloud.google.com/go/datastore"
        "context"
        "os"
                                                 )

///
/// 謖・ｮ壹＠縺溷慍蝓溘・繧ｹ繧ｱ繧ｸ繝･繝ｼ繝ｫ繧定｡ｨ遉ｺ
///

func D_schedule_keyin(w http.ResponseWriter, r *http.Request) {

    var idmy int64

    var g type2.D_District

    var g2 type2.D_District_Temp

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

    ctx := context.Background()
//	c := appengine.NewContext(r)

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return
    }


/// 蝨ｰ蛹ｺ諠・ｱ繝輔ぃ繧､繝ｫ縺九ｉ蝨ｰ蛹ｺNO 繧竪ET

	idw , err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    id := int64(idw)
//    fmt.Fprintf( w, "sky d_schedule_keyin : id %v\n", id )  // 繝・ヰ繝・け

    key := datastore.IDKey("D_District", id, nil)

    if err := client.Get(ctx, key , &g ) ; err != nil {
//    key := datastore.NewKey(c, "D_District", "", id, nil)
//    if err := datastore.Get(c, key ,&g); err!= nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

///
/// 繧ｫ繝ｬ繝ｳ繝医・蝨ｰ蛹ｺ諠・ｱ繧偵い繝・・繝・・繝・///

    g2.District_No   = g.District_No
    g2.District_Name = g.District_Name
//    fmt.Fprintf( w, "sky d_schedule_keyin : g2.District_No %v\n", g2.District_No )  // 繝・ヰ繝・け

// temporary-file繧偵う繝九す繝｣繝ｩ繧､繧ｺ

    _ = datastore2.Datastore_sgh( "D_District_Temp" ,"initialize" ,idmy , w , r  )

     new_key := datastore.IncompleteKey("D_District_Temp", nil)

    if _, err = client.Put(ctx, new_key, &g2 ); err != nil {
//	if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "D_District_Temp", nil), &g2); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

/// 繧ｹ繧ｱ繧ｸ繝･繝ｼ繝ｫ繧定｡ｨ遉ｺ ///

   process.D_schedule_showall(w , r , g.District_No )

}

