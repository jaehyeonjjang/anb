import 'dart:async';
import 'dart:convert';
import 'dart:io';
import 'dart:ui' as ui;

import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:localstorage/localstorage.dart';
import 'package:patrol/controllers/auth_controller.dart';
import 'package:patrol/controllers/blueprint_controller.dart';
import 'package:patrol/models/blueprint.dart';

class Point {
  List<Offset> items = <Offset>[];

  int width;
  LineColor color;
  DrawType type;
  int icon;
  int number;
  bool selected = false;
  bool grouped = false;

  String part = '';
  String member = '';
  String shape = '';
  String weight = '';
  String length = '';
  String count = '';
  String progress = '';
  String remark = '';
  int order = 0;
  List<String> images = <String>[];
  List<String> onlineimages = <String>[];

  Point(
      {required this.items,
      required this.color,
      required this.width,
      required this.type,
      required this.icon,
      required this.number,
      required this.part,
      required this.member,
      required this.shape,
      required this.weight,
      required this.length,
      required this.count,
      required this.progress,
      required this.remark,
      required this.order,
      required this.images,
      required this.onlineimages});

  factory Point.fromJson(Map<String, dynamic> json) {
    final items = (json['items'] as List<dynamic>)
        .map((item) => Offset(item['dx'] as double, item['dy'] as double))
        .toList();
    return Point(
        items: items,
        width: json['width'] as int,
        color: LineColor.values[(json['color'] as int) - 1],
        type: DrawType.values[(json['type'] as int) - 1],
        icon: json['icon'] as int,
        number: json['number'] as int,
        part: json['part'] as String,
        member: json['member'] as String,
        shape: json['shape'] as String,
        weight: json['weight'] as String,
        length: json['length'] as String,
        count: json['count'] as String,
        progress: json['progress'] as String,
        remark: json['remark'] as String,
        order: json['order'] as int,
        images: (json['images'] as List<dynamic>)
            .map((item) => item as String)
            .toList(),
        onlineimages: (json['onlineimages'] as List<dynamic>)
            .map((item) => item as String)
            .toList());
  }

  Map<String, dynamic> toJson() {
    return {
      'items': items.map((item) => {'dx': item.dx, 'dy': item.dy}).toList(),
      'width': width,
      'color': color.code,
      'type': type.code,
      'icon': icon,
      'number': number,
      'part': part,
      'member': member,
      'shape': shape,
      'weight': weight,
      'length': length,
      'count': count,
      'progress': progress,
      'remark': remark,
      'order': order,
      'images': images,
      'onlineimages': onlineimages
    };
  }

  add(Offset offset) {
    items.add(offset);
  }
}

enum LineColor {
  black(1, '검정'),
  red(2, '빨강'),
  blue(3, '파랑'),
  green(4, '초록'),
  lightblue(5, '하늘');

  final int code;
  final String name;
  const LineColor(this.code, this.name);
}

enum DrawType {
  curve(1, '자유곡선'),
  line(2, '직선'),
  multiline(3, '연속선'),
  icon(4, '아이콘'),
  number(5, '순번');

  final int code;
  final String name;
  const DrawType(this.code, this.name);
}

enum Mode {
  draw(1, '그리기'),
  zoom(2, '줌'),
  select(3, '선택'),
  selectEnd(3, '선택'),
  move(4, '이동');

  final int code;
  final String name;
  const Mode(this.code, this.name);
}

enum UndoType {
  draw(1, '그리기'),
  delete(2, '지우기');

  final int code;
  final String name;
  const UndoType(this.code, this.name);
}

class Undo {
  List<Point> points = <Point>[];
  UndoType type = UndoType.draw;

  Undo({required this.points, required this.type});
}

class PainterController extends GetxController {
  Blueprint blueprint = Blueprint();

  bool _autosave = false;
  final _toolbox = true.obs;
  final _groupbox = false.obs;
  final _databox = false.obs;
  final _points = [].obs;
  final _undos = [].obs;
  final _works = [].obs;
  final _isLoaded = false.obs;

  final _current = (-1).obs;

  final _number = 1.obs;

  final _mode = Rx<Mode>(Mode.draw);
  final _zoom = 1.0.obs;

  final _sx = 0.0.obs;
  final _sy = 0.0.obs;
  final _currentSx = 0.0.obs;
  final _currentSy = 0.0.obs;
  final _startSx = 0.0.obs;
  final _startSy = 0.0.obs;
  final _endSx = 0.0.obs;
  final _endSy = 0.0.obs;

