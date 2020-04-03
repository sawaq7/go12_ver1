package csv_delete

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
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

//    fmt.Fprintf( w, "csv_delete start \n" )  // 繝・ヰ繝・け

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

    delidw , err := strconv.Atoi(r.FormValue("id"))

	if err  != nil {
//	   fmt.Fprintf( w, "csv_delete :error delidw %v\n", delidw )  // 繝・ヰ繝・け
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    delid := int64(delidw)

//    fmt.Fprintf( w, "csv_delete : delidw %v\n", delidw )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "csv_delete : delid %v\n", delid )  // 繝・ヰ繝・け

//	key := datastore.NewKey(c, "Csv_Inf", "", delid, nil)
	key := datastore.IDKey("Csv_Inf", delid, nil)

// 繝・・繧ｿ繧ｹ繝医い縺九ｉ1繝ｬ繧ｳ繝ｼ繝牙炎髯､


//	if err := datastore.Delete(c, key); err != nil {
    if err := client.Delete(ctx, key); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


/// 繝｢繝九ち繝ｼ縲陦ｨ遉ｺ ///

	process3.Csv_inf ( w , r )

//	fmt.Fprintf( w, "csv_delete : normal end \n" )  // 繝・ヰ繝・け




}
