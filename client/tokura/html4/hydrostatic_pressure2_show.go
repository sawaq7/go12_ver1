package html4

const Hydrostatic_pressure2_show = `

<!DOCTYPE html>

<html>

   <head>
      <meta charset="UTF-8">
      <title>静水圧 culculate</title>
      <link rel="stylesheet" href="css/tokura.css" type="text/css">
   </head>

<body>
  <section>
  <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
  <p>&nbsp;&nbsp;&nbsp; etc. input information                  </p>

  <p>&nbsp;&nbsp;&nbsp; 単位容積重釁EρE  &nbsp;&nbsp;&nbsp;  {{.Omega|html}} �E�E/㎡�E�E/p>
  <p>&nbsp;&nbsp;&nbsp; U字管�E��E直征ED1  &nbsp;&nbsp;&nbsp;  {{.D1|html}} �E�ｍ！E </p>
  <p>&nbsp;&nbsp;&nbsp; U字管 2の直征ED2  &nbsp;&nbsp;&nbsp;  {{.D2|html}}  �E�ｍ！E </p>
  <p>&nbsp;&nbsp;&nbsp; 荷釁EP           &nbsp;&nbsp;&nbsp;   {{.P|html}} �E�E�E�E   </p>
  <p>&nbsp;&nbsp;&nbsp; 高度差は H       &nbsp;&nbsp;&nbsp;   {{.H|html}}  �E�ｍ！E  </p>

  <p>&nbsp;&nbsp;&nbsp; the result of calculate                 </p>

  <p>&nbsp;&nbsp;&nbsp; 忁E��な荷釁EP2    &nbsp;&nbsp;&nbsp;  {{.P2|html}} �E�E�E�E  </p>
  </section>
</body>

</html>
`