  final _width = 0.0.obs;
  final _height = 0.0.obs;
  final _imageWidth = 0.0.obs;
  final _imageHeight = 0.0.obs;

  final _perX = 0.0.obs;
  final _perY = 0.0.obs;

  final _color = Rx<LineColor>(LineColor.lightblue);
  final _line = 1.obs;
  final _type = Rx<DrawType>(DrawType.number);

  final _index = 1.obs;

  final _lineStart = false.obs;
  final _iconZoom = 100.0.obs;

  final _visibleGroup = false.obs;

  final _modified = false.obs;

  reset() {
    toolbox = true;
    groupbox = false;
    databox = false;
    points = [];
    _undos.clear();
    _undos.refresh();
    _works.clear();
    _works.refresh();

    isLoaded = false;

    current = -1;

    number = 1;

    mode = Mode.draw;
    zoom = 1.0;

    sx = 0.0;
    sy = 0.0;
    currentSx = 0.0;
    currentSy = 0.0;
    startSx = 0.0;
    startSy = 0.0;
    endSx = 0.0;
    endSy = 0.0;

    perX = 0.0;
    perY = 0.0;

    color = LineColor.lightblue;
    line = 1;
    type = DrawType.number;

    index = 1;

    lineStart = false;
    _iconZoom.value = 100.0;

    visibleGroup = false;

    modified = false;
  }

  bool get modified => _modified.value;
  set modified(value) {
    _modified.value = value;

    if (value == true) {
      if (_autosave == true) {
        runSave++;

        //save();
      }
    }
  }

  get points => _points;
  set points(value) => _points.value = value;

  int get current => _current.value;
  set current(value) => _current.value = value;

  int get number => _number.value;
  set number(value) => _number.value = value;

  double get iconZoom {
    return _iconZoom.value;
  }

  set iconZoom(value) {
    _iconZoom.value = value;
    modified = true;
  }

  int get index => _index.value;
  set index(value) => _index.value = value;

  bool get lineStart => _lineStart.value;
  set lineStart(value) => _lineStart.value = value;
  LineColor get color => _color.value;
  set color(value) => _color.value = value;
  int get line => _line.value;
  set line(value) => _line.value = value;
  DrawType get type => _type.value;
  set type(value) => _type.value = value;
  bool get toolbox => _toolbox.value;
  set toolbox(value) => _toolbox.value = value;
  bool get groupbox => _groupbox.value;
  set groupbox(value) => _groupbox.value = value;
  bool get databox => _databox.value;
  set databox(value) => _databox.value = value;
  double get perX => _perX.value;
  set perX(value) => _perX.value = value;
  double get perY => _perY.value;
  set perY(value) => _perY.value = value;
  bool get isLoaded => _isLoaded.value;
  set isLoaded(value) => _isLoaded.value = value;
  Mode get mode => _mode.value;
  set mode(value) => _mode.value = value;
  double get sx => _sx.value;
  set sx(value) => _sx.value = value;
  double get sy => _sy.value;
  set sy(value) => _sy.value = value;
  double get currentSx => _currentSx.value;
  set currentSx(value) => _currentSx.value = value;
  double get currentSy => _currentSy.value;
  set currentSy(value) => _currentSy.value = value;
  double get startSx => _startSx.value;
  set startSx(value) => _startSx.value = value;
  double get startSy => _startSy.value;
  set startSy(value) => _startSy.value = value;
  double get endSx => _endSx.value;
  set endSx(value) => _endSx.value = value;
  double get endSy => _endSy.value;
  set endSy(value) => _endSy.value = value;
  double get zoom => _zoom.value;
  set zoom(value) => _zoom.value = value;
  double get width => _width.value;
  set width(value) => _width.value = value;
  double get height => _height.value;
  set height(value) => _height.value = value;
  double get imageWidth => _imageWidth.value;
  set imageWidth(value) => _imageWidth.value = value;
  double get imageHeight => _imageHeight.value;
  set imageHeight(value) => _imageHeight.value = value;

  bool get visibleGroup => _visibleGroup.value;
  set visibleGroup(value) => _visibleGroup.value = value;

  late ui.Image image;
  var events = StreamController<int>();

