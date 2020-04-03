package d_district_area_show

import (

//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"
//	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2/check"
	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2"
	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
	    "github.com/sawaq7/go12_ver1/general/type5"
//	    "strconv"
//	    "time"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                  )

func D_district_area_show(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "d_district_area_show start \n" )  // 繝・ヰ繝・け

	var g type2.D_Area

	var idmy int64

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

///  temporary-file繧医ｊ縲∝慍蛹ｺNO繝ｻ蝨ｰ蛹ｺ蜷阪ｒGET

    flexible_out := datastore2.Datastore_sgh( "D_District_Temp" ,"check" ,idmy , w , r  )

    value2, _ := flexible_out.([]type5.General_Work)

    g.District_No = value2[0].Int64_Work
    g.District_Name = value2[0].String_Work

//	fmt.Fprintf( w, "d_district_area_show : g.District_No %v\n", g.District_No )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "d_district_area_show : g.District_Name %v\n", g.District_Name )  // 繝・ヰ繝・け

//  繧ｨ繝ｪ繧｢NO縲縺ｮ譌｢蟄倥・MAX蛟､繧竪ET

    count := datastore2.Datastore_sgh( "D_Area","check" ,g.District_No , w , r  )

// 遨ｺ繧､繝ｳ繧ｿ繝ｼ繝輔ぉ繧､繧ｹ螟画焚繧医ｊ繝舌Μ繝･繝ｼ蛟､繧偵ご繝・ヨ

    value, _ := count.(int64)

//	fmt.Fprintf( w, "d_district_area_show count %v   \n" , count  )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "d_district_area_show district_no %v   \n" , district_no  )  // 繝・ヰ繝・け

    g.Area_No  = g.Area_No + int64(value + 1)
    g.Area_Name   = r.FormValue("area_name")
	g.Area_Detail = r.FormValue("area_detail")
    g.Course_No = g.District_No * 100 + g.Area_No

//    fmt.Fprintf( w, "d_district_area_show : g.Area_No %v\n", g.Area_No )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "d_district_area_show : g.Area_Name %v\n", g.Area_Name )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "d_district_area_show : g.Area_Detail %v\n", g.Area_Detail )  // 繝・ヰ繝・け

/// 繝・・繧ｿ繧ｹ繝医い繝ｼ縺ｫ繝・・繧ｿ繧偵そ繝・ヨ ///

    new_key := datastore.IncompleteKey("D_Area", nil)

    if _, err = client.Put(ctx, new_key, &g ); err != nil {
//	if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "D_Area", nil), &g); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

/// 繝｢繝九ち繝ｼ縲蜀崎｡ｨ遉ｺ ///

	process.D_district_area(w , r ,g.District_No)

//	fmt.Fprintf( w, "d_district_area_show : normal end \n" )  // 繝・ヰ繝・け

}
