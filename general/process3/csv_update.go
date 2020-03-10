package process3

import (

//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
	    "general/type5"
//	    "fmt"

        "os"
        "cloud.google.com/go/datastore"
        "context"
                                                )


func Csv_update(w http.ResponseWriter, r *http.Request ,updid int64) {

//     IN    w      : レスポンスライター
//     IN    r      : リクエストパラメータ
//     IN  updid　  : 修正するデータストアのレコードのid

//    fmt.Fprintf( w, "csv_update start \n" )  // デバック
    var csv_inf type5.Csv_Inf

///
///   プロジェクト名をゲット
///
    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return
    }

	key := datastore.IDKey("Csv_Inf", updid, nil)

    if err := client.Get(ctx, key , &csv_inf ) ; err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    csv_inf.Column1 = r.FormValue("column1")
    csv_inf.Column2 = r.FormValue("column2")
    csv_inf.Column3 = r.FormValue("column3")
    csv_inf.Column4 = r.FormValue("column4")
    csv_inf.Column5 = r.FormValue("column5")
    csv_inf.Column6 = r.FormValue("column6")
    csv_inf.Column7 = r.FormValue("column7")
    csv_inf.Column8 = r.FormValue("column8")
    csv_inf.Column9 = r.FormValue("column9")
    csv_inf.Column10 = r.FormValue("column10")



//	fmt.Fprintf( w, "csv_update : csv_inf.Column1 %v\n", csv_inf.Column1 )  // デバック

    if _, err := client.Put(ctx, key, &csv_inf ); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

}