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
/// 蝨ｰ蛹ｺ縺ｮ繧ｨ繝ｪ繧｢繝・・繧ｿ繧偵ご繝・ヨ縺吶ｋ
///

func Guest_temp( w http.ResponseWriter, r *http.Request )  ([]type6.Guest_Temp ) {

//     IN    w      縲縲: 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r      縲縲: 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ

//     OUT guest_temp_slice  : 讒矩菴薙窶昴お繝ｪ繧｢諠・ｱ窶昴・繧ｹ繝ｩ繧､繧ｹ

//    fmt.Fprintf( w, "trans.guest_temp start \n" )  // 繝・ヰ繝・け

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
//		fmt.Fprintf( w, "d_district_area_show err \n" ,err)  // 繝・ヰ繝・け
		return	nil
	}

    keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	for pos, guest_tempw := range guest_temp {

//	  fmt.Fprintf( w, "trans.guest_temp guest_tempw %v\n" ,guest_tempw)  // 繝・ヰ繝・け

      guest_temp_slice = append(guest_temp_slice, type6.Guest_Temp { keys_wk[pos]            ,
                                                                     guest_tempw.Guest_No    ,
                                                                     guest_tempw.Guest_Name    })

	}

//	fmt.Fprintf( w, "trans.guest_temp : guest_temp_slice %v\n", guest_temp_slice )  // 繝・ヰ繝・け

//    fmt.Fprintf( w, "trans.guest_temp : normal end \n" )  // 繝・ヰ繝・け

    return	guest_temp_slice
}

