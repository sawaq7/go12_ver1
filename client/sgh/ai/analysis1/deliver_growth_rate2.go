package analysis1

import (

//        "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
	    "fmt"
        "basic/date1"
//	    "client/sgh/datastore2/trans"
	    "client/sgh/datastore2"
//	    "client/sgh/datastore2/sort"
	    "client/sgh/type2"
	    "general/type5"

//	    "basic/type3"
	    "time"
//	    "strings"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                )

///
///    最小二乗法で、荷物の成長率を算出する
///    (式はAIファイルに登録）
///

func Deliver_growth_rate2( course_no int64 ,w http.ResponseWriter, r *http.Request ) {

//     IN  course_no  : コースNo
//     IN    w      　: レスポンスライター
//     IN    r      　: リクエストパラメータ

//    var f_dmy ,f_dmy2 ,f_dmy3 float64

    var sgh_ai type2.Sgh_Ai   // AIファイル用のフォーマットを指定

//    fmt.Fprintf( w, "analysis1.deliver_growth_rate2 start \n" )  // デバック

// データストアーから、該当するコースNo.のデータをGET

//     deliver_view := trans.Deliver ( 0 ,course_no ,w ,r ) /// セレクトデータをＧＥＴ
//     deliver_view2 := sort.Deliver ( w ,deliver_view )       /// 2重ソート(日付・号車）

     general_work := make([]type5.General_Work, 2)
     general_work[0].Int64_Work = 0          // 地区情報
     general_work[1].Int64_Work = course_no  //　コースNO

     deliver_view := datastore2.Datastore_sgh( "Deliver" ,"trans" ,general_work , w , r  )

     // 空インターフェイス変数よりバリュー値をゲット

     value, _ := deliver_view.([]type2.Deliver)

     sgh_ai.Date_Basic = "2017/01/01"
     sgh_ai.Date_Basic_Real = date1.Date_realdata_get( w  ,sgh_ai.Date_Basic )   // タイムデータ作成

     date_w := time.Now()        // 日付をゲット

//     fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : date_w %v\n", date_w.Year(), date_w.Month(),date_w.Day( ) ) // デバック

     date_max := fmt.Sprintf("%04d/%02d/%02d" ,date_w.Year(),  date_w.Month(),date_w.Day())

//     date_min := fmt.Sprintf("%04d/%02d/%02d" ,date_w.Year()-1, date_w.Month(),date_w.Day())
     date_min := fmt.Sprintf( "2017/01/01" )

///
/// 最小二乗法の作業データの初期化
///

     nn        := 0
     siguma_x  := 0.
     siguma_y  := 0.
     siguma_xx := 0.
     siguma_yy := 0.
     siguma_xy := 0.

///
/// 最小二乗法のデータを作成する
///　

     for _, deliverw := range value {

       if deliverw.Date <=  date_max &&         // 範囲内のデータかチェック
          deliverw.Date > date_min     {

          nn ++

          date_data := date1.Date_realdata_get( w  ,deliverw.Date )   // タイムデータ作成

          date_sub := date_data.Sub(sgh_ai.Date_Basic_Real)  // 基準日との差を計算　

//          fmt.Fprintf( w, "analysis1.deliver_growth_rate : NUM %v\n", num ) // デバック
//          fmt.Fprintf( w, "analysis1.deliver_growth_rate : XX %f\n", xx ) // デバック
//          fmt.Fprintf( w, "analysis1.deliver_growth_rate : RATERATE %f\n", num/xx ) // デバック

          siguma_x   = siguma_x  + float64(date_sub/(3600000000000*24))
          siguma_y   = siguma_y  + float64(deliverw.Number)
          siguma_xx  = siguma_xx + float64(date_sub/(3600000000000*24)) *  float64(date_sub/(3600000000000*24))
          siguma_yy  = siguma_yy + float64(deliverw.Number)      *  float64(deliverw.Number)
          siguma_xy  = siguma_xy + float64(date_sub/(3600000000000*24)) *  float64(deliverw.Number)

//          fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : X %v\n", float64(date_sub/10000000000000) )  // デバック
//          fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : Y %v\n", float64(deliverw.Number*10) )  // デバック
//          fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : XX %v\n", float64(date_sub/10000000000000) *  float64(date_sub/10000000000000) )  // デバック
//          fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : YY %v\n",  float64(deliverw.Number*10) *  float64(deliverw.Number*10))  // デバック
	   }
	}

//	fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : SIGUMA.X %f\n"  ,  siguma_x )  // デバック
//    fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : SIGUMA.Y %f\n"  ,  siguma_y )  // デバック
//    fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : SIGUMA.XX %f\n" ,  siguma_xx)  // デバック
//    fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : SIGUMA.YY %f\n" ,  siguma_yy )  // デバック
//    fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : SIGUMA.XY %f\n" ,  siguma_xy )  // デバック

///
/// 最小二乗法の勾配と切片を計算する
///

    aa := ( float64(nn) * siguma_xy - siguma_x * siguma_y ) / (  float64(nn) * siguma_xx - siguma_x * siguma_x )

    bb := ( siguma_xx * siguma_y - siguma_xy * siguma_x ) / (  float64(nn) * siguma_xx - siguma_x * siguma_x )

//    fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : AA %f\n" ,  aa )  // デバック
//    fmt.Fprintf( w, "analysis1.deliver_growth_rate2 : BB %f\n" ,  bb )  // デバック

///
/// 条件式をAIファイルに登録する
/// (切片bbは省略）

     for pos, deliverw := range value {
       if pos == 0  {      // AIファイル用の情報をセット

          sgh_ai.Course_No     = course_no
          sgh_ai.District_No   = deliverw.District_No
          sgh_ai.District_Name = deliverw.District_Name
          sgh_ai.Area_No       = deliverw.Area_No
	      sgh_ai.Area_Name     = deliverw.Area_Name

       }
	 }

    sgh_ai.Ex_Type       = "function-001"
    sgh_ai.Expression    = fmt.Sprintf( "Y=%fX+%f",aa ,bb)  // 条件式の作成
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


//    fmt.Fprintf( w, "analysis1.deliver_growth_rate2 normal end \n" )  // デバック
}
