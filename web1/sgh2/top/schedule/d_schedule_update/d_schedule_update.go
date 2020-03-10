package d_schedule_update

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
	    "client/sgh/process"

	    "client/sgh/type2"
	    "strconv"
	    "basic/date1"

	    "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                  )

func D_schedule_update(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "d_schedule_update start \n" )  // デバック

	var g type2.D_Schedule

	updidw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {

//	   fmt.Fprintf( w, "d_schedule_update :error updidw %v\n", updidw )  // デバック

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

//    fmt.Fprintf( w, "d_schedule_update : updidw %v\n", updidw )  // デバック
//    fmt.Fprintf( w, "d_schedule_update : updid %v\n", updid )  // デバック



	g.Course_01 = r.FormValue("course_no_01")  // 01号車の担当者ゲット
	g.Course_02 = r.FormValue("course_no_02")  // 02号車の担当者ゲット
	g.Course_03 = r.FormValue("course_no_03")  // 03号車の担当者ゲット
	g.Course_04 = r.FormValue("course_no_04")  // 04号車の担当者ゲット
    g.Date   = r.FormValue("date")               // 日付をゲット

    g.Date_Real = date1.Date_realdata_get( w  ,g.Date )   // タイムデータ作成

    district_now , err := strconv.Atoi(r.FormValue("district_no"))
	if err  != nil {

//	   fmt.Fprintf( w, "d_schedule_update :error updidw %v\n", updidw )  // デバック

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

//   	fmt.Fprintf( w, "d_schedule_update : g.Area_Name %v\n", g.Area_Name )  // デバック
//    fmt.Fprintf( w, "d_schedule_update : g.Number %v\n", g.Number )  // デバック
//    fmt.Fprintf( w, "d_schedule_update : g.Date %v\n", g.Date )  // デバック

/// モニター　再表示 ///

	process.D_schedule_showall(w , r ,g.District_No)

//	fmt.Fprintf( w, "d_schedule_update : normal end \n" )  // デバック




}
