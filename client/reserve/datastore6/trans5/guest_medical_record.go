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

func Guest_medical_record( guest_no int64 ,w http.ResponseWriter, r *http.Request )  ([]type6.Guest_Medical_Record ) {

//     IN  guest_no  　 : ゲスチEO.
//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ

//     OUT guest_medical_record_slice  : 構造体　”エリア惁E��”�Eスライス

//    fmt.Fprintf( w, "trans.guest_medical_record start \n" )  // チE��チE��
//    fmt.Fprintf( w, "trans.guest_medical_record guest_no \n" ,guest_no)  // チE��チE��

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

    query := datastore.NewQuery("Guest_Medical_Record").Order("Date")
//	q := datastore.NewQuery("Guest_Medical_Record").Order("Date")

    count, err := client.Count(ctx, query)
//	count, err := q.Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

	guest_medical_record      := make([]type6.Guest_Medical_Record, 0, count)
	guest_medical_record_slice := make([]type6.Guest_Medical_Record, 0)

    keys, err := client.GetAll(ctx, query , &guest_medical_record)
//	keys, err := q.GetAll(c, &guest_medical_record)
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

	for pos, guest_medical_recordw := range guest_medical_record {

//	  fmt.Fprintf( w, "trans.guest_medical_record guest_medical_recordw %v\n" ,guest_medical_recordw)  // チE��チE��

///  機�EによりチェチE��頁E��をセチE��

      if guest_no == guest_medical_recordw.Guest_No {

         i_count ++

         guest_medical_recordw.Line_No = i_count

         guest_medical_record_slice = append(guest_medical_record_slice, type6.Guest_Medical_Record { keys_wk[pos],
                                                          guest_medical_recordw.Line_No     ,
                                                          guest_medical_recordw.Date        ,
                                                          guest_medical_recordw.Guest_No    ,
                                                          guest_medical_recordw.Guest_Name  ,
                                                          guest_medical_recordw.Text           })

      }
	}

    return	guest_medical_record_slice
}

