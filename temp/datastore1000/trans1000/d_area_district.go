package trans1000

import (


	    "net/http"
//	    "fmt"
//	    "html/template"

	    "github.com/sawaq7/go12_ver1/temp/type1000"

	    "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                )

func D_area_district(w http.ResponseWriter, r *http.Request , district_no int64)  ([]type1000.D_Area ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　district_no : 地区No

//     OUT d_area_view  :

//    fmt.Fprintf( w, "trans.d_area_district district_no \n" ,district_no)

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

    query := datastore.NewQuery("D_Area").Order("Area_No")

    count, err := client.Count(ctx, query)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

	d_area      := make([]type1000.D_Area, 0, count)
	d_area_view := make([]type1000.D_Area, 0)

    keys, err := client.GetAll(ctx, query , &d_area)

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

         d_area_view = append(d_area_view, type1000.D_Area { keys_wk[pos]      ,
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

