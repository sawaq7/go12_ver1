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
///     ãEEã¿ã¹ãã¢ã¼ã¢ã¯ã»ã¹ã®ã¡ã¤ã³ã«ã¼ãã³ Eãfor sghã)
///


func Datastore_sgh( fname string ,function string ,flexible_in interface{} ,w http.ResponseWriter, r *http.Request )  (flexible_out interface{} ) {

//     IN    fname       : ãEEã¿ã¹ãã¢ã®ãã¡ã¤ã«åE//     IN    function    : ãã¡ã³ã¯ã·ã§ã³ã
//        ãããããããããEãtrans ,check ,initialize ,sortãetc
//     IN flexible_inã  : åE¨®ã¤ã³ããããã¼ã¿ãEå¥ç´åç§EE//     IN    w      ãã : ã¬ã¹ãã³ã¹ã©ã¤ã¿ã¼
//     IN    r      ãã : ãªã¯ã¨ã¹ããã©ã¡ã¼ã¿

//     out flexible_out  : åE¨®ã¢ã¦ããEãEãEEã¿ãEå¥ç´åç§EE
//    fmt.Fprintf( w, "datastore_sgh start \n" )  // ãEãE¯
//    fmt.Fprintf( w, "datastore_sgh function \n" ,function )  // ãEãE¯
//    fmt.Fprintf( w, "datastore_sgh fname \n" ,fname )  // ãEãE¯

///
///  ãã¡ã³ã¯ã·ã§ã³ã»ãã¡ã¤ã«åã«ããåE¨®å¦çEåE²ãã¦è¡ããE///

	switch function {

///
///  ãã©ã³ã¹ã®å ´åE///

      case "trans" :     // trans ã®å ´åE
        switch fname {

          case "Deliver" :  // è©²å½ããéEéæå ±ãã²ãE

            // ç©ºã¤ã³ã¿ã¼ãã§ã¤ã¹å¤æ°ããããªã¥ã¼å¤ãã²ãE

            value, _ := flexible_in.([]type5.General_Work)
            flexible_out = trans.Deliver ( value[0].Int64_Work  ,value[1].Int64_Work  ,w ,r )

          break;

          case "D_Area" :  // å·è»æå ±ãã²ãE

            value, _ := flexible_in.(int64)
            flexible_out = trans.D_area ( 0 ,value ,w ,r )

          break;

          case "D_District" :  // è©²å½ããå°åºæE ±ãã²ãE

            // ç©ºã¤ã³ã¿ã¼ãã§ã¤ã¹å¤æ°ããããªã¥ã¼å¤ãã²ãE

            value, _ := flexible_in.([]type5.General_Work)
            flexible_out = trans.D_district ( value[0].Int64_Work  ,value[1].Int64_Work  ,w ,r )

          break;

          case "D_District_Temp" :     // å°åºæE ±ã®ä¸æãã¡ã¤ã«ã®å ´åE


          break;

          case "Private" :  // workerã®åäººæE ±ãã²ãE

            flexible_out = trans.Private (w ,r )

          break;

          case "Car" :  // å·è»æå ±ãã²ãE

            value, _ := flexible_in.(int64)
            flexible_out = trans.Car_district ( value ,w ,r )

          break;

          case "Sgh_Ai" :  // AIæE ±ãã²ãE

            value, _ := flexible_in.(int64)
            flexible_out = trans.Sgh_ai( value  ,w , r  )

          break;

        }
      break;

      case "trans2" :       // trans2 ã®å ´åE
        switch fname {

          case "Deliver" :  // è©²å½ããéEéæå ±ãã²ãE



          break;

          case "D_Area" :  // å·è»æå ±ãã²ãE

            value, _ := flexible_in.(int64)
            flexible_out = trans.D_area_district ( w ,r ,value )


          break;

          case "Private" :  // workerã®åäººæE ±ãã²ãE



          break;

          case "Car" :  // å·è»æå ±ãã²ãE



          break;

          case "D_District_Temp" :     // å°åºæE ±ã®ä¸æãã¡ã¤ã«ã®å ´åE


          break;
        }
      break;

///
///  åæåãEå ´åE///

      case "initialize" :

        switch fname {

          case "D_Area_Temp" :   // ä¸æãã¡ã¤ã«ããã«ã¬ã³ããEã¨ãªã¢ã®æE ±ãã²ãEã

            initialize.D_area_temp (w , r ) //  æ¢å­ãEãD_Area_Temp temporary-fileãã¯ãªã¢ã¼

          break;

          case "D_District_Temp" :     // ã¨ãªã¢æE ±ã®å ´åE
            initialize.D_district_temp (w , r ) // temporary-fileãã¤ãã·ã£ã©ã¤ãº

          break;

          case "Sgh_Ai" :     // ã¨ãªã¢æE ±ã®å ´åE
            value, _ := flexible_in.(int64)
            initialize.Sgh_ai( value ,w , r )  //   æ¢å­ãEAIæE ±ãåé¤

          break;

        }

      break;


///
///  ãã§ãE¯ã®å ´åE///

      case "check" :

        switch fname {

          case "D_Area" :   // ééå°åºã®ã¨ãªã¢ã®MAXå¤ãã²ãE

            // ç©ºã¤ã³ã¿ã¼ãã§ã¤ã¹å¤æ°ããããªã¥ã¼å¤ãã²ãE

            value, _ := flexible_in.(int64)
            flexible_out = check.D_area ( w , r  ,value )

          break;

          case "D_District_Temp" :   // ä¸æãã¡ã¤ã«ããã«ã¬ã³ããEå°åºã®æE ±ãã²ãEã

            flexible_out = check.D_district_temp (w , r )

          break;

          case "Car" :   // ééå°åºã«ã¼NOã®MAXå¤ãã²ãE

            // ç©ºã¤ã³ã¿ã¼ãã§ã¤ã¹å¤æ°ããããªã¥ã¼å¤ãã²ãE

            value, _ := flexible_in.(int64)
            flexible_out = check.Car_no_max(w , r  ,value)

          break;

        }

      break;

    }

	return flexible_out

}

