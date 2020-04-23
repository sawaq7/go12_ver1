package deliver_showall1

import (

	    "net/http"
	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"

	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
	    "strconv"

	    "cloud.google.com/go/datastore"
	    "context"
	    "os"
	    "time"

                                                  )

///
/// 　　   show delivery inf.
///

func Deliver_showall1(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "deliver_showall1 start \n" )

    var course_no int64

	var g type2.Deliver

	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)

	g.Area_Name = r.FormValue("area_name")  // get delivery area name

	number := r.FormValue("number")
//	fmt.Fprintf( w, "deliver_showall1 : number %v\n", number )

	numberw ,err := strconv.Atoi(number)
	if err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

	g.Number = int64(numberw)

	private_no := r.FormValue("private_no")
//	fmt.Fprintf( w, "deliver_showall1 : private_no %v\n", private_no )

	private_now ,err := strconv.Atoi(private_no)
	if err != nil {
//		http.Error(w,err.Error(), http.StatusInternalServerError)
//       fmt.Fprintf( w, "deliver_showall1 : a private_no must be half-width characters %v\n"  )
		return
	}

	g.Private_No = int64(private_now)   // 整数の64ビット化

	car_no := r.FormValue("car_no")
//	fmt.Fprintf( w, "deliver_showall1 : car_no %v\n", car_no )

	car_now ,err := strconv.Atoi(car_no)
	if err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
//       fmt.Fprintf( w, "deliver_showall1 : a car_no must be half-width characters %v\n"  )
		return
	}

	g.Car_No = int64(car_now)

    date_w := time.Now()        //    set date data
    g.Date_Real = date_w
//    date_test := fmt.Sprintf("%04d%02d%02d%02d%02d%02d",
//		date_w.Year(), date_w.Month(),date_w.Day(), date_w.Hour(), date_w.Minute(), date_w.Second())
//   fmt.Fprintf( w, "deliver_showall1 : date_test %v\n", date_test )

    g.Date = fmt.Sprintf("%04d/%02d/%02d",date_w.Year(), date_w.Month(),date_w.Day())

/// 　   set area data from temp.-file

    query := datastore.NewQuery("D_Area_Temp").Order("Area_No")

	count, err := client.Count(ctx, query)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	d_area      := make([]type2.D_Area_Temp, 0, count)

    if _, err := client.GetAll(ctx, query , &d_area);  err != nil {

	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	} else{
      for pos, d_areaw := range d_area {
        if pos == 0 {
           g.Course_No     = d_areaw.District_No * 100 + d_areaw.Area_No
           course_no      = g.Course_No
           g.District_No   = d_areaw.District_No
           g.District_Name = d_areaw.District_Name
           g.Area_No       = d_areaw.Area_No
           g.Area_Name     = d_areaw.Area_Name
        }
	  }
	}

///     put delivery data in d.s.

    new_key := datastore.IncompleteKey("Deliver", nil)

    if _, err = client.Put(ctx, new_key, &g ); err != nil {

		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

//   	fmt.Fprintf( w, "deliver_showall1 : g.Area_Name %v\n", g.Area_Name )
//    fmt.Fprintf( w, "deliver_showall1 : g.Number %v\n", g.Number )
//    fmt.Fprintf( w, "deliver_showall1 : g.Date %v\n", g.Date )

    //    show delivery inf.  on web
    process.Deliver_showall2(course_no ,w , r )

//	fmt.Fprintf( w, "deliver_showall1 : normal end \n" )

}
