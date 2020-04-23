package storage3

import (

        "net/http"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/client/tokura/storage3/put1"
	    "github.com/sawaq7/go12_ver1/client/tokura/storage3/check4"
	    "github.com/sawaq7/go12_ver1/client/tokura/storage3/trans4"
	    "github.com/sawaq7/go12_ver1/client/tokura/storage3/delete1"
	    "github.com/sawaq7/go12_ver1/client/tokura/storage3/initialize3"
	    "github.com/sawaq7/go12_ver1/client/tokura/storage3/struct_set"

	    "github.com/sawaq7/go12_ver1/general/type5"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"

                                                )

///
///      Storege_tokura  main routine
///


func Storage_tokura( fname string ,function string ,flexible_in1 interface{} ,flexible_in2 interface{} ,w http.ResponseWriter, r *http.Request )  (flexible_out1 interface{} ,flexible_out2 interface{}) {

//     IN    fname       : file-name
//     IN    function    : 　
//        　　　　　　　　　trans ,check ,initialize ,sort　etc
//     IN flexible_in1　  : see attachment
//     IN flexible_in2　  : see attachment
//     IN    w      　　 : レスポンスライター
//     IN    r      　　 : リクエストパラメータ

//     out1 flexible_out1  : see attachment
//     out2 flexible_out2  : see attachment


//    fmt.Fprintf( w, "storage_tokura start \n" )
//    fmt.Fprintf( w, "storage_tokura function \n" ,function )
//    fmt.Fprintf( w, "storage_tokura fname \n" ,fname )

///
///     select process  by function and file-name
///

	switch function {

      case "trans" :
        switch fname {

          case "Water_Line" :
            value, _ := flexible_in1.(string)
            flexible_out1 = trans4.Water_line ( value , w ,r )

          break;

          case "Water2" :
            flexible_out1 = trans4.Water2 ( w ,r )

          break;

          case "Water2_Temp" :
            flexible_out1 = trans4.Water2_temp( w , r  )

          break;


        }
      break;

      case "initialize" :

        switch fname {

          case "Water2_Temp" :
            value, _ := flexible_in1.(type4.Water2_Temp)

            initialize3.Water2_temp (w , r ,value)

          break;

        }

      break;

      case "check" :

        switch fname {

          case "Water_Line" :
            value, _ := flexible_in1.(string)

            flexible_out1 = check4.Water_line_re_num( value  ,w , r  )

          break;

        }

      break;

      case "delete" :

        switch fname {

          case "Water_Line" :
            value, _ := flexible_in1.(int64)
            value2, _ := flexible_in2.(string)

            delete1.Water_line( w , r ,value , value2 )

          break;

          case "Water2" :
            delid, _ := flexible_in1.(int64)

            delete1.Water2( w , r ,delid )

          break;

        }

      break;


      case "put" :
        switch fname {

          case "Water_Line" :

            value, _ := flexible_in1.(type4.Water_Line)

            put1.Water_line( w , r ,value )

          break;

          case "Water_Slope" :

          break;

          case "Water2" :

            water2, _ := flexible_in1.( type4.Water2 )

            put1.Water2 ( w , r ,water2 )

          break;

        }

      break;

      case "put2" :
        switch fname {

          case "Water_Line" :

            value, _ := flexible_in1.([]type5.General_Work)
            value2, _ := flexible_in2.(type4.Water_Line)

            put1.Water_line_update ( w , r ,value[0].Int64_Work ,value[0].String_Work ,value2 )

          break;

          case "Water2" :

            water2, _ := flexible_in1.( type4.Water2 )

            put1.Water2_new ( w , r ,water2 )

          break;

        }

      break;

      case "put3" :
        switch fname {

          case "Water2" :

            updid, _ := flexible_in1.( int64 )
            water2, _ := flexible_in2.( type4.Water2 )

            put1.Water2_update ( w , r ,updid ,water2 )

          break;

        }

      break;


      case "put_test" :

        switch fname {

          case "Water_Line" :

            general_work_value, _ := flexible_in1.([]type5.General_Work)
            struct_colle_value, _ := flexible_in2.(type4.Struct_Colle)

//            fmt.Fprintf( w, "storage_tokura : general_work_value %v\n", general_work_value )
//            fmt.Fprintf( w, "storage_tokura : struct_colle_value.Water_Line_Slice %v\n", struct_colle_value.Water_Line_Slice )
//            fmt.Fprintf( w, "storage_tokura : struct_colle_value.Water2_Slice %v\n", struct_colle_value.Water2_Slice )

            flexible_out1 = general_work_value
            flexible_out2 = struct_colle_value

          break;

        }

      break;

      case "struct_set" :
        switch fname {

          case "Water2" :

            line, _ := flexible_in1.(string)

            flexible_out1 = struct_set.Water2( w , line )

          break;

        }

      break;

    }

	return flexible_out1 ,flexible_out2

}

