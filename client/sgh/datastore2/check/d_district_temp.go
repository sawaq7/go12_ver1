package check

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
//	    "html/template"
//	    "web/htmls/sgh"
	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
	    "github.com/sawaq7/go12_ver1/general/type5"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                )

///                         縲縲縲縲  ///
/// 謖・ｮ壹＠縺溷慍蛹ｺNO繝ｻ蝨ｰ蛹ｺ蜷阪ｒ繧ｲ繝・ヨ縺吶ｋ ///
///                         縲縲縲縲 ///

func D_district_temp(w http.ResponseWriter, r *http.Request )  ([]type5.General_Work) {


//     IN    w      縲縲     : 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r      縲縲     : 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ
//     OUT general_work_out  : 蝨ｰ蛹ｺNO繝ｻ蝨ｰ蛹ｺ蜷・
//    fmt.Fprintf( w, "check/d_district_temp start \n" )  // 繝・ヰ繝・け

    var district_no int64
    var district_name string

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

    query := datastore.NewQuery("D_District_Temp").Order("District_No")
//	q := datastore.NewQuery("D_District_Temp").Order("District_No")

    count, err := client.Count(ctx, query)
//	count, err := q.Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

//    fmt.Fprintf( w, "check/d_district_temp count \n" ,count )  // 繝・ヰ繝・け

	d_district_temp     := make([]type2.D_District_Temp, 0, count)

    if _, err := client.GetAll(ctx, query , &d_district_temp) ; err != nil {
//	if _, err := q.GetAll(c, &d_district_temp ) ; err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
    } else {
	  for _, d_district_tempw := range d_district_temp {

        district_no =    d_district_tempw.District_No
        district_name =    d_district_tempw.District_Name

//        fmt.Fprintf( w, "check/d_district_temp pos2 %v   \n" , pos2  )  // 繝・ヰ繝・け

      }
//      fmt.Fprintf( w, "check/d_district_temp district_no \n" ,district_no )  // 繝・ヰ繝・け
//      fmt.Fprintf( w, "check/d_district_temp district_name \n" ,district_name )  // 繝・ヰ繝・け
    }

    general_work_out := make([]type5.General_Work, 1)
    general_work_out[0].Int64_Work  = district_no
    general_work_out[0].String_Work = district_name

	return general_work_out
}

