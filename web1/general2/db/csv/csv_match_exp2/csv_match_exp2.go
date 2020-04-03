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

//    fmt.Fprintf( w, "csv_match_exp2 start \n" )  // 繝・ヰ繝・け

    var err error

    var csv_records  type5.Csv_Records

    csv_inf := trans3.Csv_inf ( w ,r )  ///      csv諠・ｱ繧偵ご繝・ヨ


    column_no1w := r.FormValue("column_no1")  // 蛻湧O繧偵ご繝・ヨ

	column_no1 ,err := strconv.Atoi(column_no1w)  // 謨ｴ謨ｰ蛹・	if err != nil {
//	     http.Error(w,err.Error(), http.StatusInternalServerError)
//         return
    } else {

      expression1 := r.FormValue("expression1")  // 蠑上ｒ繧ｲ繝・ヨ
	  if expression1 == " " {

//	    fmt.Fprintf( w, "csv_match_exp1 : expression1 is exsiting (no1)\n" )  // 繝・ヰ繝・け
	    return
	  }

	  expression1_sep := strings.Split ( expression1, " "  )

	  csv_records.Records[0] = operator.Operator( csv_inf  ,expression1_sep[0] ,expression1_sep[1] ,column_no1  ,w , r )

	  csv_records.Records_Num ++

	}

	column_no2w := r.FormValue("column_no2")  // 蛻湧O繧偵ご繝・ヨ

	column_no2 ,err := strconv.Atoi(column_no2w)  // 謨ｴ謨ｰ蛹・	if err != nil {
//	     http.Error(w,err.Error(), http.StatusInternalServerError)
//         return
    } else {

      expression2 := r.FormValue("expression2")  // 蠑上ｒ繧ｲ繝・ヨ
	  if expression2 == " " {

//	    fmt.Fprintf( w, "csv_match_exp2 : expression2 is exsiting (no2)\n" )  // 繝・ヰ繝・け
	    return
	  }

	  expression2_sep := strings.Split ( expression2, " "  )

	  csv_records.Records[1] = operator.Operator( csv_inf  ,expression2_sep[0] ,expression2_sep[1] ,column_no2  ,w , r )

	  csv_records.Records_Num ++

	}

    column_no3w := r.FormValue("column_no3")  // 蛻湧O繧偵ご繝・ヨ

	column_no3 ,err := strconv.Atoi(column_no3w)  // 謨ｴ謨ｰ蛹・	if err != nil {
//	     http.Error(w,err.Error(), http.StatusInternalServerError)
//         return
    } else {

      expression3 := r.FormValue("expression3")  // 蠑上ｒ繧ｲ繝・ヨ
	  if expression3 == " " {

//	    fmt.Fprintf( w, "csv_match_exp3 : expression3 is exsiting (no3)\n" )  // 繝・ヰ繝・け
	    return
	  }

	  expression3_sep := strings.Split ( expression3, " "  )

	  csv_records.Records[2] = operator.Operator( csv_inf  ,expression3_sep[0] ,expression3_sep[1] ,column_no3  ,w , r )

	  csv_records.Records_Num ++

	}
	logical := r.FormValue("logical")  // 隲也炊螟画焚繧偵ご繝・ヨ

//    _ = logical

//	fmt.Fprintf( w, "csv_match_exp2 : logical %v\n", logical )  // 繝・ヰ繝・け

    csv_inf4 := operator.Operator2( csv_records  ,logical ,w , r  )

///
///縲縲web 縺ｫcsv諠・ｱ繧定｡ｨ遉ｺ
///

     monitor := template.Must( template.New("html").Parse( html5.Csv_match_exp )) // 繝・Φ繝励Ξ繝ｼ繝医・繝倥ャ繝繝ｼ繧竪ET

     err = monitor.Execute ( w, csv_inf4 )
	 if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	 }

//	fmt.Fprintf( w, "csv_match_exp2 : normal end \n" )  // 繝・ヰ繝・け

}

