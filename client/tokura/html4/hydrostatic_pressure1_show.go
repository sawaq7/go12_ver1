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

      <p>&nbsp;&nbsp;&nbsp; key-in various data of hydrostatic-pressure  </p>
      <p>&nbsp;&nbsp;&nbsp;  bulk density ω！E/㎡ &nbsp;&nbsp;  type="text" name="omega" /> </p>
　    <p>&nbsp;&nbsp;  Deameter of U-tube1 D1（ｍ）&nbsp;&nbsp; ：<input type="text" name="d1" /> </p>
　    <p>&nbsp;&nbsp;  Deameter of U-tube2 D2（ｍ）&nbsp;&nbsp; ：<input type="text" name="d2" /> </p>
　    <p>&nbsp;&nbsp;  Weight of the load P（t）&nbsp;&nbsp;&nbsp;&nbsp;
                   &nbsp;&nbsp;&nbsp;&nbsp;
                   &nbsp;&nbsp;&nbsp;&nbsp;      ：<input type="text" name="p" /> </p>

      <p>&nbsp;&nbsp;  height difference H（ｍ） &nbsp;&nbsp;&nbsp;
                   &nbsp;&nbsp;&nbsp;&nbsp;
                   &nbsp;&nbsp;&nbsp;&nbsp; ：<input type="text" name="h" /> </p>


      <input type="submit" value="h.p.（key-in)" />
     </form>
     </section>
  </body>
</html>
`
