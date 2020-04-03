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
/// 驟埼＃繝・・繧ｿ繧定｡ｨ遉ｺ縺吶ｋ ///

func Deliver_keyin_private(w http.ResponseWriter, r *http.Request) {

//     IN    w      縲縲: 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r      縲縲: 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ

//    fmt.Fprintf( w, "deliver1_show start \n" )  // 繝・ヰ繝・け

	c := appengine.NewContext(r)

	q := datastore.NewQuery("Deliver").Order("Date")

	count, err := q.Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	deliver_view := make([]type2.Deliver, count)

    monitor := template.Must(template.New("html").Parse(html2.Deliver_keyin_private))

// 繝｢繝九ち繝ｼ縺ｫ陦ｨ遉ｺ

	err = monitor.Execute(w, deliver_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

