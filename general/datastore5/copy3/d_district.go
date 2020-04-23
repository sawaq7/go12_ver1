package copy3

import (

	    "net/http"
//	    "fmt"

        "github.com/sawaq7/go12_ver1/client/sgh/type2"

        "os"
        "cloud.google.com/go/datastore"
        "context"

                                                )

///
///     copy d.s (　name:D_District )
///

func D_district( w http.ResponseWriter, r *http.Request ,basic_name string ,copy_file string ,new_file string ) {

//     IN    w        : レスポンスライター
//     IN    r        : リクエストパラメータ
//     IN  basic_name : d.s. name of basic
//     IN  copy_file  : d.s. name which is cpied
//     IN  new_file   : ch is cpied

//    OUT  err        : error inf.

//    fmt.Fprintf( w, "copy3.d_district start \n" )
//    fmt.Fprintf( w, "copy3.d_district basic_name %v\n" ,basic_name)

///
///    get project name
///

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)

	query := datastore.NewQuery(copy_file)

    count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

                                      //   allocate work area for records
    ds_data := make([]type2.D_District, 0, count)

    if _, err := client.GetAll(ctx, query , &ds_data);  err != nil {

	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return

	} else{
      for _, ds_dataw := range ds_data {      //   put copy inf. for d.s.

        new_key := datastore.IncompleteKey( new_file , nil)

        if _, err = client.Put(ctx, new_key, &ds_dataw ); err != nil {

		  http.Error(w,err.Error(), http.StatusInternalServerError)
		  return

	    }
	  }
	}

//	fmt.Fprintf( w, "copy3.d_district normal end \n" )

    return
}

