package storage_object_list

import (

	    "net/http"
//	    "fmt"
	    "os"
	    "strconv"

        "storage2"
//        "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"

        "github.com/sawaq7/go12_ver1/general/type5"
        "github.com/sawaq7/go12_ver1/general/process3"
        "github.com/sawaq7/go12_ver1/general/datastore5/initialize"
//        "html/template"
        "log"
        "cloud.google.com/go/datastore"
        "context"

                                                  )

///
/// 繧ｪ繝悶ず繧ｧ繧ｯ繝医Μ繧ｹ繝医ｒ陦ｨ遉ｺ縺吶ｋ縲・///



func Storage_object_list(w http.ResponseWriter, r *http.Request) {

//    fmt.Printf( w, "storage_object_list start \n" )  // 繝・ヰ繝・け
// fmt.Printf("Saved %v: %v\n", taskKey, task.Description)

    var storage_b_o_temp type5.Storage_B_O_Temp

    var sdmy string

//    c := appengine.NewContext(r)
    ctx := context.Background()

/// 謖・ｮ壹＠縺殕ine-no繧竪ET縺励※謨ｴ謨ｰ蛹・///

    select_id , err := strconv.Atoi(r.FormValue("line_no"))
	if err  != nil {

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )  // 繝・ヰ繝・け

      projectID = "sample-7777"

	}

//	fmt.Fprintf( w, "storage_bucket_list :  projectID %v\n" ,  projectID  )  // 繝・ヰ繝・け

    buckets_minor , _ := storage2.Storage_basic( "list" ,projectID ,sdmy, w , r  )

    buckets, _ := buckets_minor.([]string)  // 繧､繝ｳ繧ｿ繝ｼ繝輔ぉ繧､繧ｹ蝙九ｒ蝙句､画鋤

//	buckets :=  storage2.Bucket_List ( w  ,r , projectID )

//	fmt.Fprintf( w, "storage_object_list : buckets: %v\n", buckets)
//	fmt.Fprintf( w, "storage_object_list : select_id: %v\n",select_id)

	for pos , bucketsw := range buckets {

      if select_id-1 == pos {

        client, err := datastore.NewClient(ctx, projectID)
        if err != nil {
            log.Fatalf("Failed to create client: %v", err)
        }

        initialize.Storage_b_o_temp (w , r ) //  譌｢蟄倥・縲Storage_B_O_Temp 繧ｳ繝｢繝ｳ逕ｨ縺ｮtemporary-file繧偵け繝ｪ繧｢繝ｼ

        storage_b_o_temp.Line_No =  1
        storage_b_o_temp.Project_Name = projectID
        storage_b_o_temp.Bucket_Name = bucketsw

/// 繧ｳ繝｢繝ｳ逕ｨ縺ｮtemporary-file縺ｫ繝舌こ繝・ヨ蜷阪ｒ蜀阪そ繝・ヨ

        new_key := datastore.IncompleteKey("Storage_B_O_Temp", nil)

        if _, err := client.Put(ctx, new_key, &storage_b_o_temp ); err != nil {
                log.Fatalf("Failed to save task: %v", err)
        }


//	    if _, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "Storage_B_O_Temp", nil) , &storage_b_o_temp); err != nil {
//		  http.Error(w,err.Error(), http.StatusInternalServerError)
//		  return
//	    }

	    /// 繝｢繝九ち繝ｼ縲蜀崎｡ｨ遉ｺ ///

        process3.Storage_object_show ( w , r ,projectID  ,bucketsw )

	  }

	}

//	fmt.Fprintf( w, "storage_object_list : normal end \n" )  // 繝・ヰ繝・け

}

