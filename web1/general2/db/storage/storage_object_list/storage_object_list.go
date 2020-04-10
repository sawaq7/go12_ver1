package storage_object_list

import (

	    "net/http"
//	    "fmt"
	    "os"
	    "strconv"

        "github.com/sawaq7/go12_ver1/storage2"

        "github.com/sawaq7/go12_ver1/general/type5"
        "github.com/sawaq7/go12_ver1/general/process3"
        "github.com/sawaq7/go12_ver1/general/datastore5/initialize"

        "log"
        "cloud.google.com/go/datastore"
        "context"

                                                  )

///
/// オブジェクトリストを表示する
///


func Storage_object_list(w http.ResponseWriter, r *http.Request) {

    var storage_b_o_temp type5.Storage_B_O_Temp

    var sdmy string

    ctx := context.Background()

    select_id , err := strconv.Atoi(r.FormValue("line_no"))  ///    get line no.
	if err  != nil {

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )

      projectID = "sample-7777"

	}

//	fmt.Fprintf( w, "storage_bucket_list :  projectID %v\n" ,  projectID  )

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

        initialize.Storage_b_o_temp (w , r ) //  clear Storage_B_O_Temp

        storage_b_o_temp.Line_No =  1
        storage_b_o_temp.Project_Name = projectID
        storage_b_o_temp.Bucket_Name = bucketsw

///
///   set new Storage_B_O_Temp
///

        new_key := datastore.IncompleteKey("Storage_B_O_Temp", nil)

        if _, err := client.Put(ctx, new_key, &storage_b_o_temp ); err != nil {
                log.Fatalf("Failed to save task: %v", err)
        }

///
///     show web
///

        process3.Storage_object_show ( w , r ,projectID  ,bucketsw )

	  }

	}

//	fmt.Fprintf( w, "storage_object_list : normal end \n" )

}

