package d_schedule_showall

import (

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

//    fmt.Fprintf( w, "d_schedule_showall start \n" )

    var idmy int64

	var g type2.D_Schedule

    g.Course_Num = 4
	g.Course_01 = r.FormValue("course_no_01")  //     get car no1
	g.Car_Name_01 = "332"
	g.Course_02 = r.FormValue("course_no_02")  //     get car no2
	g.Car_Name_02 = "852"
	g.Course_03 = r.FormValue("course_no_03")  //     get car no3
	g.Car_Name_03 = "765"
	g.Course_04 = r.FormValue("course_no_04")  //     get car no4
	g.Car_Name_04 = "784"

//    g.Date = time.Now()
    date_w := time.Now()
    g.Date_Real = date_w
//   g.Date = fmt.Sprintf("%04d/%02d/%02d %02d:%02d:%02d",
//		date_w.Year(), date_w.Month(),date_w.Day(), date_w.Hour(), date_w.Minute(), date_w.Second())
    g.Date = fmt.Sprintf("%04d/%02d/%02d",date_w.Year(), date_w.Month(),date_w.Day())

///     get district np

    flexible_out := datastore2.Datastore_sgh( "D_District_Temp" ,"check" ,idmy , w , r  )

    value2, _ := flexible_out.([]type5.General_Work)

    g.District_No = value2[0].Int64_Work

    district_no := g.District_No

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return
    }
    new_key := datastore.IncompleteKey("D_Schedule", nil)

    if _, err = client.Put(ctx, new_key, &g ); err != nil {

		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

//   	fmt.Fprintf( w, "d_schedule_showall : g.Date_Real %v\n", g.Date_Real.Month() )
//    fmt.Fprintf( w, "d_schedule_showall : g.Number %v\n", g.Number )
//    fmt.Fprintf( w, "d_schedule_showall : g.Date %v\n", g.Date )

///    show schedule inf. on web

	process.D_schedule_showall(w , r , district_no)

//	fmt.Fprintf( w, "d_schedule_showall : normal end \n" )

}
