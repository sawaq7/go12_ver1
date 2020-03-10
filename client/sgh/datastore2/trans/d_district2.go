package trans

import (

	    "net/http"
//	    "fmt"

	    "client/sgh/type2"
//	    "time"

        "cloud.google.com/go/datastore"
        "context"
        "os"                                      )

///                           ///
/// 地区情報をゲットする ///
///                          ///

func D_district2( w http.ResponseWriter, r *http.Request )  ([]type2.D_District_View ) {

//     IN    w      　　　　: レスポンスライター
//     IN    r      　　　　: リクエストパラメータ

//     OUT d_district_view  : 構造体　”地区情報”のスライス

//    fmt.Fprintf( w, "trans.d_district2 start \n" )  // デバック

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "trans.d_district2 :  projectID unset \n"  )  // デバック

      project_name = "sample-7777"

	}

    ctx := context.Background()

	query := datastore.NewQuery("D_District").Order("District_No")

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return  nil
    }

	count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

	d_district      := make([]type2.D_District, 0, count)
	d_district_view := make([]type2.D_District_View, 0)

    keys, err := client.GetAll(ctx, query , &d_district)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
//		fmt.Fprintf( w, "d_district2 err \n" ,err)  // デバック
		return	nil
	}

	keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	for pos, d_districtw := range d_district {

///  機能によりチェック項目をセット

        d_area_slice :=   D_area_district ( w ,r ,d_districtw.District_No )

        d_district_view = append(d_district_view, type2.D_District_View { keys_wk[pos]            ,
                                                                          d_districtw.District_No   ,
                                                                          d_districtw.District_Name ,
                                                                          d_area_slice                })
	}

    return	d_district_view
}

