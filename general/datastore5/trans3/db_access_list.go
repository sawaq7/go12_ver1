package trans3

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/general/type5"

	    "os"
	    "log"
        "cloud.google.com/go/datastore"
        "context"

                                                )

///
///  DB 繧｢繧ｯ繧ｻ繧ｹ繝ｪ繧ｹ繝医ｒ陦ｨ遉ｺ縺吶ｋ
///

func Db_access_list( w http.ResponseWriter, r *http.Request )  ([]type5.Db_Access_List2 ) {

//    fmt.Fprintf( w, "trans3.db_access_list2 start \n" )  // 繝・ヰ繝・け


///
///   繝励Ο繧ｸ繧ｧ繧ｯ繝亥錐繧偵ご繝・ヨ
///
    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )  // 繝・ヰ繝・け

      project_name = "sample-7777"

	}
    ctx := context.Background()
//    c := appengine.NewContext(r)

	query := datastore.NewQuery("Db_Access_List2")

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       log.Fatalf("Failed to create client: %v", err)
    }

    count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return  nil
	}
//	count, err := q.Count(c)

	db_access_list2      := make([]type5.Db_Access_List2, 0, count)

	db_access_list2_view := make([]type5.Db_Access_List2, 0)

    keys, err := client.GetAll(ctx, query , &db_access_list2 )
//	keys, err := q.GetAll(c, &db_access_list2)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)

       return	nil
	}

	 keys_wk := make([]int64, count)

	 for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

     }

	for pos, db_access_list2w := range db_access_list2 {

	   db_access_list2w.Line_No = int64(pos+1)

       db_access_list2_view = append(db_access_list2_view, type5.Db_Access_List2 { keys_wk[pos]               ,
                                                                        db_access_list2w.Line_No      ,
                                                                        db_access_list2w.Db_Type      ,
                                                                        db_access_list2w.Access_Type  ,
                                                                        db_access_list2w.Project_Name ,
                                                                        db_access_list2w.Bucket_Name  ,
                                                                        db_access_list2w.Basic_File_Name   ,
                                                                        db_access_list2w.New_File_Name    ,

                                                                                                 })


	}

    return	db_access_list2_view
}