  startPoint(Offset offset) {
    var num = 0;
    if (type == DrawType.number) {
      for (var i = 0; i < points.length; i++) {
        if (points[i].number > num) {
          num = points[i].number;
        }
      }

      num++;
    }

    var point = Point(
        items: [],
        color: color,
        width: 1,
        type: type,
        icon: index,
        number: num,
        part: '',
        member: '',
        shape: '',
        weight: '',
        length: '',
        count: '',
        progress: '',
        remark: '',
        order: 0,
        images: [],
        onlineimages: []);

    if (type == DrawType.number) {
      point.weight = '0.2';

      if (index == 1) {
        point.member = '벽체';
      } else {
        point.member = '슬래브';
      }

      point.count = '1';
      if (points.length > 0) {
        point.part = points[points.length - 1].part;
      }
    }
    point.add(offset);
    points.add(point);

    clearUndo();
    _works.add(points.toList());
    _works.refresh();

    if (type == DrawType.number) {
      viewDatabox(points.length - 1);
    } else {
      databox = false;
    }

    if (type != DrawType.curve &&
        type != DrawType.line &&
        type != DrawType.multiline) {
      modified = true;
    }
  }

  addPoint(Offset offset) {
    points[points.length - 1].add(offset);

    clearUndo();

    _works[_works.length - 1] = points.toList();
    _works.refresh();
  }

  movePoint(Offset offset) {
    points[points.length - 1]
        .items[points[points.length - 1].items.length - 1] = offset;

    clearUndo();

    _works[_works.length - 1] = points.toList();
    _works.refresh();
  }

  pointClear() {
    points.clear();

    _works.add(points.toList());
    _works.refresh();

    modified = true;
  }

  initEvent() {
    events.close();
    events = StreamController<int>();
  }

  updateCanvas() {
    events.add(points.length);
  }

  clear() {
    pointClear();
    _works.clear();
    _works.refresh();
    _undos.clear();
    _undos.refresh();

    sx = 0.0;
    sy = 0.0;
    initZoom();
    setIndex(1);
    _iconZoom.value = 100.0;
    setMode(Mode.draw);
  }

  clearUndo() {
    _undos.clear();
    _undos.refresh();
  }

  initZoom() {
    var zoom = 1.0;

    var rate = width / height;
    var imageRate = imageWidth / imageHeight;

    if (rate > imageRate) {
      zoom = height / imageHeight;
    } else {
      zoom = width / imageWidth;
    }

    setZoom(zoom);
  }

  setZoom(zoom) {
    if (imageWidth * zoom < width && imageHeight * zoom < height) {
      var rate = width / height;
      var imageRate = imageWidth / imageHeight;

      if (rate > imageRate) {
        zoom = height / imageHeight;
      } else {
        zoom = width / imageWidth;
      }
    }

    if (zoom > 5.0) {
      zoom = 5.0;
    }

    sx = ((imageWidth * zoom * perX) - startSx) / zoom;
    sy = ((imageHeight * zoom * perY) - startSy) / zoom;

    this.zoom = zoom;

    if (imageWidth * zoom < width) {
      sx = 0.0;
    } else {
      if (sx * zoom >= imageWidth * zoom - width) {
        sx = (imageWidth * zoom - width) / zoom;
      }
    }

    if (imageHeight * zoom < height) {
      sy = 0.0;
    } else {
      if (sy * zoom >= imageHeight * zoom - height) {
        sy = (imageHeight * zoom - height) / zoom;
      }
    }

    if (sx < 0.0) {
      sx = 0.0;
    }

    if (sy < 0.0) {
      sy = 0.0;
    }

    updateCanvas();
  }

  toolboxToggle() {
    toolbox = !toolbox;
  }

  setMode(value) {
    mode = value;

    lineStart = false;

    if (mode == Mode.draw || mode == Mode.zoom || mode == Mode.move) {
      startSx = 0.0;
      startSy = 0.0;
      endSx = 0.0;
      endSy = 0.0;
    }

    if (mode == Mode.selectEnd) {
      _points.refresh();
    }

    if (mode == Mode.draw || mode == Mode.zoom) {
      clearSelection();
      databox = false;
    } else {
      var count = 0;
      for (var i = 0; i < points.length; i++) {
        var item = points[i];

        if (item.type != DrawType.number) {
          continue;
        }

        if (item.selected == true) {
          count++;
        }
      }

      if (count == 1) {
        databox = true;
      } else {
        databox = false;
      }
    }

    updateCanvas();
  }

  modeToggle() {
    if (mode == Mode.draw) {
      mode = Mode.zoom;
    } else {
      mode = Mode.draw;
    }

    lineStart = false;
  }

