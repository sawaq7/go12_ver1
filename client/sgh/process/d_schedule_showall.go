package process

import (

	    "net/http"
//	    "fmt"
//	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2/initialize"
	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2/trans"
	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2"
	    "github.com/sawaq7/go12_ver1/client/sgh/ai"
	    "github.com/sawaq7/go12_ver1/client/sgh/ai/cal2"
	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2/cal3"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/sgh/html2"
	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
	    "strconv"

        "cloud.google.com/go/datastore"
        "context"
        "os"

//	    "time"
                                                )

///
///   　　　　make deliver schdule
///


func D_schedule_showall(w http.ResponseWriter, r *http.Request, district_no int64) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　district_no : district no


//    fmt.Fprintf( w, "d_schedule_showall start \n" )
//    fmt.Fprintf( w, "d_schedule_showall district_no \n" ,district_no)

    var course_no , car_no , car_num int64

    var expected_num , ability_num float64

    var d_schedule_headline type2.D_Schedule //   allocate work area for headline

///
///      make deliver schdule
///

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

	query := datastore.NewQuery("D_Schedule").Order("Date")

	count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	d_schedule      := make([]type2.D_Schedule, 0, count)
	d_schedule_view := make([]type2.D_Schedule, 0)

    keys, err := client.GetAll(ctx, query , &d_schedule)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    keys_wk := make([]int64, count)

    for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }
