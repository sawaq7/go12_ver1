package type6

import (

//          "time"

                  )
//
// the collection of struct information for reservation
//

//
// guest information　（地区情報）
//

type Guest struct {            /// ゲスト情報

       Id            int64           //　データid
       Guest_No      int64           // ゲストNO.
	   Guest_Name    string          // ゲスト名


   }

type Guest_Temp struct {         /// 一時ファイル用

       Id          int64           //　データid
       Guest_No    int64           // ゲストNO.
	   Guest_Name  string          // ゲスト名


   }
//
// reservation information　
//

type Guest_Reserve_Minor struct {               ///  予約情報

       Id             int64          // データid
       Line_No        int64          // 行NO.
       Date           string         // 予約日
       Guest_No       int64          // ゲストNO.
	   Guest_Name     string         // ゲスト名
	   Start_Time     string         // 開始時間
	   End_Time       string         // 終了時間

   }
//
// reservation information　
//

type Guest_Reserve_View struct {               ///  表示用予約情報

       Id             int64          // データid
       Line_No        int64          // 行NO.
       Date           string         // 予約日
       Guest_No       int64          // ゲストNO.
	   Guest_Name     string         // ゲスト名
	   Start_Time     string         // 開始時間
	   End_Time       string         // 終了時間
	   File_Name      string        // ファイル名
	   Url            string        // url

   }
//
// reservation information　
//

type Guest_Medical_Record struct {   ///  カルテ情報

       Id             int64          // データid
       Line_No        int64          // 行NO.
       Date           string         // 予約日
       Guest_No       int64          // ゲストNO.
	   Guest_Name     string         // ゲスト名
	   Text           string         // 本文

   }

//
// reservation information　
//

type Guest_Medical_Xray struct {   ///  レントゲン写真

       Id             int64          // データid
       Line_No        int64          // 行NO.
       Date           string         // 撮影日
       Guest_No       int64          // ゲストNO.
	   Guest_Name     string         // ゲスト名
	   File_Name      string         // ファイル名
	   Url            string         // url

	 }

//
// reservation information　
//

type Guest_Payment struct {               ///  予約情報

       Id             int64          // データid
       Line_No        int64          // 行NO.
       Date           string         // 予約日
       Guest_No       int64          // ゲストNO.
	   Guest_Name     string         // ゲスト名
	   Item           string         // 支払項目
	   Amount         int64          // 支払金額

   }