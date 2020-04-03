package trans3

import (

//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/general/type5"
//	    "time"

        "cloud.google.com/go/datastore"
        "context"
        "os"
                                                )

///
///    æE®ããcolumnNOã®csvæE ±ãã²ãEãã
///

func Csv_inf_column ( w http.ResponseWriter, r *http.Request ,column_no int )  ( []string ) {

//     IN    w      ãããã: ã¬ã¹ãã³ã¹ã©ã¤ã¿ã¼
//     IN    r      ãããã: ãªã¯ã¨ã¹ããã©ã¡ã¼ã¿
//     IN  column_noã      : æ½åºãããE¡NOã®æE ±

//     OUT csv_inf_column  : csvæE ±âE
//    fmt.Fprintf( w, "trans.csv_inf_column start \n" )  // ãEãE¯

    var string_wk string

///
///   ãã­ã¸ã§ã¯ãåãã²ãE
///
    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )  // ãEãE¯

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

///  æ©èEã«ãããã§ãE¯é E®ãã»ãE
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

//    fmt.Fprintf( w, "trans.csv_inf_column normal end )  // ãEãE¯

    return	csv_inf_column
}

