package d_schedule_showall

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"

	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2"
	    "github.com/sawaq7/go12_ver1/general/type5"

//	    "strconv"
	    "time"

	    "cloud.google.com/go/datastore"
        "context"
        "os"
                                                  )

func D_schedule_showall(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "d_schedule_showall start \n" )  // 繝・ヰ繝・け


    var idmy int64

	var g type2.D_Schedule

    g.Course_Num = 4
	g.Course_01 = r.FormValue("course_no_01")  // 332蜿ｷ霆翫・諡・ｽ楢・ご繝・ヨ
	g.Car_Name_01 = "332"
	g.Course_02 = r.FormValue("course_no_02")  // 852蜿ｷ霆翫・諡・ｽ楢・ご繝・ヨ
	g.Car_Name_02 = "852"
	g.Course_03 = r.FormValue("course_no_03")  // 765蜿ｷ霆翫・諡・ｽ楢・ご繝・ヨ
	g.Car_Name_03 = "765"
	g.Course_04 = r.FormValue("course_no_04")  // 784蜿ｷ霆翫・諡・ｽ楢・ご繝・ヨ
	g.Car_Name_04 = "784"

//    g.Date = time.Now()         // 譌･莉倥ｒ繧ｻ繝・ヨ
    date_w := time.Now()        // 譌･莉倥ｒ繧ｻ繝・ヨ
    g.Date_Real = date_w
//   g.Date = fmt.Sprintf("%04d/%02d/%02d %02d:%02d:%02d",
//		date_w.Year(), date_w.Month(),date_w.Day(), date_w.Hour(), date_w.Minute(), date_w.Second())
    g.Date = fmt.Sprintf("%04d/%02d/%02d",date_w.Year(), date_w.Month(),date_w.Day())

//蝨ｰ蛹ｺNO繧竪ET

    flexible_out := datastore2.Datastore_sgh( "D_District_Temp" ,"check" ,idmy , w , r  )

    value2, _ := flexible_out.([]type5.General_Work)

    g.District_No = value2[0].Int64_Work

    district_no := g.District_No

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}
    ctx := context.Background()
//    c := appengine.NewContext(r)

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return
    }
    new_key := datastore.IncompleteKey("D_Schedule", nil)

    if _, err = client.Put(ctx, new_key, &g ); err != nil {
//	if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "D_Schedule", nil), &g); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

//   	fmt.Fprintf( w, "d_schedule_showall : g.Date_Real %v\n", g.Date_Real.Month() )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "d_schedule_showall : g.Number %v\n", g.Number )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "d_schedule_showall : g.Date %v\n", g.Date )  // 繝・ヰ繝・け

/// 繝｢繝九ち繝ｼ縲蜀崎｡ｨ遉ｺ ///

	process.D_schedule_showall(w , r , district_no)

//	fmt.Fprintf( w, "d_schedule_showall : normal end \n" )  // 繝・ヰ繝・け

}
