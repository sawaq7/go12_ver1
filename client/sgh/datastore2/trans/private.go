package trans

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
//	    "html/template"

	    "client/sgh/type2"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"                                       )

///
///  private ファイルから　プライベート情報をゲットする
///

func Private(w http.ResponseWriter, r *http.Request )  ( private_view []type2.Private ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     OUT private_view  : 構造体　”エリア情報”のスライス

//    fmt.Fprintf( w, "trans.private start \n" )  // デバック

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
//		fmt.Fprintf( w, "d_district_area_show err \n" ,err)  // デバック
		return	nil
	}

	keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	for pos, privatew := range private {

//	  fmt.Fprintf( w, "trans.private privatew %v\n" ,privatew)  // デバック

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

