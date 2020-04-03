package storage_object_delete

import (

	"net/http"
	"strconv"
//	"fmt"

	"github.com/sawaq7/go12_ver1/general/type5"
	"github.com/sawaq7/go12_ver1/general/process3"
	"storage2"
//	"github.com/sawaq7/go12_ver1/storage2/get"

    "cloud.google.com/go/datastore"

	"os"
	"context"
	"log"
                                            )

///
///   謖・ｮ壹＠縺溘が繝悶ず繧ｧ繧ｯ繝医ｒ蜑企勁縺吶ｋ
///

func Storage_object_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "storage_object_delete start \n" )  // 繝・ヰ繝・け

    var bucket ,filename ,project string

    project = os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project == "" {

      project = "sample-7777"

	}

	ctx := context.Background()

    line_no := r.FormValue("line_no")

//    fmt.Fprintf( w, "storage_object_show : line_no %v\n", line_no )  // 繝・ヰ繝・け

	select_id ,_ := strconv.Atoi(line_no)

//    fmt.Fprintf( w, "storage_object_show : select_id %v\n", select_id )  // 繝・ヰ繝・け

//
///   繝舌こ繝・ヨ蜷阪ｒ繧ｲ繝・ヨ
///

	query := datastore.NewQuery("Storage_B_O_Temp")

    client, err := datastore.NewClient(ctx, project)
    if err != nil {
       log.Fatalf("Failed to create client: %v", err)
    }

    count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	storage_b_o_temp      := make([]type5.Storage_B_O_Temp, 0, count)

    if _, err = client.GetAll(ctx, query , &storage_b_o_temp) ; err != nil {

	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return

	} else{

      for pos, storage_b_o_tempw := range storage_b_o_temp {

        if pos == 0 {
           project   = storage_b_o_tempw.Project_Name    // 繝励Ο繧ｸ繧ｧ繧ｯ繝亥錐繧偵ご繝・ヨ
           bucket    = storage_b_o_tempw.Bucket_Name    // 繝舌こ繝・ヨ蜷阪ｒ繧ｲ繝・ヨ

        }
	  }
	}

///
///   繝輔ぃ繧､繝ｫ蜷阪ｒ繧ｲ繝・ヨ
///

	objects :=  storage2.Object_List ( w  ,r , bucket )

//    fmt.Fprintf( w, "storage_object_show : objects: %v\n", objects )

    for pos2 , objectsw := range objects {

      if pos2+1 == select_id {

        filename = objectsw

      }

    }

//	storage2.File_Delete ( w , r  ,bucket ,filename  )

	storage2.Storage_basic( "delete" ,bucket ,filename , w , r  )

///
/// 繝｢繝九ち繝ｼ縲陦ｨ遉ｺ ///
///

    process3.Storage_object_show ( w , r ,project  ,bucket )

}
