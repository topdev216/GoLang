// http://lbsyun.baidu.com/index.php?title=lbscloud/api/geodata
/*
 * 百度地图全局配置信息
 */
var map = new BMap.Map("allmap");          // 创建地图实例
map.centerAndZoom(new BMap.Point(116.403694, 39.947552), 16);  // 初始化地图，设置中心点坐标和地图级别
map.enableScrollWheelZoom(true);
// 添加全景控件
map.addControl(new BMap.PanoramaControl());
// 左上角，添加比例尺
map.addControl(new BMap.ScaleControl({anchor: BMAP_ANCHOR_TOP_LEFT}));
// 左上角，添加默认缩放平移控件
map.addControl(new BMap.NavigationControl());








/*
 * 标注工具配置
 */
var markText = [];
markText.push('<div>属性信息</div>');
markText.push('<div><span>名称</span><input type="text" name="title" id="title" /></div>');
markText.push('<div><span>地址</span><input type="text" name="address" id="address" /></div>');
markText.push('<div><span>中介</span><input type="text" name="hagent" id="hagent" /></div>');
markText.push('<div><input type="hidden" name="longitude" id="longitude" /></div>');
markText.push('<div><input type="hidden" name="latitude" id="latitude" /></div>');
markText.push('<div><span>看房</span><textarea name="hvisits" id="hvisits"></textarea></div>');
markText.push('<div><span>备注</span><textarea name="hremark" id="hremark"></textarea></div>');
markText.push('<div><button type="button" onclick="postForm();">保存</button><button type="button" onclick="resetForm()">重置</button><span id="notify"></span></div>');
var markWindow = new BMap.InfoWindow(markText.join(''), {offset: new BMap.Size(0, -10)});
var curMarker = null;  // 记录当前添加的marker

// 标注工具
var markTool = new BMapLib.MarkerTool(map, {autoClose: true});
markTool.addEventListener('markend', function(e) {
	let mkr = e.marker;
	mkr.openInfoWindow(markWindow);
	curMarker = mkr;
});

// 添加右键菜单：标注工具
var menu = new BMap.ContextMenu();
menu.addItem(new BMap.MenuItem('添加标注', function(){
	markTool.open();
}, 100));
map.addContextMenu(menu);

// 去除表单HTML标签
function filtHtml(a) {
	return a.replace(/&/g, "&amp;").replace(/</g, "&lt;").replace(/>/g, "&gt;").replace(/"/g, "&quot;");
}

// 清空表单
function resetForm() {
	document.getElementById('title').value = '';
	document.getElementById('address').value = '';
	document.getElementById('hagent').value = '';
	document.getElementById('longitude').value = '';
	document.getElementById('latitude').value = '';
	document.getElementById('hvisits').value = '';
	document.getElementById('hremark').value = '';
	document.getElementById('notify').innerHTML = '';
}

// 提交表单数据
function postForm() {
	let title = filtHtml(document.getElementById('title').value);
	let address = filtHtml(document.getElementById('address').value);
	let hagent = filtHtml(document.getElementById('hagent').value);
	let longitude = filtHtml(document.getElementById('longitude').value);
	let latitude = filtHtml(document.getElementById('latitude').value);
	let hvisits = filtHtml(document.getElementById('hvisits').value);
	let hremark = filtHtml(document.getElementById('hremark').value);

	if (!title || !address || !hagent || !hvisits || !hremark) {
		document.getElementById('notify').innerHTML = '表单没填完';
		return;
	}

	if (curMarker) {
		let lbl = new BMap.Label(title, {offset: new BMap.Size(1, 1)});
		lbl.setStyle({border: 'solid 1px gray', color: 'red'});
		curMarker.setLabel(lbl);
		curMarker.setTitle(title);
	}

	if (markWindow.isOpen()) {
		resetForm();
		map.closeInfoWindow();
	}
}




/*
 * 添加图层
 */
var customLayer;
function addCustomLayer(keyword) {
	if (customLayer) {
		map.removeTileLayer(customLayer);
	}
	customLayer=new BMap.CustomLayer({
		geotableId: 146844,
		q: '', //检索关键字
		tags: '', //空格分隔的多字符串
		filter: '' //过滤条件,参考http://developer.baidu.com/map/lbs-geosearch.htm#.search.nearby
	});
	map.addTileLayer(customLayer);
	customLayer.addEventListener('hotspotclick',callback);
}

function callback(e)//单击热点图层
{
		var customPoi = e.customPoi;//poi的默认字段
		var contentPoi = e.content;//poi的自定义字段
		var content = '<p style="width:280px;margin:0;line-height:20px;">地址：' + customPoi.address + '</p>';
		var searchInfoWindow = new BMapLib.SearchInfoWindow(map, content, {
			title: customPoi.title, //标题
			width: 300, //宽度
			height: 60, //高度
			panel : "panel", //检索结果面板
			enableAutoPan : true, //自动平移
			enableSendToPhone: true, //是否显示发送到手机按钮
			searchTypes :[
				BMAPLIB_TAB_SEARCH,   //周边检索
				BMAPLIB_TAB_TO_HERE  //到这里去
			]
		});
		var point = new BMap.Point(customPoi.point.lng, customPoi.point.lat);
		searchInfoWindow.open(point);
}
document.getElementById("open").onclick = function(){
	addCustomLayer();
};
document.getElementById("open").click();
document.getElementById("close").onclick = function(){
	if (customLayer) {
		map.removeTileLayer(customLayer);
	}
};








/*
 * 显示行政区划
 */
/*
var bdary = new BMap.Boundary();
bdary.get("北京市东城区", function(rs){       //获取行政区域
	map.clearOverlays();        //清除地图覆盖物
	var count = rs.boundaries.length; //行政区域的点有多少个
	if (count === 0) {
		return;
	}
  	var pointArray = [];
	for (var i = 0; i < count; i++) {
		var ply = new BMap.Polygon(rs.boundaries[i], {strokeWeight: 1, strokeColor: "#ff0000"}); //建立多边形覆盖物
		map.addOverlay(ply);  //添加覆盖物
		pointArray = pointArray.concat(ply.getPath());
	}
	map.setViewport(pointArray);    //调整视野
});
*/




