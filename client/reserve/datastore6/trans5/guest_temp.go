package trans5

import (

//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
//	    "html/template"
//	    "web/htmls/sgh"
	    "github.com/sawaq7/go12_ver1/client/reserve/type6"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"                                       )

///
/// 地区のエリアチE�EタをゲチE��する
///

func Guest_temp( w http.ResponseWriter, r *http.Request )  ([]type6.Guest_Temp ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ

//     OUT guest_temp_slice  : 構造体　”エリア惁E��”�Eスライス

//    fmt.Fprintf( w, "trans.guest_temp start \n" )  // チE��チE��

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

    query := datastore.NewQuery("Guest_Temp")
//	q := datastore.NewQuery("Guest_Temp")

    count, err := client.Count(ctx, query)
//	count, err := q.Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

	guest_temp      := make([]type6.Guest_Temp, 0, count)
	guest_temp_slice := make([]type6.Guest_Temp, 0)

    keys, err := client.GetAll(ctx, query , &guest_temp)
//	keys, err := q.GetAll(c, &guest_temp)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
//		fmt.Fprintf( w, "d_district_area_show err \n" ,err)  // チE��チE��
		return	nil
	}

    keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	for pos, guest_tempw := range guest_temp {

//	  fmt.Fprintf( w, "trans.guest_temp guest_tempw %v\n" ,guest_tempw)  // チE��チE��

      guest_temp_slice = append(guest_temp_slice, type6.Guest_Temp { keys_wk[pos]            ,
                                                                     guest_tempw.Guest_No    ,
                                                                     guest_tempw.Guest_Name    })

	}

//	fmt.Fprintf( w, "trans.guest_temp : guest_temp_slice %v\n", guest_temp_slice )  // チE��チE��

//    fmt.Fprintf( w, "trans.guest_temp : normal end \n" )  // チE��チE��

    return	guest_temp_slice
}

