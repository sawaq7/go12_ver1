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
///     クライアント�EストレチE��のアクセスルーチン �E�　for tokura　)
///     Storage_クライアント名


func Storage_tokura( fname string ,function string ,flexible_in1 interface{} ,flexible_in2 interface{} ,w http.ResponseWriter, r *http.Request )  (flexible_out1 interface{} ,flexible_out2 interface{}) {

//     IN    fname       : チE�Eタストアのファイル吁E//     IN    function    : ファンクション　
//        　　　　　　　　　�E�　trans ,check ,initialize ,sort　etc
//     IN flexible_in1　  : 吁E��インプットデータ　�E�別紙参照�E�E//     IN flexible_in2　  : 吁E��インプットデータ　�E�別紙参照�E�E//     IN    w      　　 : レスポンスライター
//     IN    r      　　 : リクエストパラメータ

//     out flexible_out1  : 吁E��アウト�EチE��チE�Eタ　�E�別紙参照�E�E//     out flexible_out2  : 吁E��アウト�EチE��チE�Eタ　�E�別紙参照�E�E
//    fmt.Fprintf( w, "storage_tokura start \n" )  // チE��チE��
//    fmt.Fprintf( w, "storage_tokura function \n" ,function )  // チE��チE��
//    fmt.Fprintf( w, "storage_tokura fname \n" ,fname )  // チE��チE��

///
///  ファンクション・ファイル名により吁E��処琁E��刁E��して行う、E///

	switch function {

///
///  トランスの場吁E///

      case "trans" :     // trans の場吁E
        switch fname {

          case "Water_Line" :     // 水路ラインファイルの場吁E
            value, _ := flexible_in1.(string)
            flexible_out1 = trans4.Water_line ( value , w ,r )

          break;

          case "Water2" :     // 水路ファイルの場吁E
            flexible_out1 = trans4.Water2 ( w ,r )

          break;

          case "Water2_Temp" :     // 水路チE��ポラリーファイルの場吁E
            flexible_out1 = trans4.Water2_temp( w , r  )

          break;


        }
      break;

///
///  初期化�E場吁E///

      case "initialize" :

        switch fname {

          case "Water2_Temp" :     // 水路チE��ポラリーファイルの場吁E
            value, _ := flexible_in1.(type4.Water2_Temp)

            initialize3.Water2_temp (w , r ,value)

          break;

        }

      break;


///
///  チェチE��の場吁E///

      case "check" :

        switch fname {

          case "Water_Line" :  // 水路ラインファイルの場吁E
            value, _ := flexible_in1.(string)

            flexible_out1 = check4.Water_line_re_num( value  ,w , r  )

          break;

        }

      break;

///
///  削除の場吁E///

      case "delete" :

        switch fname {

          case "Water_Line" :  // 水路ラインファイルの場吁E
            value, _ := flexible_in1.(int64)
            value2, _ := flexible_in2.(string)

            delete1.Water_line( w , r ,value , value2 )

          break;

          case "Water2" :  // 水路ファイルの場吁E
            delid, _ := flexible_in1.(int64)

            delete1.Water2( w , r ,delid )

          break;

        }

      break;
///
///      プット�E場吁E///

      case "put" :     // アチE�EチE�EチE
        switch fname {

          case "Water_Line" :     // 水路ラインファイルの場吁E
            // 空インターフェイス変数よりバリュー値をゲチE��

            value, _ := flexible_in1.(type4.Water_Line)

            put1.Water_line( w , r ,value )

          break;

          case "Water_Slope" :     // 導水勾配線ファイルの場吁E

          break;

          case "Water2" :        // 水路ファイル

            water2, _ := flexible_in1.( type4.Water2 )

            put1.Water2 ( w , r ,water2 )

          break;

        }

      break;

      case "put2" :     // put2 の場吁E
        switch fname {

          case "Water_Line" :     // 水路ラインファイルの場吁E
            // 空インターフェイス変数よりバリュー値をゲチE��

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

      case "put3" :     // put3 の場吁E
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

          case "Water_Line" :     // 水路ラインファイルの場吁E
            // 空インターフェイス変数よりバリュー値をゲチE��

            general_work_value, _ := flexible_in1.([]type5.General_Work)
            struct_colle_value, _ := flexible_in2.(type4.Struct_Colle)

//            fmt.Fprintf( w, "storage_tokura : general_work_value %v\n", general_work_value )
//            fmt.Fprintf( w, "storage_tokura : struct_colle_value.Water_Line_Slice %v\n", struct_colle_value.Water_Line_Slice )  // チE��チE��
//            fmt.Fprintf( w, "storage_tokura : struct_colle_value.Water2_Slice %v\n", struct_colle_value.Water2_Slice )  // チE��チE��

            flexible_out1 = general_work_value
            flexible_out2 = struct_colle_value

          break;

        }

      break;

      case "struct_set" :        // 構造体セチE��の場吁E
        switch fname {

          case "Water2" :     // 水路ファイルの場吁E
            // 空インターフェイス変数よりバリュー値をゲチE��

            line, _ := flexible_in1.(string)

            flexible_out1 = struct_set.Water2( w , line )

          break;

        }

      break;


    }

	return flexible_out1 ,flexible_out2

}

