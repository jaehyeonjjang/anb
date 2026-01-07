var MODE_NONE = 0;
var MODE_DRAW = 1;
var MODE_SELECT = 2;
var MODE_MOVE = 3;
var MODE_GROUP = 21;
var MODE_PAN = 4;
var MODE_ZOOM = 22;

var TYPE_CURVE = 1;
var TYPE_STRAIGHT = 2;
var TYPE_CURVE2 = 15;
var TYPE_STRAIGHT2 = 16;
var TYPE_CURVE3 = 51;
var TYPE_STRAIGHT3 = 52;
var TYPE_CURVE4 = 53;
var TYPE_STRAIGHT4 = 54;

var TYPE_ARROW = 17;
var TYPE_ARROWV = 18;
var TYPE_ARROWFREE = 19;
var TYPE_ARROWCHAIN = 22;
var TYPE_NUMBER = 20;
var TYPE_NUMBER2 = 21;
var TYPE_MEMO = 30;
var TYPE_LENGTH = 40;
var TYPE_AREA = 41;
var TYPE_LENGTH2 = 42;


var TYPE_LINE1 = 500;
var TYPE_LINE2 = 501;
var TYPE_LINE3 = 502;
var TYPE_LINE4 = 503;

var TYPE_CIRCLE1 = 510;
var TYPE_CIRCLE2 = 511;
var TYPE_CIRCLE3 = 512;

var TYPE_TRIANGLE1 = 520;
var TYPE_TRIANGLE2 = 521;
var TYPE_TRIANGLE3 = 522;

var TYPE_BOX1 = 530;
var TYPE_BOX2 = 531;
var TYPE_BOX3 = 532;

var TYPE_CRACK1 = 540;
var TYPE_CRACK2 = 541;
var TYPE_CRACK3 = 542;
var TYPE_CRACK4 = 543;

var TYPE_POINT1 = 550;
var TYPE_POINT4 = 551;

var TYPE_CROSS1 = 560;

var TYPE_DIAMOND1 = 570;
var TYPE_DIAMOND2 = 571;
var TYPE_DIAMOND3 = 572;

var TYPE_CIRCLETRIANGLE1 = 580;
var TYPE_CIRCLETRIANGLE2 = 581;
var TYPE_CIRCLETRIANGLE3 = 582;

var TYPE_FILLCIRCLE1 = 590;
var TYPE_FILLCIRCLE2 = 591;
var TYPE_FILLCIRCLE3 = 592;


var _colorMap = {};
_colorMap[TYPE_LINE1] = {type: 'line', color: '#0000FF', size: 4.0};
_colorMap[TYPE_LINE2] = {type: 'line', color: '#FF0000', size: 4.0};
_colorMap[TYPE_LINE3] = {type: 'line', color: '#000000', size: 4.0};
_colorMap[TYPE_LINE4] = {type: 'line', color: '#2596be', size: 4.0};
_colorMap[TYPE_CIRCLE1] = {type: 'circle', color: '#0000FF', size: 4.0};
_colorMap[TYPE_CIRCLE2] = {type: 'circle', color: '#0000FF', size: 4.0};
_colorMap[TYPE_CIRCLE3] = {type: 'circle', color: '#0000FF', size: 2.0};
_colorMap[TYPE_TRIANGLE1] = {type: 'triangle', color: '#0000FF', size: 4.0};
_colorMap[TYPE_TRIANGLE2] = {type: 'triangle', color: '#0000FF', size: 4.0};
_colorMap[TYPE_TRIANGLE3] = {type: 'triangle', color: '#0000FF', size: 2.0};
_colorMap[TYPE_BOX1] = {type: 'box', color: '#FF0000', size: 4.0};
_colorMap[TYPE_BOX2] = {type: 'box', color: '#FF0000', size: 4.0};
_colorMap[TYPE_BOX3] = {type: 'box', color: '#FF0000', size: 2.0};
_colorMap[TYPE_CRACK1] = {type: 'crack', color: '#0000FF', size: 4.0, direction: 0};
_colorMap[TYPE_CRACK2] = {type: 'crack', color: '#0000FF', size: 4.0, direction: 1};
_colorMap[TYPE_CRACK3] = {type: 'crack', color: '#0000FF', size: 4.0, direction: 2};
_colorMap[TYPE_CRACK4] = {type: 'crack', color: '#0000FF', size: 4.0, direction: 3};
_colorMap[TYPE_POINT1] = {type: 'point', color: '#FF0000', size: 2.0};
_colorMap[TYPE_POINT4] = {type: 'point', color: '#0000FF', size: 2.0};
_colorMap[TYPE_CROSS1] = {type: 'cross', color: '#0000FF', size: 4.0};
_colorMap[TYPE_DIAMOND1] = {type: 'diamond', color: '#0000FF', size: 4.0};
_colorMap[TYPE_DIAMOND2] = {type: 'diamond', color: '#0000FF', size: 4.0};
_colorMap[TYPE_DIAMOND3] = {type: 'diamond', color: '#0000FF', size: 2.0};
_colorMap[TYPE_CIRCLETRIANGLE1] = {type: 'circletriangle', color: '#0000FF', size: 4.0};
_colorMap[TYPE_CIRCLETRIANGLE2] = {type: 'circletriangle', color: '#0000FF', size: 4.0};
_colorMap[TYPE_CIRCLETRIANGLE3] = {type: 'circletriangle', color: '#0000FF', size: 2.0};
_colorMap[TYPE_FILLCIRCLE1] = {type: 'fillcircle', color: '#FF0000', size: 4.0};
_colorMap[TYPE_FILLCIRCLE2] = {type: 'fillcircle', color: '#FF0000', size: 4.0};
_colorMap[TYPE_FILLCIRCLE3] = {type: 'fillcircle', color: '#FF0000', size: 2.0};

var TYPE_EXTRASTATUS = 300;

var _adjustX = -23;
var _adjustY = -118;
var _w;
var _h;
var _imgWidth = 0;
var _imgHeight = 0;
var _bg = '';
var _bgReload = true;
var _zoom = 1.0;
var _imageZoom = 5.0;
var _arrowZoom = 3.0;

var canvas;
var ctx;

var color = '#0000ff';
var oldx;
var oldy;
var down = 0;

var _mode = 0;

var _items = [];
var _items2 = [];
var _type = 1;

var _undos = [];
var _redos = [];
var _imgs = [];
var imgRemark;
var _memo = false;

var _r = 3.7;

var _modified = false;
var _x1;
var _y1;

var _singleTarget = 0;

var _area = false;

function getImageName(str) {
  if (_serverMode == 'mobile') {
    return str.replace('webdata/', '');
  } else {
    return str;
  }
}

function getTypeColor(type) {
  var color = '#FF0000';
  if (type == TYPE_NUMBER2)
    color = '#0000FF';
  else if (type == TYPE_CURVE)
    color = '#0000FF';
  else if (type == TYPE_STRAIGHT)
    color = '#0000FF';
  else if (type == TYPE_CURVE2)
    color = '#FF0000';
  else if (type == TYPE_STRAIGHT2)
    color = '#FF0000';
  else if (type == TYPE_CURVE3)
    color = '#000000';
  else if (type == TYPE_STRAIGHT3)
    color = '#000000';
  else if (type == TYPE_CURVE4)
    color = '#15a419';
  else if (type == TYPE_STRAIGHT4)
    color = '#15a419';

  return color;
}

