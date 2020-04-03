package trans

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
//	    "html/template"

	    "github.com/sawaq7/go12_ver1/client/sgh/type2"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"                                       )

///
///  private 繝輔ぃ繧､繝ｫ縺九ｉ縲繝励Λ繧､繝吶・繝域ュ蝣ｱ繧偵ご繝・ヨ縺吶ｋ
///

func Private(w http.ResponseWriter, r *http.Request )  ( private_view []type2.Private ) {

//     IN    w      縲縲: 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r      縲縲: 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ
//     OUT private_view  : 讒矩菴薙窶昴お繝ｪ繧｢諠・ｱ窶昴・繧ｹ繝ｩ繧､繧ｹ

//    fmt.Fprintf( w, "trans.private start \n" )  // 繝・ヰ繝・け

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)

    query := datastore.NewQuery("Private").Order("Worker_No")
//	q := datastore.NewQuery("Private").Order("Worker_No")

    count, err := client.Count(ctx, query)
//	count, err := q.Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

	private      := make([]type2.Private, 0, count)
	private_view = make([]type2.Private, 0)

    keys, err := client.GetAll(ctx, query , &private)
//	keys, err := q.GetAll(c, &private)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
//		fmt.Fprintf( w, "d_district_area_show err \n" ,err)  // 繝・ヰ繝・け
		return	nil
	}

	keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	for pos, privatew := range private {

//	  fmt.Fprintf( w, "trans.private privatew %v\n" ,privatew)  // 繝・ヰ繝・け

      private_view = append(private_view, type2.Private { keys_wk[pos]            ,
                                                          privatew.Worker_No      ,
                                                          privatew.Worker_Name    ,
                                                          privatew.Worker_Type    ,
                                                          privatew.Worker_Salary  ,
                                                          privatew.Worker_Twh     ,
                                                          privatew.Worker_H_Pay   ,
                                                          privatew.Number_Total   ,
                                                          privatew.Time_Total     ,
                                                          privatew.Productivity     })

	}

    return	private_view
}

