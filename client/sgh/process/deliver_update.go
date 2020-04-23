package process

import (

	    "net/http"
	    "strconv"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/basic/date1"
	    "github.com/sawaq7/go12_ver1/client/sgh/type2"

	    "cloud.google.com/go/datastore"
	    "context"
	    "os"

                                                )

///
///      update deliver inf. for each private in d.s.
///

func Deliver_update_single(w http.ResponseWriter, r *http.Request ,updid int64) {

//     IN    w      : レスポンスライター
//     IN    r      : リクエストパラメータ
//     IN  updid　  : uodate id

//    fmt.Fprintf( w, "deliver_update_single start \n" )
    var g type2.Deliver

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)

    key := datastore.IDKey("Deliver", updid, nil)

    if err := client.Get(ctx, key , &g ) ; err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    g.Area_Name = r.FormValue("area_name")

    g.Date   = r.FormValue("date")

    g.Date_Real = date1.Date_realdata_get( w  ,g.Date )   // make time data

	number := r.FormValue("number")

//	fmt.Fprintf( w, "deliver_update_single : number %v\n", number )

	numberw ,err := strconv.Atoi(number)
	if err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)

		return
	}

	g.Number = int64(numberw)

	private_no := r.FormValue("private_no")         //   get private no
//	fmt.Fprintf( w, "deliver_update_single : private_no %v\n", private_no )

	private_now ,err := strconv.Atoi(private_no)  // make an integer
	if err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)

	}

	g.Private_No = int64(private_now)   // make an integer64

	car_no := r.FormValue("car_no")         //   get car no
//	fmt.Fprintf( w, "deliver_update_single : car_no %v\n", car_no )

	car_now ,err := strconv.Atoi(car_no)  // make an integer
	if err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)

	}

	g.Car_No = int64(car_now)   // 整数の64ビット化

//	fmt.Fprintf( w, "deliver_update_single : g.Area_Name %v\n", g.Area_Name )
//	fmt.Fprintf( w, "deliver_update_single : g.Course_No %v\n", g.Course_No )
//	fmt.Fprintf( w, "deliver_update_single : g.Private_No %v\n", g.Private_No )
//	fmt.Fprintf( w, "deliver_update_single : g.Car_No %v\n", g.Car_No )
//	fmt.Fprintf( w, "deliver_update_single : g.Number %v\n", g.Number )

    if _, err = client.Put(ctx, key, &g ); err != nil {

		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

}

