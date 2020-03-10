package html4

const Hydrostatic_pressure1_show = `
<!DOCTYPE html>
<html>
   <head>
      <meta charset="utf-8">
	  <title>(株)HTML5</title>
	  <link rel="stylesheet" href="css/tokura.css" type="text/css">
   </head>
   <body>
     <section>
       <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">

     <form method="GET" action="/hydrostatic_pressure1_excute">

      <p>&nbsp;&nbsp;&nbsp; 静水圧の各種データを入力してください。   </p>
      <p>&nbsp;&nbsp;&nbsp;  単位容積重量 ω（t/㎡）&nbsp;&nbsp;&nbsp; ：<input type="text" name="omega" /> </p>
　    <p>&nbsp;&nbsp;  U字管１の直径 D1（ｍ）&nbsp;&nbsp;  ：<input type="text" name="d1" /> </p>
　    <p>&nbsp;&nbsp;  U字管２の直径 D2（ｍ) &nbsp;&nbsp;  ：<input type="text" name="d2" /> </p>
　    <p>&nbsp;&nbsp;  荷重の重さ P（t）&nbsp;&nbsp;&nbsp;&nbsp;
                   &nbsp;&nbsp;&nbsp;&nbsp;
                   &nbsp;&nbsp;&nbsp;&nbsp;      ：<input type="text" name="p" /> </p>
　    <p>&nbsp;&nbsp;  高度差は H（ｍ） &nbsp;&nbsp;&nbsp;
                   &nbsp;&nbsp;&nbsp;&nbsp;
                   &nbsp;&nbsp;&nbsp;&nbsp; ：<input type="text" name="h" /> </p>

      <input type="submit" value="静水圧（key-in)" />
     </form>
     </section>
  </body>
</html>
`
