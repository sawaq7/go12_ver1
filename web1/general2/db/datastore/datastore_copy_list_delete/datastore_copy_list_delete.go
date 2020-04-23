package datastore_copy_list_delete

import (

	"net/http"
	"strconv"
//	"fmt"
	"github.com/sawaq7/go12_ver1/general/process3"

	"cloud.google.com/go/datastore"

	"os"
	"context"

//	"github.com/sawaq7/go12_ver1/general/type5"
                                                          )

func Datastore_copy_list_delete (w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "datastore_copy_list_delete start \n" )



    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )

      projectID = "sample-7777"

	}

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

///
///       show copy list inf. on web
///

    process3.Datastore_copy_list_keyin(w , r )

}
