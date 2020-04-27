package datastore4

import (

        "net/http"
//	    "fmt"

        "github.com/sawaq7/go12_ver1/client/sgh/datastore2/check"
        "github.com/sawaq7/go12_ver1/client/sgh/datastore2/initialize"

	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2/trans"
	    "github.com/sawaq7/go12_ver1/client/tokura/datastore4/trans2"
	    "github.com/sawaq7/go12_ver1/client/tokura/datastore4/initialize2"
//	    "github.com/sawaq7/go12_ver1/general/type5"

                                                )

///
///     Datastore_tokura  main routine
///


func Datastore_tokura( fname string ,function string ,flexible_in interface{} ,w http.ResponseWriter, r *http.Request )  (flexible_out interface{} ) {

//     IN    fname       : file-name
//     IN    function    : 　
//        　　　　　　　　　trans ,check ,initialize ,sort　etc
//     IN flexible_in　  : see attachment
//     IN    w      : response-writer
//     IN    r      : request- paramete

//     out flexible_out  : see attachment

//    fmt.Fprintf( w, "datastore_tokura start \n" )
//    fmt.Fprintf( w, "datastore_tokura function \n" ,function )
//    fmt.Fprintf( w, "datastore_tokura fname \n" ,fname )

///
///     select process  by function and file-name
///

	switch function {

      case "trans" :
        switch fname {

          case "Water_Line" :
            value, _ := flexible_in.(string)

            flexible_out = trans2.Water_line ( 1  ,value , w ,r )

          break;

          case "Water_Slope" :
            flexible_out = trans2.Water_slope ( w ,r )

          break;

          case "Water2" :

            flexible_out = trans2.Water2 ( w ,r )

          break;

          case "Car" :

            value, _ := flexible_in.(int64)
            flexible_out = trans.Car_district ( value ,w ,r )

          break;

          case "Sgh_Ai" :

            value, _ := flexible_in.(int64)
            flexible_out = trans.Sgh_ai( value  ,w , r  )

          break;

        }
      break;

      case "trans2" :
        switch fname {

          case "Deliver" :



          break;

          case "D_Area" :

            value, _ := flexible_in.(int64)
            flexible_out = trans.D_area_district ( w ,r ,value )


          break;

          case "Private" :



          break;

          case "Car" :



          break;

          case "D_District_Temp" :


          break;
        }
      break;

      case "initialize" :

        switch fname {

          case "Water2_Temp" :

            initialize2.Water2_temp (w , r )

          break;

          case "D_District_Temp" :
            initialize.D_district_temp (w , r )

          break;

          case "Sgh_Ai" :
            value, _ := flexible_in.(int64)
            initialize.Sgh_ai( value ,w , r )

          break;

        }

      break;

      case "check" :

        switch fname {

          case "D_Area" :

            value, _ := flexible_in.(int64)
            flexible_out = check.D_area ( w , r  ,value )

          break;

          case "D_District_Temp" :

            flexible_out = check.D_district_temp (w , r )

          break;

          case "Car" :

            value, _ := flexible_in.(int64)
            flexible_out = check.Car_no_max(w , r  ,value)

          break;

        }

      break;

    }

	return flexible_out

}

