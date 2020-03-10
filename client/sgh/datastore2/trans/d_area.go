package trans

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
//	    "html/template"

	    "client/sgh/type2"
//	    "time"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                )

///
/// 地区のエリアデータをゲットする
///

func D_area(funct int64 ,some_no int64 ,w http.ResponseWriter, r *http.Request )  ([]type2.D_Area ) {

//     IN  funct  　　　: ファンクション
//     　　　　　：０  地区NO
//     　　　　　：１  カーNO
//     　　　　　：２  プライベートNO
//     IN  some_no  　　: 各種NO
//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ

//     OUT d_area_view  : 構造体　”エリア情報”のスライス

//    fmt.Fprintf( w, "trans.d_area start \n" )  // デバック
//    fmt.Fprintf( w, "trans.d_area funct \n" ,funct )  // デバック
//    fmt.Fprintf( w, "trans.d_area some_no \n" ,some_no)  // デバック

    var check_no int64

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

    query := datastore.NewQuery("D_Area").Order("Area_No")
//	q := datastore.NewQuery("D_Area").Order("Area_No")

//	count, err := q.Count(c)
	count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

	d_area      := make([]type2.D_Area, 0, count)
	d_area_view := make([]type2.D_Area, 0)


//	keys, err := q.GetAll(c, &d_area)
	keys, err := client.GetAll(ctx, query , &d_area)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
//		fmt.Fprintf( w, "d_district_area_show err \n" ,err)  // デバック
		return	nil
	}

	keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	for pos, d_areaw := range d_area {

//	  fmt.Fprintf( w, "trans.d_area d_areaw %v\n" ,d_areaw)  // デバック

///  機能によりチェック項目をセット

	  if funct == 0 {   // 地区NOの場合

	     check_no = d_areaw.District_No

	  }else if funct == 1 {   // カーNOの場合

	     check_no = 1

	  }else if funct == 2 {   // 個人NOの場合

	     check_no = 2

	  }
      if  some_no == 0 {

         d_area_view = append(d_area_view, type2.D_Area { keys_wk[pos]      ,
                                                          d_areaw.Course_No      ,
                                                          d_areaw.District_No    ,
                                                          d_areaw.District_Name  ,
                                                          d_areaw.Area_No        ,
                                                          d_areaw.Area_Name      ,
                                                          d_areaw.Area_Detail    ,
                                                          d_areaw.Number_Total   ,
                                                          d_areaw.Time_Total     ,
                                                          d_areaw.Productivity   ,
                                                          d_areaw.Car_No           })


      }else if some_no == check_no {

         d_area_view = append(d_area_view, type2.D_Area { keys_wk[pos]      ,
                                                          d_areaw.Course_No      ,
                                                          d_areaw.District_No    ,
                                                          d_areaw.District_Name  ,
                                                          d_areaw.Area_No        ,
                                                          d_areaw.Area_Name      ,
                                                          d_areaw.Area_Detail    ,
                                                          d_areaw.Number_Total   ,
                                                          d_areaw.Time_Total     ,
                                                          d_areaw.Productivity   ,
                                                          d_areaw.Car_No           })

      }
	}

    return	d_area_view
}

