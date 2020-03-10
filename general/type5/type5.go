package type5

import (

	     "time"
	     "cloud.google.com/go/storage"
	                                   )

///
///    ツール コマンド用　フォーマット集
///

///
///    データストア　コピー　リスト
///

type  Ds_Copy_List    struct           {

          Id             int64    //　データid
	      Basic_Name     string   // 基本のデータストア名     ＊＊ 未使用のため廃止予定　＊＊
	      Copy_Name      string   // コピー元のデータストア名
	      New_Name       string   // ニューデータストア名

}

///
///    データベース　アクセス　リスト
///

type  Db_Access_List    struct           {

          Id              int64       //　データid
          Line_No         int64    // 行NO.
          Db_Type         string      // データベースタイプ
                                      //  ds : データストア
                                      //  sr : ストレッジ

          Access_Type        string   // アクセスタイプ
                                      //  copy
                                      //  rename

          Project_Name     string     // プロジェクト名
	      Bucket_Name     string      // バケット名
	      Basic_File_Name    string   // 基本のファイル名
	      New_File_Name      string   // ニューファイル名

}

///
///    データベース　アクセス　リスト
///

type  Db_Access_List2    struct           {

          Id              int64       //　データid
          Line_No         int64    // 行NO.
          Db_Type         string      // データベースタイプ
                                      //  ds : データストア
                                      //  sr : ストレッジ

          Access_Type        string   // アクセスタイプ
                                      //  copy
                                      //  rename

          Project_Name     string     // プロジェクト名
	      Bucket_Name     string      // バケット名
	      Basic_File_Name    string   // 基本のファイル名
	      New_File_Name      string   // ニューファイル名

}

///
///           record for csv file
///

type  Csv_Inf    struct           {

          Id            int64
          Line_No       int64    // 行NO.
          File_Name     string   // ファイル名
          Column_Num    int64    // 列数
	      Column1       string   // 列１
	      Column2       string   // 列２
	      Column3       string   // 列３
	      Column4       string   // 列４
	      Column5       string   // 列５
	      Column6       string   // 列６
	      Column7       string   // 列７
	      Column8       string   // 列８
	      Column9       string   // 列９
	      Column10      string   // 列１０

}

///
///
///

type  Interpret    struct           {

//          Id            int64
//          Line_No       int64    // 行NO.

	      Ex_All        string   // 計算式ALL
	      Ex_Parts      string   // 計算式Parts

}
///
///           record for csv file
///

type  Csv_Records    struct           {

      Records_Num    int64   // csvレコードの構造体の数

//          Id            int64
//          Line_No       int64       // 行NO.
      Records[10]    []Csv_Inf   // csvレコードの構造体

}

///
///   画像ファイル表示
///

type  Image_Show    struct           {

          Id              int64    // データid
	      File_Name       string   // ファイル名
	      Url             string   // url
}
///
///  ストレッジのバケット・オブジェクト（ファイル）の表示用
///

type  Storage_B_O_View    struct           {

          Line_No         int64    // 行NO.
          Project_Name     string   // プロジェクト名
	      Bucket_Name     string   // バケット名
	      Object_Name     string   // オブジェクト名
	      Created       time.Time  //作成時間

}

///
///  ストレッジのバケット・オブジェクト（ファイル）のコモン用(データストア）
///

type  Storage_B_O_Temp    struct           {

          Id              int64    // データid
          Line_No         int64    // 行NO.
          Project_Name     string   // プロジェクト名
	      Bucket_Name     string   // バケット名
	      Object_Name     string   // オブジェクト名
	      Created       time.Time  //作成時間

}

type  Storage_B_O    struct           {

          Id              int64    // データid
          Line_No         int64    // 行NO.
          Project_Name     string   // プロジェクト名
	      Bucket_Name     string   // バケット名
	      Object_Name     string   // オブジェクト名
	      Created       time.Time  //作成時間

}
///
///  ワーク用
///

type  General_Work    struct           {

          Int64_Work     int64           // int型ワークエリア
          Float64_Work   float64         // float型ワークエリア
	      String_Work    string          // string型ワークエリア
	      Sw_Work        *storage.Writer // ストレッジライター型

}