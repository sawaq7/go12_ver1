package csv_show_test1

import (

	    "net/http"
	    "fmt"


//	    "github.com/sawaq7/go12_ver1/storage2/get"
	    "storage2"
	    "strconv"
	    "strings"
	    "io"
	    "bufio"


	    "github.com/sawaq7/go12_ver1/general/type5"

	    "github.com/sawaq7/go12_ver1/general/datastore5/initialize"


        "cloud.google.com/go/datastore"
        "context"
        "os"


                                                  )


func Csv_show_test1(w http.ResponseWriter, r *http.Request) {

    fmt.Fprintf( w, "csv_show_test start \n" )  // γEγE―

    var bucket ,filename string

    var index   int64

    var column int


    line_no := r.FormValue("line_no")

//    fmt.Fprintf( w, "csv_show : line_no %v\n", line_no )  // γEγE―

	select_id ,_ := strconv.Atoi(line_no)

//    fmt.Fprintf( w, "csv_show : select_id %v\n", select_id )  // γEγE―


///
///   γγ±γEεγγ²γE
///
     projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )  // γEγE―

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
           bucket    = storage_b_o_tempw.Bucket_Name    // γγ±γEεγγ²γE

        }
	  }
	}

//    fmt.Fprintf( w, "csv_show : bucket2: %v\n", bucket2 )

///
///   γγ‘γ€γ«εγγ²γE
///

	objects :=  storage2.Object_List ( w  ,r , bucket )

//    fmt.Fprintf( w, "csv_show : objects: %v\n", objects )

    for pos2 , objectsw := range objects {

      if pos2 == select_id-1 {

        filename = objectsw


      }

    }

///
///      csvγγ‘γ€γ«γγ€γγ·γ£γ©γ€γΊ
///

    initialize.Csv_inf (w , r )

///
///      csvγγ‘γ€γ«γζE ±γγ²γEγγ¦θ‘¨η€Ί
///

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)  // γ€γ³γΏγΌγγ§γ€γΉεγεε€ζ

    defer reader.Close()

    csv_reader := bufio.NewReaderSize(reader, 4096)

    index = 0 // γ¬γ³γΌγγ«γ¦γ³γΏγΌγγ€γγ·γ£γ©γ€γΊ

    for {

        index ++     // γ¬γ³γΌγγ«γ¦γ³γΏγΌγγ«γ¦γ³γE
	    record ,err  := csv_reader.ReadString('\n')

	    fmt.Fprintf( w, "csv_show : record %v\n", record )  // γEγE―

	    if err == io.EOF {

	      break

		}
		if err != nil {

		  http.Error(w, err.Error(), http.StatusInternalServerError)
	      return

		}

		record = strings.Replace( record, ",", " ", -1)     /// εΊεEζE­γε€ζ΄

		if index == 1 {   // εζ°γγ²γE

		  column = strings.Count( record ," ") + 1

		  fmt.Fprintf( w, "csv_show : column %v\n", column )  // γEγE―

		}




///    γ―γΌγ―γ¨γͺγ’(γEEγΏγΉγγ’Eγ«csvζE ±γγ»γE


	}

//	fmt.Fprintf( w, "csv_show : normal end \n" )  // γEγE―

}

