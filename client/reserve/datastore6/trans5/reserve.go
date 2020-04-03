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
///    繧ｲ繧ｹ繝医Μ繧ｹ繝医ｒ繧ｲ繝・ヨ縺吶ｋ
///

func Reserve( w http.ResponseWriter, r *http.Request )  ([]type6.Guest ) {

//     IN    w      縲縲縲縲: 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r      縲縲縲縲: 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ

//     OUT guest_slice  : 讒矩菴薙窶昴ご繧ｹ繝医Μ繧ｹ繝遺昴・繧ｹ繝ｩ繧､繧ｹ

//    fmt.Fprintf( w, "trans5.reserve start \n" )  // 繝・ヰ繝・け

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
//		fmt.Fprintf( w, "reserve err \n" ,err)  // 繝・ヰ繝・け
		return	nil
	}

	keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	for pos, guestw := range guest {

///  讖溯・縺ｫ繧医ｊ繝√ぉ繝・け鬆・岼繧偵そ繝・ヨ

         guest_slice = append(guest_slice, type6.Guest { keys_wk[pos]      ,
                                                         guestw.Guest_No   ,
                                                         guestw.Guest_Name    })


	}

    return	guest_slice

}

