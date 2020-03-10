package trans5

import (

//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
//	    "html/template"
//	    "web/htmls/sgh"
	    "client/reserve/type6"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                )

///
///    ゲストリストをゲットする
///

func Reserve( w http.ResponseWriter, r *http.Request )  ([]type6.Guest ) {

//     IN    w      　　　　: レスポンスライター
//     IN    r      　　　　: リクエストパラメータ

//     OUT guest_slice  : 構造体　”ゲストリスト”のスライス

//    fmt.Fprintf( w, "trans5.reserve start \n" )  // デバック

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

    query := datastore.NewQuery("Guest").Order("Guest_No")
//	q := datastore.NewQuery("Guest").Order("Guest_No")

    count, err := client.Count(ctx, query)
//	count, err := q.Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

	guest      := make([]type6.Guest, 0, count)
	guest_slice := make([]type6.Guest, 0)

    keys, err := client.GetAll(ctx, query , &guest)
//	keys, err := q.GetAll(c, &guest)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
//		fmt.Fprintf( w, "reserve err \n" ,err)  // デバック
		return	nil
	}

	keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	for pos, guestw := range guest {

///  機能によりチェック項目をセット

         guest_slice = append(guest_slice, type6.Guest { keys_wk[pos]      ,
                                                         guestw.Guest_No   ,
                                                         guestw.Guest_Name    })


	}

    return	guest_slice

}

