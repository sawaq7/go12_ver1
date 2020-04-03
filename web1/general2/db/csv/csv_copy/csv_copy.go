package csv_copy

import (

//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/general/process3"

	    "github.com/sawaq7/go12_ver1/general/type5"
	    "strconv"
//	    "time"

        "os"

        "cloud.google.com/go/datastore"
        "context"
                                                  )

func Csv_copy(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "csv_copy start \n" )  // 繝・ヰ繝・け

    var csv_inf type5.Csv_Inf

///
///   繝励Ο繧ｸ繧ｧ繧ｯ繝亥錐繧偵ご繝・ヨ
///
    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )  // 繝・ヰ繝・け

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
//    fmt.Fprintf( w, "csv_copy  : id %v\n", id )  // 繝・ヰ繝・け

	copy_idw ,_ := strconv.Atoi(id)
	copy_id := int64(copy_idw)

//    fmt.Fprintf( w, "csv_copy  : copy_idw %v\n", copy_idw )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "csv_copy  : copy_id %v\n", copy_id )  // 繝・ヰ繝・け



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
///    繝｢繝九ち繝ｼ縲蜀崎｡ｨ遉ｺ
///

	process3.Csv_inf ( w , r )

//	fmt.Fprintf( w, "csv_copy : normal end \n" )  // 繝・ヰ繝・け




}