///
///      get area inf in d.s.
///

	d_area := datastore2.Datastore_sgh( "D_Area","trans2" ,district_no , w , r  )

	// get value from interface data

    d_area_value, _ := d_area.([]type2.D_Area)

	count2 := 0
	ability_num = 0.

	for pos, d_schedulew := range d_schedule {
     if district_no == d_schedulew.District_No {

        count2 ++

        if count2 == 1 {

          for _, d_area_valuew := range d_area_value {

            course_no = district_no * 100 + d_area_valuew.Area_No

///
///      renewal expression for expecting deliver number
///

            datastore2.Datastore_sgh( "Sgh_Ai","initialize" ,course_no , w , r  )

            ai.Ai_sgh(course_no ,w , r )              //　make expression for each area and put itin d.s.

///
///        calculate Productivity for each area
///

            d_area_valuew.Number_Total ,d_area_valuew.Time_Total ,d_area_valuew.Productivity = cal3.Deliver_course_no( course_no ,w , r   ) // エリア生産性を算�E

            if d_area_valuew.Productivity >= 0. {

              ability_num = ability_num + d_area_valuew.Productivity * 10.

		    }

//            fmt.Fprintf( w, "d_schedule_showall 配達能劁E %f\n" ,ability_num )
//            fmt.Fprintf( w, "d_schedule_showall 配達能劁E-P %f\n" ,d_area_valuew.Productivity )

            key := datastore.IDKey("D_Area", d_area_valuew.Id, nil)

            if _, err := client.Put(ctx, key, &d_area_valuew ); err != nil {

              http.Error(w,err.Error(), http.StatusInternalServerError)
		      return
		    }

          }

        }

///
///        calculate deliver number which is expected for each area
///

        for _, d_area_valuew := range d_area_value {

          course_no = district_no * 100 + d_area_valuew.Area_No

          expected_num = cal2.Ai_sgh_analysis( w , r ,course_no ,d_schedulew.Date )
          d_schedulew.Expected_Num = d_schedulew.Expected_Num + expected_num

        }

///
///      judge schedule
///
        ability_per := ability_num / d_schedulew.Expected_Num  * 100.

//        fmt.Fprintf( w, "d_schedule_showall 配達能劁E %f\n" ,ability_num )
//        fmt.Fprintf( w, "d_schedule_showall 予想個数 %f\n" ,d_schedulew.Expected_Num )
//        fmt.Fprintf( w, "d_schedule_showall 判定　パ�EセンチE%f\n" ,ability_per )

        if ability_per >= 100. {

           d_schedulew.Judge = "this member can excute this job"

	    }else if ability_per  >= 80. {

           d_schedulew.Judge = "you should change deliver's member for this job"

	    }else  {

	       d_schedulew.Judge = "I'm afraid ,this member can't excute this job"

	    }

///
///        make head line
///
        if count2 == 1 {

          car_district := trans.Car_district( district_no ,w , r  )           //   get car inf.

//          fmt.Fprintf( w, "d_schedule_showall car number \n" ,len(car_district))

          car_num = int64 ( len(car_district) )
          d_schedule_headline.Id = 77     //    set　number "77" for juding headline
          d_schedule_headline.Course_Num = car_num

          for pos2, car_districtw := range car_district {

///
///  　　　make expression for each car and put itin d.s.
///
            car_no ,_ =strconv.ParseInt( car_districtw.Car_Name ,10 ,64)   // make an integer64

            car_districtw.Number_Total ,car_districtw.Time_Total ,car_districtw.Productivity = cal3.Deliver_car_no ( car_no ,w , r   )

            key := datastore.IDKey("Car", car_districtw.Id, nil)

            if _, err := client.Put(ctx, key, &car_districtw ); err != nil {

		      http.Error(w,err.Error(), http.StatusInternalServerError)
		      return
		    }

            if pos2 == 0 {                       // set car name

              d_schedule_headline.Car_Name_01 = car_districtw.Car_Name

	        }else if pos2 == 1 {

	          d_schedule_headline.Car_Name_02 = car_districtw.Car_Name

	        }else if pos2 == 2 {

	          d_schedule_headline.Car_Name_03 = car_districtw.Car_Name

	        }else if pos2 == 3 {

	          d_schedule_headline.Car_Name_04 = car_districtw.Car_Name

	        }else if pos2 == 4 {

	          d_schedule_headline.Car_Name_05 = car_districtw.Car_Name

            }else if pos2 == 5 {

              d_schedule_headline.Car_Name_06 = car_districtw.Car_Name

	        }else if pos2 == 6 {

	          d_schedule_headline.Car_Name_07 = car_districtw.Car_Name

	        }else if pos2 == 7 {

	          d_schedule_headline.Car_Name_08 = car_districtw.Car_Name

	        }else if pos2 == 8 {

	          d_schedule_headline.Car_Name_09 = car_districtw.Car_Name

	        }else if pos2 == 9 {

	          d_schedule_headline.Car_Name_10 = car_districtw.Car_Name

	        }


          }
///
///     set head line
///
          d_schedule_view = append(d_schedule_view, type2.D_Schedule {  d_schedule_headline.Id        ,

                                                                    d_schedule_headline.District_No   ,
                                                                    d_schedule_headline.District_Name ,
                                                                    d_schedule_headline.Date          ,
                                                                    d_schedule_headline.Date_Real     ,
                                                                    d_schedule_headline.Expected_Num  ,
                                                                    d_schedule_headline.Judge         ,
                                                                    d_schedule_headline.Course_Num    ,
                                                                    d_schedule_headline.Course_01     ,
                                                                    d_schedule_headline.Car_Name_01   ,
                                                                    d_schedule_headline.Course_02     ,
                                                                    d_schedule_headline.Car_Name_02   ,
                                                                    d_schedule_headline.Course_03     ,
                                                                    d_schedule_headline.Car_Name_03   ,
                                                                    d_schedule_headline.Course_04     ,
                                                                    d_schedule_headline.Car_Name_04   ,
                                                                    d_schedule_headline.Course_05     ,
                                                                    d_schedule_headline.Car_Name_05   ,
                                                                    d_schedule_headline.Course_06     ,
                                                                    d_schedule_headline.Car_Name_06   ,
                                                                    d_schedule_headline.Course_07     ,
                                                                    d_schedule_headline.Car_Name_07   ,
                                                                    d_schedule_headline.Course_08     ,
                                                                    d_schedule_headline.Car_Name_08   ,
                                                                    d_schedule_headline.Course_09     ,
                                                                    d_schedule_headline.Car_Name_09   ,
                                                                    d_schedule_headline.Course_10     ,
                                                                    d_schedule_headline.Car_Name_10        })


        }
///
///      set schedule inf.
///
        d_schedulew.Course_Num = car_num     //  cource number

        d_schedule_view = append(d_schedule_view, type2.D_Schedule { keys_wk[pos]     ,

                                                                    d_schedulew.District_No   ,
                                                                    d_schedulew.District_Name ,
                                                                    d_schedulew.Date          ,
                                                                    d_schedulew.Date_Real     ,
                                                                    d_schedulew.Expected_Num  ,
                                                                    d_schedulew.Judge         ,
                                                                    d_schedulew.Course_Num    ,
                                                                    d_schedulew.Course_01     ,
                                                                    d_schedulew.Car_Name_01   ,
                                                                    d_schedulew.Course_02     ,
                                                                    d_schedulew.Car_Name_02   ,
                                                                    d_schedulew.Course_03     ,
                                                                    d_schedulew.Car_Name_03   ,
                                                                    d_schedulew.Course_04     ,
                                                                    d_schedulew.Car_Name_04   ,
                                                                    d_schedulew.Course_05     ,
                                                                    d_schedulew.Car_Name_05   ,
                                                                    d_schedulew.Course_06     ,
                                                                    d_schedulew.Car_Name_06   ,
                                                                    d_schedulew.Course_07     ,
                                                                    d_schedulew.Car_Name_07   ,
                                                                    d_schedulew.Course_08     ,
                                                                    d_schedulew.Car_Name_08   ,
                                                                    d_schedulew.Course_09     ,
                                                                    d_schedulew.Car_Name_09   ,
                                                                    d_schedulew.Course_10     ,
                                                                    d_schedulew.Car_Name_10        })


//            fmt.Fprintf( w, "d_schedule_showall pos %v   \n" , pos  )
      }
	}

//  set template

       monitor := template.Must(template.New("html").Parse(html2.D_schedule_showall_04))

     if car_num >= 5 {

       monitor = template.Must(template.New("html").Parse(html2.D_schedule_showall_05))

     }

//     show schedule inf. on web

	err = monitor.Execute(w, d_schedule_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

