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
/// 蝨ｰ蛹ｺ縺ｮ繧ｨ繝ｪ繧｢謨ｰ繧偵ご繝・ヨ縺吶ｋ ///
///        test test test                  ///

func Deliver2( w http.ResponseWriter, r *http.Request )  ([]type2.Deliver ) {

//     IN    w      縲縲: 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r      縲縲: 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ

//     OUT deliver_view : 讒矩菴薙窶晞・驕疲ュ蝣ｱ窶昴・繧ｹ繝ｩ繧､繧ｹ

//    fmt.Fprintf( w, "trans.Deliver2 start \n" )  // 繝・ヰ繝・け

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
//		fmt.Fprintf( w, "d_district_area_show err \n" ,err)  // 繝・ヰ繝・け
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

