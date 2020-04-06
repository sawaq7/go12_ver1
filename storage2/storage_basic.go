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

//     IN    function    : ファンクション　
//        　　　　　　　　　�E�　trans ,check ,initialize ,sort　etc
//     IN flexible_in1　  : 吁E��インプットデータ　�E�別紙参照�E�E//     IN flexible_in2　  : 吁E��インプットデータ　�E�別紙参照�E�E//     IN    w      　　 : レスポンスライター
//     IN    r      　　 : リクエストパラメータ

//     out flexible_out1  : 吁E��アウト�EチE��チE�Eタ　�E�別紙参照�E�E//     out flexible_out2  : 吁E��アウト�EチE��チE�Eタ　�E�別紙参照�E�E
//    fmt.Fprintf( w, "storage_basic start \n" )  // チE��チE��
//    fmt.Fprintf( w, "storage_basic function \n" ,function )  // チE��チE��

///
///  ファンクションにより吁E��処琁E��刁E��して行う、E///

	switch function {

///
///  オープンの場吁E///

      case "open" :     // trans の場吁E
         value1, _ := flexible_in1.(string)
         value2, _ := flexible_in2.(string)

         flexible_out1 = File_Open ( w ,r ,value1 ,value2 )

      break;

///
///  クリエイト�E場吁E///

      case "create" :

         value1, _ := flexible_in1.(string)
         value2, _ := flexible_in2.(string)

         flexible_out1 = File_Create ( w ,r ,value1 ,value2 )


      break;


///
///  チE��ート�E場吁E///

      case "delete" :

         value1, _ := flexible_in1.(string)
         value2, _ := flexible_in2.(string)

         File_Delete ( w , r ,value1 ,value2 )

      break;

///
///  コピ�Eの場吁E///

      case "copy" :

         value1, _ := flexible_in1.(string)
         value2, _ := flexible_in2.([]type5.General_Work)

         File_Copy ( w , r ,value1 ,value2[0].String_Work ,value2[1].String_Work )

      break;

///
///  リネ�Eムの場吁E///

      case "rename" :

         value1, _ := flexible_in1.(string)
         value2, _ := flexible_in2.([]type5.General_Work)

         File_Rename ( w , r ,value1 ,value2[0].String_Work ,value2[1].String_Work )

      break;

///
///  ライト�E場吁E�E�Etring垁E行を書く！E///

      case "write" :

         value1, _ := flexible_in1.(*storage.Writer)
         value2, _ := flexible_in2.([]string)

         File_write ( w ,value1 ,value2 )

      break;

///
///  ライト２�E場吁E�E�構造体型1行を書く！E///

      case "write2" :

         value1, _ := flexible_in1.([]type5.General_Work)

         File_Write_Struct ( w ,value1[0].Sw_Work ,value1[0].Int64_Work ,flexible_in2 )

      break;

///
///  リスト�E場吁E�E�バケチE��リストを出す！E///

      case "list" :

         value1, _ := flexible_in1.(string)

         flexible_out1 = Bucket_List ( w ,r , value1 )

      break;

///
///  リスチEの場吁E�E�オブジェクトリストを出す！E///

      case "list2" :

         value1, _ := flexible_in1.(string)

         flexible_out1 = Object_List ( w ,r , value1 )

      break;

///
///  リスチEの場吁E�E�オブジェクトリスチE詳細�E�を出す！E///

      case "list3" :

         value1, _ := flexible_in1.(string)

         flexible_out1 = Object_List_Detail ( w ,r , value1 )

      break;

///
///  ショウ1の場吁E�E�グラフ表示�E�！E///

      case "show1" :

         value1, _ := flexible_in1.(string)
         value2, _ := flexible_in2.(string)

         get.Image_file_show ( w ,r , value1 ,value2 )

      break;

//
///  ショウ2の場吁E�E�グラフ表示2�E�E///

      case "show2" :

         value1, _ := flexible_in1.(type5.Image_Show)

         get.Image_file_show2 ( w ,r , value1 )

      break;


    }

	return flexible_out1 ,flexible_out2

}
