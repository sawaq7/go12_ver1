package html5


   const Image_file_show = `
   <!DOCTYPE html>
   <html>
     <head>
        <meta charset="UTF-8">
        <title>show image file</title>
        <link rel="stylesheet" href="css/member1_show.css" type="text/css">
     </head>
     <body>
       <article>
       </article>

       <section id="main">


           <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
           <h2 align="center">List Of Deliverly Situation</h2>


            <tr>
             <form  >
               <td>

                 <!-- <img src="{{.Url}}" width="400" height="400"      /> -->

                 <iframe src="{{.Url}}" width="400" height="400"  />



               </td>
             </form>


           </tr>

           </table>

        </section>


     </body>
   </html>
   `

