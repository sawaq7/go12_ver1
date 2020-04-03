package deliver_showall1

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
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

func Deliver_showall1(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "deliver_showall1 start \n" )  // ใEใEฏ

    var course_no int64

	var g type2.Deliver

	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)

	g.Area_Name = r.FormValue("area_name")  // ้้ใจใชใขๅใใฒใE

	number := r.FormValue("number")         // ้้ๅๆฐใใฒใE
//	fmt.Fprintf( w, "deliver_showall1 : number %v\n", number )  // ใEใEฏ

	numberw ,err := strconv.Atoi(number)  // ้้ๅๆฐใฎๆดๆฐๅE	if err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

	g.Number = int64(numberw)   // ๆดๆฐใฎ64ใใใๅ

	private_no := r.FormValue("private_no")         // ๅไบบNoใใฒใE
//	fmt.Fprintf( w, "deliver_showall1 : private_no %v\n", private_no )  // ใEใEฏ

	private_now ,err := strconv.Atoi(private_no)  // ๅไบบNoใฎๆดๆฐๅE	if err != nil {
//		http.Error(w,err.Error(), http.StatusInternalServerError)
//       fmt.Fprintf( w, "deliver_showall1 : a private_no must be half-width characters %v\n"  )
		return
	}

	g.Private_No = int64(private_now)   // ๆดๆฐใฎ64ใใใๅ

	car_no := r.FormValue("car_no")         // ๅไบบNoใใฒใE
//	fmt.Fprintf( w, "deliver_showall1 : car_no %v\n", car_no )  // ใEใEฏ

	car_now ,err := strconv.Atoi(car_no)  // ๅไบบNoใฎๆดๆฐๅE	if err != nil {
//		http.Error(w,err.Error(), http.StatusInternalServerError)
//       fmt.Fprintf( w, "deliver_showall1 : a car_no must be half-width characters %v\n"  )
		return
	}

	g.Car_No = int64(car_now)   // ๆดๆฐใฎ64ใใใๅ

    date_w := time.Now()        // ๆฅไปใใปใE
    g.Date_Real = date_w
//    date_test := fmt.Sprintf("%04d%02d%02d%02d%02d%02d",
//		date_w.Year(), date_w.Month(),date_w.Day(), date_w.Hour(), date_w.Minute(), date_w.Second())
//   fmt.Fprintf( w, "deliver_showall1 : date_test %v\n", date_test )  // ใEใEฏ

    g.Date = fmt.Sprintf("%04d/%02d/%02d",date_w.Year(), date_w.Month(),date_w.Day())

/// ไธๆใใกใคใซใใใใจใชใขๆE ฑใใปใEใ///

    query := datastore.NewQuery("D_Area_Temp").Order("Area_No")

//	count, err := q.Count(c)
	count, err := client.Count(ctx, query)
//	q := datastore.NewQuery("D_Area_Temp").Order("Area_No")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	d_area      := make([]type2.D_Area_Temp, 0, count)

    if _, err := client.GetAll(ctx, query , &d_area);  err != nil {
//	if _, err := q.GetAll(c, &d_area);  err != nil {

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

/// ใEEใฟในใใขใผใซใEEใฟใใปใE ///

    new_key := datastore.IncompleteKey("Deliver", nil)

    if _, err = client.Put(ctx, new_key, &g ); err != nil {
//	if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Deliver", nil), &g); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

//   	fmt.Fprintf( w, "deliver_showall1 : g.Area_Name %v\n", g.Area_Name )  // ใEใEฏ
//    fmt.Fprintf( w, "deliver_showall1 : g.Number %v\n", g.Number )  // ใEใEฏ
//    fmt.Fprintf( w, "deliver_showall1 : g.Date %v\n", g.Date )  // ใEใEฏ

/// ใขใใฟใผใๅ่กจ็คบ ///

    process.Deliver_showall2(course_no ,w , r )

//	fmt.Fprintf( w, "deliver_showall1 : normal end \n" )  // ใEใEฏ




}
