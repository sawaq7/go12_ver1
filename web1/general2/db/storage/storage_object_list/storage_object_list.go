package storage_object_list

import (

	    "net/http"
//	    "fmt"
	    "os"
	    "strconv"

        "storage2"
//        "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"

        "general/type5"
        "general/process3"
        "general/datastore5/initialize"
//        "html/template"
        "log"
        "cloud.google.com/go/datastore"
        "context"

                                                  )

///
/// オブジェクトリストを表示する。
///



func Storage_object_list(w http.ResponseWriter, r *http.Request) {

//    fmt.Printf( w, "storage_object_list start \n" )  // デバック
// fmt.Printf("Saved %v: %v\n", taskKey, task.Description)

    var storage_b_o_temp type5.Storage_B_O_Temp

    var sdmy string

//    c := appengine.NewContext(r)
    ctx := context.Background()

/// 指定したline-noをGETして整数化 ///

    select_id , err := strconv.Atoi(r.FormValue("line_no"))
	if err  != nil {

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )  // デバック

      projectID = "sample-7777"

	}

//	fmt.Fprintf( w, "storage_bucket_list :  projectID %v\n" ,  projectID  )  // デバック

    buckets_minor , _ := storage2.Storage_basic( "list" ,projectID ,sdmy, w , r  )

    buckets, _ := buckets_minor.([]string)  // インターフェイス型を型変換

//	buckets :=  storage2.Bucket_List ( w  ,r , projectID )

//	fmt.Fprintf( w, "storage_object_list : buckets: %v\n", buckets)
//	fmt.Fprintf( w, "storage_object_list : select_id: %v\n",select_id)

	for pos , bucketsw := range buckets {

      if select_id-1 == pos {

        client, err := datastore.NewClient(ctx, projectID)
        if err != nil {
            log.Fatalf("Failed to create client: %v", err)
        }

        initialize.Storage_b_o_temp (w , r ) //  既存の　Storage_B_O_Temp コモン用のtemporary-fileをクリアー

        storage_b_o_temp.Line_No =  1
        storage_b_o_temp.Project_Name = projectID
        storage_b_o_temp.Bucket_Name = bucketsw

/// コモン用のtemporary-fileにバケット名を再セット

        new_key := datastore.IncompleteKey("Storage_B_O_Temp", nil)

        if _, err := client.Put(ctx, new_key, &storage_b_o_temp ); err != nil {
                log.Fatalf("Failed to save task: %v", err)
        }


//	    if _, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "Storage_B_O_Temp", nil) , &storage_b_o_temp); err != nil {
//		  http.Error(w,err.Error(), http.StatusInternalServerError)
//		  return
//	    }

	    /// モニター　再表示 ///

        process3.Storage_object_show ( w , r ,projectID  ,bucketsw )

	  }

	}

//	fmt.Fprintf( w, "storage_object_list : normal end \n" )  // デバック

}

