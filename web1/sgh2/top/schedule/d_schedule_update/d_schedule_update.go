package d_schedule_update

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"

	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
	    "strconv"
	    "github.com/sawaq7/go12_ver1/basic/date1"

	    "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                  )

func D_schedule_update(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "d_schedule_update start \n" )  // 繝・ヰ繝・け

	var g type2.D_Schedule

	updidw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {

//	   fmt.Fprintf( w, "d_schedule_update :error updidw %v\n", updidw )  // 繝・ヰ繝・け

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

//    fmt.Fprintf( w, "d_schedule_update : updidw %v\n", updidw )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "d_schedule_update : updid %v\n", updid )  // 繝・ヰ繝・け



	g.Course_01 = r.FormValue("course_no_01")  // 01蜿ｷ霆翫・諡・ｽ楢・ご繝・ヨ
	g.Course_02 = r.FormValue("course_no_02")  // 02蜿ｷ霆翫・諡・ｽ楢・ご繝・ヨ
	g.Course_03 = r.FormValue("course_no_03")  // 03蜿ｷ霆翫・諡・ｽ楢・ご繝・ヨ
	g.Course_04 = r.FormValue("course_no_04")  // 04蜿ｷ霆翫・諡・ｽ楢・ご繝・ヨ
    g.Date   = r.FormValue("date")               // 譌･莉倥ｒ繧ｲ繝・ヨ

    g.Date_Real = date1.Date_realdata_get( w  ,g.Date )   // 繧ｿ繧､繝繝・・繧ｿ菴懈・

    district_now , err := strconv.Atoi(r.FormValue("district_no"))
	if err  != nil {

//	   fmt.Fprintf( w, "d_schedule_update :error updidw %v\n", updidw )  // 繝・ヰ繝・け

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    g.District_No = int64(district_now)

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)

    key := datastore.IDKey("D_Schedule", updid, nil)

    if _, err := client.Put(ctx, key, &g ); err != nil {
//    key := datastore.NewKey(c, "D_Schedule", "", updid, nil)
//	if _, err := datastore.Put(c, key, &g); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

//   	fmt.Fprintf( w, "d_schedule_update : g.Area_Name %v\n", g.Area_Name )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "d_schedule_update : g.Number %v\n", g.Number )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "d_schedule_update : g.Date %v\n", g.Date )  // 繝・ヰ繝・け

/// 繝｢繝九ち繝ｼ縲蜀崎｡ｨ遉ｺ ///

	process.D_schedule_showall(w , r ,g.District_No)

//	fmt.Fprintf( w, "d_schedule_update : normal end \n" )  // 繝・ヰ繝・け




}
