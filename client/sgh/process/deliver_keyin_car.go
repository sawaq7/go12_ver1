package process

import (
	    "google.golang.org/appengine"
	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
	    "html/template"
	    "client/sgh/html2"
	    "client/sgh/type2"
//	    "time"
                                                )
/// 配達データを表示する ///

func Deliver_keyin_car(w http.ResponseWriter, r *http.Request) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ

//    fmt.Fprintf( w, "deliver_keyin_car start \n" )  // デバック

var line_counter int64

	c := appengine.NewContext(r)

	q := datastore.NewQuery("Deliver").Order("Date")

	count, err := q.Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

// import struct for accessing datastore get from client/sgh/type2/sgh.go

	deliver      := make([]type2.Deliver, 0, count)
	deliver_view := make([]type2.Deliver, 0)

	keys, err := q.GetAll(c, &deliver)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	line_counter = 0

    for pos, deliverw := range deliver {

      line_counter ++
      deliverw.Line_No = line_counter

      deliver_view = append(deliver_view, type2.Deliver { keys[pos].IntID()                ,
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

// テンプレートのヘッダーをGET

     monitor := template.Must(template.New("html").Parse(html2.Deliver_keyin_car))

// モニターに表示

	err = monitor.Execute(w, deliver_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

