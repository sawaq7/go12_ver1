package html6

   const Appri_start = `
   <!DOCTYPE html>
   <html>
	<head>
		<meta charset="utf-8">
		<title>sawa Q</title>
		<link rel="stylesheet" href="css/index.css" type="text/css">
		<style type="text/css"> /* マウスポインタの設定！Entense用 */
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
			<div><img src="images/logo.png" width="320" height="32" alt="ロゴ"></div>
		</header>
		<nav>
			<ul>
                <li><a href="calculate.html">calculate</a></li>
				<li><a href="https://sample-7777.appspot.com/reserve_index">Medical_System</a></li>
				<li><a href="https://sample-7777.appspot.com/sgh_index">Sgh Management System</a></li>
				<li><a href="https://sample-7777.appspot.com/tokura_index">Tokura Scientific Calculation System</a></li>
				<li><a href="https://sample-7777.appspot.com/general_index">General Soft</a></li>
				<li><a href="contact.html">question</a></li>
				<li><a href="nakamura.html">members nakamura/a></li>

			</ul>
		</nav>
		<article>
			<h1>電力事情につ</h1>
			<p>
				現在、力供給に不安があるため、各事業部は国冁E�E
				チEタセンターだけでなく海外タセンターへの
				移行も視野に入れるがでてきました、E			</p>
			<p>
				今後リスク刁つ早急に対処しなければぁません、E			</p>
		</article>
		<section id="main">
			<h1>新啁に関するお知らせ</h1>
			<section>
				<h2>sawaQコンバター ver 2</h2>
				<p>近日中にアナウンスする予定です、E/p>
			</section>
			<section>
				<h2>sawaQ API辞書</h2>
				<p>好評発売中。現在、キャンペ中につき半額セール実施中!</p>
			</section>
		</section>
		<aside>
			<div><img src="images/banner.png" width="200" height="200" alt="庁E��"></div>
			<div><a href="member.html" >
			        <img src="photo/bird.jpg"  width="90" height="100" alt="庁E��">
			     </a>
			     <img src="photo/bird3.jpg"   width="90" height="100" alt="庁E��">
			     <img src="photo/bird.jpg" width="90" height="100" data-title="some animals"
                                                                  data-caption="parakeet" class="intense">
                 <img src="images/illust_finger.png" width="90" height="100" data-title="some animals"
                                                                  data-caption="parakeet" class="intense">
			</div>
		</aside>
		<footer>
			<p><small>2011 &copy; (株)sawaQ</small></p>
			<p>、E00-0014 東京都十E��田区永田町1-7-1</p>
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

