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

func Deliver_keyin_private(w http.ResponseWriter, r *http.Request) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ

//    fmt.Fprintf( w, "deliver1_show start \n" )  // デバック

	c := appengine.NewContext(r)

	q := datastore.NewQuery("Deliver").Order("Date")

	count, err := q.Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	deliver_view := make([]type2.Deliver, count)

    monitor := template.Must(template.New("html").Parse(html2.Deliver_keyin_private))

// モニターに表示

	err = monitor.Execute(w, deliver_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

