package trans1000

import (
	    "google.golang.org/appengine"
	    "google.golang.org/appengine/datastore"
	    "net/http"
	    "fmt"
//	    "html/template"
//	    "web/htmls/sgh"

	    "github.com/sawaq7/go12_ver1/temp/type1000"
//	    "time"
                                                )

///
/// 謖・ｮ壹＠縺溷慍蛹ｺ縺ｮ繧ｨ繝ｪ繧｢繧偵ご繝・ヨ
///

func D_area_district_3d(w http.ResponseWriter, r *http.Request , district_no int64)  ([]type1000.D_Area_2 ) {

//     IN    w      縲縲: 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r      縲縲: 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ
//     IN 縲district_no : 蝨ｰ蛹ｺNo

//     OUT d_area_view  : 讒矩菴薙窶昴お繝ｪ繧｢諠・ｱ窶昴・繧ｹ繝ｩ繧､繧ｹ

    fmt.Fprintf( w, "trans.d_area_district_3d district_no \n" ,district_no)  // 繝・ヰ繝・け

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

