package datastore_copy_list_delete

import (

	"net/http"
	"strconv"
//	"fmt"
	"general/process3"

	"cloud.google.com/go/datastore"

	"os"
	"context"

//	"general/type5"
                                            )

func Datastore_copy_list_delete (w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "datastore_copy_list_delete start \n" )  // デバック



    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )  // デバック

      projectID = "sample-7777"

	}

//	c := appengine.NewContext(r)
	ctx := context.Background()

	delidw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    delid := int64(delidw)

//	     _ = line_no2

    client, err := datastore.NewClient(ctx, projectID)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    if err := client.Delete(ctx, datastore.IDKey("Ds_Copy_List", delid, nil)); err != nil {

       http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
//	  if err := datastore.Delete(c, key  ); err != nil {

//         fmt.Fprintf( w, "datastore_copy_list_delete : err %v\n", err )  // デバック

//		 http.Error(w, err.Error(), http.StatusInternalServerError)
//		 return
//	  }

///
///     モニター　再表示
///

//    process3.Db_access_list(w , r )
    process3.Datastore_copy_list_keyin(w , r )

}
