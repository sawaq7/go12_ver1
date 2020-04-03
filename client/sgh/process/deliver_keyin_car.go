package process

import (
	    "google.golang.org/appengine"
	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/sgh/html2"
	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
//	    "time"
                                                )
/// ιιγEEγΏγθ‘¨η€Ίγγ ///

func Deliver_keyin_car(w http.ResponseWriter, r *http.Request) {

//     IN    w      γγ: γ¬γΉγγ³γΉγ©γ€γΏγΌ
//     IN    r      γγ: γͺγ―γ¨γΉγγγ©γ‘γΌγΏ

//    fmt.Fprintf( w, "deliver_keyin_car start \n" )  // γEγE―

var line_counter int64

	c := appengine.NewContext(r)

	q := datastore.NewQuery("Deliver").Order("Date")

	count, err := q.Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

// import struct for accessing datastore get from github.com/sawaq7/go12_ver1/client/sgh/type2/sgh.go

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

// γE³γγ¬γΌγγEγγγγΌγGET

     monitor := template.Must(template.New("html").Parse(html2.Deliver_keyin_car))

// γ’γγΏγΌγ«θ‘¨η€Ί

	err = monitor.Execute(w, deliver_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

