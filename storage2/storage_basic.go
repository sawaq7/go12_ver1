package storage2

import (

        "net/http"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/general/type5"
	    "github.com/sawaq7/go12_ver1/storage2/get"

//	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
        "cloud.google.com/go/storage"

                                        )

///
///     ã¹ãã¬ãE¸ã¢ã¯ã»ã¹ã®åºæ¬ã«ã¼ãã³
///


func Storage_basic( function string ,flexible_in1 interface{} ,flexible_in2 interface{} ,w http.ResponseWriter, r *http.Request )  (flexible_out1 interface{} ,flexible_out2 interface{}) {

//     IN    function    : ãã¡ã³ã¯ã·ã§ã³ã
//        ãããããããããEãtrans ,check ,initialize ,sortãetc
//     IN flexible_in1ã  : åE¨®ã¤ã³ããããã¼ã¿ãEå¥ç´åç§EE//     IN flexible_in2ã  : åE¨®ã¤ã³ããããã¼ã¿ãEå¥ç´åç§EE//     IN    w      ãã : ã¬ã¹ãã³ã¹ã©ã¤ã¿ã¼
//     IN    r      ãã : ãªã¯ã¨ã¹ããã©ã¡ã¼ã¿

//     out flexible_out1  : åE¨®ã¢ã¦ããEãEãEEã¿ãEå¥ç´åç§EE//     out flexible_out2  : åE¨®ã¢ã¦ããEãEãEEã¿ãEå¥ç´åç§EE
//    fmt.Fprintf( w, "storage_basic start \n" )  // ãEãE¯
//    fmt.Fprintf( w, "storage_basic function \n" ,function )  // ãEãE¯

///
///  ãã¡ã³ã¯ã·ã§ã³ã«ããåE¨®å¦çEåE²ãã¦è¡ããE///

	switch function {

///
///  ãªã¼ãã³ã®å ´åE///

      case "open" :     // trans ã®å ´åE
         value1, _ := flexible_in1.(string)
         value2, _ := flexible_in2.(string)

         flexible_out1 = File_Open ( w ,r ,value1 ,value2 )

      break;

///
///  ã¯ãªã¨ã¤ããEå ´åE///

      case "create" :

         value1, _ := flexible_in1.(string)
         value2, _ := flexible_in2.(string)

         flexible_out1 = File_Create ( w ,r ,value1 ,value2 )


      break;


///
///  ãEªã¼ããEå ´åE///

      case "delete" :

         value1, _ := flexible_in1.(string)
         value2, _ := flexible_in2.(string)

         File_Delete ( w , r ,value1 ,value2 )

      break;

///
///  ã³ããEã®å ´åE///

      case "copy" :

         value1, _ := flexible_in1.(string)
         value2, _ := flexible_in2.([]type5.General_Work)

         File_Copy ( w , r ,value1 ,value2[0].String_Work ,value2[1].String_Work )

      break;

///
///  ãªããEã ã®å ´åE///

      case "rename" :

         value1, _ := flexible_in1.(string)
         value2, _ := flexible_in2.([]type5.General_Work)

         File_Rename ( w , r ,value1 ,value2[0].String_Work ,value2[1].String_Work )

      break;

///
///  ã©ã¤ããEå ´åEEEtringåEè¡ãæ¸ãï¼E///

      case "write" :

         value1, _ := flexible_in1.(*storage.Writer)
         value2, _ := flexible_in2.([]string)

         File_write ( w ,value1 ,value2 )

      break;

///
///  ã©ã¤ãï¼ãEå ´åEEæ§é ä½å1è¡ãæ¸ãï¼E///

      case "write2" :

         value1, _ := flexible_in1.([]type5.General_Work)

         File_Write_Struct ( w ,value1[0].Sw_Work ,value1[0].Int64_Work ,flexible_in2 )

      break;

///
///  ãªã¹ããEå ´åEEãã±ãEãªã¹ããåºãï¼E///

      case "list" :

         value1, _ := flexible_in1.(string)

         flexible_out1 = Bucket_List ( w ,r , value1 )

      break;

///
///  ãªã¹ãEã®å ´åEEãªãã¸ã§ã¯ããªã¹ããåºãï¼E///

      case "list2" :

         value1, _ := flexible_in1.(string)

         flexible_out1 = Object_List ( w ,r , value1 )

      break;

///
///  ãªã¹ãEã®å ´åEEãªãã¸ã§ã¯ããªã¹ãEè©³ç´°Eãåºãï¼E///

      case "list3" :

         value1, _ := flexible_in1.(string)

         flexible_out1 = Object_List_Detail ( w ,r , value1 )

      break;

///
///  ã·ã§ã¦1ã®å ´åEEã°ã©ãè¡¨ç¤ºEï¼E///

      case "show1" :

         value1, _ := flexible_in1.(string)
         value2, _ := flexible_in2.(string)

         get.Image_file_show ( w ,r , value1 ,value2 )

      break;

//
///  ã·ã§ã¦2ã®å ´åEEã°ã©ãè¡¨ç¤º2EE///

      case "show2" :

         value1, _ := flexible_in1.(type5.Image_Show)

         get.Image_file_show2 ( w ,r , value1 )

      break;


    }

	return flexible_out1 ,flexible_out2

}
