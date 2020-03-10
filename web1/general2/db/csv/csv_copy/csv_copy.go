package csv_copy

import (

//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"

	    "general/process3"

	    "general/type5"
	    "strconv"
//	    "time"

        "os"

        "cloud.google.com/go/datastore"
        "context"
                                                  )

func Csv_copy(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "csv_copy start \n" )  // デバック

    var csv_inf type5.Csv_Inf

///
///   プロジェクト名をゲット
///
    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )  // デバック

      project_name = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return
    }

	id := r.FormValue("id")
//    fmt.Fprintf( w, "csv_copy  : id %v\n", id )  // デバック

	copy_idw ,_ := strconv.Atoi(id)
	copy_id := int64(copy_idw)

//    fmt.Fprintf( w, "csv_copy  : copy_idw %v\n", copy_idw )  // デバック
//    fmt.Fprintf( w, "csv_copy  : copy_id %v\n", copy_id )  // デバック



//	key := datastore.NewKey(c, "Csv_Inf", "", copy_id, nil)

    key := datastore.IDKey("Csv_Inf", copy_id, nil)

    if err := client.Get(ctx, key , &csv_inf ) ; err != nil {
//	if  err := datastore.Get(c, key,  &csv_inf); err != nil {

		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

    new_key := datastore.IncompleteKey("Csv_Inf", nil)

    if _, err := client.Put(ctx, new_key, &csv_inf ); err != nil {
      http.Error(w,err.Error(), http.StatusInternalServerError)
    }

//    if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Csv_Inf", nil), &csv_inf ); err != nil {
//		http.Error(w,err.Error(), http.StatusInternalServerError)
//		return
//	}

///
///    モニター　再表示
///

	process3.Csv_inf ( w , r )

//	fmt.Fprintf( w, "csv_copy : normal end \n" )  // デバック




}
