package trans

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
//	    "html/template"

	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
//	    "time"

        "cloud.google.com/go/datastore"
        "context"
        "os"
                                                )

///                           ///
/// 蝨ｰ蛹ｺ縺ｮ繧ｨ繝ｪ繧｢謨ｰ繧偵ご繝・ヨ縺吶ｋ ///
///                          ///

func Car_district( district_no int64  ,w http.ResponseWriter, r *http.Request )  ( car_district_view []type2.Car ) {

//     IN  district_no  : 蝨ｰ蝓櫻o
//     IN    w      縲縲: 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r      縲縲: 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ

//     OUT car_district_view  : 讒矩菴薙窶昴お繝ｪ繧｢諠・ｱ窶昴・繧ｹ繝ｩ繧､繧ｹ

//    fmt.Fprintf( w, "trans.car_district start \n" )  // 繝・ヰ繝・け

     project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

    ctx := context.Background()
//	c := appengine.NewContext(r)

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return
    }

    query := datastore.NewQuery("Car").Order("Car_No")
//    q := datastore.NewQuery("Car").Order("Car_No")

    count, err := client.Count(ctx, query)
//	count, err := q.Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

	car_district      := make([]type2.Car, 0, count)
	car_district_view = make([]type2.Car, 0)

    keys, err := client.GetAll(ctx, query , &car_district)
//	keys, err := q.GetAll(c, &car_district)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
//		fmt.Fprintf( w, "d_district_area_show err \n" ,err)  // 繝・ヰ繝・け
		return	nil
	}

    keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	for pos, car_districtw := range car_district {

//	  fmt.Fprintf( w, "trans.car_district car_districtw %v\n" ,car_districtw)  // 繝・ヰ繝・け


      if district_no == car_districtw.District_No {

         car_district_view = append(car_district_view, type2.Car { keys_wk[pos]   ,
                                                          car_districtw.District_No    ,
                                                          car_districtw.District_Name  ,
                                                          car_districtw.Car_No         ,
                                                          car_districtw.Car_Name       ,
                                                          car_districtw.Car_Explain    ,
                                                          car_districtw.Number_Total   ,
                                                          car_districtw.Time_Total     ,
                                                          car_districtw.Productivity      })

      }
	}

    return	car_district_view
}