  setColor(value) {
    color = value;

    lineStart = false;
  }

  setType(value) {
    type = value;

    lineStart = false;
  }

  undo() {
    if (_works.length <= 1) {
      return;
    }

    databox = false;

    var last = _works[_works.length - 1];
    _undos.add(last);
    _works.removeLast();

    _undos.refresh();
    _works.refresh();

    if (_works.isEmpty) {
      points.clear();
    } else {
      points = [..._works[_works.length - 1]];
      modified = true;
    }

    updateCanvas();
  }

  redo() {
    if (_undos.isEmpty) {
      return;
    }

    databox = false;

    var last = _undos[_undos.length - 1];
    _works.add(last);
    points = [...last];
    modified = true;
    _undos.removeLast();

    _works.refresh();
    _undos.refresh();

    updateCanvas();
  }

  setIndex(value) {
    index = value;

    if (value == 1) {
      setColor(LineColor.red);
      setType(DrawType.number);
    } else if (value == 2) {
      setColor(LineColor.blue);
      setType(DrawType.number);
    } else if (value == 31) {
      setColor(LineColor.lightblue);
      setType(DrawType.curve);
    } else if (value == 32) {
      setColor(LineColor.red);
      setType(DrawType.curve);
    } else if (value == 33) {
      setColor(LineColor.green);
      setType(DrawType.curve);
    } else if (value == 41) {
      setColor(LineColor.lightblue);
      setType(DrawType.line);
    } else if (value == 42) {
      setColor(LineColor.red);
      setType(DrawType.line);
    } else if (value == 43) {
      setColor(LineColor.green);
      setType(DrawType.line);
    } else {
      setType(DrawType.icon);
    }
  }

  setIconZoom(zoom) {
    if (zoom <= 0.0) {
      return;
    }

    iconZoom = zoom;

    updateCanvas();
  }

  clearSelection() {
    for (var i = 0; i < points.length; i++) {
      var point = points[i];

      point.selected = false;
      point.grouped = false;
    }

    updateCanvas();
  }

  deleteSelection() {
    var flag = true;
    var find = false;

    while (flag == true) {
      flag = false;

      for (var i = 0; i < points.length; i++) {
        var point = points[i];

        if (point.selected != true) {
          continue;
        }

        points.removeAt(i);
        flag = true;
        find = true;
        break;
      }

      databox = false;
    }

    _works.add(points.toList());
    _works.refresh();

    setMode(Mode.selectEnd);
    groupSort();
    updateCanvas();

    if (find == true) {
      modified = true;
    }
  }

  deleteSelectionWithoutNumber() {
    var flag = true;
    var find = false;

    while (flag == true) {
      flag = false;

      for (var i = 0; i < points.length; i++) {
        var point = points[i];

        if (point.type == DrawType.number) {
          continue;
        }

        if (point.selected != true) {
          continue;
        }

        points.removeAt(i);
        flag = true;
        find = true;

        break;
      }
    }

    databox = false;

    _works.add(points.toList());
    _works.refresh();

    setMode(Mode.selectEnd);
    updateCanvas();

    if (find == true) {
      modified = true;
    }
  }

  groupSort() {
    var flag = true;

    while (flag == true) {
      flag = false;

      for (var i = 0; i < points.length; i++) {
        var find = false;

        for (var j = 0; j < points.length; j++) {
          var item2 = points[j];

          if (item2.number == i + 1) {
            find = true;
            break;
          }
        }

        if (find == true) {
          continue;
        }

        for (var j = 0; j < points.length; j++) {
          var item2 = points[j];

          if (item2.number > i + 1) {
            flag = true;
            points[j].number--;
          }
        }

        break;
      }
    }
  }

  groupApply() {
    var max = 0;

    for (var i = 0; i < points.length; i++) {
      var item = points[i];

      if (item.grouped == false) {
        if (item.number > max) {
          max = item.number;
        }

        continue;
      }

      points[i].number = 0;
    }

    var number = max + 1;

    for (var i = 0; i < points.length; i++) {
      var flag = false;

      for (var j = 0; j < points.length; j++) {
        var item2 = points[j];

        if (item2.number == i + 1) {
          flag = true;
          break;
        }
      }

      if (flag == true) {
        continue;
      }

      number = i + 1;
      break;
    }

    for (var i = 0; i < points.length; i++) {
      var item = points[i];

      if (item.grouped == false) {
        continue;
      }

      points[i].number = number;
    }

    groupSort();

    updatePoints();

    modified = true;
  }

