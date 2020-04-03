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

//    fmt.Fprintf( w, "csv_delete start \n" )  // チE��チE��

///
///   プロジェクト名をゲチE��
///
    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )  // チE��チE��

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
//	   fmt.Fprintf( w, "csv_delete :error delidw %v\n", delidw )  // チE��チE��
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    delid := int64(delidw)

//    fmt.Fprintf( w, "csv_delete : delidw %v\n", delidw )  // チE��チE��
//    fmt.Fprintf( w, "csv_delete : delid %v\n", delid )  // チE��チE��

//	key := datastore.NewKey(c, "Csv_Inf", "", delid, nil)
	key := datastore.IDKey("Csv_Inf", delid, nil)

// チE�Eタストアから1レコード削除


//	if err := datastore.Delete(c, key); err != nil {
    if err := client.Delete(ctx, key); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


/// モニター　表示 ///

	process3.Csv_inf ( w , r )

//	fmt.Fprintf( w, "csv_delete : normal end \n" )  // チE��チE��




}
