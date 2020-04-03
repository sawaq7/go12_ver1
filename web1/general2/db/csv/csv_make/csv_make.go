package csv_make

import (
	    "storage2"
//	    "fmt"
	    "net/http"

	    "github.com/sawaq7/go12_ver1/general/type5"
//	    "github.com/sawaq7/go12_ver1/general/process3"
	    "github.com/sawaq7/go12_ver1/general/datastore5/trans3"
//	    "strings"
        "github.com/sawaq7/go12_ver1/general/html5"
        "html/template"

        "os"

        "cloud.google.com/go/datastore"
        "context"

                                                  )

func Csv_make(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "csv_make start \n" )  // γEγE―

    var bucket ,filename  ,project_name string

    var csv_inf_first type5.Csv_Inf

///
/// ε₯εγγΌγΏγGET γ
///

    filename = r.FormValue("file_name")  // γγ₯γΌγγ‘γ€γ«εγγ²γE

///
///   γγ­γΈγ§γ―γεγγ²γE
///

    project_name = os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return
    }

	query := datastore.NewQuery("Storage_B_O_Temp")

	count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	storage_b_o_temp      := make([]type5.Storage_B_O_Temp, 0, count)

    if _, err := client.GetAll(ctx, query , &storage_b_o_temp );  err != nil {

	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return

	} else{

      for pos, storage_b_o_tempw := range storage_b_o_temp {

        if pos == 0 {

           project_name       = storage_b_o_tempw.Project_Name    // γγ­γΈγ§γ―γεγγ²γE
           bucket   = storage_b_o_tempw.Bucket_Name    // γγ±γEεγγ²γE

        }
	  }
	}

///
/// γγγcsvγγ‘γ€γ«γδ½ζEγγγ
///

     writer := storage2.File_Create2( w ,r ,bucket ,filename ,"text/plain" )
//     writer := storage2.File_Create2( w ,r ,bucket ,filename ,"text/csv" )

	defer writer.Close()

	csv_inf := trans3.Csv_inf ( w ,r )  ///      csvζE ±γγ²γE

	colum_num := int ( csv_inf[0].Column_Num )  // εζ°γγ²γE
	filename2 := csv_inf[0].File_Name            // γγ‘γ€γ«εγγ²γE
	first_id  :=  csv_inf[0].Id


//	fmt.Fprintf( w, "csv_make : colum_num %v\n", colum_num )  // γEγE―
//	fmt.Fprintf( w, "csv_make : record_num %v\n", len(csv_inf) )  // γEγE―

	record := make ( []string ,colum_num )   //γγ¬γ³γΌγγEγ―γΌγ―γ¨γͺγ’γη’ΊδΏE
///
///    csvγγ‘γ€γ«γδ½ζE
///

    for _ , csv_infw := range csv_inf {

      for ii := 0 ; ii < colum_num ; ii++ {  //γγ¬γ³γΌγγγ»γE

        switch ii {

          case 0 :

            record[ii] = csv_infw.Column1

          break;

          case 1 :

            record[ii] = csv_infw.Column2

          break;

          case 2 :

            record[ii] = csv_infw.Column3

          break;

          case 3 :

            record[ii] = csv_infw.Column4

          break;

          case 4 :

            record[ii] = csv_infw.Column5

          break;

          case 5 :

            record[ii] = csv_infw.Column6

          break;

          case 6 :

            record[ii] = csv_infw.Column7

          break;

          case 7 :

            record[ii] = csv_infw.Column8

          break;

          case 8 :

            record[ii] = csv_infw.Column9

          break;

          case 9 :

            record[ii] = csv_infw.Column10

          break;

        }
      }

//      fmt.Fprintf( w, "csv_make : record %v\n", record )  // γEγE―

      storage2.File_Write_Csv2 ( w  ,writer ,record )  // csvγ¬γ³γΌγγζΈγθΎΌγ
//      storage2.File_Write_Csv ( w  ,writer ,record )  // csvγ¬γ³γΌγγζΈγθΎΌγ

	}

///
///γγγγ‘γ€γ«εγmodify
///

    if filename != filename2 {

       key := datastore.IDKey("Csv_Inf", first_id, nil)

       if err := client.Get(ctx, key , &csv_inf_first ) ; err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	   }

       csv_inf_first.File_Name = filename

       if _, err := client.Put(ctx, key, &csv_inf_first ); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	   }
    }

///
///γγweb γ«csvζE ±γθ‘¨η€Ί
///

     csv_inf_view := trans3.Csv_inf ( w ,r )  ///      csvζE ±γγ²γE

     monitor := template.Must( template.New("html").Parse( html5.Csv_show )) // γE³γγ¬γΌγγEγγγγΌγGET

     err = monitor.Execute ( w, csv_inf_view )
	 if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	 }

//	fmt.Fprintf( w, "csv_make normal end \n" )  // γEγE―

}
