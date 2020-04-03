package trans1000

import (

//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
//	    "html/template"

	    "github.com/sawaq7/go12_ver1/temp/type1000"

	    "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                )

///
/// æE®ããå°åºã®ã¨ãªã¢ãã²ãE
///

func D_area_district(w http.ResponseWriter, r *http.Request , district_no int64)  ([]type1000.D_Area ) {

//     IN    w      ãã: ã¬ã¹ãã³ã¹ã©ã¤ã¿ã¼
//     IN    r      ãã: ãªã¯ã¨ã¹ããã©ã¡ã¼ã¿
//     IN ãdistrict_no : å°åºNo

//     OUT d_area_view  : æ§é ä½ãâã¨ãªã¢æE ±âãEã¹ã©ã¤ã¹

//    fmt.Fprintf( w, "trans.d_area_district district_no \n" ,district_no)  // ãEãE¯

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

    count, err := client.Count(ctx, query)
//	count, err := q.Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

	d_area      := make([]type1000.D_Area, 0, count)
	d_area_view := make([]type1000.D_Area, 0)

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

