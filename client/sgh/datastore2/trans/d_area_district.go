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

///
/// 謖・ｮ壹＠縺溷慍蛹ｺ縺ｮ繧ｨ繝ｪ繧｢繧偵ご繝・ヨ
///

func D_area_district(w http.ResponseWriter, r *http.Request , district_no int64)  ([]type2.D_Area ) {

//     IN    w      縲縲: 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r      縲縲: 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ
//     IN 縲district_no : 蝨ｰ蛹ｺNo

//     OUT d_area_view  : 讒矩菴薙窶昴お繝ｪ繧｢諠・ｱ窶昴・繧ｹ繝ｩ繧､繧ｹ

//    fmt.Fprintf( w, "trans.d_area_district district_no \n" ,district_no)  // 繝・ヰ繝・け

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

    ctx := context.Background()
//	c := appengine.NewContext(r)

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return nil
    }

    query := datastore.NewQuery("D_Area").Order("Area_No")
//	q := datastore.NewQuery("D_Area").Order("Area_No")

    count, err := client.Count(ctx, query)
//	count, err := q.Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

	d_area      := make([]type2.D_Area, 0, count)
	d_area_view := make([]type2.D_Area, 0)

    keys, err := client.GetAll(ctx, query , &d_area)
//	keys, err := q.GetAll(c, &d_area)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

	keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	for pos, d_areaw := range d_area {

      if district_no == d_areaw.District_No {

         d_area_view = append(d_area_view, type2.D_Area { keys_wk[pos]           ,
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

