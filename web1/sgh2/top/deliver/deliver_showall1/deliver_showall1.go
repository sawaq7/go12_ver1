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

//    fmt.Fprintf( w, "deliver_showall1 start \n" )  // チE��チE��

    var course_no int64

	var g type2.Deliver

	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)

	g.Area_Name = r.FormValue("area_name")  // 配達エリア名をゲチE��

	number := r.FormValue("number")         // 配達個数をゲチE��
//	fmt.Fprintf( w, "deliver_showall1 : number %v\n", number )  // チE��チE��

	numberw ,err := strconv.Atoi(number)  // 配達個数の整数匁E	if err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

	g.Number = int64(numberw)   // 整数の64ビット化

	private_no := r.FormValue("private_no")         // 個人NoをゲチE��
//	fmt.Fprintf( w, "deliver_showall1 : private_no %v\n", private_no )  // チE��チE��

	private_now ,err := strconv.Atoi(private_no)  // 個人Noの整数匁E	if err != nil {
//		http.Error(w,err.Error(), http.StatusInternalServerError)
//       fmt.Fprintf( w, "deliver_showall1 : a private_no must be half-width characters %v\n"  )
		return
	}

	g.Private_No = int64(private_now)   // 整数の64ビット化

	car_no := r.FormValue("car_no")         // 個人NoをゲチE��
//	fmt.Fprintf( w, "deliver_showall1 : car_no %v\n", car_no )  // チE��チE��

	car_now ,err := strconv.Atoi(car_no)  // 個人Noの整数匁E	if err != nil {
//		http.Error(w,err.Error(), http.StatusInternalServerError)
//       fmt.Fprintf( w, "deliver_showall1 : a car_no must be half-width characters %v\n"  )
		return
	}

	g.Car_No = int64(car_now)   // 整数の64ビット化

    date_w := time.Now()        // 日付をセチE��
    g.Date_Real = date_w
//    date_test := fmt.Sprintf("%04d%02d%02d%02d%02d%02d",
//		date_w.Year(), date_w.Month(),date_w.Day(), date_w.Hour(), date_w.Minute(), date_w.Second())
//   fmt.Fprintf( w, "deliver_showall1 : date_test %v\n", date_test )  // チE��チE��

    g.Date = fmt.Sprintf("%04d/%02d/%02d",date_w.Year(), date_w.Month(),date_w.Day())

/// 一時ファイルより、エリア惁E��をセチE��　///

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

/// チE�EタストアーにチE�EタをセチE�� ///

    new_key := datastore.IncompleteKey("Deliver", nil)

    if _, err = client.Put(ctx, new_key, &g ); err != nil {
//	if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Deliver", nil), &g); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

//   	fmt.Fprintf( w, "deliver_showall1 : g.Area_Name %v\n", g.Area_Name )  // チE��チE��
//    fmt.Fprintf( w, "deliver_showall1 : g.Number %v\n", g.Number )  // チE��チE��
//    fmt.Fprintf( w, "deliver_showall1 : g.Date %v\n", g.Date )  // チE��チE��

/// モニター　再表示 ///

    process.Deliver_showall2(course_no ,w , r )

//	fmt.Fprintf( w, "deliver_showall1 : normal end \n" )  // チE��チE��




}
