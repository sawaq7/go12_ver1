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
/// 謖・ｮ壹＠縺滓律莉倥・莠郁ｦ壽ュ蝣ｱ繧偵ご繝・ヨ縺吶ｋ
///

func Guest_reserve_minor2( reserve_date string ,w http.ResponseWriter, r *http.Request )  ([]type6.Guest_Reserve_Minor ) {

//     IN  reserve_date : 莠育ｴ・律
//     IN    w      縲縲: 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r      縲縲: 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ

//     OUT guest_reserve_minor2_slice  : 讒矩菴薙窶昴お繝ｪ繧｢諠・ｱ窶昴・繧ｹ繝ｩ繧､繧ｹ

//    fmt.Fprintf( w, "trans.guest_reserve_minor2 start \n" )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "trans.guest_reserve_minor2 reserve_date \n" ,reserve_date)  // 繝・ヰ繝・け

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
//		fmt.Fprintf( w, "d_district_area_show err \n" ,err)  // 繝・ヰ繝・け
		return	nil
	}

    keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	i_count = 0

	for pos, guest_reserve_minorw := range guest_reserve_minor {

//	  fmt.Fprintf( w, "trans.guest_reserve_minor guest_reserve_minorw %v\n" ,guest_reserve_minorw)  // 繝・ヰ繝・け

///  讖溯・縺ｫ繧医ｊ繝√ぉ繝・け鬆・岼繧偵そ繝・ヨ

      if reserve_date == guest_reserve_minorw.Date {

         i_count ++

         guest_reserve_minorw.Line_No = i_count

         guest_reserve_minor_slice = append(guest_reserve_minor_slice, type6.Guest_Reserve_Minor { keys_wk[pos]   ,

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

