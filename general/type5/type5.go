package type5

import (

	     "time"
	     "cloud.google.com/go/storage"
	                                   )

///
///    ãEEã« ã³ãã³ãç¨ããã©ã¼ãããé
///

///
///    ãEEã¿ã¹ãã¢ãã³ããEããªã¹ãE///

type  Ds_Copy_List    struct           {

          Id             int64    //ããEEã¿id
	      Basic_Name     string   // åºæ¬ã®ãEEã¿ã¹ãã¢åE    Eï¼Eæªä½¿ç¨ã®ããå»E­¢äºå®ãEï¼E	      Copy_Name      string   // ã³ããEåEEãEEã¿ã¹ãã¢åE	      New_Name       string   // ãã¥ã¼ãEEã¿ã¹ãã¢åE
}

///
///    ãEEã¿ããEã¹ãã¢ã¯ã»ã¹ããªã¹ãE///

type  Db_Access_List    struct           {

          Id              int64       //ããEEã¿id
          Line_No         int64    // è¡NO.
          Db_Type         string      // ãEEã¿ããEã¹ã¿ã¤ãE                                      //  ds : ãEEã¿ã¹ãã¢
                                      //  sr : ã¹ãã¬ãE¸

          Access_Type        string   // ã¢ã¯ã»ã¹ã¿ã¤ãE                                      //  copy
                                      //  rename

          Project_Name     string     // ãã­ã¸ã§ã¯ãå
	      Bucket_Name     string      // ãã±ãEåE	      Basic_File_Name    string   // åºæ¬ã®ãã¡ã¤ã«åE	      New_File_Name      string   // ãã¥ã¼ãã¡ã¤ã«åE
}

///
///    ãEEã¿ããEã¹ãã¢ã¯ã»ã¹ããªã¹ãE///

type  Db_Access_List2    struct           {

          Id              int64       //ããEEã¿id
          Line_No         int64    // è¡NO.
          Db_Type         string      // ãEEã¿ããEã¹ã¿ã¤ãE                                      //  ds : ãEEã¿ã¹ãã¢
                                      //  sr : ã¹ãã¬ãE¸

          Access_Type        string   // ã¢ã¯ã»ã¹ã¿ã¤ãE                                      //  copy
                                      //  rename

          Project_Name     string     // ãã­ã¸ã§ã¯ãå
	      Bucket_Name     string      // ãã±ãEåE	      Basic_File_Name    string   // åºæ¬ã®ãã¡ã¤ã«åE	      New_File_Name      string   // ãã¥ã¼ãã¡ã¤ã«åE
}

///
///           record for csv file
///

type  Csv_Inf    struct           {

          Id            int64
          Line_No       int64    // è¡NO.
          File_Name     string   // ãã¡ã¤ã«åE          Column_Num    int64    // åæ°
	      Column1       string   // åï¼E	      Column2       string   // åï¼E	      Column3       string   // åï¼E	      Column4       string   // åï¼E	      Column5       string   // åï¼E	      Column6       string   // åï¼E	      Column7       string   // åï¼E	      Column8       string   // åï¼E	      Column9       string   // åï¼E	      Column10      string   // åï¼ï¼E
}

///
///
///

type  Interpret    struct           {

//          Id            int64
//          Line_No       int64    // è¡NO.

	      Ex_All        string   // è¨ç®å¼ALL
	      Ex_Parts      string   // è¨ç®å¼Parts

}
///
///           record for csv file
///

type  Csv_Records    struct           {

      Records_Num    int64   // csvã¬ã³ã¼ããEæ§é ä½ãEæ°

//          Id            int64
//          Line_No       int64       // è¡NO.
      Records[10]    []Csv_Inf   // csvã¬ã³ã¼ããEæ§é ä½E
}

///
///   ç»åãã¡ã¤ã«è¡¨ç¤º
///

type  Image_Show    struct           {

          Id              int64    // ãEEã¿id
	      File_Name       string   // ãã¡ã¤ã«åE	      Url             string   // url
}
///
///  ã¹ãã¬ãE¸ã®ãã±ãEã»ãªãã¸ã§ã¯ãï¼ãã¡ã¤ã«EãEè¡¨ç¤ºç¨
///

type  Storage_B_O_View    struct           {

          Line_No         int64    // è¡NO.
          Project_Name     string   // ãã­ã¸ã§ã¯ãå
	      Bucket_Name     string   // ãã±ãEåE	      Object_Name     string   // ãªãã¸ã§ã¯ãå
	      Created       time.Time  //ä½æEæé

}

///
///  ã¹ãã¬ãE¸ã®ãã±ãEã»ãªãã¸ã§ã¯ãï¼ãã¡ã¤ã«EãEã³ã¢ã³ç¨(ãEEã¿ã¹ãã¢EE///

type  Storage_B_O_Temp    struct           {

          Id              int64    // ãEEã¿id
          Line_No         int64    // è¡NO.
          Project_Name     string   // ãã­ã¸ã§ã¯ãå
	      Bucket_Name     string   // ãã±ãEåE	      Object_Name     string   // ãªãã¸ã§ã¯ãå
	      Created       time.Time  //ä½æEæé

}

type  Storage_B_O    struct           {

          Id              int64    // ãEEã¿id
          Line_No         int64    // è¡NO.
          Project_Name     string   // ãã­ã¸ã§ã¯ãå
	      Bucket_Name     string   // ãã±ãEåE	      Object_Name     string   // ãªãã¸ã§ã¯ãå
	      Created       time.Time  //ä½æEæé

}
///
///  ã¯ã¼ã¯ç¨
///

type  General_Work    struct           {

          Int64_Work     int64           // intåã¯ã¼ã¯ã¨ãªã¢
          Float64_Work   float64         // floatåã¯ã¼ã¯ã¨ãªã¢
	      String_Work    string          // stringåã¯ã¼ã¯ã¨ãªã¢
	      Sw_Work        *storage.Writer // ã¹ãã¬ãE¸ã©ã¤ã¿ã¼åE
}
