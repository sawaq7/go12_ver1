package trans

import (

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
///   　　　get deliver inf.
///


func Deliver(funct int64 ,some_no int64 ,w http.ResponseWriter, r *http.Request )  (deliver2 []type2.Deliver ) {

//     IN  funct  　　　: ファンクション
//     　　　　　1:  area no
//     　　　　　2:  car no
//     　　　　　3:  private no

//     IN  some_no  　　: if flag
//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ

//     OUT deliver_view : slice of struct ( Deliver )

//    fmt.Fprintf( w, "trans.Deliver start \n" )

    var check_no  int64

    var line_counter int64

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)

    query := datastore.NewQuery("Deliver").Order("Date")

    count, err := client.Count(ctx, query)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

	deliver      := make([]type2.Deliver, 0, count)
	deliver_view := make([]type2.Deliver, 0)

    keys, err := client.GetAll(ctx, query , &deliver)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
//		fmt.Fprintf( w, "d_district_area_show err \n" ,err)
        return	nil
	}

	keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	line_counter = 0

	for pos, deliverw := range deliver {

///  branch flag set from function key

	  if funct == 0 {             //  when func. eq. area
	     check_no = deliverw.Course_No

	  }else if funct == 1 {      //  when func. eq. car
	     check_no = deliverw.Car_No

	  }else if funct == 2 {      //  when func. eq. private
	     check_no = deliverw.Private_No

	  }
      if  some_no == 0 {

         line_counter ++

         deliverw.Id      = keys_wk[pos]
         deliverw.Line_No = line_counter

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

         deliverw.Id      = keys_wk[pos]
         deliverw.Line_No = line_counter

         deliver_view = append ( deliver_view, deliverw )

      }
	}

///
///    sort deliver by double sort
///           key1 : Date  , key2 : Car_No

    deliver2 = sort.Deliver( w ,deliver_view  )

    return	deliver2

}

