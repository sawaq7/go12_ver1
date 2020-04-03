package process3

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
//	    "html/template"
	    "github.com/sawaq7/go12_ver1/general/datastore5/reformat"

//	    "github.com/sawaq7/go12_ver1/general/html5"
	    "github.com/sawaq7/go12_ver1/general/type5"
	    "github.com/sawaq7/go12_ver1/general/datastore5/trans3"
	    "github.com/sawaq7/go12_ver1/general/datastore5/set1"
	    "storage2"
	    "io"
	    "strings"
	    "bufio"
	    "cloud.google.com/go/datastore"
        "context"
        "os"
                                                )

///
///     æE®ãããã¡ã¤ã«ã®1åãEãEEã¿ãè¿½å ãã
///

func Csv_column_join(w http.ResponseWriter, r *http.Request , filename string ,column_no int ) {

//     IN    w      ãã: ã¬ã¹ãã³ã¹ã©ã¤ã¿ã¼
//     IN    r      ãã: ãªã¯ã¨ã¹ããã©ã¡ã¼ã¿
//     IN  filename ã  : ãã¡ã¤ã«åE//     IN  column_noã  : è¿½å ããè¡NO

//    fmt.Fprintf( w, "process3.csv_column_join start \n" )  // ãEãE¯

    var index   int64

    var column   int

    var str_work[10] string

    var bucket string

    csv_inf_join := make( []string , 0)

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )  // ãEãE¯

      project_name = "sample-7777"

	}
    ctx := context.Background()

	query := datastore.NewQuery("Storage_B_O_Temp")

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return
    }

    count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	storage_b_o_temp      := make([]type5.Storage_B_O_Temp, 0, count)

    if _, err := client.GetAll(ctx, query , &storage_b_o_temp); err != nil {

	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return

	} else{

      for pos, storage_b_o_tempw := range storage_b_o_temp {

        if pos == 0 {

//           project_name       = storage_b_o_tempw.Project_Name    // ãã­ã¸ã§ã¯ãåãã²ãE
           bucket   = storage_b_o_tempw.Bucket_Name    // ãã±ãEåãã²ãE
//           filename = storage_b_o_tempw.Object_Name    // ããEã·ãE¯ãã¡ã¤ã«åãã²ãE

        }
	  }
	}

///
/// ãããcsvæE ±ãä¿®æ­£ããã
///

	csv_inf := trans3.Csv_inf ( w ,r )  ///      csvæE ±ãã²ãE

    csv_inf2 := reformat.Csv_inf ( 1,csv_inf[0].Column_Num+1 ,csv_inf ,w ,r )
                                                           /// ãã©ã¼ãããã1åæ¡å¼µãã
///
///      è¿½å ããcsvæE ±ãã²ãE
///

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)  // ã¤ã³ã¿ã¼ãã§ã¤ã¹åãåå¤æ

    defer reader.Close()

    csv_reader := bufio.NewReaderSize(reader, 4096)

    index = 0 // ã¬ã³ã¼ãã«ã¦ã³ã¿ã¼ãã¤ãã·ã£ã©ã¤ãº

    for {

        index ++     // ã¬ã³ã¼ãã«ã¦ã³ã¿ã¼ãã«ã¦ã³ãE
	    record ,err  := csv_reader.ReadString('\n')

//	    fmt.Fprintf( w, "csv_show : record %v\n", record )  // ãEãE¯

	    if err == io.EOF {

	      break

		}
		if err != nil {

		  http.Error(w, err.Error(), http.StatusInternalServerError)
	      return

		}
		if index == 1 {   // åæ°ãã²ãE

		  column = strings.Count( record ,",") + 1

//		  fmt.Fprintf( w, "csv_show : column %v\n", column )  // ãEãE¯

		}

		str := strings.Split ( record, ","  )

//		fmt.Fprintf( w, "csv_show : str %v\n", str )  // ãEãE¯

		for ii := 0 ; ii < column ; ii++ {

           str_work[ii] = str[ii]

        }

        csv_inf_join = append( csv_inf_join ,str_work[column_no-1] )
                                                      ///    è¿½å ããcsvæE ±ãã»ãE

    }

    csv_inf_new := set1.Csv_inf (  csv_inf2 ,csv_inf_join ,int(csv_inf2[0].Column_Num) , w ,r )
                                                        /// è¿½å ãããEEã¿1åãã»ãE

///
/// ããããEEã¿ã¹ãã¢ã«ãcsvæE ±ãåEã»ãEããã
///

    for _, csv_inf_neww := range csv_inf_new {

//   	  fmt.Fprintf( w, "process3.csv_column_join csv_inf_neww %v\n", csv_inf_neww )  // ãEãE¯

      key := datastore.IDKey("Csv_Inf", csv_inf_neww.Id, nil)

      if _, err := client.Put(ctx, key, &csv_inf_neww ); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	  }

    }

//	fmt.Fprintf( w, "process3.csv_column_join normal end \n" )  // ãEãE¯


}

