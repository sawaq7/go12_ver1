package type4

///
///    豌ｴ霍ｯ繝輔ぃ繧､繝ｫ・・繧ｹ繝医Ξ繝・ず繝ｻ繝輔Μ繝ｼ・・///

type  Struct_Colle    struct  {

	      Water2_Slice   []Water2
	   Water_Line_Slice  []Water_Line

}

///
///    豌ｴ霍ｯ繝輔ぃ繧､繝ｫ・・繧ｹ繝医Ξ繝・ず繝ｻ繝輔Μ繝ｼ・・///

type  Water    struct           {
	      No             string   // 繝・・繧ｿID
	      Name           string   // 豌ｴ霍ｯ蜷・	      High           string   // 豌ｴ霍ｯ鬮・	    Roughness_factor string   // 邊礼ｲ剃ｿよ焚
}

///
///    豌ｴ霍ｯ繝輔ぃ繧､繝ｫ・・///

type  Water2    struct           {
          Id             int64      //縲繝・・繧ｿid  * 繧ｹ繝医Ξ繝・ず縺ｮ蝣ｴ蜷医・繝ｬ繧ｳ繝ｼ繝丑O
	      Name           string     // 豌ｴ霍ｯ蜷・	      High           float64    // 豌ｴ霍ｯ鬮・	    Roughness_Factor float64    // 邊礼ｲ剃ｿよ焚
}
///
///    豌ｴ霍ｯ繝輔ぃ繧､繝ｫ・・繝・・繧ｿ繧ｹ繝医い繝ｻ繝・Φ繝昴Λ繝ｪ繝ｼ・・///

type  Water2_Temp    struct           {       //  繝・・繧ｿ縺ｯ1繝ｬ繧ｳ繝ｼ繝・          Id             int64      //縲繝・・繧ｿid
	      Name           string     // 豌ｴ霍ｯ蜷・	      High           float64    // 豌ｴ霍ｯ鬮・	    Roughness_Factor float64    // 邊礼ｲ剃ｿよ焚
}
///
///    豌ｴ霍ｯ繝ｩ繧､繝ｳ繝輔ぃ繧､繝ｫ(繝・・繧ｿ繧ｹ繝医い・・///

type  Water_Line    struct           {
          Id              int64    // 繝・・繧ｿid縲縲* 繧ｹ繝医Ξ繝・ず縺ｮ蝣ｴ蜷医・縲∵ｰｴ霍ｯ蜊倅ｽ阪・繝ｬ繧ｳ繝ｼ繝丑O
	      Name            string   // 豌ｴ霍ｯ蜷・	      Section         string   // 蛹ｺ髢灘錐
	      Friction_Factor float64  // 鞫ｩ謫ｦ菫よ焚
	      Velocity        float64  // 騾溷ｺｦ
	      Pipe_Diameter   float64  // 邂｡蠕・	      Pipe_Length     float64  // 邂｡髟ｷ
}

///
///    蟆取ｰｴ蜍ｾ驟咲ｷ壹ヵ繧｡繧､繝ｫ(繧ｰ繝ｩ繝包ｼ・///
///  *** 逕ｻ蜒上ヵ繧｡繧､繝ｫ陦ｨ遉ｺ逕ｨ縲struct "Image_Show" 縺ｨ縲蜷後ヵ繧ｩ繝ｼ繝槭ャ繝・
type  Water_Slope    struct           {

          Id              int64    // 繝・・繧ｿid
	      File_Name       string   // 繝輔ぃ繧､繝ｫ蜷・	      Url             string   // url

}
///
///    hydro-static information
///

type Seisui struct {

      Omega string
      D1    string
      D2    string
      P     string
      H     string
     P2    float64

}
