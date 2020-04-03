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
	    "os"
                                                )

///
/// 地区のエリアチE�EタをゲチE��する
///

func Guest_payment( guest_no int64 ,w http.ResponseWriter, r *http.Request )  ([]type6.Guest_Payment ) {

//     IN  guest_no  　 : ゲスチEO.
//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ

//     OUT guest_payment_slice  : 構造体　”エリア惁E��”�Eスライス

//    fmt.Fprintf( w, "trans.guest_payment start \n" )  // チE��チE��
//    fmt.Fprintf( w, "trans.guest_payment guest_no \n" ,guest_no)  // チE��チE��

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
		return nil
	}

    query := datastore.NewQuery("Guest_Payment").Order("Date")
//	q := datastore.NewQuery("Guest_Payment").Order("Date")

    count, err := client.Count(ctx, query)
//	count, err := q.Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

	guest_payment      := make([]type6.Guest_Payment, 0, count)
	guest_payment_slice := make([]type6.Guest_Payment, 0)

    keys, err := client.GetAll(ctx, query , &guest_payment)
//	keys, err := q.GetAll(c, &guest_payment)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
//		fmt.Fprintf( w, "d_district_area_show err \n" ,err)  // チE��チE��
		return	nil
	}

    keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	i_count = 0

	for pos, guest_paymentw := range guest_payment {

//	  fmt.Fprintf( w, "trans.guest_payment guest_paymentw %v\n" ,guest_paymentw)  // チE��チE��

///  機�EによりチェチE��頁E��をセチE��

      if guest_no == guest_paymentw.Guest_No {

         i_count ++

         guest_paymentw.Line_No = i_count

         guest_payment_slice = append(guest_payment_slice, type6.Guest_Payment { keys_wk[pos]   ,

                                                          guest_paymentw.Line_No     ,
                                                          guest_paymentw.Date        ,
                                                          guest_paymentw.Guest_No    ,
                                                          guest_paymentw.Guest_Name  ,
                                                          guest_paymentw.Item        ,
                                                          guest_paymentw.Amount           })

      }
	}

    return	guest_payment_slice
}

