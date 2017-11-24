<head>
	<style>
	
		* {
			font-size: 22px;
			font-family: 'Helvetica Neue';
			font-weight: 300;
			-webkit-font-smoothing: antialiased;
			-moz-osx-font-smoothing: grayscale;
			white-space: nowrap;
			color: #fff;
		}
		@font-face {
			font-family: 'Helvetica Neue';
			src: url('HelveticaNeue.eot');
			src: url('HelveticaNeue.eot?#iefix') format('embedded-opentype'),
				url('HelveticaNeue.woff') format('woff'),
				url('HelveticaNeue.ttf') format('truetype');
			font-weight: normal;
			font-style: normal;
		}

		body {
			margin: 0;
			width: 100vw;
			height: 100vh;
			background: black;	
		}
		#screen{
			margin: 200px;
			background-image: url(http://iphonewalls.net/wp-content/uploads/2014/09/Snow%20Mountain%20Chalet%20Aurora%20Milky%20Way%20Stars%20iPhone%206%20Wallpaper.jpg);
			position: relative;
			width: 750px;
			height: 1334px;
			border-radius: 10px;
		}
		#dockrow {
			display: block;
			position: absolute;
			bottom: 4px;
			left: 54px;
		}
		.app {
			margin-right: 50px;
		}
		.app img{
			width: 120px;
			height: 120px;
			border-radius: 30px;
		}
		.appspan {
			color: #fff;
			margin-right: 50px;
		}
		#dock {
			height: 192px;
			width: 750px;
			position: absolute;
			bottom: 0;
			left: 0;
			background-color: rgba(81, 142, 228, 0.96);
			white-space: nowrap;
		}
		.rowpadding {
			height: 5px;
			display: block;
		}
	</style>
</head>
<body>
	<div id="screen">
		<div id="statusbar">
		
		<center><span><?php echo date("g:i  A", time()); ?></span></center>
		</div>
	
		<div id="dock">
			<table id="dockrow">
			<tr>
				<td><div class="app">
					<img src="dialer.png"></img>
				</div></td>
				<td><div class="app">
					<img src="safari.png"></img>
				</div></td>
				<td><div class="app">
					<img src="music.png"></img>
				</div></td>
				<td><div class="app">
					<img src="camera.png"></img>
				</div></td>
			</tr>
			<tr class="rowpadding"><td></td></tr>
			<tr>
				<td><center><span class="appspan">Phone</span></center></td>
				<td><center><span class="appspan">Safari</span></center></td>
				<td><center><span class="appspan">Apple Music</span></center></td>
				<td><center><span class="appspan">Camera</span></center></td>
			</tr>
			</table>
		</div>
	</div>
</body>