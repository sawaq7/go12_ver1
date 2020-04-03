package csv_match_exp2

import (

	    "net/http"
//	    "fmt"

	    "html/template"
        "strconv"
	    "github.com/sawaq7/go12_ver1/general/html5"
	    "github.com/sawaq7/go12_ver1/general/type5"
	    "github.com/sawaq7/go12_ver1/general/datastore5/trans3"
	    "github.com/sawaq7/go12_ver1/general/interpret/operator"
	    "strings"

                                                  )

func Csv_match_exp2(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "csv_match_exp2 start \n" )  // γEγE―

    var err error

    var csv_records  type5.Csv_Records

    csv_inf := trans3.Csv_inf ( w ,r )  ///      csvζE ±γγ²γE


    column_no1w := r.FormValue("column_no1")  // εNOγγ²γE

	column_no1 ,err := strconv.Atoi(column_no1w)  // ζ΄ζ°εE	if err != nil {
//	     http.Error(w,err.Error(), http.StatusInternalServerError)
//         return
    } else {

      expression1 := r.FormValue("expression1")  // εΌγγ²γE
	  if expression1 == " " {

//	    fmt.Fprintf( w, "csv_match_exp1 : expression1 is exsiting (no1)\n" )  // γEγE―
	    return
	  }

	  expression1_sep := strings.Split ( expression1, " "  )

	  csv_records.Records[0] = operator.Operator( csv_inf  ,expression1_sep[0] ,expression1_sep[1] ,column_no1  ,w , r )

	  csv_records.Records_Num ++

	}

	column_no2w := r.FormValue("column_no2")  // εNOγγ²γE

	column_no2 ,err := strconv.Atoi(column_no2w)  // ζ΄ζ°εE	if err != nil {
//	     http.Error(w,err.Error(), http.StatusInternalServerError)
//         return
    } else {

      expression2 := r.FormValue("expression2")  // εΌγγ²γE
	  if expression2 == " " {

//	    fmt.Fprintf( w, "csv_match_exp2 : expression2 is exsiting (no2)\n" )  // γEγE―
	    return
	  }

	  expression2_sep := strings.Split ( expression2, " "  )

	  csv_records.Records[1] = operator.Operator( csv_inf  ,expression2_sep[0] ,expression2_sep[1] ,column_no2  ,w , r )

	  csv_records.Records_Num ++

	}

    column_no3w := r.FormValue("column_no3")  // εNOγγ²γE

	column_no3 ,err := strconv.Atoi(column_no3w)  // ζ΄ζ°εE	if err != nil {
//	     http.Error(w,err.Error(), http.StatusInternalServerError)
//         return
    } else {

      expression3 := r.FormValue("expression3")  // εΌγγ²γE
	  if expression3 == " " {

//	    fmt.Fprintf( w, "csv_match_exp3 : expression3 is exsiting (no3)\n" )  // γEγE―
	    return
	  }

	  expression3_sep := strings.Split ( expression3, " "  )

	  csv_records.Records[2] = operator.Operator( csv_inf  ,expression3_sep[0] ,expression3_sep[1] ,column_no3  ,w , r )

	  csv_records.Records_Num ++

	}
	logical := r.FormValue("logical")  // θ«ηε€ζ°γγ²γE

//    _ = logical

//	fmt.Fprintf( w, "csv_match_exp2 : logical %v\n", logical )  // γEγE―

    csv_inf4 := operator.Operator2( csv_records  ,logical ,w , r  )

///
///γγweb γ«csvζE ±γθ‘¨η€Ί
///

     monitor := template.Must( template.New("html").Parse( html5.Csv_match_exp )) // γE³γγ¬γΌγγEγγγγΌγGET

     err = monitor.Execute ( w, csv_inf4 )
	 if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	 }

//	fmt.Fprintf( w, "csv_match_exp2 : normal end \n" )  // γEγE―

}

