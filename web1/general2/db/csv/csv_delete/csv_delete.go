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

//    fmt.Fprintf( w, "csv_delete start \n" )  // γEγE―

///
///   γγ­γΈγ§γ―γεγγ²γE
///
    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )  // γEγE―

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
//	   fmt.Fprintf( w, "csv_delete :error delidw %v\n", delidw )  // γEγE―
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    delid := int64(delidw)

//    fmt.Fprintf( w, "csv_delete : delidw %v\n", delidw )  // γEγE―
//    fmt.Fprintf( w, "csv_delete : delid %v\n", delid )  // γEγE―

//	key := datastore.NewKey(c, "Csv_Inf", "", delid, nil)
	key := datastore.IDKey("Csv_Inf", delid, nil)

// γEEγΏγΉγγ’γγ1γ¬γ³γΌγει€


//	if err := datastore.Delete(c, key); err != nil {
    if err := client.Delete(ctx, key); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


/// γ’γγΏγΌγθ‘¨η€Ί ///

	process3.Csv_inf ( w , r )

//	fmt.Fprintf( w, "csv_delete : normal end \n" )  // γEγE―




}
