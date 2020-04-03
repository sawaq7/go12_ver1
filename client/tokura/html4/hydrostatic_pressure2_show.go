package html4

const Hydrostatic_pressure2_show = `

<!DOCTYPE html>

<html>

   <head>
      <meta charset="UTF-8">
      <title>髱呎ｰｴ蝨ｧ culculate</title>
      <link rel="stylesheet" href="css/tokura.css" type="text/css">
   </head>

<body>
  <section>
  <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
  <p>&nbsp;&nbsp;&nbsp; etc. input information                  </p>

  <p>&nbsp;&nbsp;&nbsp; 蜊倅ｽ榊ｮｹ遨埼㍾驥・ﾏ・  &nbsp;&nbsp;&nbsp;  {{.Omega|html}} ・・/緕｡・・/p>
  <p>&nbsp;&nbsp;&nbsp; U蟄礼ｮ｡・代・逶ｴ蠕・D1  &nbsp;&nbsp;&nbsp;  {{.D1|html}} ・茨ｽ搾ｼ・ </p>
  <p>&nbsp;&nbsp;&nbsp; U蟄礼ｮ｡ 2縺ｮ逶ｴ蠕・D2  &nbsp;&nbsp;&nbsp;  {{.D2|html}}  ・茨ｽ搾ｼ・ </p>
  <p>&nbsp;&nbsp;&nbsp; 闕ｷ驥・P           &nbsp;&nbsp;&nbsp;   {{.P|html}} ・・・・   </p>
  <p>&nbsp;&nbsp;&nbsp; 鬮伜ｺｦ蟾ｮ縺ｯ H       &nbsp;&nbsp;&nbsp;   {{.H|html}}  ・茨ｽ搾ｼ・  </p>

  <p>&nbsp;&nbsp;&nbsp; the result of calculate                 </p>

  <p>&nbsp;&nbsp;&nbsp; 蠢・ｦ√↑闕ｷ驥・P2    &nbsp;&nbsp;&nbsp;  {{.P2|html}} ・・・・  </p>
  </section>
</body>

</html>
`
