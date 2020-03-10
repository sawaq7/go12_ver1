package storage2

import (

        "net/http"
//	    "fmt"

	    "general/type5"
	    "storage2/get"

//	    "client/tokura/suiri/type4"
        "cloud.google.com/go/storage"

                                        )

///
///     ストレッジアクセスの基本ルーチン
///


func Storage_basic( function string ,flexible_in1 interface{} ,flexible_in2 interface{} ,w http.ResponseWriter, r *http.Request )  (flexible_out1 interface{} ,flexible_out2 interface{}) {

//     IN    function    : ファンクション　
//        　　　　　　　　　＊　trans ,check ,initialize ,sort　etc
//     IN flexible_in1　  : 各種インプットデータ　（別紙参照）
//     IN flexible_in2　  : 各種インプットデータ　（別紙参照）
//     IN    w      　　 : レスポンスライター
//     IN    r      　　 : リクエストパラメータ

//     out flexible_out1  : 各種アウトプットデータ　（別紙参照）
//     out flexible_out2  : 各種アウトプットデータ　（別紙参照）

//    fmt.Fprintf( w, "storage_basic start \n" )  // デバック
//    fmt.Fprintf( w, "storage_basic function \n" ,function )  // デバック

///
///  ファンクションにより各種処理を分岐して行う。
///

	switch function {

///
///  オープンの場合
///

      case "open" :     // trans の場合

         value1, _ := flexible_in1.(string)
         value2, _ := flexible_in2.(string)

         flexible_out1 = File_Open ( w ,r ,value1 ,value2 )

      break;

///
///  クリエイトの場合
///

      case "create" :

         value1, _ := flexible_in1.(string)
         value2, _ := flexible_in2.(string)

         flexible_out1 = File_Create ( w ,r ,value1 ,value2 )


      break;


///
///  デリートの場合
///

      case "delete" :

         value1, _ := flexible_in1.(string)
         value2, _ := flexible_in2.(string)

         File_Delete ( w , r ,value1 ,value2 )

      break;

///
///  コピーの場合
///

      case "copy" :

         value1, _ := flexible_in1.(string)
         value2, _ := flexible_in2.([]type5.General_Work)

         File_Copy ( w , r ,value1 ,value2[0].String_Work ,value2[1].String_Work )

      break;

///
///  リネームの場合
///

      case "rename" :

         value1, _ := flexible_in1.(string)
         value2, _ := flexible_in2.([]type5.General_Work)

         File_Rename ( w , r ,value1 ,value2[0].String_Work ,value2[1].String_Work )

      break;

///
///  ライトの場合 （string型1行を書く）
///

      case "write" :

         value1, _ := flexible_in1.(*storage.Writer)
         value2, _ := flexible_in2.([]string)

         File_write ( w ,value1 ,value2 )

      break;

///
///  ライト２の場合 （構造体型1行を書く）
///

      case "write2" :

         value1, _ := flexible_in1.([]type5.General_Work)

         File_Write_Struct ( w ,value1[0].Sw_Work ,value1[0].Int64_Work ,flexible_in2 )

      break;

///
///  リストの場合 （バケットリストを出す）
///

      case "list" :

         value1, _ := flexible_in1.(string)

         flexible_out1 = Bucket_List ( w ,r , value1 )

      break;

///
///  リスト2の場合 （オブジェクトリストを出す）
///

      case "list2" :

         value1, _ := flexible_in1.(string)

         flexible_out1 = Object_List ( w ,r , value1 )

      break;

///
///  リスト3の場合 （オブジェクトリスト(詳細）を出す）
///

      case "list3" :

         value1, _ := flexible_in1.(string)

         flexible_out1 = Object_List_Detail ( w ,r , value1 )

      break;

///
///  ショウ1の場合 （グラフ表示１）
///

      case "show1" :

         value1, _ := flexible_in1.(string)
         value2, _ := flexible_in2.(string)

         get.Image_file_show ( w ,r , value1 ,value2 )

      break;

//
///  ショウ2の場合 （グラフ表示2）
///

      case "show2" :

         value1, _ := flexible_in1.(type5.Image_Show)

         get.Image_file_show2 ( w ,r , value1 )

      break;


    }

	return flexible_out1 ,flexible_out2

}
