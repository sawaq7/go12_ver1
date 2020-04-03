package trans

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
//	    "html/template"

	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2/sort"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                )

///
///   　　　　　該当する�E達情報をゲチE��する
///


func Deliver(funct int64 ,some_no int64 ,w http.ResponseWriter, r *http.Request )  (deliver2 []type2.Deliver ) {

//     IN  funct  　　　: ファンクション
//     　　　　　�E�！E 地区NO
//     　　　　　�E�！E カーNO
//     　　　　　�E�！E プライベ�EチEO
//     IN  some_no  　　: 吁E��NO
//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ

//     OUT deliver_view : 構造体　”�E達情報”�Eスライス

//    fmt.Fprintf( w, "trans.Deliver start \n" )  // チE��チE��

    var check_no  int64

    var line_counter int64

///
///     配達惁E��をゲチE��する
///

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)

    query := datastore.NewQuery("Deliver").Order("Date")
//	q := datastore.NewQuery("Deliver").Order("Date")

    count, err := client.Count(ctx, query)
//	count, err := q.Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

	deliver      := make([]type2.Deliver, 0, count)
	deliver_view := make([]type2.Deliver, 0)


//	keys, err := q.GetAll(c, &deliver)
    keys, err := client.GetAll(ctx, query , &deliver)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
//		fmt.Fprintf( w, "d_district_area_show err \n" ,err)  // チE��チE��
        return	nil
	}

	keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	line_counter = 0

	for pos, deliverw := range deliver {

///  機�EによりチェチE��頁E��をセチE��

	  if funct == 0 {   // コースNOの場吁E
	     check_no = deliverw.Course_No

	  }else if funct == 1 {   // カーNOの場吁E
	     check_no = deliverw.Car_No

	  }else if funct == 2 {   // 個人NOの場吁E
	     check_no = deliverw.Private_No

	  }
      if  some_no == 0 {

         line_counter ++

         deliverw.Id      = keys_wk[pos]    //  チE�EタストアidをセチE��
         deliverw.Line_No = line_counter         //  行NOをセチE��

         deliver_view = append ( deliver_view, deliverw )

//         deliver_view = append(deliver_view, type2.Deliver { keys[pos].IntID()             ,
//                                                                    deliverw.Line_No       ,
//                                                                    deliverw.Course_No     ,
//                                                                    deliverw.District_No   ,
//                                                                    deliverw.District_Name ,
//                                                                    deliverw.Area_No       ,
//                                                                    deliverw.Area_Name     ,
//                                                                    deliverw.Date          ,
//                                                                    deliverw.Date_Real     ,
//                                                                    deliverw.Car_No        ,
//                                                                    deliverw.Private_No    ,
//                                                                    deliverw.Car_Division  ,
//                                                                    deliverw.Number          })


      }else if some_no == check_no {

         line_counter ++

         deliverw.Id      = keys_wk[pos]    //  チE�EタストアidをセチE��
         deliverw.Line_No = line_counter         //  行NOをセチE��

         deliver_view = append ( deliver_view, deliverw )


      }
	}

///
/// 配達惁E��を、E重sortする
///           key1 : Date  , key2 : Car_No

    deliver2 = sort.Deliver( w ,deliver_view  )

    return	deliver2

}

