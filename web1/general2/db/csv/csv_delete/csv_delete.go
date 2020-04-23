package csv_delete

import (

	    "net/http"
//	    "fmt"

        "github.com/sawaq7/go12_ver1/general/process3"

	    "strconv"

//	    "time"
        "os"

        "cloud.google.com/go/datastore"
        "context"
                                                  )

func Csv_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "csv_delete start \n" )

///
///   get project name
///

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )
      project_name = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return
    }

    delidw , err := strconv.Atoi(r.FormValue("id"))

	if err  != nil {
//	   fmt.Fprintf( w, "csv_delete :error delidw %v\n", delidw )
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    delid := int64(delidw)

//    fmt.Fprintf( w, "csv_delete : delidw %v\n", delidw )
//    fmt.Fprintf( w, "csv_delete : delid %v\n", delid )

	key := datastore.IDKey("Csv_Inf", delid, nil)

//    delete one record in d.s.

    if err := client.Delete(ctx, key); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


/// show csv inf. on web

	process3.Csv_inf ( w , r )

//	fmt.Fprintf( w, "csv_delete : normal end \n" )




}
