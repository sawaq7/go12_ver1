package type5

import (

	     "time"
	     "cloud.google.com/go/storage"
	                                   )

///
///    繝・・繝ｫ 繧ｳ繝槭Φ繝臥畑縲繝輔か繝ｼ繝槭ャ繝磯寔
///

///
///    繝・・繧ｿ繧ｹ繝医い縲繧ｳ繝斐・縲繝ｪ繧ｹ繝・///

type  Ds_Copy_List    struct           {

          Id             int64    //縲繝・・繧ｿid
	      Basic_Name     string   // 蝓ｺ譛ｬ縺ｮ繝・・繧ｿ繧ｹ繝医い蜷・    ・奇ｼ・譛ｪ菴ｿ逕ｨ縺ｮ縺溘ａ蟒・ｭ｢莠亥ｮ壹・奇ｼ・	      Copy_Name      string   // 繧ｳ繝斐・蜈・・繝・・繧ｿ繧ｹ繝医い蜷・	      New_Name       string   // 繝九Η繝ｼ繝・・繧ｿ繧ｹ繝医い蜷・
}

///
///    繝・・繧ｿ繝吶・繧ｹ縲繧｢繧ｯ繧ｻ繧ｹ縲繝ｪ繧ｹ繝・///

type  Db_Access_List    struct           {

          Id              int64       //縲繝・・繧ｿid
          Line_No         int64    // 陦君O.
          Db_Type         string      // 繝・・繧ｿ繝吶・繧ｹ繧ｿ繧､繝・                                      //  ds : 繝・・繧ｿ繧ｹ繝医い
                                      //  sr : 繧ｹ繝医Ξ繝・ず

          Access_Type        string   // 繧｢繧ｯ繧ｻ繧ｹ繧ｿ繧､繝・                                      //  copy
                                      //  rename

          Project_Name     string     // 繝励Ο繧ｸ繧ｧ繧ｯ繝亥錐
	      Bucket_Name     string      // 繝舌こ繝・ヨ蜷・	      Basic_File_Name    string   // 蝓ｺ譛ｬ縺ｮ繝輔ぃ繧､繝ｫ蜷・	      New_File_Name      string   // 繝九Η繝ｼ繝輔ぃ繧､繝ｫ蜷・
}

///
///    繝・・繧ｿ繝吶・繧ｹ縲繧｢繧ｯ繧ｻ繧ｹ縲繝ｪ繧ｹ繝・///

type  Db_Access_List2    struct           {

          Id              int64       //縲繝・・繧ｿid
          Line_No         int64    // 陦君O.
          Db_Type         string      // 繝・・繧ｿ繝吶・繧ｹ繧ｿ繧､繝・                                      //  ds : 繝・・繧ｿ繧ｹ繝医い
                                      //  sr : 繧ｹ繝医Ξ繝・ず

          Access_Type        string   // 繧｢繧ｯ繧ｻ繧ｹ繧ｿ繧､繝・                                      //  copy
                                      //  rename

          Project_Name     string     // 繝励Ο繧ｸ繧ｧ繧ｯ繝亥錐
	      Bucket_Name     string      // 繝舌こ繝・ヨ蜷・	      Basic_File_Name    string   // 蝓ｺ譛ｬ縺ｮ繝輔ぃ繧､繝ｫ蜷・	      New_File_Name      string   // 繝九Η繝ｼ繝輔ぃ繧､繝ｫ蜷・
}

///
///           record for csv file
///

type  Csv_Inf    struct           {

          Id            int64
          Line_No       int64    // 陦君O.
          File_Name     string   // 繝輔ぃ繧､繝ｫ蜷・          Column_Num    int64    // 蛻玲焚
	      Column1       string   // 蛻暦ｼ・	      Column2       string   // 蛻暦ｼ・	      Column3       string   // 蛻暦ｼ・	      Column4       string   // 蛻暦ｼ・	      Column5       string   // 蛻暦ｼ・	      Column6       string   // 蛻暦ｼ・	      Column7       string   // 蛻暦ｼ・	      Column8       string   // 蛻暦ｼ・	      Column9       string   // 蛻暦ｼ・	      Column10      string   // 蛻暦ｼ托ｼ・
}

