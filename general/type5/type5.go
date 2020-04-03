package type5

import (

	     "time"
	     "cloud.google.com/go/storage"
	                                   )

///
///    チE�Eル コマンド用　フォーマット集
///

///
///    チE�Eタストア　コピ�E　リスチE///

type  Ds_Copy_List    struct           {

          Id             int64    //　チE�Eタid
	      Basic_Name     string   // 基本のチE�Eタストア吁E    �E�！E未使用のため廁E��予定　�E�！E	      Copy_Name      string   // コピ�E允E�EチE�Eタストア吁E	      New_Name       string   // ニューチE�Eタストア吁E
}

///
///    チE�Eタベ�Eス　アクセス　リスチE///

type  Db_Access_List    struct           {

          Id              int64       //　チE�Eタid
          Line_No         int64    // 行NO.
          Db_Type         string      // チE�Eタベ�EスタイチE                                      //  ds : チE�Eタストア
                                      //  sr : ストレチE��

          Access_Type        string   // アクセスタイチE                                      //  copy
                                      //  rename

          Project_Name     string     // プロジェクト名
	      Bucket_Name     string      // バケチE��吁E	      Basic_File_Name    string   // 基本のファイル吁E	      New_File_Name      string   // ニューファイル吁E
}

///
///    チE�Eタベ�Eス　アクセス　リスチE///

type  Db_Access_List2    struct           {

          Id              int64       //　チE�Eタid
          Line_No         int64    // 行NO.
          Db_Type         string      // チE�Eタベ�EスタイチE                                      //  ds : チE�Eタストア
                                      //  sr : ストレチE��

          Access_Type        string   // アクセスタイチE                                      //  copy
                                      //  rename

          Project_Name     string     // プロジェクト名
	      Bucket_Name     string      // バケチE��吁E	      Basic_File_Name    string   // 基本のファイル吁E	      New_File_Name      string   // ニューファイル吁E
}

///
///           record for csv file
///

type  Csv_Inf    struct           {

          Id            int64
          Line_No       int64    // 行NO.
          File_Name     string   // ファイル吁E          Column_Num    int64    // 列数
	      Column1       string   // 列！E	      Column2       string   // 列！E	      Column3       string   // 列！E	      Column4       string   // 列！E	      Column5       string   // 列！E	      Column6       string   // 列！E	      Column7       string   // 列！E	      Column8       string   // 列！E	      Column9       string   // 列！E	      Column10      string   // 列１！E
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

      Records_Num    int64   // csvレコード�E構造体�E数

//          Id            int64
//          Line_No       int64       // 行NO.
      Records[10]    []Csv_Inf   // csvレコード�E構造佁E
}

///
///   画像ファイル表示
///

type  Image_Show    struct           {

          Id              int64    // チE�Eタid
	      File_Name       string   // ファイル吁E	      Url             string   // url
}
///
///  ストレチE��のバケチE��・オブジェクト（ファイル�E��E表示用
///

type  Storage_B_O_View    struct           {

          Line_No         int64    // 行NO.
          Project_Name     string   // プロジェクト名
	      Bucket_Name     string   // バケチE��吁E	      Object_Name     string   // オブジェクト名
	      Created       time.Time  //作�E時間

}

///
///  ストレチE��のバケチE��・オブジェクト（ファイル�E��Eコモン用(チE�Eタストア�E�E///

type  Storage_B_O_Temp    struct           {

          Id              int64    // チE�Eタid
          Line_No         int64    // 行NO.
          Project_Name     string   // プロジェクト名
	      Bucket_Name     string   // バケチE��吁E	      Object_Name     string   // オブジェクト名
	      Created       time.Time  //作�E時間

}

type  Storage_B_O    struct           {

          Id              int64    // チE�Eタid
          Line_No         int64    // 行NO.
          Project_Name     string   // プロジェクト名
	      Bucket_Name     string   // バケチE��吁E	      Object_Name     string   // オブジェクト名
	      Created       time.Time  //作�E時間

}
///
///  ワーク用
///

type  General_Work    struct           {

          Int64_Work     int64           // int型ワークエリア
          Float64_Work   float64         // float型ワークエリア
	      String_Work    string          // string型ワークエリア
	      Sw_Work        *storage.Writer // ストレチE��ライター垁E
}
