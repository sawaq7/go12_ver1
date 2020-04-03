package html6

   const Appri_start = `
   <!DOCTYPE html>
   <html>
	<head>
		<meta charset="utf-8">
		<title>sawa Q</title>
		<link rel="stylesheet" href="css/index.css" type="text/css">
		<style type="text/css"> /* 繝槭え繧ｹ繝昴う繝ｳ繧ｿ縺ｮ險ｭ螳夲ｼ・ntense逕ｨ・・/
            .intense {
            cursor: url("./plus_cursor.png"), pointer; /* 繝槭え繧ｹ繝昴う繝ｳ繧ｿ繧呈欠螳・*/
            display: inline-block;   /* 讓ｪ譁ｹ蜷代↓荳ｦ縺ｹ繧区欠螳・*/
            margin: 0px 5px 5px 0px; /* 蜻ｨ蝗ｲ縺ｮ菴咏區驥・蜿ｳ縺ｨ荳九↓5px縺壹▽) */
            }
        </style>
		<script src="js/main.js"></script>
		<script src="js/intense.js"> /* intense逕ｨ */ </script>

	</head>
	<body>


	    <header>
			<div><img src="images/logo.png" width="320" height="32" alt="繝ｭ繧ｴ"></div>
		</header>
		<nav>
			<ul>
                <li><a href="calculate.html">calculate</a></li>
				<li><a href="https://sample-7777.appspot.com/reserve_index">Medical_System</a></li>
				<li><a href="https://sample-7777.appspot.com/sgh_index">Sgh Management System</a></li>
				<li><a href="https://sample-7777.appspot.com/tokura_index">Tokura Scientific Calculation System</a></li>
				<li><a href="https://sample-7777.appspot.com/general_index">General Soft</a></li>
				<li><a href="contact.html">縺雁撫縺・粋繧上○</a></li>
				<li><a href="nakamura.html">members 荳ｭ譚・/a></li>

			</ul>
		</nav>
		<article>
			<h1>髮ｻ蜉帑ｺ区ュ縺ｫ縺､縺・※</h1>
			<p>
				迴ｾ蝨ｨ縲・崕蜉帑ｾ帷ｵｦ縺ｫ荳榊ｮ峨′縺ゅｋ縺溘ａ縲∝推莠区･ｭ驛ｨ縺ｯ蝗ｽ蜀・・
				繝・・繧ｿ繧ｻ繝ｳ繧ｿ繝ｼ縺縺代〒縺ｪ縺乗ｵｷ螟悶・繝・・繧ｿ繧ｻ繝ｳ繧ｿ繝ｼ縺ｸ縺ｮ
				遘ｻ陦後ｂ隕夜㍽縺ｫ蜈･繧後ｋ蠢・ｦ√′縺ｧ縺ｦ縺阪∪縺励◆縲・			</p>
			<p>
				莉雁ｾ後・繝ｪ繧ｹ繧ｯ蛻・淵縺ｫ縺､縺・※譌ｩ諤･縺ｫ蟇ｾ蜃ｦ縺励↑縺代ｌ縺ｰ縺・￠縺ｾ縺帙ｓ縲・			</p>
		</article>
		<section id="main">
			<h1>譁ｰ蝠・刀縺ｫ髢｢縺吶ｋ縺顔衍繧峨○</h1>
			<section>
				<h2>sawaQ繧ｳ繝ｳ繝舌・繧ｿ繝ｼ ver 2</h2>
				<p>霑第律荳ｭ縺ｫ繧｢繝翫え繝ｳ繧ｹ縺吶ｋ莠亥ｮ壹〒縺吶・/p>
			</section>
			<section>
				<h2>sawaQ API霎樊嶌</h2>
				<p>螂ｽ隧慕匱螢ｲ荳ｭ縲ら樟蝨ｨ縲√く繝｣繝ｳ繝壹・繝ｳ荳ｭ縺ｫ縺､縺榊濠鬘阪そ繝ｼ繝ｫ螳滓命荳ｭ!</p>
			</section>
		</section>
		<aside>
			<div><img src="images/banner.png" width="200" height="200" alt="蠎・相"></div>
			<div><a href="member.html" >
			        <img src="photo/bird.jpg"  width="90" height="100" alt="蠎・相">
			     </a>
			     <img src="photo/bird3.jpg"   width="90" height="100" alt="蠎・相">
			     <img src="photo/bird.jpg" width="90" height="100" data-title="some animals"
                                                                  data-caption="parakeet" class="intense">
                 <img src="images/illust_finger.png" width="90" height="100" data-title="some animals"
                                                                  data-caption="parakeet" class="intense">
			</div>
		</aside>
		<footer>
			<p><small>2011 &copy; (譬ｪ)sawaQ</small></p>
			<p>縲・00-0014 譚ｱ莠ｬ驛ｽ蜊・ｻ｣逕ｰ蛹ｺ豌ｸ逕ｰ逕ｺ1-7-1</p>
		</footer>
		<script> /* intense逕ｨ */
		   window.onload = function() {
           var elements = document.querySelectorAll( '.intense' );
           Intense( elements );
           }
        </script>

	</body>
  </html>

`

