package storage3

import (

        "net/http"
//	    "fmt"

	    "client/tokura/storage3/put1"
	    "client/tokura/storage3/check4"
	    "client/tokura/storage3/trans4"
	    "client/tokura/storage3/delete1"
	    "client/tokura/storage3/initialize3"
	    "client/tokura/storage3/struct_set"

	    "general/type5"
	    "client/tokura/suiri/type4"

                                                )

///
///     クライアントのストレッジのアクセスルーチン （　for tokura　)
///     Storage_クライアント名


func Storage_tokura( fname string ,function string ,flexible_in1 interface{} ,flexible_in2 interface{} ,w http.ResponseWriter, r *http.Request )  (flexible_out1 interface{} ,flexible_out2 interface{}) {

//     IN    fname       : データストアのファイル名
//     IN    function    : ファンクション　
//        　　　　　　　　　＊　trans ,check ,initialize ,sort　etc
//     IN flexible_in1　  : 各種インプットデータ　（別紙参照）
//     IN flexible_in2　  : 各種インプットデータ　（別紙参照）
//     IN    w      　　 : レスポンスライター
//     IN    r      　　 : リクエストパラメータ

//     out flexible_out1  : 各種アウトプットデータ　（別紙参照）
//     out flexible_out2  : 各種アウトプットデータ　（別紙参照）

//    fmt.Fprintf( w, "storage_tokura start \n" )  // デバック
//    fmt.Fprintf( w, "storage_tokura function \n" ,function )  // デバック
//    fmt.Fprintf( w, "storage_tokura fname \n" ,fname )  // デバック

///
///  ファンクション・ファイル名により各種処理を分岐して行う。
///

	switch function {

///
///  トランスの場合
///

      case "trans" :     // trans の場合

        switch fname {

          case "Water_Line" :     // 水路ラインファイルの場合

            value, _ := flexible_in1.(string)
            flexible_out1 = trans4.Water_line ( value , w ,r )

          break;

          case "Water2" :     // 水路ファイルの場合

            flexible_out1 = trans4.Water2 ( w ,r )

          break;

          case "Water2_Temp" :     // 水路テンポラリーファイルの場合

            flexible_out1 = trans4.Water2_temp( w , r  )

          break;


        }
      break;

///
///  初期化の場合
///

      case "initialize" :

        switch fname {

          case "Water2_Temp" :     // 水路テンポラリーファイルの場合

            value, _ := flexible_in1.(type4.Water2_Temp)

            initialize3.Water2_temp (w , r ,value)

          break;

        }

      break;


///
///  チェックの場合
///

      case "check" :

        switch fname {

          case "Water_Line" :  // 水路ラインファイルの場合

            value, _ := flexible_in1.(string)

            flexible_out1 = check4.Water_line_re_num( value  ,w , r  )

          break;

        }

      break;

///
///  削除の場合
///

      case "delete" :

        switch fname {

          case "Water_Line" :  // 水路ラインファイルの場合

            value, _ := flexible_in1.(int64)
            value2, _ := flexible_in2.(string)

            delete1.Water_line( w , r ,value , value2 )

          break;

          case "Water2" :  // 水路ファイルの場合

            delid, _ := flexible_in1.(int64)

            delete1.Water2( w , r ,delid )

          break;

        }

      break;
///
///      プットの場合
///

      case "put" :     // アップデート

        switch fname {

          case "Water_Line" :     // 水路ラインファイルの場合

            // 空インターフェイス変数よりバリュー値をゲット

            value, _ := flexible_in1.(type4.Water_Line)

            put1.Water_line( w , r ,value )

          break;

          case "Water_Slope" :     // 導水勾配線ファイルの場合


          break;

          case "Water2" :        // 水路ファイル

            water2, _ := flexible_in1.( type4.Water2 )

            put1.Water2 ( w , r ,water2 )

          break;

        }

      break;

      case "put2" :     // put2 の場合

        switch fname {

          case "Water_Line" :     // 水路ラインファイルの場合

            // 空インターフェイス変数よりバリュー値をゲット

            value, _ := flexible_in1.([]type5.General_Work)
            value2, _ := flexible_in2.(type4.Water_Line)

            put1.Water_line_update ( w , r ,value[0].Int64_Work ,value[0].String_Work ,value2 )

          break;

          case "Water2" :        // 水路ファイル

            water2, _ := flexible_in1.( type4.Water2 )

            put1.Water2_new ( w , r ,water2 )

          break;

        }

      break;

      case "put3" :     // put3 の場合

        switch fname {

          case "Water2" :        // 水路ファイル

            updid, _ := flexible_in1.( int64 )
            water2, _ := flexible_in2.( type4.Water2 )

            put1.Water2_update ( w , r ,updid ,water2 )

          break;

        }

      break;


      case "put_test" :              // test用

        switch fname {

          case "Water_Line" :     // 水路ラインファイルの場合

            // 空インターフェイス変数よりバリュー値をゲット

            general_work_value, _ := flexible_in1.([]type5.General_Work)
            struct_colle_value, _ := flexible_in2.(type4.Struct_Colle)

//            fmt.Fprintf( w, "storage_tokura : general_work_value %v\n", general_work_value )
//            fmt.Fprintf( w, "storage_tokura : struct_colle_value.Water_Line_Slice %v\n", struct_colle_value.Water_Line_Slice )  // デバック
//            fmt.Fprintf( w, "storage_tokura : struct_colle_value.Water2_Slice %v\n", struct_colle_value.Water2_Slice )  // デバック

            flexible_out1 = general_work_value
            flexible_out2 = struct_colle_value

          break;

        }

      break;

      case "struct_set" :        // 構造体セットの場合

        switch fname {

          case "Water2" :     // 水路ファイルの場合

            // 空インターフェイス変数よりバリュー値をゲット

            line, _ := flexible_in1.(string)

            flexible_out1 = struct_set.Water2( w , line )

          break;

        }

      break;


    }

	return flexible_out1 ,flexible_out2

}