  groupCancel() {
    var max = 0;

    for (var i = 0; i < points.length; i++) {
      var item = points[i];

      if (item.grouped == false) {
        if (item.number > max) {
          max = item.number;
        }

        continue;
      }

      points[i].number = 0;
    }

    for (var i = 0; i < points.length; i++) {
      var item = points[i];

      if (item.grouped == false) {
        continue;
      }

      max++;
      points[i].number = max;
    }

    groupSort();

    updatePoints();
    modified = true;
  }

  updatePoints() {
    points.refresh();
    updateCanvas();
  }

  viewDatabox(int pos) {
    databox = true;
    current = pos;
  }

  setData(int pos, String value) {
    if (pos == 1) {
      points[current].part = value;
    } else if (pos == 2) {
      points[current].member = value;
    } else if (pos == 3) {
      points[current].shape = value;
    } else if (pos == 4) {
      points[current].weight = value;
    } else if (pos == 5) {
      points[current].length = value;
    } else if (pos == 6) {
      points[current].count = value;
    } else if (pos == 7) {
      points[current].progress = value;
    } else if (pos == 8) {
      points[current].remark = value;
    }

    updatePoints();
    modified = true;
  }

  setCurrent(Point value) {
    points[current] = value;

    updatePoints();
  }

  Point getCurrent() {
    if (current == -1) {
      return Point(
          items: [],
          color: LineColor.black,
          icon: 0,
          number: 0,
          type: DrawType.line,
          width: 0,
          part: '',
          member: '',
          shape: '',
          weight: '',
          length: '',
          count: '',
          progress: '',
          remark: '',
          order: 0,
          images: [],
          onlineimages: []);
    }

    return points[current];
  }

  save() async {
    print('save');

    Map<String, dynamic> items = {
      'id': blueprint.id,
      'points': points.map((item) => item.toJson()).toList(),
      'zoom': zoom,
      'iconzoom': iconZoom
    };

    final LocalStorage storage = LocalStorage('periodic.json');
    final str = json.encode(items);

    final ready = await storage.ready;
    if (ready) {
      try {
        await storage.setItem('data_${blueprint.id}', str);
      } catch (e) {
        print('storage error');
        print(e);
      }
    } else {
      print('not ready');
    }

    modified = false;
    final blueprintController = Get.find<BlueprintController>();
    blueprintController.modified = true;

    if (_autosave == false) {
      databox = false;
    }
  }

  void load() async {
    final autoController = Get.find<AuthController>();
    _autosave = autoController.autosave;

    final LocalStorage storage = LocalStorage('periodic.json');

    await storage.ready;
    final str = await storage.getItem('data_${blueprint.id}');

    if (str == null || str == '') {
      modified = false;
      _works.add(points.toList());
      _works.refresh();
      return;
    }

    final j = json.decode(str);

    zoom = j['zoom'];
    _iconZoom.value = j['iconzoom'];

    var items = j['points'].map<Point>((json) => Point.fromJson(json)).toList();

    for (var i = 0; i < items.length; i++) {
      List<String> newImages = [];

      for (var j = 0; j < items[i].images.length; j++) {
        var image = items[i].images[j];
        if (File(image).existsSync() == true) {
          newImages.add(image);
        }
      }

      items[i].images = newImages;
    }

    _points.value = items;
    _works.add(points.toList());
    _works.refresh();
    _points.refresh();
    updateCanvas();
    modified = false;
  }

  removeDataimage(pos) {
    removeDataimageIndex(current, pos);
  }

  removeDataimageIndex(index, pos) {
    Point point = points[index];
    point.images.removeAt(pos);

    updatePoints();
    modified = true;
  }

  removePoint(index) {
    points.removeAt(index);
    groupSort();
    updatePoints();
    modified = true;
  }

  isUndo() {
    if (_works.length > 1) {
      return true;
    } else {
      return false;
    }
  }

  isRedo() {
    if (_undos.isEmpty) {
      return false;
    } else {
      return true;
    }
  }

  final _runSave = 1.obs;
  int get runSave => _runSave.value;
  set runSave(value) => _runSave.value = value;
  @override
  void onInit() async {
    super.onInit();

    debounce(_runSave, (_) {
      save();
    }, time: const Duration(milliseconds: 300));
  }  
}
