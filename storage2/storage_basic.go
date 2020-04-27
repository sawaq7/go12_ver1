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
///    storage's basic routine
///


func Storage_basic( function string ,flexible_in1 interface{} ,flexible_in2 interface{} ,w http.ResponseWriter, r *http.Request )  (flexible_out1 interface{} ,flexible_out2 interface{}) {

//     IN    function    : trans ,check ,initialize ,sort　etc　        　　　　　　　　　
//     IN flexible_in1　  : see attachment
//     IN flexible_in2　  : see attachment
//     IN    w      : response-writer
//     IN    r      : request- paramete

//     out1 flexible_out1  : see attachment
//     out2 flexible_out2  : see attachment

//    fmt.Fprintf( w, "storage_basic start \n" )
//    fmt.Fprintf( w, "storage_basic function \n" ,function )

///
///     select process  by function
///

	switch function {

      case "open" :
         value1, _ := flexible_in1.(string)
         value2, _ := flexible_in2.(string)

         flexible_out1 = File_Open ( w ,r ,value1 ,value2 )

      break;

      case "create" :

         value1, _ := flexible_in1.(string)
         value2, _ := flexible_in2.(string)

         flexible_out1 = File_Create ( w ,r ,value1 ,value2 )

      break;

      case "delete" :

         value1, _ := flexible_in1.(string)
         value2, _ := flexible_in2.(string)

         File_Delete ( w , r ,value1 ,value2 )

      break;

      case "copy" :

         value1, _ := flexible_in1.(string)
         value2, _ := flexible_in2.([]type5.General_Work)

         File_Copy ( w , r ,value1 ,value2[0].String_Work ,value2[1].String_Work )

      break;

      case "rename" :

         value1, _ := flexible_in1.(string)
         value2, _ := flexible_in2.([]type5.General_Work)

         File_Rename ( w , r ,value1 ,value2[0].String_Work ,value2[1].String_Work )

      break;

      case "write" :   // write line-data

         value1, _ := flexible_in1.(*storage.Writer)
         value2, _ := flexible_in2.([]string)

         File_write ( w ,value1 ,value2 )

      break;

      case "write2" :    //   write struct-data

         value1, _ := flexible_in1.([]type5.General_Work)

         File_Write_Struct ( w ,value1[0].Sw_Work ,value1[0].Int64_Work ,flexible_in2 )

      break;

      case "list" :   //  get bucket-list

         value1, _ := flexible_in1.(string)

         flexible_out1 = Bucket_List ( w ,r , value1 )

      break;

      case "list2" :    // get object-list

         value1, _ := flexible_in1.(string)

         flexible_out1 = Object_List ( w ,r , value1 )

      break;

      case "list3" :     // get object-list which is detail

         value1, _ := flexible_in1.(string)

         flexible_out1 = Object_List_Detail ( w ,r , value1 )

      break;

      case "show1" :   //   show graf  on web ( type1 )

         value1, _ := flexible_in1.(string)
         value2, _ := flexible_in2.(string)

         get.Image_file_show ( w ,r , value1 ,value2 )

      break;

      case "show2" :    //   show graf  on web ( type2 )

         value1, _ := flexible_in1.(type5.Image_Show)

         get.Image_file_show2 ( w ,r , value1 )

      break;


    }

	return flexible_out1 ,flexible_out2

}
