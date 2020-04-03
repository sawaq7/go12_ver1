package trans

import (

//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
//	    "html/template"
//	    "web/htmls/sgh"
	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
//	    "time"

        "cloud.google.com/go/datastore"
        "context"
        "os"                                      )

///                           ///
/// 地区のエリア数をゲチE��する ///
///                          ///

func D_district(funct int64 ,some_no int64 ,w http.ResponseWriter, r *http.Request )  ([]type2.D_District ) {

//     IN  funct  　　　　　: ファンクション
//     　　　　　�E�！E 地区NO
//     　　　　　�E�！E カーNO
//     　　　　　�E�！E プライベ�EチEO
//     IN  some_no  　　　　:
//     IN    w      　　　　: レスポンスライター
//     IN    r      　　　　: リクエストパラメータ

//     OUT d_district_view  : 構造体　”地区惁E��”�Eスライス

//    fmt.Fprintf( w, "trans.d_district start \n" )  // チE��チE��
//    fmt.Fprintf( w, "trans.d_district some_no \n" ,some_no)  // チE��チE��

//    var check_no int64

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )  // チE��チE��

      project_name = "sample-7777"

	}
    ctx := context.Background()
//    c := appengine.NewContext(r)

	query := datastore.NewQuery("D_District").Order("District_No")

//	q := datastore.NewQuery("D_District").Order("District_No")

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return  nil
    }

//	count, err := q.Count(c)
	count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

	d_district      := make([]type2.D_District, 0, count)
	d_district_view := make([]type2.D_District, 0)

    keys, err := client.GetAll(ctx, query , &d_district)
//	keys, err := q.GetAll(c, &d_district)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
//		fmt.Fprintf( w, "d_district_area_show err \n" ,err)  // チE��チE��
		return	nil
	}

	keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	for pos, d_districtw := range d_district {

///  機�EによりチェチE��頁E��をセチE��

         d_district_view = append(d_district_view, type2.D_District { keys_wk[pos]            ,
                                                                    d_districtw.District_No   ,
                                                                    d_districtw.District_Name            })

	}

    return	d_district_view
}

