package trans1000

import (
	    "google.golang.org/appengine"
	    "google.golang.org/appengine/datastore"
	    "net/http"
	    "fmt"
//	    "html/template"
//	    "web/htmls/sgh"

	    "temp/type1000"
//	    "time"
                                                )

///
/// 指定した地区のエリアをゲット
///

func D_area_district_3d(w http.ResponseWriter, r *http.Request , district_no int64)  ([]type1000.D_Area_2 ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　district_no : 地区No

//     OUT d_area_view  : 構造体　”エリア情報”のスライス

    fmt.Fprintf( w, "trans.d_area_district_3d district_no \n" ,district_no)  // デバック

    c := appengine.NewContext(r)

	q := datastore.NewQuery("D_Area").Order("Area_No")

	count, err := q.Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

	d_area      := make([]type1000.D_Area, 0, count)
	d_area_view := make([]type1000.D_Area_2, 0)

    d_area_small_slice := make([]type1000.D_Area_Small, 1)
    d_area_small_slice[0].Area_Name = "area_name_test"
    d_area_small_slice[0].Area_Small_Name = "area_small_name_test"

	keys, err := q.GetAll(c, &d_area)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}
	for pos, d_areaw := range d_area {

      if district_no == d_areaw.District_No {

         d_area_small_slice := make([]type1000.D_Area_Small, 1)
         d_area_small_slice[0].Area_Name = "area_name_test"
         d_area_small_slice[0].Area_Small_Name = "area_small_name_test"

         d_area_view = append(d_area_view, type1000.D_Area_2 { keys[pos].IntID()      ,
                                                          d_areaw.Course_No      ,
                                                          d_areaw.District_No    ,
                                                          d_areaw.District_Name  ,
                                                          d_areaw.Area_No        ,
                                                          d_areaw.Area_Name      ,
                                                          d_areaw.Area_Detail    ,
                                                          d_areaw.Number_Total   ,
                                                          d_areaw.Time_Total     ,
                                                          d_areaw.Productivity   ,
                                                          d_areaw.Car_No         ,
                                                          d_area_small_slice        })



      }
	}

    return	d_area_view
}

