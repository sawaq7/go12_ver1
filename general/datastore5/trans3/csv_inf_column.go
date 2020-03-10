package trans3

import (

//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"

	    "general/type5"
//	    "time"

        "cloud.google.com/go/datastore"
        "context"
        "os"
                                                )

///
///    指定したcolumnNOのcsv情報をゲットする
///

func Csv_inf_column ( w http.ResponseWriter, r *http.Request ,column_no int )  ( []string ) {

//     IN    w      　　　　: レスポンスライター
//     IN    r      　　　　: リクエストパラメータ
//     IN  column_no　      : 抽出したい行NOの情報

//     OUT csv_inf_column  : csv情報”

//    fmt.Fprintf( w, "trans.csv_inf_column start \n" )  // デバック

    var string_wk string

///
///   プロジェクト名をゲット
///
    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )  // デバック

      project_name = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return nil
    }

//	q := datastore.NewQuery("Csv_Inf").Order("Line_No")
	query := datastore.NewQuery("Csv_Inf").Order("Line_No")

    count, err := client.Count(ctx, query)
//	count, err := q.Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

	csv_inf      := make([]type5.Csv_Inf, 0, count)
	csv_inf_column := make([]string, 0)

    _, err = client.GetAll(ctx, query , &csv_inf)
//	_, err = q.GetAll(c, &csv_inf)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

	for _, csv_infw := range csv_inf {

///  機能によりチェック項目をセット
     switch column_no {

          case 1 :

            string_wk = csv_infw.Column1

          break;

          case 2 :

            string_wk = csv_infw.Column2

          break;

          case 3 :

            string_wk = csv_infw.Column3

          break;

          case 4 :

            string_wk = csv_infw.Column4

          break;

          case 5 :

            string_wk = csv_infw.Column5

          break;

          case 6 :

            string_wk = csv_infw.Column6

          break;

          case 7 :

            string_wk = csv_infw.Column7

          break;

          case 8 :

            string_wk = csv_infw.Column8

          break;

          case 9 :

            string_wk = csv_infw.Column9

          break;

          case 10 :

            string_wk = csv_infw.Column10

          break;

      }

      csv_inf_column = append( csv_inf_column, string_wk )

	}

//    fmt.Fprintf( w, "trans.csv_inf_column normal end )  // デバック

    return	csv_inf_column
}

