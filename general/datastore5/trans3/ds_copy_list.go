package trans3

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"

	    "general/type5"

	    "cloud.google.com/go/datastore"
        "context"
        "os"

                                                )

///                           ///
/// 地区のエリア数をゲットする ///
///                          ///

func Ds_copy_list( w http.ResponseWriter, r *http.Request )  ([]type5.Ds_Copy_List ) {

//    fmt.Fprintf( w, "trans.Ds_copy_list start \n" )  // デバック

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "trans.Ds_copy_list :  projectID unset \n"  )  // デバック

      project_name = "sample-7777"

	}
    ctx := context.Background()
//    c := appengine.NewContext(r)

//	q := datastore.NewQuery("Ds_Copy_List")
	query := datastore.NewQuery("Ds_Copy_List")

	client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return  nil
    }

    count, err := client.Count(ctx, query)
//	count, err := q.Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

// import struct for accessing datastore get from client/sgh/type2/sgh.go

	ds_copy_list      := make([]type5.Ds_Copy_List, 0, count)

	ds_copy_list_view := make([]type5.Ds_Copy_List, 0)

    keys, err := client.GetAll(ctx, query , &ds_copy_list)
//	keys, err := q.GetAll(c, &ds_copy_list)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
//		fmt.Fprintf( w, "d_district_area_show err \n" ,err)  // デバック
       return	nil
	}

	keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	for pos, ds_copy_listw := range ds_copy_list {


         ds_copy_list_view = append(ds_copy_list_view, type5.Ds_Copy_List { keys_wk[pos] ,
                                                                    ds_copy_listw.Basic_Name  ,
                                                                    ds_copy_listw.Copy_Name   ,
                                                                    ds_copy_listw.New_Name    ,

                                                                                                 })


	}

    return	ds_copy_list_view
}

