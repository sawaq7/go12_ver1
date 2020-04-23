package analysis1

import (


	    "net/http"
	    "fmt"
        "github.com/sawaq7/go12_ver1/basic/date1"
//	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2/trans"
	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2"
//	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2/sort"
	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
	    "github.com/sawaq7/go12_ver1/general/type5"

//	    "github.com/sawaq7/go12_ver1/basic/type3"
	    "time"
//	    "strings"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                )

///
///    expect delivery number by least squares method
///

func Deliver_growth_rate2( course_no int64 ,w http.ResponseWriter, r *http.Request ) {

//     IN  course_no  : コースNo
//     IN    w      　: レスポンスライター
//     IN    r      　: リクエストパラメータ

    var sgh_ai type2.Sgh_Ai
//    fmt.Fprintf( w, "analysis1.deliver_growth_rate2 start \n" )

     general_work := make([]type5.General_Work, 2)
     general_work[0].Int64_Work = 0
     general_work[1].Int64_Work = course_no

     deliver_view := datastore2.Datastore_sgh( "Deliver" ,"trans" ,general_work , w , r  )

     //   get value from interface data

     value, _ := deliver_view.([]type2.Deliver)

     sgh_ai.Date_Basic = "2017/01/01"
     sgh_ai.Date_Basic_Real = date1.Date_realdata_get( w  ,sgh_ai.Date_Basic )   //  make date data by time type

     date_w := time.Now()

//     fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : date_w %v\n", date_w.Year(), date_w.Month(),date_w.Day( ) )

     date_max := fmt.Sprintf("%04d/%02d/%02d" ,date_w.Year(),  date_w.Month(),date_w.Day())

//     date_min := fmt.Sprintf("%04d/%02d/%02d" ,date_w.Year()-1, date_w.Month(),date_w.Day())
     date_min := fmt.Sprintf( "2017/01/01" )

     nn        := 0
     siguma_x  := 0.
     siguma_y  := 0.
     siguma_xx := 0.
     siguma_yy := 0.
     siguma_xy := 0.

///
///  make expression by least squares method
///　

     for _, deliverw := range value {

       if deliverw.Date <=  date_max &&         //  check whether or not in range
          deliverw.Date > date_min     {

          nn ++

          date_data := date1.Date_realdata_get( w  ,deliverw.Date )   //  make time data

          date_sub := date_data.Sub(sgh_ai.Date_Basic_Real)  //  calcurate difference with the basic date

//          fmt.Fprintf( w, "analysis1.deliver_growth_rate : NUM %v\n", num )
//          fmt.Fprintf( w, "analysis1.deliver_growth_rate : XX %f\n", xx )
//          fmt.Fprintf( w, "analysis1.deliver_growth_rate : RATERATE %f\n", num/xx )

          siguma_x   = siguma_x  + float64(date_sub/(3600000000000*24))
          siguma_y   = siguma_y  + float64(deliverw.Number)
          siguma_xx  = siguma_xx + float64(date_sub/(3600000000000*24)) *  float64(date_sub/(3600000000000*24))
          siguma_yy  = siguma_yy + float64(deliverw.Number)      *  float64(deliverw.Number)
          siguma_xy  = siguma_xy + float64(date_sub/(3600000000000*24)) *  float64(deliverw.Number)

//          fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : X %v\n", float64(date_sub/10000000000000) )
//          fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : Y %v\n", float64(deliverw.Number*10) )
//          fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : XX %v\n", float64(date_sub/10000000000000) *  float64(date_sub/10000000000000) )
//          fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : YY %v\n",  float64(deliverw.Number*10) *  float64(deliverw.Number*10))
	   }
	}

//	fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : SIGUMA.X %f\n"  ,  siguma_x )
//    fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : SIGUMA.Y %f\n"  ,  siguma_y )
//    fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : SIGUMA.XX %f\n" ,  siguma_xx)
//    fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : SIGUMA.YY %f\n" ,  siguma_yy )
//    fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : SIGUMA.XY %f\n" ,  siguma_xy )

///
/// 最小二乗法�E勾配と刁E��を計算すめE///

    aa := ( float64(nn) * siguma_xy - siguma_x * siguma_y ) / (  float64(nn) * siguma_xx - siguma_x * siguma_x )

    bb := ( siguma_xx * siguma_y - siguma_xy * siguma_x ) / (  float64(nn) * siguma_xx - siguma_x * siguma_x )

//    fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : AA %f\n" ,  aa )
//    fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : BB %f\n" ,  bb )

///
///    register expression in ai file
///

     for pos, deliverw := range value {
       if pos == 0  {      //   set header in ai file

          sgh_ai.Course_No     = course_no
          sgh_ai.District_No   = deliverw.District_No
          sgh_ai.District_Name = deliverw.District_Name
          sgh_ai.Area_No       = deliverw.Area_No
	      sgh_ai.Area_Name     = deliverw.Area_Name

       }
	 }

    sgh_ai.Ex_Type       = "function-001"
    sgh_ai.Expression    = fmt.Sprintf( "Y=%fX+%f",aa ,bb)  // make expression
    sgh_ai.Item_Num      = 2
	sgh_ai.Item1_Name    = "*"
	sgh_ai.Item1_Factor  = aa
	sgh_ai.Item2_Name    = "+"
	sgh_ai.Item2_Factor  = bb

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return
    }

    new_key := datastore.IncompleteKey("Sgh_Ai", nil)

    if _, err = client.Put(ctx, new_key, &sgh_ai ); err != nil {

		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}


//    fmt.Fprintf( w, "analysis1.deliver_growth_rate2 normal end \n" )
}
