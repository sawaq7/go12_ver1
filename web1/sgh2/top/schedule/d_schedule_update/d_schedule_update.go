package d_schedule_update

import (

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

//    fmt.Fprintf( w, "d_schedule_update start \n" )

	var g type2.D_Schedule

	updidw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {

//	   fmt.Fprintf( w, "d_schedule_update :error updidw %v\n", updidw )

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

//    fmt.Fprintf( w, "d_schedule_update : updidw %v\n", updidw )
//    fmt.Fprintf( w, "d_schedule_update : updid %v\n", updid )



	g.Course_01 = r.FormValue("course_no_01")  //     get car no2
	g.Course_02 = r.FormValue("course_no_02")  //     get car no2
	g.Course_03 = r.FormValue("course_no_03")  //     get car no3
	g.Course_04 = r.FormValue("course_no_04")  //     get car no4
    g.Date   = r.FormValue("date")

 //   change data-type from string-type to date-type
    g.Date_Real = date1.Date_realdata_get( w  ,g.Date )

    district_now , err := strconv.Atoi(r.FormValue("district_no"))
	if err  != nil {

//	   fmt.Fprintf( w, "d_schedule_update :error updidw %v\n", updidw )

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    g.District_No = int64(district_now)

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)

    key := datastore.IDKey("D_Schedule", updid, nil)

    if _, err := client.Put(ctx, key, &g ); err != nil {

		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

//   	fmt.Fprintf( w, "d_schedule_update : g.Area_Name %v\n", g.Area_Name )
//    fmt.Fprintf( w, "d_schedule_update : g.Number %v\n", g.Number )
//    fmt.Fprintf( w, "d_schedule_update : g.Date %v\n", g.Date )

///    show schedule inf. on web

	process.D_schedule_showall(w , r ,g.District_No)

//	fmt.Fprintf( w, "d_schedule_update : normal end \n" )

}
