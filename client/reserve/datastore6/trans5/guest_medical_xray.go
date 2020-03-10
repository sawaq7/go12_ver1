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

func Guest_medical_xray( guest_no int64 ,w http.ResponseWriter, r *http.Request )  ([]type6.Guest_Medical_Xray ) {

//     IN  guest_no  　 : ゲストNO.
//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ

//     OUT guest_medical_xray_slice  : 構造体　”エリア情報”のスライス

//    fmt.Fprintf( w, "trans.guest_medical_xray start \n" )  // デバック
//    fmt.Fprintf( w, "trans.guest_medical_xray guest_no \n" ,guest_no)  // デバック

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

    query := datastore.NewQuery("Guest_Medical_Xray").Order("Date")
//	q := datastore.NewQuery("Guest_Medical_Xray").Order("Date")

    count, err := client.Count(ctx, query)
//	count, err := q.Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

	guest_medical_xray      := make([]type6.Guest_Medical_Xray, 0, count)
	guest_medical_xray_slice := make([]type6.Guest_Medical_Xray, 0)

    keys, err := client.GetAll(ctx, query , &guest_medical_xray)
//	keys, err := q.GetAll(c, &guest_medical_xray)
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

	for pos, guest_medical_xrayw := range guest_medical_xray {

//	  fmt.Fprintf( w, "trans.guest_medical_xray guest_medical_xrayw %v\n" ,guest_medical_xrayw)  // デバック

///  機能によりチェック項目をセット

      if guest_no == guest_medical_xrayw.Guest_No {

         i_count ++

         guest_medical_xrayw.Line_No = i_count

         guest_medical_xray_slice = append(guest_medical_xray_slice, type6.Guest_Medical_Xray { keys_wk[pos],
                                                          guest_medical_xrayw.Line_No     ,
                                                          guest_medical_xrayw.Date        ,
                                                          guest_medical_xrayw.Guest_No    ,
                                                          guest_medical_xrayw.Guest_Name  ,
                                                          guest_medical_xrayw.File_Name   ,
                                                          guest_medical_xrayw.Url            })

      }
	}

    return	guest_medical_xray_slice
}