function getImagePathCamera(str) {
  //return "data:image/jpeg;base64," + str;
  if (_serverMode == 'mobile')
    return str;
  else
    return getImagePath(str);
}

    
function getImagePath(str) {
  if (str == null)
    str = '';
  if (_serverMode == 'mobile') {
    if (_filetype == 5 || _filetype == FILETYPE_FLOORDRAW)
      return 'assets/img/floor.png';
    else {
      var ret = cordova.file.dataDirectory + str.replace('webdata/', '');
      ret = ret.replace(/^file:\/\//, '');

      return ret;
    }
  } else {
    return _serverAddr + '/' + str;
  }
}

function getImageViewPath(str) {
  if (str == null)
    str = '';
  if (_serverMode == 'mobile') {
    var ret = cordova.file.dataDirectory + str.replace('webdata/', '');
    ret = ret.replace(/^file:\/\//, '');

    return ret;
  } else {
    return str;
  }
}

function initCanvas() {  
  canvas = document.getElementById("canvas");
  ctx = canvas.getContext("2d");

  if (_serverMode != 'mobile') {
    canvas.onmouseup = function (event) {
      onMouseUp(event);
    };

    canvas.onmousedown = function (event) {
      onMouseDown(event);
    };

    canvas.onmousemove = function (event) {
      onMouseMove(event);
    };
  }

  canvas.addEventListener("touchstart", function(event) {
    if (_serverMode != 'mobile') {
      if (_mode == MODE_DRAW) {
        if (_type != TYPE_MEMO) {
          if (!isLine(_type) && !isArrow(_type) && _type != TYPE_LENGTH && _type != TYPE_AREA && _type != TYPE_LENGTH2)
            return;
        }
      }
    }

    onMouseDown(event);
  }, false);

  canvas.addEventListener("touchmove", function(event) {
    if (_serverMode != 'mobile') {
      if (_mode == MODE_DRAW) {
        if (_type != TYPE_MEMO) {
          if (!isLine(_type) && !isArrow(_type) && _type != TYPE_LENGTH && _type != TYPE_AREA && _type != TYPE_LENGTH2)
            return;
        }
      }
    }

    onMouseMove(event);
  }, false);

  canvas.addEventListener("touchend", function(event) {
    if (_serverMode != 'mobile') {
      if (_mode == MODE_DRAW) {
        if (_type != TYPE_MEMO) {
          if (!isLine(_type) && !isArrow(_type) && _type != TYPE_LENGTH && _type != TYPE_AREA && _type != TYPE_LENGTH2)
            return;
        }
      }
    }

    onMouseUp(event);
  }, false);

  canvas.addEventListener("contextmenu", function(event) {
    event.preventDefault();
    _area = false;
    down = 0;
     _mouseDown = false;

    viewAreaPopup();
  }, false);


  imgRemark = new Image();
  imgRemark.src = "assets/img/remark.png";

  for (var i = 1; i <= 12; i++) {
    var img = new Image();
    img.src = "assets/img/i_crack_" + pad(i, 2) + ".png";

    _imgs.push(img);
  }

  for (var i = 81; i <= 87; i++) {
    var img = new Image();
    img.src = "assets/img/i_crack_" + i + ".png";

    _imgs.push(img);
  }

  resize();

  $(window).resize(function() {
    resize();
  });

  window.requestAnimationFrame(redrawAnimation);

  clearScreen();
}

function makeUndo() {
  var items = JSON.parse(JSON.stringify(_items));
  for (var j = 0; j < items.length; j++) {
    items[j].selected = false;
  }
  _undos.push(items);
}

function initModified() {
  _modified = false;
}

function modified() {
  _modified = true;
}

function isModified() {
  return _modified;
}

function initUndo() {
  _redos = [];

  for (var j = 0; j < _items.length; j++) {
    _items[j].selected = false;
  }
}

function clearScreen() {
  ctx.fillStyle = '#666666';
  ctx.fillRect(0, 0, _w, _h);
}

function drawNumber(i, number, x, y, color) {
  ctx.beginPath();

  var fs = 4.0 * _imageZoom * _zoom;
  ctx.font = fs + 'px 나눔고딕';

  ctx.setLineDash([0]);
  if (i == _items.length) {
    if (color == null)
      ctx.strokeStyle = '#FF0000';
    else
      ctx.strokeStyle = color;
  } else {
    if (_mode != MODE_GROUP) {
      if (color == null)
        ctx.strokeStyle = '#FF0000';
      else
        ctx.strokeStyle = color;
    } else {
      if (_items[i].selected == true && _saveFlag == false)
        ctx.strokeStyle = '#00cc00';
      else {
        if (_items[i].mode == TYPE_NUMBER)
          ctx.strokeStyle = '#FF0000';
        else
          ctx.strokeStyle = '#0000FF';
      }
    }
  }

  x = parseFloat(x);
  y = parseFloat(y);


  var r = _r * _zoom;

  var x1 = x * _zoom - r/2;
  var y1 = y * _zoom - r/2;

  ctx.arc(x1, y1, r * _imageZoom, 0, Math.PI * 2, true);
  ctx.fillStyle = '#ffffff';
  ctx.fill();

  if (number > 9) {
    ctx.strokeText(number, x1 - 2.5 * _imageZoom * _zoom, y1 + 1.4 * _imageZoom * _zoom);
  } else {
    ctx.strokeText(number, x1 - 1.0 * _imageZoom * _zoom, y1 + 1.4 * _imageZoom * _zoom);
  }

  ctx.stroke();
}

function byte_length(str) {
  var count = 0;
  var ch = '';

  for(var i = 0; i < str.length; i++) {
    ch = str.charAt(i);
    if(escape(ch).length == 6) {
      count ++;
    }

    count ++;
  }

  return count;
}

function drawExtrastatus(str) {
  if (str == undefined || str == null || str == '') {
    return;
  }

  ctx.beginPath();

  var fs = 12.0 * _zoom;
  ctx.font = fs + 'px 나눔고딕';

  ctx.setLineDash([0]);
  ctx.strokeStyle = '#FF0000';

  var x = 10;
  var y = 10;

  ctx.strokeText(str, (x + 10) * _zoom, (y + 17) * _zoom);
  ctx.stroke();

  ctx.strokeRect(x * _zoom, y * _zoom, fs * (byte_length(str) / 2) + 20 * _zoom, fs + 14 * _zoom);
  ctx.stroke();
}

function isArrow(mode) {
  if (mode == TYPE_ARROW || mode == TYPE_ARROWV || mode == TYPE_ARROWFREE || mode == TYPE_ARROWCHAIN || mode == TYPE_LENGTH)
    return true;
  else
    return false;
}

function isNumber(mode) {
  if (mode == TYPE_NUMBER || mode == TYPE_NUMBER2)
    return true;
  else
    return false;
}

function isMemo(mode) {
  if (mode == TYPE_MEMO)
    return true;
  else
    return false;
}

function isLine(mode) {
  if (isStraightLine(mode) || isCurveLine(mode))
    return true;
  else
    return false;
}

function isStraightLine(mode) {
  if (mode == TYPE_STRAIGHT || mode == TYPE_STRAIGHT2 || mode == TYPE_STRAIGHT3 || mode == TYPE_STRAIGHT4)
    return true;
  else
    return false;
}

function isCurveLine(mode) {
  if (mode == TYPE_CURVE || mode == TYPE_CURVE2 || mode == TYPE_CURVE3 || mode == TYPE_CURVE4)
    return true;
  else
    return false;
}

function isGroupable(mode) {
  if (isNumber(mode) || isArrow(mode) || mode == TYPE_AREA || mode == TYPE_LENGTH2)
    return true;
  else
    return false;
}

function drawSlote(i, number, mode, color, x1, y1, x2, y2, x3, y3) {
  color = '#FF0000';

  ctx.beginPath();
  ctx.setLineDash([0]);
  ctx.strokeStyle = color;

  if (mode == TYPE_ARROW) {
    if (y1 == y2) {
      ctx.moveTo(x1 * _zoom, y1 * _zoom);
      ctx.lineTo(x2 * _zoom, y2 * _zoom);
    } else {
      ctx.moveTo(x1 * _zoom, y1 * _zoom);
      ctx.lineTo(x2 * _zoom, y1 * _zoom);
      ctx.lineTo(x2 * _zoom, y2 * _zoom);
    }
  } else if (mode == TYPE_ARROWV) {
    if (x1 == x2) {
      ctx.moveTo(x1 * _zoom, y1 * _zoom);
      ctx.lineTo(x2 * _zoom, y2 * _zoom);
    } else {
      ctx.moveTo(x1 * _zoom, y1 * _zoom);
      ctx.lineTo(x1 * _zoom, y2 * _zoom);
      ctx.lineTo(x2 * _zoom, y2 * _zoom);
    }
  } else {
    ctx.moveTo(x1 * _zoom, y1 * _zoom);
    ctx.lineTo(x2 * _zoom, y2 * _zoom);
  }

  ctx.stroke();

  ctx.fillStyle = color;

  ctx.beginPath();
  var div = 3.0;
  if (mode == TYPE_ARROW) {
    if (x2 > x1) {
      ctx.moveTo(x1 * _zoom, y1 * _zoom);
      ctx.lineTo((x1 + 4.0/div * _arrowZoom) * _zoom, (y1 - 3.0/div * _arrowZoom) * _zoom);
      ctx.lineTo((x1 + 4.0/div * _arrowZoom) * _zoom, (y1 + 3.0/div * _arrowZoom) * _zoom);
      ctx.lineTo(x1 * _zoom, y1 * _zoom);
      ctx.fill();
    } else if (x1 > x2) {
      ctx.moveTo(x1 * _zoom, y1 * _zoom);
      ctx.lineTo((x1 - 4.0/div * _arrowZoom) * _zoom, (y1 - 3.0/div * _arrowZoom) * _zoom);
      ctx.lineTo((x1 - 4.0/div * _arrowZoom) * _zoom, (y1 + 3.0/div * _arrowZoom) * _zoom);
      ctx.lineTo(x1 * _zoom, y1 * _zoom);
      ctx.fill();
    }
  } else if (mode == TYPE_ARROWV) {
    if (y2 > y1) {
      ctx.moveTo(x1 * _zoom, y1 * _zoom);
      ctx.lineTo((x1 + 3.0/div * _arrowZoom) * _zoom, (y1 + 4.0/div * _arrowZoom) * _zoom);
      ctx.lineTo((x1 - 3.0/div * _arrowZoom) * _zoom, (y1 + 4.0/div * _arrowZoom) * _zoom);
      ctx.lineTo(x1 * _zoom, y1 * _zoom);
      ctx.fill();
    } else if (y1 > y2) {
      ctx.moveTo(x1 * _zoom, y1 * _zoom);
      ctx.lineTo((x1 + 3.0/div * _arrowZoom) * _zoom, (y1 - 4.0/div * _arrowZoom) * _zoom);
      ctx.lineTo((x1 - 3.0/div * _arrowZoom) * _zoom, (y1 - 4.0/div * _arrowZoom) * _zoom);
      ctx.lineTo(x1 * _zoom, y1 * _zoom);
      ctx.fill();
    }
  } else if (mode == TYPE_ARROWFREE || mode == TYPE_ARROWCHAIN) {
    var dy = y2 - y1;
    var dx = x2 - x1;
    var angle = Math.atan(dy/dx) * (180.0/Math.PI);

    angle += 180;

    if(dx < 0.0) {
      angle += 180.0;
    } else {
      if(dy<0.0) angle += 360.0;
    }

    var r = 4.0/div * _arrowZoom;
    var x;
    var y;

    x = Math.cos(angle * Math.PI / 180) * r;
    y = Math.sin(angle * Math.PI / 180) * r;

    ctx.moveTo((x1 + x) * _zoom, (y1 + y) * _zoom);

    angle += 120;

    x = Math.cos(angle * Math.PI / 180) * r;
    y = Math.sin(angle * Math.PI / 180) * r;

    ctx.lineTo((x1 + x) * _zoom, (y1 + y) * _zoom);

    angle += 120;

    x = Math.cos(angle * Math.PI / 180) * r;
    y = Math.sin(angle * Math.PI / 180) * r;

    ctx.lineTo((x1 + x) * _zoom, (y1 + y) * _zoom);

    angle += 120;

    x = Math.cos(angle * Math.PI / 180) * r;
    y = Math.sin(angle * Math.PI / 180) * r;

    ctx.lineTo((x1 + x) * _zoom, (y1 + y) * _zoom);

    ctx.fill();
  }
  ctx.stroke();

  var dx = 0;
  var dy = 0;
  
  if (mode == TYPE_LENGTH) {
    dy = 15;
  } else if (mode == TYPE_ARROW) {
    dy -= _r * _imageZoom / 2;
    dx -= _r * _imageZoom / 2;
  } else if (mode == TYPE_ARROWV) {
    dy -= _r * _imageZoom / 2;
    dx -= _r * _imageZoom / 2;    
  } else if (mode == TYPE_ARROWFREE) {
    dy -= _r * _imageZoom / 2;
    dx -= _r * _imageZoom / 2;
  } else if (mode == TYPE_ARROWCHAIN) {
    dy -= _r * _imageZoom / 2;
    dx -= _r * _imageZoom / 2;

    x2 = x3;
    y2 = y3;
  }

  dx = 0;
  dy = 0;

  if (mode == TYPE_LENGTH) {
    //drawNumber(i, number, x2 - 20 * _zoom - 20, y2 - 20 * _zoom - 20, null);
    drawNumber(i, number, (x2 + x1) / 2, (y2 + y1) / 2, null);    
  } else {
    drawNumber(i, number, x2 + dx, y2 + dy, null);
  }
}

function redraw() {
  ctx.setLineDash([0]);
  var max = 0;
  for (var i = 0; i < _items.length; i++) {
    var mode = _items[i].mode;
    if (!isNumber(mode) && !isArrow(mode))
      continue;

    if (_items[i].group > max)
      max = _items[i].group;
  }

  for (var i = 0; i < _items.length; i++) {
    var item = _items[i];

    if (_memo == true) {
      if (item.mode != 30) {
        continue;
      }
    } else {
      if (item.mode == 30) {
        continue;
      }
    }

    if (isLine(item.mode) || isMemo(item.mode) || item.mode == TYPE_AREA || item.mode == TYPE_LENGTH2  || item.mode == TYPE_ARROWCHAIN) {
      var oldx1 = 0;
      var oldy1 = 0;

      var maxx = -9999;
      var maxy = -9999;
      var minx = 99999;
      var miny = 99999;

      for (var j = 0; j < item.points.length; j++) {
        var point = item.points[j];
        if (j == 0) {
          oldx1 = point.x;
          oldy1 = point.y;

          if (_area == true && (_type == TYPE_AREA || _type == TYPE_LENGTH2) && i == _items.length - 1) {
            if (item.mode == TYPE_AREA || item.mode == TYPE_LENGTH2) {
              ctx.beginPath();
              ctx.setLineDash([0]);
              ctx.arc(oldx1 * _zoom, oldy1 * _zoom, 5, 0, Math.PI * 2, true);
              ctx.fillStyle = '#ff0000';
              ctx.fill();
              ctx.stroke();
            }
          }
        }

        ctx.beginPath();
        ctx.setLineDash([0]);
        if (item.mode == 1 || item.mode == 2) {
          ctx.lineWidth = 1;
          ctx.strokeStyle = '#0000FF';
        } else if (item.mode == 15 || item.mode == 16) {
          ctx.lineWidth = 1;
          ctx.strokeStyle = '#FF0000';
        } else if (item.mode == TYPE_STRAIGHT3 || item.mode == TYPE_CURVE3) {
          ctx.lineWidth = 2;
          ctx.strokeStyle = '#000000';
        } else if (item.mode == TYPE_STRAIGHT4 || item.mode == TYPE_CURVE4) {
          ctx.lineWidth = 2;
          ctx.strokeStyle = '#15a419';
        } else if (item.mode == 30) {
          ctx.lineWidth = 2;
          ctx.strokeStyle = '#000000';
        } else if (item.mode == TYPE_AREA || item.mode == TYPE_LENGTH2 || item.mode == TYPE_ARROWCHAIN) {
          ctx.lineWidth = 1;
          ctx.strokeStyle = '#FF0000';
        }

        ctx.moveTo(oldx1 * _zoom, oldy1 * _zoom);
        ctx.lineTo(point.x * _zoom, point.y * _zoom);

        oldx1 = point.x;
        oldy1 = point.y;

        ctx.stroke();

        if (item.mode == TYPE_AREA || item.mode == TYPE_LENGTH2 || item.mode == TYPE_ARROWCHAIN) {
          if (point.x < minx)
            minx = point.x;

          if (point.x > maxx)
            maxx = point.x;

          if (point.y < miny)
            miny = point.y;

          if (point.y > maxy)
            maxy = point.y;
        }

        ctx.lineWidth = 1;
      }

      if (item.mode == TYPE_AREA && (_area == false || (_area == true && i != _items.length - 1))) {
        var ax = minx + (maxx - minx) / 2;
        var ay = miny + (maxy - miny) / 2;
        
        if (item.points.length > 1) {
          var point1 = item.points[0];
          var point2 = item.points[1];
          ax = (point1.x + point2.x) / 2;
          ay = (point1.y + point2.y) / 2;
        }
        
        drawNumber(i, item.number, ax, ay, '#FF0000');
      } else if (item.mode == TYPE_LENGTH2 && (_area == false || (_area == true && i != _items.length - 1))) {
        var ax = minx + (maxx - minx) / 2;
        var ay = miny + (maxy - miny) / 2;

        if (item.points.length > 1) {
          var point1 = item.points[0];
          var point2 = item.points[1];
          ax = (point1.x + point2.x) / 2;
          ay = (point1.y + point2.y) / 2;
        }
        drawNumber(i, item.number, ax, ay, '#FF0000');
      } else if (item.mode == TYPE_ARROWCHAIN) {
        if (item.points.length > 1) {
          drawSlote(i, item.number, item.mode, '#FF0000', item.points[0].x, item.points[0].y, item.points[1].x, item.points[1].y, item.points[item.points.length - 1].x, item.points[item.points.length - 1].y);
        }

        /*
        var ax = minx + (maxx - minx) / 2;
        var ay = miny + (maxy - miny) / 2;
        drawNumber(i, item.number, ax, ay, '#FF0000');
        */
      }
    } else if (isArrow(item.mode)) {
      if (item.points.length > 1)
        drawSlote(i, item.number, item.mode, '#FF0000', item.points[0].x, item.points[0].y, item.points[1].x, item.points[1].y);
    } else if (isNumber(item.mode)) {
      if (item.mode == TYPE_NUMBER)
        drawNumber(i, item.number, item.x, item.y, null);
      else
        drawNumber(i, item.number, item.x, item.y, '#0000FF');
    } else if (item.mode == 0) {
    } else if (item.mode == TYPE_EXTRASTATUS) {
      drawExtrastatus(item.name);
    } else if (item.mode > 2) {
      drawImage(item.mode, item.x, item.y);
    }

    if (item.selected == true) {
      var minx = 9999;
      var miny = 9999;
      var maxx = -1;
      var maxy = -1;

      if (isLine(item.mode) || isMemo(item.mode) || item.mode == TYPE_AREA || item.mode == TYPE_LENGTH2 || item.mode == TYPE_ARROWCHAIN) {
        for (var j = 0; j < item.points.length; j++) {
          var point = item.points[j];

          if (point.x > maxx)
            maxx = point.x;

          if (point.x < minx)
            minx = point.x;

          if (point.y > maxy)
            maxy = point.y;

          if (point.y < miny)
            miny = point.y;
        }
      } else if (isArrow(item.mode)) {
        for (var j = 0; j < item.points.length; j++) {
          var point = item.points[j];

          if (point.x > maxx)
            maxx = point.x;

          if (point.x < minx)
            minx = point.x;

          if (point.y > maxy)
            maxy = point.y;

          if (point.y < miny)
            miny = point.y;
        }

        minx -= 10;
        miny -= 10;
        maxx += 10;
        maxy += 10;
      } else if (isNumber(item.mode)) {
        minx = item.x - 4 * _imageZoom;
        miny = item.y - 4 * _imageZoom;
        maxx = item.x + 4 * _imageZoom;
        maxy = item.y + 4 * _imageZoom;
      } else {
        minx = item.x - 4 * _imageZoom;
        miny = item.y - 4 * _imageZoom;
        maxx = item.x + 4 * _imageZoom;
        maxy = item.y + 4 * _imageZoom;
      }

      ctx.stroke();

      if (_saveFlag == false) {
        ctx.strokeStyle = '#00aa00';
        ctx.setLineDash([6]);
        ctx.strokeRect(minx * _zoom, miny * _zoom, (maxx - minx) * _zoom, (maxy - miny) * _zoom);
      }
    }
  }

  if (_filetype == 4) {
    var bg = '#ff0000';
    var bg2 = '#0000ff';


    var cnt = 0;
    var x = 145;
    var my = 70;
    var font = parseInt(11 * _zoom) + 'px 나눔고딕';
    var step = 20;
    var w = _imgWidth;

    ctx.beginPath();
    ctx.font = font;
    ctx.setLineDash([0]);
    ctx.strokeStyle = '#ff0000';
    ctx.strokeText('수직', (w - x + 100) * _zoom, 20 * _zoom);
    ctx.stroke();

    ctx.beginPath();
    ctx.font = font;
    ctx.setLineDash([0]);
    ctx.strokeStyle = '#0000ff';
    ctx.strokeText('수평', (w - x + 100) * _zoom, 40 * _zoom);
    ctx.stroke();


    var pos = 0;
    for (var i = 0; i < _items.length; i++) {
      var item = _items[i];
      if (!isNumber(item.mode)) {
        continue;
      }

      pos++;

      if (pos % 2 != 0)
        continue;

      if (Array.isArray(item.memo)) {
        var flag = false;
        for (var j = 0; j < item.memo.length; j++) {
          var memo = item.memo[j];

          if (memo.floor == '' || memo.t == '')
            continue;

          flag = true;
          break;
        }

        if (flag == true) {
          var col;

          if (item.mode == TYPE_NUMBER)
            col = bg;
          else
            col = bg2;

          ctx.beginPath();
          ctx.font = font;
          ctx.setLineDash([0]);
          ctx.strokeStyle = col;
          ctx.strokeText('측정위치', (w - x) * _zoom, (my + cnt * step) * _zoom);
          ctx.stroke();

          drawNumber(_items.length, pos-1, (w - x + 62), (my + cnt * step + 4), col);

          ctx.beginPath();
          ctx.font = font;
          ctx.setLineDash([0]);
          ctx.strokeStyle = col;
          ctx.strokeText('측정위치', (w - x + 70) * _zoom, (my + cnt * step) * _zoom);
          ctx.stroke();

          drawNumber(_items.length, pos, (w - x + 70 + 62), (my + cnt * step + 4), col);

          cnt++;

          ctx.beginPath();
          ctx.font = font;
          ctx.setLineDash([0]);
          ctx.strokeStyle = col;
          ctx.strokeText('층', (w - x) * _zoom, (my + cnt * step) * _zoom);
          ctx.stroke();

          ctx.beginPath();
          ctx.font = font;
          ctx.setLineDash([0]);
          ctx.strokeStyle = col;
          ctx.strokeText('T', (w - x + 50) * _zoom, (my + cnt * step) * _zoom);
          ctx.stroke();

          ctx.beginPath();
          ctx.font = font;
          ctx.setLineDash([0]);
          ctx.strokeStyle = col;
          ctx.strokeText('층', (w - x + 70) * _zoom, (my + cnt * step) * _zoom);
          ctx.stroke();

          ctx.beginPath();
          ctx.font = font;
          ctx.setLineDash([0]);
          ctx.strokeStyle = col;
          ctx.strokeText('T', (w - x + 70 + 50) * _zoom, (my + cnt * step) * _zoom);
          ctx.stroke();

          cnt++;
          for (var j = 0; j < item.memo.length; j++) {
            var memo = item.memo[j];

            if (memo.floor == '' || memo.t == '')
              continue;

            var t2 = '';
            try {
              t2 = parseInt(memo.t) + 1;
            } catch (e) {
            }

            ctx.beginPath();
            ctx.font = font;
            ctx.setLineDash([0]);
            ctx.strokeStyle = col;
            ctx.strokeText(memo.floor, (w - x) * _zoom, (my + cnt * step) * _zoom);
            ctx.stroke();

            ctx.beginPath();
            ctx.font = font;
            ctx.setLineDash([0]);
            ctx.strokeStyle = col;
            ctx.strokeText(parseInt(memo.t), (w - x + 50) * _zoom, (my + cnt * step) * _zoom);
            ctx.stroke();



            ctx.beginPath();
            ctx.font = font;
            ctx.setLineDash([0]);
            ctx.strokeStyle = col;
            ctx.strokeText(memo.floor, (w - x + 70) * _zoom, (my + cnt * step) * _zoom);
            ctx.stroke();

            ctx.beginPath();
            ctx.font = font;
            ctx.setLineDash([0]);
            ctx.strokeStyle = col;
            ctx.strokeText(t2, (w - x + 70 + 50) * _zoom, (my + cnt * step) * _zoom);
            ctx.stroke();

            cnt++;
          }
        }
      }

      my += 10;
    }
  }



  if (_filetype == 3) {
    var bg = '#ffcccc';
    var bg2 = '#ccccff';

    var b = 0;
    var r = 0;

    for (var i = 0; i < _items.length; i++) {
      var item = _items[i];

      if (item.mode == 13) {
        b++;
      } else if (item.mode == 14) {
        r++;
      }
    }

    if (b == 1 && r == 1) {
      bg = '#ff0000';
      bg2 = '#0000ff';
    } else if (b >= 2) {
      bg = '#0000ff';
      bg2 = '#0000ff';
    } else if (r >= 2) {
      bg = '#ff0000';
      bg2 = '#ff0000';
    } else {
      return;
    }

    var cnt = 0;
    var x = 140;
    var w = _imgWidth;
    var font = parseInt(11 * _zoom) + 'px 나눔고딕';

    if (b >= 1) {
      ctx.beginPath();
      ctx.font = font;
      ctx.setLineDash([0]);
      ctx.strokeStyle = '#0000ff';
      ctx.strokeText('수평', (w - x + 60) * _zoom, 50 * _zoom);
      ctx.stroke();
    }


    if (r >= 1) {
      ctx.beginPath();
      ctx.font = font;
      ctx.setLineDash([0]);
      ctx.strokeStyle = '#ff0000';
      ctx.strokeText('수직', (w - x + 60) * _zoom, 80 * _zoom);
      ctx.stroke();
    }


    ctx.beginPath();
    ctx.font = font;
    ctx.setLineDash([0]);
    ctx.strokeStyle = bg;
    ctx.strokeText('층', (w - x) * _zoom, (120 + cnt * 30) * _zoom);
    ctx.stroke();


    ctx.beginPath();
    ctx.font = font;
    ctx.setLineDash([0]);
    ctx.strokeStyle = bg;
    ctx.strokeText('SH', (w - x + 50) * _zoom, (120 + cnt * 30) * _zoom);
    ctx.stroke();

    ctx.beginPath();
    ctx.font = font;
    ctx.setLineDash([0]);
    ctx.strokeStyle = bg;
    ctx.strokeText('N', (w - x + 100) * _zoom, (120 + cnt * 30) * _zoom);
    ctx.stroke();

    for (var i = 0; i < _items2.length; i++) {
      var item = _items2[i];

      if (item.name == '')
        continue;

      ctx.beginPath();
      ctx.font = font;
      ctx.setLineDash([0]);
      ctx.strokeStyle = bg;
      ctx.strokeText(item.name, (w - x) * _zoom, (150 + cnt * 30) * _zoom);
      ctx.stroke();


      ctx.beginPath();
      ctx.font = font;
      ctx.setLineDash([0]);
      ctx.strokeStyle = bg;
      ctx.strokeText(item.content + ',' + item.width, (w - x + 50) * _zoom, (150 + cnt * 30) * _zoom);
      ctx.stroke();

      ctx.beginPath();
      ctx.font = font;
      ctx.setLineDash([0]);
      ctx.strokeStyle = bg;
      ctx.strokeText(item.length, (w - x + 100) * _zoom, (150 + cnt * 30) * _zoom);
      ctx.stroke();

      cnt++;
    }

    cnt++;
    cnt++;


    ctx.beginPath();
    ctx.font = font;
    ctx.setLineDash([0]);
    ctx.strokeStyle = bg2;
    ctx.strokeText('층', (w - x) * _zoom, (120 + cnt * 30) * _zoom);
    ctx.stroke();


    ctx.beginPath();
    ctx.font = font;
    ctx.setLineDash([0]);
    ctx.strokeStyle = bg2;
    ctx.strokeText('SH', (w - x + 50) * _zoom, (120 + cnt * 30) * _zoom);
    ctx.stroke();

    ctx.beginPath();
    ctx.font = font;
    ctx.setLineDash([0]);
    ctx.strokeStyle = bg2;
    ctx.strokeText('N', (w - x + 100) * _zoom, (120 + cnt * 30) * _zoom);
    ctx.stroke();

    for (var i = 0; i < _items2.length; i++) {
      var item = _items2[i];

      if (item.name == '')
        continue;

      ctx.beginPath();
      ctx.font = font;
      ctx.setLineDash([0]);
      ctx.strokeStyle = bg2;
      ctx.strokeText(item.name, (w - x) * _zoom, (150 + cnt * 30) * _zoom);
      ctx.stroke();


      ctx.beginPath();
      ctx.font = font;
      ctx.setLineDash([0]);
      ctx.strokeStyle = bg2;
      ctx.strokeText((parseInt(item.content)+2) + ',' + (parseInt(item.width)+2), (w - x + 50) * _zoom, (150 + cnt * 30) * _zoom);
      ctx.stroke();

      ctx.beginPath();
      ctx.font = font;
      ctx.setLineDash([0]);
      ctx.strokeStyle = bg2;
      ctx.strokeText((parseInt(item.length)+1), (w - x + 100) * _zoom, (150 + cnt * 30) * _zoom);
      ctx.stroke();

      cnt++;
    }
  }

  if ((_mode == MODE_SELECT || _mode == MODE_GROUP || _mode == MODE_ZOOM) && down == 1) {
    var x1 = oldx;
    var y1 = oldy;
    var x2 = _x1;
    var y2 = _y1;
    var temp;

    if (x1 > x2) {
      temp = x1;
      x1 = x2;
      x2 = temp;
    }

    if (y1 > y2) {
      temp = y1;
      y1 = y2;
      y2 = temp;
    }

    ctx.stroke();
    ctx.strokeStyle = '#00aa00';
    ctx.setLineDash([6]);
    ctx.strokeRect(x1 * _zoom, y1 * _zoom, (x2 - x1) * _zoom, (y2 - y1) * _zoom);
  }
}

var _originalZoom = 0;

function redrawAnimation() {
  if (_bg != '') {
    var img = new Image();
    //img.crossOrigin = 'Anonymous';

    var filename = getImagePath(_bg);

    img.src = filename;

    img.onload = function() {
      clearScreen();

      if (_bgReload == true) {
        var zoom = _w / img.width;
        if (img.height * zoom > _h) {
          zoom = _h / img.height;
        }

        if (_originalZoom == 0) {
          _originalZoom = zoom;
        }

        _zoom = zoom;

        if (_filetype == 5 || _filetype == FILETYPE_FLOORDRAW)
          _zoom = 1.0;


        _imgWidth = img.width;
        _imgHeight = img.height;

        resizeCanvas(true);

        ctx.drawImage(img, 0, 0, img.width * _zoom, img.height * _zoom);

        _bgReload = false;

        //$('#canvas').css('width', parseInt(_imgWidth * _zoom) + 'px');
        //$('#canvas').css('height', parseInt(_imgHeight * _zoom) + 'px');
      } else {
        ctx.drawImage(img, 0, 0, img.width * _zoom, img.height * _zoom);
      }

      /*
        if (_filetype == 1) {
        var w = 398 * _zoom / 2;
        var h = 860 * _zoom / 2;

        var x = _w * _zoom - w - 30;
        var y = _h * _zoom - h - 30;

        ctx.drawImage(imgRemark, x, y, w, h);
        ctx.strokeStyle = '#000000';
        ctx.rect(x, y, w, h);
        ctx.stroke();
        }
      */

      redraw();
    };
    
  } else {
    clearScreen();
    redraw();
  }

  window.requestAnimationFrame(redrawAnimation);
}


function drawVectorImage2(pos, x, y) {
  if (pos > 600)
    pos -= 100;

  var item = _colorMap[pos];

  var zoom = _zoom;
  if (item.type == 'line') {
    var size = item.size * _imageZoom;
    x -= (size / 2);
    y -= (size / 2);
    drawVectorLine(x, y, zoom, item);
  } else if (item.type == 'circle') {
    var size = item.size * _imageZoom / 2;
    x = x + size/2;
    y = y + size/2;
    drawVectorCircle(x, y, zoom, item);
  } else if (item.type == 'fillcircle') {
    var size = item.size * _imageZoom / 2;
    x = x + size/2;
    y = y + size/2;
    drawVectorFillcircle(x, y, zoom, item);
  } else if (item.type == 'triangle') {
    var size = item.size * _imageZoom;
    x -= (size / 2);
    y -= (size / 2);
    drawVectorTriangle(x, y, zoom, item);
  } else if (item.type == 'box') {
    var size = item.size * _imageZoom;
    x -= (size / 2);
    y -= (size / 2);
    drawVectorBox(x, y, zoom, item);
  } else if (item.type == 'crack') {
    var size = item.size * _imageZoom / 2;
    x = x + size;
    y = y + size;
    drawVectorCrack(x, y, zoom, item, pos);
  } else if (item.type == 'point') {
    drawVectorPoint(x, y, zoom, item);
  } else if (item.type == 'cross') {
    var size = item.size * _imageZoom;
    x -= (size / 2);
    y -= (size / 2);
    drawVectorCross(x, y, zoom, item);
  } else if (item.type == 'diamond') {
    var size = item.size * _imageZoom;
    x -= (size / 2);
    y -= (size / 2);
    drawVectorDiamond(x, y, zoom, item);    
  } else if (item.type == 'circletriangle') {
    var size = item.size * _imageZoom / 2;
    x = x + size/2;
    y = y + size/2;
    drawVectorCircletriangle(x, y, zoom, item);
  }
}

function drawVectorLine(x, y, zoom, item) {
  var size = item.size * _imageZoom;
  ctx.beginPath();
  ctx.setLineDash([0]);
  ctx.lineWidth = 1;
  ctx.strokeStyle = item.color;
  
  var x1 = (x + size) * zoom;
  var y1 = y * zoom;
  var x2 = x * zoom;
  var y2 = (y + size) * zoom;
  
  ctx.moveTo(x1, y1);
  ctx.lineTo(x2, y2);

  ctx.stroke();  
}

function drawVectorCircle(x, y, zoom, item) {
  var size = item.size * _imageZoom / 2;
  var tx = (x - size/2) * zoom;
  var ty = (y - size/2) * zoom;

  ctx.beginPath();
  ctx.setLineDash([0]);
  ctx.lineWidth = 1;
  ctx.strokeStyle = item.color;

  ctx.arc(tx, ty, size, 0, Math.PI * 2, true);
  

  ctx.stroke();  
}

function drawVectorFillcircle(x, y, zoom, item) {
  var size = item.size * _imageZoom / 2;
  var tx = (x - size/2) * zoom;
  var ty = (y - size/2) * zoom;

  ctx.beginPath();
  ctx.setLineDash([0]);
  ctx.lineWidth = 1;
  ctx.strokeStyle = item.color;

  ctx.arc(tx, ty, size, 0, Math.PI * 2, true);
  ctx.fillStyle = '#ff0000';
  ctx.fill();  

  ctx.stroke();  
}

function drawVectorTriangle(x, y, zoom, item) {
  var size = item.size * _imageZoom;
  ctx.beginPath();
  ctx.setLineDash([0]);
  ctx.lineWidth = 1;
  ctx.strokeStyle = item.color;
  
  ctx.moveTo((x + size/2) * zoom, y * zoom);
  ctx.lineTo((x + size) * zoom, (y + size) * zoom);
  ctx.lineTo(x * zoom, (y + size) * zoom);
  ctx.lineTo((x + size/2) * zoom, y * zoom);

  ctx.fillStyle = item.color;
  ctx.fill();
  
  ctx.stroke();    
}

function drawVectorBox(x, y, zoom, item) {
  var size = item.size * _imageZoom;
  ctx.beginPath();
  ctx.setLineDash([0]);
  ctx.lineWidth = 1;
  ctx.strokeStyle = item.color;
  
  ctx.moveTo(x * zoom, y * zoom);
  ctx.lineTo((x + size) * zoom, y * zoom);
  ctx.lineTo((x + size) * zoom, (y + size) * zoom);
  ctx.lineTo(x * zoom, (y + size) * zoom);
  ctx.lineTo(x * zoom, y * zoom);  

  ctx.fillStyle = item.color;
  ctx.fill();
  
  ctx.stroke();  
}

function drawVectorCrack(x, y, zoom, item, pos) {
  var size = item.size * _imageZoom;
  var tx = (x - size/2) * zoom;
  var ty = (y - size/2) * zoom;

  ctx.beginPath();
  ctx.setLineDash([0]);
  ctx.lineWidth = 1;
  ctx.strokeStyle = item.color;

  ctx.arc(tx, ty, size/2, 0, Math.PI * 2, true);

  x -= size;
  y -= size;

  if (pos == TYPE_CRACK1) {        
    ctx.moveTo((x + size * 1.25) * zoom, (y - size * 0.25) * zoom);
    ctx.lineTo((x - size * 0.25) * zoom, (y + size * 1.25) * zoom);    
  } else if (pos == TYPE_CRACK2) {
    ctx.moveTo((x + size/2) * zoom, (y - size/2) * zoom);
    ctx.lineTo((x + size/2) * zoom, (y + size * 1.5) * zoom);
  } else if (pos == TYPE_CRACK3) {
    ctx.moveTo((x - size * 0.25) * zoom, (y - size * 0.25) * zoom);
    ctx.lineTo((x + size * 1.25) * zoom, (y + size * 1.25) * zoom);    
  } else if (pos == TYPE_CRACK4) {
    ctx.moveTo((x - size/2) * zoom, (y + size/2) * zoom);
    ctx.lineTo((x + size * 1.5) * zoom, (y + size/2) * zoom);
  }

  ctx.stroke();        
}

function drawVectorPoint(x, y, zoom, item) {
}

function drawVectorCross(x, y, zoom, item) {
  ctx.beginPath();
  ctx.setLineDash([0]);
  ctx.lineWidth = 1;
  ctx.strokeStyle = item.color;
  
  var x1 = (x + item.size * _imageZoom) * zoom;
  var y1 = y * zoom;
  var x2 = x * zoom;
  var y2 = (y + item.size * _imageZoom) * zoom;
  
  ctx.moveTo(x1, y1);
  ctx.lineTo(x2, y2);

  x1 = x * zoom;
  y1 = y * zoom;
  x2 = (x + item.size * _imageZoom) * zoom;
  y2 = (y + item.size * _imageZoom) * zoom;

  ctx.moveTo(x1, y1);
  ctx.lineTo(x2, y2);

  ctx.stroke();
}

function drawVectorDiamond(x, y, zoom, item) {
  var size = item.size * _imageZoom;
  ctx.beginPath();
  ctx.setLineDash([0]);
  ctx.lineWidth = 1;
  ctx.strokeStyle = item.color;
  
  ctx.moveTo((x + size/2) * zoom, y * zoom);
  ctx.lineTo((x + size) * zoom, (y + size/2) * zoom);
  ctx.lineTo((x + size/2) * zoom, (y + size) * zoom);
  ctx.lineTo(x * zoom, (y + size/2) * zoom);
  ctx.lineTo((x + size/2) * zoom, y * zoom);

  ctx.stroke();    
}

function drawVectorCircletriangle(x, y, zoom, item) {
  var size = item.size * _imageZoom;
  var tx = (x - size/2/2) * zoom;
  var ty = (y - size/2/2) * zoom;

  ctx.beginPath();
  ctx.setLineDash([0]);
  ctx.lineWidth = 1;
  ctx.strokeStyle = item.color;

  ctx.arc(tx, ty, size/2, 0, Math.PI * 2, true);

  x -= size;
  y -= size;

  size *= 1.4;
  
  ctx.moveTo((x + size/2) * zoom, y * zoom);
  ctx.lineTo((x + size) * zoom, (y + size) * zoom);
  ctx.lineTo(x * zoom, (y + size) * zoom);
  ctx.lineTo((x + size/2) * zoom, y * zoom);

  ctx.stroke();      
}

function drawVectorImage(pos, x, y) {
  var item = _colorMap[pos];
  
  var zoom = _zoom;
  if (item.type == 'line') {
    drawVectorLine(x, y, zoom, item);
  } else if (item.type == 'circle') {
    drawVectorCircle(x, y, zoom, item);
  } else if (item.type == 'fillcircle') {
    drawVectorFillcircle(x, y, zoom, item);
  } else if (item.type == 'triangle') {
    drawVectorTriangle(x, y, zoom, item);
  } else if (item.type == 'box') {
    drawVectorBox(x, y, zoom, item);
  } else if (item.type == 'crack') {
    drawVectorCrack(x, y, zoom, item, pos);
  } else if (item.type == 'point') {
    drawVectorPoint(x, y, zoom, item);
  } else if (item.type == 'cross') {
    drawVectorCross(x, y, zoom, item);
  } else if (item.type == 'diamond') {
    drawVectorDiamond(x, y, zoom, item);
  } else if (item.type == 'circletriangle') {
    drawVectorCircletriangle(x, y, zoom, item);
  }
}

function drawImage(pos, x, y) {
  if (pos >= 80 && pos <= 90) {
    pos = pos - 69;
  } else {
    if (pos > 600) {
      drawVectorImage2(pos, x, y);
      return;
    } else if (pos > 500) {
      drawVectorImage(pos, x, y);
      return;
    }

    pos = pos - 3;
  }

  var size = 9.0 * _imageZoom;

  var img = _imgs[pos];
  ctx.drawImage(img, (x - size/2) * _zoom, (y - size/2) * _zoom, size * _zoom, size * _zoom);
}

function undo() {
  closeSinglePopup();

  if (_undos.length > 0) {
    if (_undos.length == 1)
      _items = [];
    else
      _items = _undos[_undos.length - 1];

    _redos.push(_undos[_undos.length - 1]);
    _undos.splice(_undos.length - 1, 1);

    modified();
  }
}

function redo() {
  closeSinglePopup();

  if (_redos.length > 0) {
    _undos.push(_redos[_redos.length - 1]);
    _redos.splice(_redos.length - 1, 1);

    _items = _undos[_undos.length - 1];

    modified();
  }
}

function defaultZoom() {
  if (_originalZoom == 0)
    zoom(1);    
  else
    zoom(_originalZoom);

  updateLastScale();
  resizeCanvas(true);
}

function selectMode(mode) {
  clearArea();

  closeSinglePopup();

  if (mode == MODE_MOVE || mode == MODE_GROUP) {
    //makeUndo();
  } else {
    initUndo();
  }

  if (_mode == MODE_ZOOM && mode == MODE_ZOOM) {
    defaultZoom();
    return;
  }    

  _mode = mode;

  if (mode != 3) {
    for (var j = 0; j < _items.length; j++) {
      _items[j].selected = false;
    }
  }

  $('#btnNumber').css('color', 'red');
  $('#btnNumber').css('background-color', 'white');
  $('#btnNumber2').css('color', 'blue');
  $('#btnNumber2').css('background-color', 'white');

  setButton();

  if (_type == 13) {
    $('#btnCarbonation1').addClass('btn-select');
  } else if (_type == 14) {
    $('#btnCarbonation2').addClass('btn-select');
  } else if (_type == TYPE_ARROW) {
    $('#btnSlope1').addClass('btn-select');
  } else if (_type == TYPE_ARROWV) {
    $('#btnSlope2').addClass('btn-select');
  } else if (_type == TYPE_ARROWFREE) {
    $('#btnSlope3').addClass('btn-select');
    } else if (_type == TYPE_ARROWCHAIN) {
    $('#btnSlope4').addClass('btn-select');
  } else if (_type == TYPE_LENGTH) {
    $('#btnLength').addClass('btn-select');
  } else if (_type == TYPE_AREA) {
    $('#btnArea').addClass('btn-select');
  } else if (_type == TYPE_LENGTH2) {
    $('#btnLength2').addClass('btn-select');
  } else if (_type == TYPE_NUMBER) {
    $('#btnNumber').css('color', 'white');
    $('#btnNumber').css('background-color', 'red');
    $('#btnNumber2').css('color', 'blue');
    $('#btnNumber2').css('background-color', 'white');
  } else if (_type == TYPE_NUMBER2) {
    $('#btnNumber2').css('color', 'white');
    $('#btnNumber2').css('background-color', 'blue');
    $('#btnNumber').css('color', 'red');
    $('#btnNumber').css('background-color', 'white');
  } else if (_type == 21) {
    $('#btnGroup').addClass('btn-select');
  } else if (_type == TYPE_MEMO) {
    if (mode == 1)
      $('#btnMemoDraw').addClass('btn-select');
  }

  if (mode != MODE_DRAW) {
    $('#btnNumber').css('color', 'red');
    $('#btnNumber').css('background-color', 'white');
    $('#btnNumber2').css('color', 'blue');
    $('#btnNumber2').css('background-color', 'white');
  }


  if (mode == MODE_SELECT) {
    $('#btnSelect').addClass('btn-select');
  } else if (mode == MODE_MOVE) {
    $('#btnMove').addClass('btn-select');
  }

  if (mode == MODE_PAN) {
    $('#btnPan').addClass('btn-select');
    $('#btnZoom').removeClass('btn-select');
  } else if (mode == MODE_ZOOM) {
    $('#btnPan').removeClass('btn-select');
    $('#btnZoom').addClass('btn-select');
  } else {
    $('#btnPan').removeClass('btn-select');
    $('#btnZoom').removeClass('btn-select');
  }

  if (mode != MODE_GROUP)
    $('#divGroup').hide();

  if (mode == MODE_SELECT || mode == MODE_GROUP || mode == MODE_ZOOM) {
    _x1 = -1;
    _y1 = -1;
    oldx = -1;
    oldy = -1;
  }

  if (mode == MODE_SELECT || mode == MODE_MOVE || mode == MODE_ZOOM) {
    $('#btnCarbonation1').removeClass('btn-select');
    $('#btnCarbonation2').removeClass('btn-select');
    $('#btnSlope1').removeClass('btn-select');
    $('#btnSlope2').removeClass('btn-select');
    $('#btnSlope3').removeClass('btn-select');
    $('#btnSlope4').removeClass('btn-select');
    $('#btnNumber').removeClass('btn-select');
    $('#btnNumber2').removeClass('btn-select');
    $('#btnGroup').removeClass('btn-select');

    $('#btnLength').removeClass('btn-select');
    $('#btnLength2').removeClass('btn-select');
    $('#btnArea').removeClass('btn-select');
  }

  if (mode != MODE_DRAW) {
     resetType();    
  }
}

function inCheck(x1, y1, x2, y2, ptx, pty) {
  if (ptx >= x1 && ptx <= x2 && pty >= y1 && pty <= y2) {
    return true;
  } else {
    return false;
  }
}

function insertItem(mode, points, x, y, color, selected, group) {
  makeUndo();

  var number = 0;
  if (isGroupable(mode)) {
    var max = 0;
    for (var i = 0; i < _items.length; i++) {
      if (_items[i].number > max)
        max = _items[i].number;
    }

    number = max + 1;
  }

  var item = {
    mode: parseInt(_type),
    type: parseInt(_type),
    points: points,
    x: x,
    y: y,
    color:color,
    selected: selected,
    group: group,
    number: number,

    name: '',
    content: '',
    width: '0.2',
    length: '',
    count: '',
    progress: '',
    remark: '',
    imagename: '',
    filename: ''
  };

  _items.push(item);
  modified();
}

function onMouseUp(event) {
  if (_mode == 0)
    return;

  if (down == 0)
    return;

  if (_pinch == true)
    return;

  if (_mouseDown == true) {
    _mouseDown = false;
    mouseDownProcess();
  }

  if (_filetype == 6)
    viewAreaPopup();

  var x1 = getX(event);
  var y1 = getY(event);

  _x1 = x1;
  _y1 = y1;

  if (_mode == MODE_PAN) {
    var dx = oldx - _x1;
    var dy = oldy - _y1;

    var nx = container.scrollLeft + dx * _zoom;
    var ny = container.scrollTop + dy * _zoom;

    if (nx < 0) nx = 0;
    if (ny < 0) ny = 0;

    if (nx >= _imgWidth * _zoom) nx = _imgWidth * _zoom - 1;
    if (ny >= _imgHeight * _zoom) ny = _imgHeight * _zoom - 1;

    container.scrollLeft = nx;
    container.scrollTop = ny;

    oldx = _x1;
    oldy = _y1;
  }

  if (_mode == MODE_ZOOM) {
    var x1 = oldx;
    var y1 = oldy;
    var x2 = _x1;
    var y2 = _y1;

    var sx = Math.abs(x2 - x1);
    var sy = Math.abs(y2 - y1);

    if (sx < 10 && sy < 10) {
      defaultZoom();
    } else {
      var per1;
      var per2;
      var z = _zoom;

      if (sx == 0) {
        z = _h / sy;
      } else if (sy == 0) {
        z = _w / sx;
      } else {
        per1 = _w / _h;
        per2 = sx / sy;

        if (per2 > per1) {
          z = _w / sx;
        } else {
          z = _h / sy;
        }
      }

      if (z > 10.0)
        z = 10.0;

      if (z < 0.5) {
        z = 0.5;
      }

      var px = x1;
      var py = y1;

      if (x1 > x2) {
        px = x2;
      }

      //px += sx / 2;

      if (y1 > y2) {
        py = y2;
      }

      //py += sy / 2;j

      var dx = px * z;
      var dy = py * z;
      zoom(z);
      updateLastScale();
      resizeCanvas(true);

      container.scrollLeft = dx;
      container.scrollTop = dy;
    }

    down = 0;
    
    return;
  }
  
  if (_mode == MODE_SELECT || _mode == MODE_GROUP) {
    if (_mode == MODE_SELECT) {
      for (var j = 0; j < _items.length; j++) {
        _items[j].selected = false;
      }
    }

    var x1 = oldx;
    var y1 = oldy;
    var x2 = _x1;
    var y2 = _y1;
    var temp;

    if (x1 > x2) {
      temp = x1;
      x1 = x2;
      x2 = temp;
    }

    if (y1 > y2) {
      temp = y1;
      y1 = y2;
      y2 = temp;
    }

    for (var i = _items.length - 1; i >= 0; i--) {
      var item = _items[i];

      var minx = 9999;
      var miny = 9999;
      var maxx = -1;
      var maxy = -1;

      if (_mode == MODE_GROUP) {
        if (!isNumber(item.mode) && !isArrow(item.mode)) {
          continue;
        }
      }

      if (isLine(item.mode) || isMemo(item.mode) || isArrow(item.mode) || item.mode == TYPE_AREA || item.mode == TYPE_LENGTH2 || item.mode == TYPE_ARROWCHAIN) {
        for (var j = 0; j < item.points.length; j++) {
          var point = item.points[j];

          if (point.x > maxx)
            maxx = point.x;

          if (point.x < minx)
            minx = point.x;

          if (point.y > maxy)
            maxy = point.y;

          if (point.y < miny)
            miny = point.y;
        }
      } else if (isNumber(item.mode)) {
        minx = item.x - 18;
        miny = item.y - 18;
        maxx = item.x;
        maxy = item.y;
      } else {
        minx = item.x;
        miny = item.y;
        maxx = item.x + 20;
        maxy = item.y + 20;
      }

      //var y2 = y;
      //if (inCheck(minx, miny, maxx, maxy, x, y2) == true) {

      var flag = false;

      if ((x1 >= minx && x1 <= maxx || x2 >= minx && x2 <= maxx || minx >= x1 && minx <= x2 || maxx >= x1 && maxx <= x2) &&
          (y1 >= miny && y1 <= maxy || y2 >= miny && y2 <= maxy || miny >= y1 && miny <= y2 || maxy >= y1 && maxy <= y2)) {
        flag = true;
      }

      if (flag == true) {
        if (_mode == MODE_GROUP) {
          if ($('#chk' + i).is(':checked')) {
            $('#chk' + i).prop("checked", false);
            _items[i].selected = false;
          } else {
            $('#chk' + i).prop("checked", true);
            _items[i].selected = true;
          }

        } else {
          _items[i].selected = true;
        }
      }
    }
  }

  var count = 0;
  var target = 0;
  if (_mode == MODE_SELECT) {
    for (var i = 0; i < _items.length; i++) {
      if (!isNumber(_items[i].mode))
        continue;

      if (_items[i].selected != true)
        continue;

      target = i;
      count++;
    }

    if (count == 1) {
      viewSinglePopup(target);
    } else {
      closeSinglePopup();
    }
  }

  down = 0;
}

function getX(event) {
  var x1 = getTouchX(event);
  x1 += _adjustX;
  x1 += $('#canvasframe').scrollLeft();
  x1 = parseFloat(x1) / parseFloat(_zoom);

  return x1;
}

function getY(event) {
  var y1 = getTouchY(event);
  y1 += _adjustY;
  y1 += $('#canvasframe').scrollTop();
  y1 = parseFloat(y1) / parseFloat(_zoom);

  return y1;
}

var _mouseDown = false;
var _mouseDownX;
var _mouseDowny;

function onMouseDown(event) {
  if (_mode == 0)
    return;

  if (_pinch == true)
    return;

  down = 1;

  var x1 = getX(event);
  var y1 = getY(event);

  closeSinglePopup();

  oldx = x1;
  oldy = y1;

  if (_mode == MODE_SELECT || _mode == MODE_GROUP || _mode == MODE_PAN || _mode == MODE_ZOOM) {
    _x1 = x1;
    _y1 = y1;

    return;
  } else if (_mode == MODE_MOVE) {
    makeUndo();

    return;
  }

  _mouseDown = true;
  _mouseDownX = x1;
  _mouseDownY = y1;
}

function mouseDownProcess() {
  var x1 = _mouseDownX;
  var y1 = _mouseDownY;

  if (isNumber(_type)) {
    var x = x1 + 3;
    var y = y1;

    if (_filetype == 4) {
      if (_items.length > 0 && _items.length % 2 != 0) {
        if (_items[_items.length - 1].mode != _type) {
          _items[_items.length - 1].mode = _type;
        }
      }
    }

    insertItem(_type, [], x, y, getTypeColor(_type), false, 0);

    viewSinglePopup(-1);
    return;
  } else if (_type == TYPE_LENGTH2 || _type == TYPE_ARROWCHAIN) {
    if (_area == false) {
      _area = true;

      insertItem(_type, [], 0, 0, getTypeColor(_type), false, 0);
      _items[_items.length - 1].points.push({x: x1, y: y1});
    } else {
      var point = _items[_items.length - 1].points[0];

      if (x1 >= point.x - 5 && x1 <= point.x + 5 && y1 >= point.y - 5 && y1 <= point.y + 5) {
        _area = false;

        var length = _items[_items.length - 1].points.length;
        _items[_items.length - 1].points[length - 1].x = point.x;
        _items[_items.length - 1].points[length - 1].y = point.y;
        down = 0;

        viewAreaPopup();

        return;
      }
    }
  } else if (isLine(_type) || isArrow(_type) || isMemo(_type)) {
    insertItem(_type, [], 0, 0, getTypeColor(_type), false, 0);
  } else if (_type == TYPE_AREA) {
    if (_area == false) {
      _area = true;

      insertItem(_type, [], 0, 0, getTypeColor(_type), false, 0);
      _items[_items.length - 1].points.push({x: x1, y: y1});
    } else {
      var point = _items[_items.length - 1].points[0];

      if (x1 >= point.x - 5 && x1 <= point.x + 5 && y1 >= point.y - 5 && y1 <= point.y + 5) {
        _area = false;

        var length = _items[_items.length - 1].points.length;
        _items[_items.length - 1].points[length - 1].x = point.x;
        _items[_items.length - 1].points[length - 1].y = point.y;
        down = 0;

        viewAreaPopup();

        return;
      }
    }
  } else {
    var x = x1;
    var y = y1;

    insertItem(_type, [], x, y, getTypeColor(_type), false, 0);

    return;
  }

  if (isCurveLine(_type) || isMemo(_type)) {
    _items[_items.length - 1].points.push({x: x1, y: y1});
    modified();
  } else if (isStraightLine(_type)) {
    _items[_items.length - 1].points.push({x: x1, y: y1});
    _items[_items.length - 1].points.push({x: x1, y: y1});
    modified();
  } else if (_type == TYPE_ARROWCHAIN) {
    _items[_items.length - 1].points.push({x: x1, y: y1});
  } else if (isArrow(_type)) {
    _items[_items.length - 1].points.push({x: x1, y: y1});
    _items[_items.length - 1].points.push({x: x1, y: y1});
    modified();
  } else if (_type == TYPE_AREA) {
    _items[_items.length - 1].points.push({x: x1, y: y1});
  } else if (_type == TYPE_LENGTH2) {
    _items[_items.length - 1].points.push({x: x1, y: y1});
  }
}

var _lastTouchX = 0;
var _lastTouchY = 0;

function getTouchX(event) {
  if (event.changedTouches == undefined || event.changedTouches == null) {
    if (event.type == 'touchmove' || event.type == 'touchstart' || event.type == 'touchend') {
      if (event.touches.length == 0)
        return _x1;
      else {
        if (event.type == 'touchstart') {
          return event.touches[0].clientX;
        } else {
          return event.touches[event.touches.length - 1].clientX;
        }
      }
    }    
  } else {
    if (event.type == 'touchmove' || event.type == 'touchstart' || event.type == 'touchend') {
      if (event.changedTouches.length == 0)
        return _x1;
      else {
        if (event.type == 'touchstart') {
          return event.changedTouches[0].clientX;
        } else {
          return event.changedTouches[event.changedTouches.length - 1].clientX;
        }
      }
    }    
  }
  
  return event.x;
}

function getTouchY(event) {
  if (event.changedTouches == undefined || event.changedTouches == null) {
    if (event.type == 'touchmove' || event.type == 'touchstart' || event.type == 'touchend') {
      if (event.touches.length == 0)
        return _y1;
      else {
        if (event.type == 'touchstart') {
          return event.touches[0].clientY;
        } else {
          return event.touches[event.touches.length - 1].clientY;
        }
      }
    }
  } else {
    if (event.type == 'touchmove' || event.type == 'touchstart' || event.type == 'touchend') {
      if (event.changedTouches.length == 0)
        return _y1;
      else {
        if (event.type == 'touchstart') {
          return event.changedTouches[0].clientY;
        } else {
          return event.changedTouches[event.changedTouches.length - 1].clientY;
        }
      }
    }
  }
  
  return event.y;
}

function onMouseMove(event) {
  if (_mode == 0)
    return;

  if (_pinch == true)
    return;

  if (_mouseDown == true) {
    _mouseDown = false;
    mouseDownProcess();
  }

  var x1 = getX(event);
  var y1 = getY(event);

  _x1 = x1;
  _y1 = y1;

  if (_type == TYPE_AREA && _area == true) {
    var length = _items[_items.length - 1].points.length;
    _items[_items.length - 1].points[length - 1].x = x1;
    _items[_items.length - 1].points[length - 1].y = y1;
  }

  if (_type == TYPE_LENGTH2 && _area == true) {
    var length = _items[_items.length - 1].points.length;
    _items[_items.length - 1].points[length - 1].x = x1;
    _items[_items.length - 1].points[length - 1].y = y1;
  }

  if (_type == TYPE_ARROWCHAIN && _area == true) {
    var length = _items[_items.length - 1].points.length;
    _items[_items.length - 1].points[length - 1].x = x1;
    _items[_items.length - 1].points[length - 1].y = y1;
  }

  if (down == 0)
    return;

  if (_mode == MODE_PAN) {
    /*
    var dx = oldx - _x1;
    var dy = oldy - _y1;

    container.scrollLeft += dx * _zoom;
    container.scrollTop += dy * _zoom;

    oldx = _x1;
    oldy = _y1;
    */

    return;
  }

  if (_mode == MODE_SELECT || _mode == MODE_GROUP|| _mode == MODE_ZOOM) {
    return;
  }

  if (_mode == MODE_MOVE) {
    var diffx = oldx - x1;
    var diffy = oldy - y1;

    for (var i = 0; i < _items.length; i++) {
      if (_items[i].selected != true)
        continue;

      var m = _items[i].mode;

      if (isLine(m) || isMemo(m) || isArrow(m)) {
        for (var j = 0; j < _items[i].points.length; j++) {
          _items[i].points[j].x -= diffx;
          _items[i].points[j].y -= diffy;
        }


      } else {
        _items[i].x -= diffx;
        _items[i].y -= diffy;
      }
    }

    oldx = x1;
    oldy = y1;

    modified();
    return;
  }

  if (isCurveLine(_type) || isMemo(_type)) {
    oldx = x1;
    oldy = y1;
    _items[_items.length - 1].points.push({x: x1, y: y1});
    modified();
  } else if (isStraightLine(_type)) {
    var length = _items[_items.length - 1].points.length;
    _items[_items.length - 1].points[length - 1].x = x1;
    _items[_items.length - 1].points[length - 1].y = y1;
    modified();
  } else if (isArrow(_type)) {
    var length = _items[_items.length - 1].points.length;
    _items[_items.length - 1].points[length - 1].x = x1;
    _items[_items.length - 1].points[length - 1].y = y1;
    modified();
  }
};

function selectType(pos) {
  if (pos == 20 || pos == 21) {
    resetType();
  }
  
  if (_mode == MODE_DRAW) {
    if (pos == _type) {
      return;
    }
  } else {
    _type = 0;
  }

  clearArea();

  initUndo();

  _type = pos;

  selectMode(1);

  if (pos == 15 || pos == 16) {
    pos = pos - 14;
    color = '#Ff0000';
  } else if (pos == 51 || pos == 52) {
    pos = pos - 50;
    color = '#000000';
  } else {
    color = '#0000ff';
  }

  _mode = 1;

  ctx.setLineDash([0]);

  if (pos == 13 || pos == 14 || pos == 17 || pos == 18 || pos == 19 || pos == TYPE_LENGTH || pos == TYPE_AREA || pos == TYPE_LENGTH2 || pos == TYPE_MEMO || pos == TYPE_NUMBER || pos == TYPE_NUMBER2) {
    //$('#msCombo').val('');
  }

  if (pos < 13 || pos == 15 || pos == 16 || pos == 51 || pos == 52 || pos == 53 || pos || 54 || pos > 80) {
  } else {
    resetType();
  }
}

function clearArea() {
  if (_area == true) {
    _items.splice(_items.length - 1, 1);
    _area = false;
  }
}

function deleteItemOnlyImage() {
  closeSinglePopup();

  makeUndo();

  if (_filetype == 4) {
    for (var i = 0; i < _items.length; i++) {
      if (_items[i].selected == true) {
        if (i % 2 == 0) {
          if (i < _items.length - 1) {
            _items[i + 1].selected = true;
          }
        } else {
          _items[i - 1].selected = true;
        }
      }
    }
  }

  var flag = false;
  do {
    flag = false;
    for (var i = 0; i < _items.length; i++) {
      if (_items[i].selected == true && _items[i].type != TYPE_NUMBER && _items[i].type != TYPE_NUMBER2) {
        _items.splice(i, 1);
        flag = true;
        break;
      }
    }
  } while (flag == true);

  modified();

  initUndo();

  makeGroupNumber();

  if (_filetype == 6) {
    viewAreaPopup();
  }
}

function deleteItem() {
  closeSinglePopup();

  makeUndo();

  if (_filetype == 4) {
    for (var i = 0; i < _items.length; i++) {
      if (_items[i].selected == true) {
        if (i % 2 == 0) {
          if (i < _items.length - 1) {
            _items[i + 1].selected = true;
          }
        } else {
          _items[i - 1].selected = true;
        }
      }
    }
  }

  var flag = false;
  do {
    flag = false;
    for (var i = 0; i < _items.length; i++) {
      if (_items[i].selected == true) {
        _items.splice(i, 1);
        flag = true;
        break;
      }
    }
  } while (flag == true);

  modified();

  initUndo();

  makeGroupNumber();

  if (_filetype == 6) {
    viewAreaPopup();
  }
}

function resizeCanvas(flag) {
  var w = parseInt(_imgWidth * _zoom);
  var h = parseInt(_imgHeight * _zoom);

  if (flag == true) {
    $('#canvas').css('width', w + 'px');
    $('#canvas').css('height', h + 'px');
  }

  canvas.width = w;
  canvas.height = h;

  lastScale = _zoom;
  curWidth = _imgWidth*_zoom;
  curHeight = _imgHeight*_zoom;
}

function zoomUp() {
  _zoom += 0.5;

  if (_zoom > 10.0)
    _zoom = 10.0;
  
  updateLastScale();
  resizeCanvas(true);
}

function zoomDown() {
  if (_zoom <= 0.5)
    return;

  _zoom -= 0.5;
  updateLastScale();
  resizeCanvas(true);
}

function imagezoomUp() {
  _imageZoom += 1.0;

  if (_imageZoom > 20.0)
    _imageZoom = 20.0;

  updateLastScale();
  resizeCanvas(true);
}

function imagezoomDown() {
  if (_imageZoom <= 1.0)
    return;

  _imageZoom -= 1.0;

  updateLastScale();
  resizeCanvas(true);
}

function arrowzoomUp() {
  _arrowZoom += 1.0;

  if (_arrowZoom > 20.0)
    _arrowZoom = 20.0;

  updateLastScale();
  resizeCanvas(true);
}

function arrowzoomDown() {
  if (_arrowZoom <= 1.0)
    return;

  _arrowZoom -= 1.0;

  updateLastScale();
  resizeCanvas(true);
}

function zoom(value) {
  _zoom = value;
}

function resize() {
  var w = $(window).width() - 40;
  var h = $(window).height() - 140;

  /*
    $('body').width(w + 'px');
    $('body').height(h + 'px');
  */

  _w = w;
  _h = h;

  $('#canvasframe').width(w + 'px');
  $('#canvasframe').height(h + 'px');

  w = w * _zoom;
  h = h * _zoom;

  w -= 18;
  h -= 21;

  $('#canvas').width(w + 'px');
  $('#canvas').height(h + 'px');

  canvas.width = w;
  canvas.height = h;


  w = 500;
  h = 400;
  $('#txtMemo').width(w + 'px');
  $('#txtMemo').height(h + 'px');
}

function insertItem2(points) {
  insertItem(52, points, 0, 0, getTypeColor(52), false, 0);
  _items[_items.length - 1].mode = 52;
  _items[_items.length - 1].type = 52;
}

function clickDrawFormat(pos) {
  if (pos == 1) {
    _items = [];
    insertItem2([{"x":252,"y":10},{"x":177,"y":100}]);
    insertItem2([{"x":179,"y":96},{"x":179,"y":96}]);
    insertItem2([{"x":176,"y":99},{"x":-3,"y":100}]);
    insertItem2([{"x":384,"y":264},{"x":384,"y":264}]);
    insertItem2([{"x":375,"y":265},{"x":595,"y":266}]);
    insertItem2([{"x":595,"y":266},{"x":595,"y":266}]);
    insertItem2([{"x":377,"y":264},{"x":260,"y":398}]);
  } else if (pos == 2) {
    _items = [];
    insertItem2([{"x":242,"y":193},{"x":242,"y":193}]);
    insertItem2([{"x":-1,"y":254},{"x":249,"y":251}]);
    insertItem2([{"x":249,"y":251},{"x":375,"y":390}]);
    insertItem2([{"x":377,"y":85},{"x":597,"y":86}]);
    insertItem2([{"x":597,"y":86},{"x":597,"y":86}]);
    insertItem2([{"x":376,"y":89},{"x":376,"y":89}]);
    insertItem2([{"x":376,"y":84},{"x":376,"y":153}]);
    insertItem2([{"x":376,"y":153},{"x":593 ,"y":153}]);
  }
}
