package datastore2

import (

        "net/http"
//	    "fmt"

        "github.com/sawaq7/go12_ver1/client/sgh/datastore2/check"
        "github.com/sawaq7/go12_ver1/client/sgh/datastore2/initialize"

	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2/trans"
	    "github.com/sawaq7/go12_ver1/general/type5"

                                                )

///
///    datastore access routine
///


func Datastore_sgh( fname string ,function string ,flexible_in interface{} ,w http.ResponseWriter, r *http.Request )  (flexible_out interface{} ) {

//     IN    fname       : file name
//     IN    function    :
//        　　　　　　　　　trans ,check ,initialize ,sort　etc
//     IN flexible_in　  : flexible data　(see attachment)
//     IN    w      　　 : レスポンスライター
//     IN    r      　　 : リクエストパラメータ

//     out flexible_out  : flexible data　(see attachment)
//    fmt.Fprintf( w, "datastore_sgh start \n" )
//    fmt.Fprintf( w, "datastore_sgh function \n" ,function )
//    fmt.Fprintf( w, "datastore_sgh fname \n" ,fname )

///
///    select process by function and file-name

	switch function {

///
///  トランスの場吁E///

      case "trans" :
        switch fname {

          case "Deliver" :

            //  get value from interface data

            value, _ := flexible_in.([]type5.General_Work)
            flexible_out = trans.Deliver ( value[0].Int64_Work  ,value[1].Int64_Work  ,w ,r )

          break;

          case "D_Area" :

            value, _ := flexible_in.(int64)
            flexible_out = trans.D_area ( 0 ,value ,w ,r )

          break;

          case "D_District" :

            value, _ := flexible_in.([]type5.General_Work)
            flexible_out = trans.D_district ( value[0].Int64_Work  ,value[1].Int64_Work  ,w ,r )

          break;

          case "D_District_Temp" :


          break;

          case "Private" :

            flexible_out = trans.Private (w ,r )

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

          case "D_Area_Temp" :

            initialize.D_area_temp (w , r )

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