package trans

import (
	    "google.golang.org/appengine"
	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
//	    "html/template"

	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
//	    "time"
                                                )

///                           ///
/// 地区のエリア数をゲチE��する ///
///        test test test                  ///

func Deliver2( w http.ResponseWriter, r *http.Request )  ([]type2.Deliver ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ

//     OUT deliver_view : 構造体　”�E達情報”�Eスライス

//    fmt.Fprintf( w, "trans.Deliver2 start \n" )  // チE��チE��

    var line_counter int64

    c := appengine.NewContext(r)

	q := datastore.NewQuery("Deliver").Order("Date")

	count, err := q.Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return	nil
	}

// import struct for accessing datastore get from github.com/sawaq7/go12_ver1/client/sgh/type2/sgh.go

	deliver      := make([]type2.Deliver, 0, count)

	deliver_view := make([]type2.Deliver, 0)

	keys, err := q.GetAll(c, &deliver)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
//		fmt.Fprintf( w, "d_district_area_show err \n" ,err)  // チE��チE��
		return	nil
	}

	line_counter = 0

	for pos, deliverw := range deliver {

         line_counter ++
         deliverw.Line_No = line_counter
         deliver_view = append(deliver_view, type2.Deliver { keys[pos].IntID()             ,
                                                                    deliverw.Line_No       ,
                                                                    deliverw.Course_No     ,
                                                                    deliverw.District_No   ,
                                                                    deliverw.District_Name ,
                                                                    deliverw.Area_No       ,
                                                                    deliverw.Area_Name     ,
                                                                    deliverw.Date          ,
                                                                    deliverw.Date_Real     ,
                                                                    deliverw.Car_No        ,
                                                                    deliverw.Private_No    ,
                                                                    deliverw.Car_Division  ,
                                                                    deliverw.Number          })


	}

    return	deliver_view
}

