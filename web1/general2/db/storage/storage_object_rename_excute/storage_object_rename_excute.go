package storage_object_rename_excute

import (
	    "storage2"
//	    "fmt"
	    "net/http"
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "github.com/sawaq7/go12_ver1/general/type5"
	    "github.com/sawaq7/go12_ver1/general/process3"

	    "os"
	    "log"
        "cloud.google.com/go/datastore"
        "context"

                                                  )

func Storage_object_rename_excute(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "storage_object_rename_excute start \n" )  // ãEãE¯

    var project_name ,bucket_name ,basic_file_name string

    var db_access_list2 type5.Db_Access_List2      //D.B. ã¢ã¯ã»ã¹ãªã¹ããEç¨ã®ã¯ã¼ã¯ã¨ãªã¢ãç¢ºä¿E
///
/// å¥åãã¼ã¿ãGET ã
///

    new_file_name := r.FormValue("new_file_name")  // ãã¥ã¼ãã¡ã¤ã«åãã²ãE

///
///   ãã­ã¸ã§ã¯ãåãã²ãE
///
    project_name = os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )  // ãEãE¯

      project_name = "sample-7777"

	}
    ctx := context.Background()
//    c := appengine.NewContext(r)

    query := datastore.NewQuery("Storage_B_O_Temp")

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       log.Fatalf("Failed to create client: %v", err)
    }

//	q := datastore.NewQuery("Storage_B_O_Temp")

//	count, err := q.Count(c)
    count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	storage_b_o_temp      := make([]type5.Storage_B_O_Temp, 0, count)

    if _, err = client.GetAll(ctx, query , &storage_b_o_temp) ; err != nil {
//	if _, err := q.GetAll(c, &storage_b_o_temp);  err != nil {

	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return

	} else{

      for pos, storage_b_o_tempw := range storage_b_o_temp {

        if pos == 0 {

           project_name       = storage_b_o_tempw.Project_Name    // ãã­ã¸ã§ã¯ãåãã²ãE
           bucket_name        = storage_b_o_tempw.Bucket_Name    // ãã±ãEåãã²ãE
           basic_file_name    = storage_b_o_tempw.Object_Name    // ããEã·ãE¯ãã¡ã¤ã«åãã²ãE

        }
	  }
	}

///
/// ãããã¹ãã¬ãE¸ãã¡ã¤ã«ããªããEã ããã
///

    storage2.File_Rename ( w , r  ,bucket_name ,basic_file_name ,new_file_name  )

///
/// ããããã¢ã¯ã»ã¹ãªã¹ãã«ç»é²ã
///

///  åE¨®ã¢ã¯ã»ã¹æE ±ãã»ãE

    db_access_list2.Db_Type = "sr"
    db_access_list2.Access_Type = "rename"
    db_access_list2.Project_Name = project_name
    db_access_list2.Bucket_Name = bucket_name
    db_access_list2.Basic_File_Name = basic_file_name
    db_access_list2.New_File_Name = new_file_name

/// ãEEã¿ã¹ãã¢ã«ãEã¬ã³ã¼ããè¿½å 

   new_key := datastore.IncompleteKey("Db_Access_List2", nil)

   if _, err := client.Put(ctx, new_key, &db_access_list2 ); err != nil {
//   if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Db_Access_List2", nil) , &db_access_list2); err != nil {
      http.Error(w,err.Error(), http.StatusInternalServerError)
      return
   }

///
/// ãããããªãã¸ã§ã¯ããªã¹ããè¡¨ç¤ºããã
///
    process3.Storage_object_show ( w , r ,project_name  ,bucket_name )

//	fmt.Fprintf( w, "storage_object_rename_excute normal end \n" )  // ãEãE¯
}
