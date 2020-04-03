package html6

   const Medical_xray_show = `
   <!DOCTYPE html>
   <html>
     <head>
        <meta charset="UTF-8">
        <title>チE�Eタストアの更新・削除</title>
        <link rel="stylesheet" href="css/sgh/d_district_area.css" type="text/css">
        <style type="text/css"> /* マウスポインタの設定！Entense用�E�E/
            .intense {
            cursor: url("./plus_cursor.png"), pointer; /* マウスポインタを指宁E*/
            display: inline-block;   /* 横方向に並べる指宁E*/
            margin: 0px 5px 5px 0px; /* 周囲の余白釁E右と下に5pxずつ) */
            }
        </style>
		<script src="js/main.js"></script>
		<script src="js/intense.js"> /* intense用 */ </script>
     </head>
     <body>
       <header>



       </header>
       <nav>


	   </nav>
       <article>
         <table border="2" cellpadding="12" align="center" bgcolor="#00ced1">
         <h2 align="center">Current Guest</h2>

           <tr> <th>guest-no</th> <th>guest-name</th> </tr>
           {{range .}}

            {{ if eq .Line_No 1}}

             <tr>

             <form method="GET" action="/d_district_area_update" >


               <td>
                  <input type="text" name="district_no" size="5" align="center" value="{{.Guest_No|html}}" />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>

               </td>
               <td>
                  <input type="text" name="district_name" size="10" align="center" value="{{.Guest_Name|html}}" />
               </td>
              </form>

            </tr>

            {{end}}

           {{end}}
           </table>


       </article>

       <section id="main">

         <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">

         <h1  align="center">Select image file which you want to upload ,please</h1>

         <tr>  <th>image-file</th>  <th>access1</th> </tr>

          <form method="POST" enctype="multipart/form-data" action="/medical_xray_upload" >

            <td >
              <input type="file" name="image" id="image" size="10" align="center">
            </td>

            <td >
              <input type="submit" size="2"  value="upload" />
            </td>

          </form>


        </table>

         <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
           <h2 align="center">Medical_Record Information</h2>

           <tr> <th>ope-date</th> <th>text</th> <th>access1</th>  </tr>
           {{range .}}

             <tr>

             <form method="GET" action="/medical_xray_delete" >

               <td>
                   <input type="text" name="date" size="12" value="{{.Date |html}}" />
              </td>

               <td>
                  <img src="{{.Url}}" width="50" height="50"
			           data-title="some animals" data-caption="parakeet1" class="intense" />

               </td>

             </form>

             <form method="GET" action="/medical_xray_delete" >

               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="delete"  />
               </td>
             </form>

           </tr>

           {{end}}
          </table>

        </section>


        <aside>

        </aside>

        <footer>

        </footer>

        <script> /* intense用 */
		   window.onload = function() {
           var elements = document.querySelectorAll( '.intense' );
           Intense( elements );
           }
        </script>

     </body>
   </html>
   `
