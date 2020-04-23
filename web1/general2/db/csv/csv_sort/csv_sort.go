package csv_sort

import (

	    "net/http"
//	    "fmt"

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

//    fmt.Fprintf( w, "sky/csv_sort start \n" )

    var err error

///
///     get column for sorting
///

    string_data := r.FormValue("sort_column")

    strings := strings2.String_no_get( w , r , string_data  )

    sort_key_no := make( []int, len(strings) )

    for pos, stringsw := range strings {

      sort_key_no[pos] ,err = strconv.Atoi(stringsw)  //  make an integer
      if err != nil {
	   http.Error(w,err.Error(), http.StatusInternalServerError)

		return
	  }

	}

//	fmt.Fprintf( w, "sky/csv_sort : sort_key_no %v\n", sort_key_no )

///
///   get project name
///

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )

      project_name = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return
    }

///
/// 　　sort csv inf.　
///

	csv_inf := trans3.Csv_inf ( w ,r )  //  get csv inf.

	csv_inf2 := sort2.Csv_inf( w ,r ,csv_inf ,sort_key_no )  //   sort csv inf.

///
/// 　　　set csv inf. in d.s.　
///

    for _, csv_inf2w := range csv_inf2 {

//   	  fmt.Fprintf( w, "process3.csv_column_join2 csv_inf2w %v\n", csv_inf2w )


      key := datastore.IDKey("Csv_Inf", csv_inf2w.Id, nil)   //　get access key

      if _, err := client.Put(ctx, key, &csv_inf2w ); err != nil {   // put csv inf. in d.s.

		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	  }

    }

///
///　　   show csv inf. on web
///

     monitor := template.Must( template.New("html").Parse( html5.Csv_show ))
                                                                //  set templateT

     err = monitor.Execute ( w, csv_inf2 )
	 if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	 }

//	fmt.Fprintf( w, "sky/csv_sort : normal end \n" )

}

