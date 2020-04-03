package analysis1

import (

//        "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
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
///    最小二乗法で、荷物の成長玁E��算�Eする
///    (式�EAIファイルに登録�E�E///

func Deliver_growth_rate2( course_no int64 ,w http.ResponseWriter, r *http.Request ) {

//     IN  course_no  : コースNo
//     IN    w      　: レスポンスライター
//     IN    r      　: リクエストパラメータ

//    var f_dmy ,f_dmy2 ,f_dmy3 float64

    var sgh_ai type2.Sgh_Ai   // AIファイル用のフォーマットを持E��E
//    fmt.Fprintf( w, "analysis1.deliver_growth_rate2 start \n" )  // チE��チE��

// チE�Eタストアーから、該当するコースNo.のチE�EタをGET

//     deliver_view := trans.Deliver ( 0 ,course_no ,w ,r ) /// セレクトデータをＧ�E��E�
//     deliver_view2 := sort.Deliver ( w ,deliver_view )       /// 2重ソーチE日付�E号車！E
     general_work := make([]type5.General_Work, 2)
     general_work[0].Int64_Work = 0          // 地区惁E��
     general_work[1].Int64_Work = course_no  //　コースNO

     deliver_view := datastore2.Datastore_sgh( "Deliver" ,"trans" ,general_work , w , r  )

     // 空インターフェイス変数よりバリュー値をゲチE��

     value, _ := deliver_view.([]type2.Deliver)

     sgh_ai.Date_Basic = "2017/01/01"
     sgh_ai.Date_Basic_Real = date1.Date_realdata_get( w  ,sgh_ai.Date_Basic )   // タイムチE�Eタ作�E

     date_w := time.Now()        // 日付をゲチE��

//     fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : date_w %v\n", date_w.Year(), date_w.Month(),date_w.Day( ) ) // チE��チE��

     date_max := fmt.Sprintf("%04d/%02d/%02d" ,date_w.Year(),  date_w.Month(),date_w.Day())

//     date_min := fmt.Sprintf("%04d/%02d/%02d" ,date_w.Year()-1, date_w.Month(),date_w.Day())
     date_min := fmt.Sprintf( "2017/01/01" )

///
/// 最小二乗法�E作業チE�Eタの初期匁E///

     nn        := 0
     siguma_x  := 0.
     siguma_y  := 0.
     siguma_xx := 0.
     siguma_yy := 0.
     siguma_xy := 0.

///
/// 最小二乗法�EチE�Eタを作�Eする
///　

     for _, deliverw := range value {

       if deliverw.Date <=  date_max &&         // 篁E��冁E�EチE�EタかチェチE��
          deliverw.Date > date_min     {

          nn ++

          date_data := date1.Date_realdata_get( w  ,deliverw.Date )   // タイムチE�Eタ作�E

          date_sub := date_data.Sub(sgh_ai.Date_Basic_Real)  // 基準日との差を計算　

//          fmt.Fprintf( w, "analysis1.deliver_growth_rate : NUM %v\n", num ) // チE��チE��
//          fmt.Fprintf( w, "analysis1.deliver_growth_rate : XX %f\n", xx ) // チE��チE��
//          fmt.Fprintf( w, "analysis1.deliver_growth_rate : RATERATE %f\n", num/xx ) // チE��チE��

          siguma_x   = siguma_x  + float64(date_sub/(3600000000000*24))
          siguma_y   = siguma_y  + float64(deliverw.Number)
          siguma_xx  = siguma_xx + float64(date_sub/(3600000000000*24)) *  float64(date_sub/(3600000000000*24))
          siguma_yy  = siguma_yy + float64(deliverw.Number)      *  float64(deliverw.Number)
          siguma_xy  = siguma_xy + float64(date_sub/(3600000000000*24)) *  float64(deliverw.Number)

//          fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : X %v\n", float64(date_sub/10000000000000) )  // チE��チE��
//          fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : Y %v\n", float64(deliverw.Number*10) )  // チE��チE��
//          fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : XX %v\n", float64(date_sub/10000000000000) *  float64(date_sub/10000000000000) )  // チE��チE��
//          fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : YY %v\n",  float64(deliverw.Number*10) *  float64(deliverw.Number*10))  // チE��チE��
	   }
	}

//	fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : SIGUMA.X %f\n"  ,  siguma_x )  // チE��チE��
//    fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : SIGUMA.Y %f\n"  ,  siguma_y )  // チE��チE��
//    fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : SIGUMA.XX %f\n" ,  siguma_xx)  // チE��チE��
//    fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : SIGUMA.YY %f\n" ,  siguma_yy )  // チE��チE��
//    fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : SIGUMA.XY %f\n" ,  siguma_xy )  // チE��チE��

///
/// 最小二乗法�E勾配と刁E��を計算すめE///

    aa := ( float64(nn) * siguma_xy - siguma_x * siguma_y ) / (  float64(nn) * siguma_xx - siguma_x * siguma_x )

    bb := ( siguma_xx * siguma_y - siguma_xy * siguma_x ) / (  float64(nn) * siguma_xx - siguma_x * siguma_x )

//    fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : AA %f\n" ,  aa )  // チE��チE��
//    fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : BB %f\n" ,  bb )  // チE��チE��

///
/// 条件式をAIファイルに登録する
/// (刁E��bbは省略�E�E
     for pos, deliverw := range value {
       if pos == 0  {      // AIファイル用の惁E��をセチE��

          sgh_ai.Course_No     = course_no
          sgh_ai.District_No   = deliverw.District_No
          sgh_ai.District_Name = deliverw.District_Name
          sgh_ai.Area_No       = deliverw.Area_No
	      sgh_ai.Area_Name     = deliverw.Area_Name

       }
	 }

    sgh_ai.Ex_Type       = "function-001"
    sgh_ai.Expression    = fmt.Sprintf( "Y=%fX+%f",aa ,bb)  // 条件式�E作�E
    sgh_ai.Item_Num      = 2
	sgh_ai.Item1_Name    = "*"
	sgh_ai.Item1_Factor  = aa
	sgh_ai.Item2_Name    = "+"
	sgh_ai.Item2_Factor  = bb

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return
    }

    new_key := datastore.IncompleteKey("Sgh_Ai", nil)

    if _, err = client.Put(ctx, new_key, &sgh_ai ); err != nil {
//    if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Sgh_Ai", nil), &sgh_ai); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}


//    fmt.Fprintf( w, "analysis1.deliver_growth_rate2 normal end \n" )  // チE��チE��
}
