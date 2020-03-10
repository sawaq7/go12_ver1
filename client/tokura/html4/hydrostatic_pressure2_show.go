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

  <p>&nbsp;&nbsp;&nbsp; 単位容積重量 ω   &nbsp;&nbsp;&nbsp;  {{.Omega|html}} （t/㎡）</p>
  <p>&nbsp;&nbsp;&nbsp; U字管１の直径 D1  &nbsp;&nbsp;&nbsp;  {{.D1|html}} （ｍ）  </p>
  <p>&nbsp;&nbsp;&nbsp; U字管 2の直径 D2  &nbsp;&nbsp;&nbsp;  {{.D2|html}}  （ｍ）  </p>
  <p>&nbsp;&nbsp;&nbsp; 荷重 P           &nbsp;&nbsp;&nbsp;   {{.P|html}} （t）    </p>
  <p>&nbsp;&nbsp;&nbsp; 高度差は H       &nbsp;&nbsp;&nbsp;   {{.H|html}}  （ｍ）   </p>

  <p>&nbsp;&nbsp;&nbsp; the result of calculate                 </p>

  <p>&nbsp;&nbsp;&nbsp; 必要な荷重 P2    &nbsp;&nbsp;&nbsp;  {{.P2|html}} （t）   </p>
  </section>
</body>

</html>
`