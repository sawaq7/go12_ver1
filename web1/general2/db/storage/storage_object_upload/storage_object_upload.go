package storage_object_upload

import (

//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
        "cloud.google.com/go/datastore"
        "cloud.google.com/go/storage"

	    "io"
	    "context"
	    "log"

        "github.com/sawaq7/go12_ver1/general/type5"
        "github.com/sawaq7/go12_ver1/general/process3"
	    "storage2"

	    "os"
	    "net/http"
//	    "fmt"

                                                  )


func Storage_object_upload(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "storage_object_upload start \n" )  // 繝・ヰ繝・け

    var bucket string

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

//    c := appengine.NewContext(r)
    ctx := context.Background()

	query := datastore.NewQuery("Storage_B_O_Temp")

	client, err := datastore.NewClient(ctx, projectID)
    if err != nil {
       log.Fatalf("Failed to create client: %v", err)
    }

//	count, err := q.Count(c)
    count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	storage_b_o_temp      := make([]type5.Storage_B_O_Temp, 0, count)

    if _, err = client.GetAll(ctx, query , &storage_b_o_temp) ; err != nil {
//	if _, err := q.GetAll(c, &storage_b_o_temp);  err != nil {

	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return

	} else{

      for pos, storage_b_o_tempw := range storage_b_o_temp {

        if pos == 0 {
           bucket    = storage_b_o_tempw.Bucket_Name    // 繝舌こ繝・ヨ蜷阪ｒ繧ｲ繝・ヨ
           projectID = storage_b_o_tempw.Project_Name   // 繝励Ο繧ｸ繧ｧ繧ｯ繝亥錐繧偵ご繝・ヨ

        }
	  }
	}

	file_data, fh, err := r.FormFile("image")

	if err != nil {
		return

	}

	w2_minor , _ := storage2.Storage_basic( "create" ,bucket ,fh.Filename , w , r  )

    w2, _ := w2_minor.(*storage.Writer)  // 繧､繝ｳ繧ｿ繝ｼ繝輔ぉ繧､繧ｹ蝙九ｒ蝙句､画鋤

//	w2 := storage2.File_Create ( w ,r ,bucket  ,fh.Filename )


/// 繧ｹ繝医Ξ繝・ず繝輔ぃ繧､繝ｫ縺ｫ譌｢蟄倥・繝輔ぃ繧､繝ｫ縺ｮ諠・ｱ繧偵さ繝斐・縲縲///

	if _, err := io.Copy(w2, file_data); err != nil {
	    http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := w2.Close(); err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

/// 繝｢繝九ち繝ｼ縲蜀崎｡ｨ遉ｺ ///

    process3.Storage_object_show ( w , r ,projectID  ,bucket )

//	fmt.Fprintf( w, "storage_object_upload : normal end \n" )  // 繝・ヰ繝・け

}


