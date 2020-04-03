package csv_show

import (

	    "net/http"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/storage2/vaccine1"
	    "storage2"
	    "strconv"
	    "strings"
	    "io"
	    "bufio"
	    "html/template"

	    "github.com/sawaq7/go12_ver1/general/type5"
	    "github.com/sawaq7/go12_ver1/general/html5"
	    "github.com/sawaq7/go12_ver1/general/datastore5/initialize"
	    "github.com/sawaq7/go12_ver1/general/datastore5/trans3"

        "cloud.google.com/go/datastore"
        "context"
        "os"
                                                  )

func Csv_show(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "csv_show start \n" )  // ãEãE¯

    var bucket ,filename string

    var index   int64

    var column  int

    var str_work[10] string

    var csv_inf type5.Csv_Inf

    var storage_b_o_temp2 type5.Storage_B_O_Temp

    line_no := r.FormValue("line_no")

//    fmt.Fprintf( w, "csv_show : line_no %v\n", line_no )  // ãEãE¯

	select_id ,_ := strconv.Atoi(line_no)

//    fmt.Fprintf( w, "csv_show : select_id %v\n", select_id )  // ãEãE¯

///
///   ãã±ãEåãã²ãE
///

     projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )  // ãEãE¯

      projectID = "sample-7777"

	}

    ctx := context.Background()

    query := datastore.NewQuery("Storage_B_O_Temp")

    client, err := datastore.NewClient(ctx, projectID)
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

    if _, err = client.GetAll(ctx, query , &storage_b_o_temp ) ;  err != nil {

	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return

	} else{

      for pos, storage_b_o_tempw := range storage_b_o_temp {

        if pos == 0 {
           bucket    = storage_b_o_tempw.Bucket_Name    // ãã±ãEåãã²ãE

        }
	  }
	}

//    fmt.Fprintf( w, "csv_show : bucket2: %v\n", bucket2 )

///
///   ãã¡ã¤ã«åãã²ãE
///

	objects :=  storage2.Object_List ( w  ,r , bucket )

//    fmt.Fprintf( w, "csv_show : objects: %v\n", objects )

    for pos2 , objectsw := range objects {

      if pos2 == select_id-1 {

        filename = objectsw

      }

    }

///
///      csvãã¡ã¤ã«ãã¤ãã·ã£ã©ã¤ãº
///

    initialize.Csv_inf (w , r )

///
///      csvãã¡ã¤ã«ãæE ±ãã²ãEãã¦è¡¨ç¤º
///

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)  // ã¤ã³ã¿ã¼ãã§ã¤ã¹åãåå¤æ

    defer reader.Close()

    csv_reader := bufio.NewReaderSize(reader, 4096)

    index = 0 // ã¬ã³ã¼ãã«ã¦ã³ã¿ã¼ãã¤ãã·ã£ã©ã¤ãº

    warn_flag := 0  // ã¯ã¼ãã³ã°ãã©ã°ã®ã¤ãã·ã£ã©ã¤ãº

    for {

	    record ,err  := csv_reader.ReadString('\n')

//	    fmt.Fprintf( w, "csv_show : record %v\n", record )  // ãEãE¯

	    if err == io.EOF {

	      break

		}
		if err != nil {

		  http.Error(w, err.Error(), http.StatusInternalServerError)
	      return

		}

        record = strings.Replace( record, ",", " ", -1)     /// åºåEæE­ãå¤æ´

        column = strings.Count( record ," ") + 1

//        fmt.Fprintf( w, "csv_show : column %v\n", strings.Count( record ," ") + 1 )  // ãEãE¯

        if  column > 1 {  //   ã¬ã³ã¼ããã¹ããEã¹ã§ãªãE ´ååEçE
          index ++     // ã¬ã³ã¼ãã«ã¦ã³ã¿ã¼ãã«ã¦ã³ãE
          str := strings.Fields(record)

//		fmt.Fprintf( w, "csv_show : str %v\n", str )  // ãEãE¯

		  for ii := 0 ; ii < column ; ii++ {

             str_work[ii] = str[ii]

          }

///    ã¯ã¼ã¯ã¨ãªã¢(ãEEã¿ã¹ãã¢Eã«csvæE ±ãã»ãE

          csv_inf.Line_No    = index
          csv_inf.File_Name  = filename
          csv_inf.Column_Num = int64(column)
	      csv_inf.Column1    = str_work[0]
	      csv_inf.Column2    = str_work[1]
	      csv_inf.Column3    = str_work[2]
	      csv_inf.Column4    = str_work[3]
	      csv_inf.Column5    = str_work[4]
	      csv_inf.Column6    = str_work[5]
	      csv_inf.Column7    = str_work[6]
	      csv_inf.Column8    = str_work[7]
	      csv_inf.Column9    = str_work[8]
	      csv_inf.Column10   = str_work[9]

          new_key := datastore.IncompleteKey("Csv_Inf", nil)

//    fmt.Fprintf(w, "storage_object_copy_excute: new_key %v\n", new_key )  // ãEãE¯

         _, err = client.Put(ctx, new_key, &csv_inf )
	     if err != nil {
		   http.Error(w,err.Error(), http.StatusInternalServerError)
		   return
	     }

       } else {

      	 warn_flag = 1

       }

	}

    initialize.Storage_b_o_temp (w , r ) //  æ¢å­ãEãStorage_B_O_Temp ã³ã¢ã³ç¨ã®temporary-fileãã¯ãªã¢ã¼

    storage_b_o_temp2.Line_No =  1
    storage_b_o_temp2.Project_Name = projectID
    storage_b_o_temp2.Bucket_Name = bucket
    storage_b_o_temp2.Object_Name = filename

/// ã³ã¢ã³ç¨ã®temporary-fileã«ãã±ãEåãåã»ãE

    new_key := datastore.IncompleteKey("Storage_B_O_Temp", nil)

    if _, err := client.Put(ctx, new_key, &storage_b_o_temp2 ); err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError )
    }

///
///ãã ãã¡ã¤ã«ã®è£æ­£ãã
///

    if warn_flag == 1 {

       vaccine1.File_Pack ( w , r ,bucket  ,filename )

	}

///
///ããweb ã«csvæE ±ãè¡¨ç¤º
///

     csv_inf_view := trans3.Csv_inf ( w ,r )  ///      csvæE ±ãã²ãE

     monitor := template.Must( template.New("html").Parse( html5.Csv_show )) // ãE³ãã¬ã¼ããEãããã¼ãGET

     err = monitor.Execute ( w, csv_inf_view )
	 if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	 }

//	fmt.Fprintf( w, "csv_show : normal end \n" )  // ãEãE¯

}

