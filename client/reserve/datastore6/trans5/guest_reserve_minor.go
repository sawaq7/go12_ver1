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
/// 地区のエリアデータをゲットする
///

func Guest_reserve_minor( guest_no int64 ,w http.ResponseWriter, r *http.Request )  ([]type6.Guest_Reserve_Minor ) {

//     IN  guest_no  　 : ゲストNO.
//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ

//     OUT guest_reserve_minor_slice  : 構造体　”エリア情報”のスライス

//    fmt.Fprintf( w, "trans.guest_reserve_minor start \n" )  // デバック
//    fmt.Fprintf( w, "trans.guest_reserve_minor guest_no \n" ,guest_no)  // デバック

    var i_count int64

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

    query := datastore.NewQuery("Guest_Reserve_Minor").Order("Date")
//	q := datastore.NewQuery("Guest_Reserve_Minor").Order("Date")

    count, err := client.Count(ctx, query)
//	count, err := q.Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

	guest_reserve_minor      := make([]type6.Guest_Reserve_Minor, 0, count)
	guest_reserve_minor_slice := make([]type6.Guest_Reserve_Minor, 0)

    keys, err := client.GetAll(ctx, query , &guest_reserve_minor)
//	keys, err := q.GetAll(c, &guest_reserve_minor)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
//		fmt.Fprintf( w, "d_district_area_show err \n" ,err)  // デバック
		return	nil
	}

    keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	i_count = 0

	for pos, guest_reserve_minorw := range guest_reserve_minor {

//	  fmt.Fprintf( w, "trans.guest_reserve_minor guest_reserve_minorw %v\n" ,guest_reserve_minorw)  // デバック

///  機能によりチェック項目をセット

      if guest_no == guest_reserve_minorw.Guest_No {

         i_count ++

         guest_reserve_minorw.Line_No = i_count

         guest_reserve_minor_slice = append(guest_reserve_minor_slice, type6.Guest_Reserve_Minor { keys_wk[pos] ,

                                                          guest_reserve_minorw.Line_No     ,
                                                          guest_reserve_minorw.Date        ,
                                                          guest_reserve_minorw.Guest_No    ,
                                                          guest_reserve_minorw.Guest_Name  ,
                                                          guest_reserve_minorw.Start_Time  ,
                                                          guest_reserve_minorw.End_Time           })

      }
	}

    return	guest_reserve_minor_slice
}

