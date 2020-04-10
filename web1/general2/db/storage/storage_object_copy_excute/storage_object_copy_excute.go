package storage_object_copy_excute

import (

	    "github.com/sawaq7/go12_ver1/storage2"
//	    "fmt"
	    "net/http"

	    "github.com/sawaq7/go12_ver1/general/type5"
	    "github.com/sawaq7/go12_ver1/general/process3"
        "os"
	    "log"
        "cloud.google.com/go/datastore"
        "context"

                                                  )

func Storage_object_copy_excute(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "storage_object_copy_excute start \n" )

    var project_name ,bucket_name ,basic_file_name string

    var db_access_list2 type5.Db_Access_List2
///
/// 入力データをGET 　
///

//    basic_file_name := r.FormValue("basic_file_name")

    new_file_name := r.FormValue("new_file_name")  // get new file name

///
///   バケチE��名をゲチE��
///
    project_name = os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )

      project_name = "sample-7777"

	}
    ctx := context.Background()

	query := datastore.NewQuery("Storage_B_O_Temp")

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       log.Fatalf("Failed to create client: %v", err)
    }

    count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	storage_b_o_temp      := make([]type5.Storage_B_O_Temp, 0, count)

    if _, err = client.GetAll(ctx, query , &storage_b_o_temp) ; err != nil {

	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return

	} else{

      for pos, storage_b_o_tempw := range storage_b_o_temp {

        if pos == 0 {

           project_name       = storage_b_o_tempw.Project_Name
           bucket_name        = storage_b_o_tempw.Bucket_Name
           basic_file_name    = storage_b_o_tempw.Object_Name

        }
	  }
	}

///
///   copy storage file　
///

    storage2.File_Copy ( w , r  ,bucket_name ,basic_file_name ,new_file_name  )

///
/// アクセスリストに登録　
///

    db_access_list2.Db_Type = "sr"
    db_access_list2.Access_Type = "copy"
    db_access_list2.Project_Name = project_name
    db_access_list2.Bucket_Name = bucket_name
    db_access_list2.Basic_File_Name = basic_file_name
    db_access_list2.New_File_Name = new_file_name

///    add new record

    new_key := datastore.IncompleteKey("Db_Access_List2", nil)

//    fmt.Fprintf(w, "storage_object_copy_excute: new_key %v\n", new_key )

    _, err = client.Put(ctx, new_key, &db_access_list2 )

//    fmt.Fprintf(w, "storage_object_copy_excute: key2 %v\n", key2 )

    if err != nil {

      http.Error(w,err.Error(), http.StatusInternalServerError)
      return
    }

///
///      show object list　
///
    process3.Storage_object_show ( w , r ,project_name  ,bucket_name )

//	fmt.Fprintf( w, "storage_object_copy_excute normal end \n" )
}
