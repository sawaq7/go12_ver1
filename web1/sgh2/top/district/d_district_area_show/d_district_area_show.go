package d_district_area_show

import (

//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
	    "client/sgh/process"
//	    "client/sgh/datastore2/check"
	    "client/sgh/datastore2"
	    "client/sgh/type2"
	    "general/type5"
//	    "strconv"
//	    "time"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                  )

func D_district_area_show(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "d_district_area_show start \n" )  // デバック

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

///  temporary-fileより、地区NO・地区名をGET

    flexible_out := datastore2.Datastore_sgh( "D_District_Temp" ,"check" ,idmy , w , r  )

    value2, _ := flexible_out.([]type5.General_Work)

    g.District_No = value2[0].Int64_Work
    g.District_Name = value2[0].String_Work

//	fmt.Fprintf( w, "d_district_area_show : g.District_No %v\n", g.District_No )  // デバック
//	fmt.Fprintf( w, "d_district_area_show : g.District_Name %v\n", g.District_Name )  // デバック

//  エリアNO　の既存のMAX値をGET

    count := datastore2.Datastore_sgh( "D_Area","check" ,g.District_No , w , r  )

// 空インターフェイス変数よりバリュー値をゲット

    value, _ := count.(int64)

//	fmt.Fprintf( w, "d_district_area_show count %v   \n" , count  )  // デバック
//	fmt.Fprintf( w, "d_district_area_show district_no %v   \n" , district_no  )  // デバック

    g.Area_No  = g.Area_No + int64(value + 1)
    g.Area_Name   = r.FormValue("area_name")
	g.Area_Detail = r.FormValue("area_detail")
    g.Course_No = g.District_No * 100 + g.Area_No

//    fmt.Fprintf( w, "d_district_area_show : g.Area_No %v\n", g.Area_No )  // デバック
//	fmt.Fprintf( w, "d_district_area_show : g.Area_Name %v\n", g.Area_Name )  // デバック
//	fmt.Fprintf( w, "d_district_area_show : g.Area_Detail %v\n", g.Area_Detail )  // デバック

/// データストアーにデータをセット ///

    new_key := datastore.IncompleteKey("D_Area", nil)

    if _, err = client.Put(ctx, new_key, &g ); err != nil {
//	if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "D_Area", nil), &g); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

/// モニター　再表示 ///

	process.D_district_area(w , r ,g.District_No)

//	fmt.Fprintf( w, "d_district_area_show : normal end \n" )  // デバック

}
