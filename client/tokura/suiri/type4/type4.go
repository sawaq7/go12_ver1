package type4

///
///     water data
///

type  Struct_Colle    struct  {

	      Water2_Slice   []Water2
	   Water_Line_Slice  []Water_Line

}

///
///    water
///

type  Water    struct           {

	      No             string   //
	      Name           string   // water name
	      High           string   // water high
	      Roughness_factor string
}

///
///    water
///

type  Water2    struct           {

          Id             int64
	      Name           string     // water name
	      High           float64    // water high
	      Roughness_Factor float64
}
///
///    water temp.-file

type  Water2_Temp    struct           {

          Id             int64
	      Name           string     // water name
	      High           float64    // water high
	      Roughness_Factor float64
}
///
///    water line
///

type  Water_Line    struct           {
          Id              int64
	      Name            string   // water name
	      Section         string   // water section
	      Friction_Factor float64
	      Velocity        float64
	      Pipe_Diameter   float64
	      Pipe_Length     float64
}

///
///    water slope line
///   the format of this  and　struct "Image_Show" is same

type  Water_Slope    struct           {

          Id              int64
	      File_Name       string
	      Url             string

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
