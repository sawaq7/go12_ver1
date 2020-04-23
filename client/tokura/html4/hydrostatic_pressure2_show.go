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

  <p>&nbsp;&nbsp;&nbsp; bulk density ω   &nbsp;&nbsp;&nbsp;  {{.Omega|html}} （t/㎡）</p>
  <p>&nbsp;&nbsp;&nbsp; Deameter of U-tube1 D1  &nbsp;&nbsp;&nbsp;  {{.D1|html}} （ｍ）  </p>
  <p>&nbsp;&nbsp;&nbsp; Deameter of U-tube2 D2  &nbsp;&nbsp;&nbsp;  {{.D2|html}}  （ｍ）  </p>
  <p>&nbsp;&nbsp;&nbsp; Weight of the load P    &nbsp;&nbsp;&nbsp;   {{.P|html}} （t）    </p>
  <p>&nbsp;&nbsp;&nbsp; height difference H     &nbsp;&nbsp;&nbsp;   {{.H|html}}  （ｍ）   </p>

  <p>&nbsp;&nbsp;&nbsp; the result of calculate                 </p>

  <p>&nbsp;&nbsp;&nbsp; the load which is required is P2    &nbsp;&nbsp;&nbsp;  {{.P2|html}} （t）   </p>
  </section>
</body>

</html>
`
