package copy3

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"

        "github.com/sawaq7/go12_ver1/client/sgh/type2"

        "os"
        "cloud.google.com/go/datastore"
        "context"

                                                )

///
/// 繝・・繧ｿ繧ｹ繝医い繧偵さ繝斐・縺吶ｋ  (縲繝吶・繧ｷ繝・け蜷阪・壹D_District縲・・///

func D_district( w http.ResponseWriter, r *http.Request ,basic_name string ,copy_file string ,new_file string ) {

//     IN    w        : 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r        : 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ
//     IN  basic_name : 蝓ｺ譛ｬ縺ｮ繝・・繧ｿ繧ｹ繝医い蜷・//     IN  copy_file  : 繧ｳ繝斐・蜈・・繝・・繧ｿ繧ｹ繝医い蜷・//     IN  new_file   : 繝九Η繝ｼ繝・・繧ｿ繧ｹ繝医い蜷・//    OUT  err        : 繧ｨ繝ｩ繝ｼ繝｡繝・そ繝ｼ繧ｸ

//    fmt.Fprintf( w, "copy3.d_district start \n" )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "copy3.d_district basic_name %v\n" ,basic_name)  // 繝・ヰ繝・け

///
///  繝励Ο繧ｸ繧ｧ繧ｯ繝・D縲繧ｲ繝・ヨ
///
    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)

	query := datastore.NewQuery(copy_file)

//	count, err := q.Count(c)
    count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

//	q := datastore.NewQuery(copy_file) /// 繧ｯ繝ｭ繝ｼ繝ｳ縺ｮ繝ｪ繝ｼ繝繝ｼ繧竪ET

//    fmt.Fprintf( w, "copy3.d_district count %v\n" ,count)  // 繝・ヰ繝・け

///
/// 繧ｯ繝ｭ繝ｼ繝ｳ諠・ｱ繧担ET縺吶ｋ繝ｯ繝ｼ繧ｯ繧ｨ繝ｪ繧｢繧堤｢ｺ菫・///

    ds_data := make([]type2.D_District, 0, count)


    if _, err := client.GetAll(ctx, query , &ds_data);  err != nil {
//	if _, err := q.GetAll(c, &ds_data);  err != nil {         // 繧ｯ繝ｭ繝ｼ繝ｳ諠・ｱ繧竪ET

	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return

	} else{
      for _, ds_dataw := range ds_data {                       //縲繝九Η繝ｼ繝輔ぃ繧､繝ｫ縺ｫ繧ｯ繝ｭ繝ｼ繝ｳ諠・ｱ繧偵そ繝・ヨ

//	    if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, new_file, nil), &ds_dataw); err != nil {
        new_key := datastore.IncompleteKey( new_file , nil)

        if _, err = client.Put(ctx, new_key, &ds_dataw ); err != nil {

		  http.Error(w,err.Error(), http.StatusInternalServerError)
		  return

	    }
	  }
	}

//	fmt.Fprintf( w, "copy3.d_district normal end \n" )  // 繝・ヰ繝・け

    return
}

