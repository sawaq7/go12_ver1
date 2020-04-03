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
///   　　　　　配達スケジュールを表示する
///


func D_schedule_showall(w http.ResponseWriter, r *http.Request, district_no int64) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　district_no : 地区No


//    fmt.Fprintf( w, "d_schedule_showall start \n" )  // チE��チE��
//    fmt.Fprintf( w, "d_schedule_showall district_no \n" ,district_no)  // チE��チE��

    var course_no , car_no , car_num int64

    var expected_num , ability_num float64

    var d_schedule_headline type2.D_Schedule // 見�Eし�Eワークエリア確俁E
///
/// スケジュールチE�EタをゲチE��
///
    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

    ctx := context.Background()
//	c := appengine.NewContext(r)

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return
    }

//	q := datastore.NewQuery("D_Schedule").Order("Date")
	query := datastore.NewQuery("D_Schedule").Order("Date")

//	count, err := q.Count(c)
	count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	d_schedule      := make([]type2.D_Schedule, 0, count)
	d_schedule_view := make([]type2.D_Schedule, 0)

    keys, err := client.GetAll(ctx, query , &d_schedule)
//	keys, err := q.GetAll(c, &d_schedule)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    keys_wk := make([]int64, count)

    for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }
///
/// エリアチE�EタをゲチE��
///

	d_area := datastore2.Datastore_sgh( "D_Area","trans2" ,district_no , w , r  )

	// 空インターフェイス変数よりバリュー値をゲチE��

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
///        荷物予想のための、計算式をリニューアル
///

            datastore2.Datastore_sgh( "Sgh_Ai","initialize" ,course_no , w , r  )

            ai.Ai_sgh(course_no ,w , r )              //　エリアごとに計算式を算�E

///
///        エリアごとの配達生産性等�EチE�Eタを算�Eし、エリアファイルにセチE��
///

            d_area_valuew.Number_Total ,d_area_valuew.Time_Total ,d_area_valuew.Productivity = cal3.Deliver_course_no( course_no ,w , r   ) // エリア生産性を算�E

            if d_area_valuew.Productivity >= 0. {

              ability_num = ability_num + d_area_valuew.Productivity * 10.

		    }

//            fmt.Fprintf( w, "d_schedule_showall 配達能劁E %f\n" ,ability_num )  // チE��チE��
//            fmt.Fprintf( w, "d_schedule_showall 配達能劁E-P %f\n" ,d_area_valuew.Productivity )  // チE��チE��

            key := datastore.IDKey("D_Area", d_area_valuew.Id, nil)

            if _, err := client.Put(ctx, key, &d_area_valuew ); err != nil {

//            key := datastore.NewKey(c, "D_Area", "", d_area_valuew.Id, nil)
//	        if _, err := datastore.Put(c, key, &d_area_valuew); err != nil {  // チE�Eタストアから1レコードアチE�EチE�EチE		      http.Error(w,err.Error(), http.StatusInternalServerError)
		      return
		    }

          }


        }

///
///        配達地区の予想荷物数を算�E(エリア単位�E配達予想数のΣ�E�E///
        for _, d_area_valuew := range d_area_value {

          course_no = district_no * 100 + d_area_valuew.Area_No

          expected_num = cal2.Ai_sgh_analysis( w , r ,course_no ,d_schedulew.Date )
          d_schedulew.Expected_Num = d_schedulew.Expected_Num + expected_num

        }

///
///     スケジュールを評価する
///
        ability_per := ability_num / d_schedulew.Expected_Num  * 100.

//        fmt.Fprintf( w, "d_schedule_showall 配達能劁E %f\n" ,ability_num )  // チE��チE��
//        fmt.Fprintf( w, "d_schedule_showall 予想個数 %f\n" ,d_schedulew.Expected_Num )  // チE��チE��
//        fmt.Fprintf( w, "d_schedule_showall 判定　パ�EセンチE%f\n" ,ability_per )  // チE��チE��

        if ability_per >= 100. {                       // 号車名をセチE��

           d_schedulew.Judge = "this member can excute this job"

	    }else if ability_per  >= 80. {

           d_schedulew.Judge = "you should change deliver's member for this job"

	    }else  {

	       d_schedulew.Judge = "I'm afraid ,this member can't excute this job"

	    }

///
///     先頭の見�Eしを作�E
///
        if count2 == 1 {

          car_district := trans.Car_district( district_no ,w , r  )           // 号車情報をゲチE��

//          fmt.Fprintf( w, "d_schedule_showall 号車数 \n" ,len(car_district))  // チE��チE��

          car_num = int64 ( len(car_district) )
          d_schedule_headline.Id = 77     // 見�Eし用のID　”７７”　をセチE��
          d_schedule_headline.Course_Num = car_num

          for pos2, car_districtw := range car_district {

///
///  　　　号車ごとの配達生産性等�EチE�Eタを算�Eし、号車ファイルにセチE��
///
            car_no ,_ =strconv.ParseInt( car_districtw.Car_Name ,10 ,64)   // 斁E���Eをint64に変換

            car_districtw.Number_Total ,car_districtw.Time_Total ,car_districtw.Productivity = cal3.Deliver_car_no ( car_no ,w , r   )

            key := datastore.IDKey("Car", car_districtw.Id, nil)

            if _, err := client.Put(ctx, key, &car_districtw ); err != nil {                                                                // 号車生産性を算�E
//            key := datastore.NewKey(c, "Car", "", car_districtw.Id, nil)
//	        if _, err := datastore.Put(c, key, &car_districtw); err != nil {
		      http.Error(w,err.Error(), http.StatusInternalServerError)
		      return
		    }

            if pos2 == 0 {                       // 号車名をセチE��

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
///     先頭の見�EしをセチE��
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
///     スケジュールチE�EタをセチE��
///
        d_schedulew.Course_Num = car_num     // コース数セチE��

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


//            fmt.Fprintf( w, "d_schedule_showall pos %v   \n" , pos  )  // チE��チE��
      }
	}

// チE��プレート�EヘッダーをGET


       monitor := template.Must(template.New("html").Parse(html2.D_schedule_showall_04))

     if car_num >= 5 {

       monitor = template.Must(template.New("html").Parse(html2.D_schedule_showall_05))

     }

// モニターに表示

	err = monitor.Execute(w, d_schedule_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

