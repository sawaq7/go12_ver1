package csv_sort

import (

	    "net/http"
//	    "fmt"
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"

	    "strconv"
	    "html/template"

	    "github.com/sawaq7/go12_ver1/general/html5"
	    "github.com/sawaq7/go12_ver1/general/strings2"
	    "github.com/sawaq7/go12_ver1/general/datastore5/trans3"
	    "github.com/sawaq7/go12_ver1/general/datastore5/sort2"

	    "cloud.google.com/go/datastore"
        "context"
        "os"

                                                  )

func Csv_sort(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "sky/csv_sort start \n" )  // γEγE―

    var err error

///
///     γ½γΌγγγγεENOγγ²γE
///

    string_data := r.FormValue("sort_column")

    strings := strings2.String_no_get( w , r , string_data  )

    sort_key_no := make( []int, len(strings) )

    for pos, stringsw := range strings {

      sort_key_no[pos] ,err = strconv.Atoi(stringsw)  // ζ΄ζ°εE	  if err != nil {
	   http.Error(w,err.Error(), http.StatusInternalServerError)

		return
	  }

	}

//	fmt.Fprintf( w, "sky/csv_sort : sort_key_no %v\n", sort_key_no )  // γEγE―

///
///   γγ­γΈγ§γ―γεγγ²γE
///

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )  // γEγE―

      project_name = "sample-7777"

	}

//    c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return
    }

///
/// γγγcsvζE ±γγ½γΌγγγγ
///

	csv_inf := trans3.Csv_inf ( w ,r )  ///  csvζE ±γγ²γE

	csv_inf2 := sort2.Csv_inf( w ,r ,csv_inf ,sort_key_no )  ///  csvζE ±γγ½γΌγγγE
///
/// γγγγEEγΏγΉγγ’γ«γcsvζE ±γεEγ»γEγγγ
///

    for _, csv_inf2w := range csv_inf2 {

//   	  fmt.Fprintf( w, "process3.csv_column_join2 csv_inf2w %v\n", csv_inf2w )  // γEγE―

//      key := datastore.NewKey(c, "Csv_Inf", "", csv_inf2w.Id, nil)  //γγ’γ―γ»γΉγ­γΌγ²γE
      key := datastore.IDKey("Csv_Inf", csv_inf2w.Id, nil)

      if _, err := client.Put(ctx, key, &csv_inf2w ); err != nil {
//      if _, err := datastore.Put(c, key, &csv_inf2w); err != nil {  // γEEγΏγΉγγ’γ«εγ»γE
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	  }

    }

///
///γγweb γ«csvζE ±γθ‘¨η€Ί
///

     monitor := template.Must( template.New("html").Parse( html5.Csv_show )) // γE³γγ¬γΌγγEγγγγΌγGET

     err = monitor.Execute ( w, csv_inf2 )
	 if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	 }

//	fmt.Fprintf( w, "sky/csv_sort : normal end \n" )  // γEγE―

}