///
///
///

type  Interpret    struct           {

//          Id            int64
//          Line_No       int64    // 陦君O.

	      Ex_All        string   // 險育ｮ怜ｼ就LL
	      Ex_Parts      string   // 險育ｮ怜ｼ襲arts

}
///
///           record for csv file
///

type  Csv_Records    struct           {

      Records_Num    int64   // csv繝ｬ繧ｳ繝ｼ繝峨・讒矩菴薙・謨ｰ

//          Id            int64
//          Line_No       int64       // 陦君O.
      Records[10]    []Csv_Inf   // csv繝ｬ繧ｳ繝ｼ繝峨・讒矩菴・
}

///
///   逕ｻ蜒上ヵ繧｡繧､繝ｫ陦ｨ遉ｺ
///

type  Image_Show    struct           {

          Id              int64    // 繝・・繧ｿid
	      File_Name       string   // 繝輔ぃ繧､繝ｫ蜷・	      Url             string   // url
}
///
///  繧ｹ繝医Ξ繝・ず縺ｮ繝舌こ繝・ヨ繝ｻ繧ｪ繝悶ず繧ｧ繧ｯ繝茨ｼ医ヵ繧｡繧､繝ｫ・峨・陦ｨ遉ｺ逕ｨ
///

type  Storage_B_O_View    struct           {

          Line_No         int64    // 陦君O.
          Project_Name     string   // 繝励Ο繧ｸ繧ｧ繧ｯ繝亥錐
	      Bucket_Name     string   // 繝舌こ繝・ヨ蜷・	      Object_Name     string   // 繧ｪ繝悶ず繧ｧ繧ｯ繝亥錐
	      Created       time.Time  //菴懈・譎る俣

}

///
///  繧ｹ繝医Ξ繝・ず縺ｮ繝舌こ繝・ヨ繝ｻ繧ｪ繝悶ず繧ｧ繧ｯ繝茨ｼ医ヵ繧｡繧､繝ｫ・峨・繧ｳ繝｢繝ｳ逕ｨ(繝・・繧ｿ繧ｹ繝医い・・///

type  Storage_B_O_Temp    struct           {

          Id              int64    // 繝・・繧ｿid
          Line_No         int64    // 陦君O.
          Project_Name     string   // 繝励Ο繧ｸ繧ｧ繧ｯ繝亥錐
	      Bucket_Name     string   // 繝舌こ繝・ヨ蜷・	      Object_Name     string   // 繧ｪ繝悶ず繧ｧ繧ｯ繝亥錐
	      Created       time.Time  //菴懈・譎る俣

}

type  Storage_B_O    struct           {

          Id              int64    // 繝・・繧ｿid
          Line_No         int64    // 陦君O.
          Project_Name     string   // 繝励Ο繧ｸ繧ｧ繧ｯ繝亥錐
	      Bucket_Name     string   // 繝舌こ繝・ヨ蜷・	      Object_Name     string   // 繧ｪ繝悶ず繧ｧ繧ｯ繝亥錐
	      Created       time.Time  //菴懈・譎る俣

}
///
///  繝ｯ繝ｼ繧ｯ逕ｨ
///

type  General_Work    struct           {

          Int64_Work     int64           // int蝙九Ρ繝ｼ繧ｯ繧ｨ繝ｪ繧｢
          Float64_Work   float64         // float蝙九Ρ繝ｼ繧ｯ繧ｨ繝ｪ繧｢
	      String_Work    string          // string蝙九Ρ繝ｼ繧ｯ繧ｨ繝ｪ繧｢
	      Sw_Work        *storage.Writer // 繧ｹ繝医Ξ繝・ず繝ｩ繧､繧ｿ繝ｼ蝙・
}
