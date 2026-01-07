import 'dart:async';
import 'dart:io';
import 'dart:ui' as ui;

import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart' show ByteData, rootBundle;
import 'package:gesture_x_detector/gesture_x_detector.dart';
import 'package:get/get.dart';
import 'package:image_size_getter/file_input.dart';
import 'package:image_size_getter/image_size_getter.dart';
import 'package:periodic/components/painter/painter_controller.dart';
import 'package:periodic/components/painter/painter_drawer.dart';
import 'package:periodic/components/painter/widgets/data_box.dart';
import 'package:periodic/components/painter/widgets/fiber_box.dart';
import 'package:periodic/components/painter/widgets/group_box.dart';
import 'package:periodic/components/painter/widgets/inclination_box.dart';
import 'package:periodic/components/painter/widgets/material_box.dart';
import 'package:periodic/components/painter/widgets/tool_box.dart';
import 'package:periodic/config/config.dart' as config;
import 'package:periodic/controllers/auth_controller.dart';
import 'package:periodic/controllers/blueprint_controller.dart';
import 'package:periodic/models/blueprint.dart';

export 'painter_controller.dart';

class Painter extends StatelessWidget {
  final PainterController c = Get.put(PainterController());

  final GlobalKey _widgetKey = GlobalKey();
  final String filename;

  final authController = Get.find<AuthController>();

  Painter({super.key, required this.filename});

  Future<ui.Size> getImageSizeIOS() {
    Completer<ui.Size> completer = Completer();
    Image image = Image.asset(c.blueprint.offlinefilename);
    image.image.resolve(const ImageConfiguration()).addListener(
      ImageStreamListener(
        (ImageInfo imageinfo, bool synchronousCall) {
          var myImage = imageinfo.image;
          ui.Size size =
              ui.Size(myImage.width.toDouble(), myImage.height.toDouble());
          completer.complete(size);
          c.updateCanvas();
        },
      ),
    );
    return completer.future;
  }

  Future<ui.Size> getImageSize() async {
    if (Platform.isAndroid) {
      try {
        final file = File(c.blueprint.offlinefilename);
        final size = ImageSizeGetter.getSize(FileInput(file));

        return ui.Size(size.width.toDouble(), size.height.toDouble());
      } catch (e) {
        return const ui.Size(0, 0);
      }
    } else {
      return getImageSizeIOS();
    }
  }

  Future<ui.Image> _loadImage(String imageAssetPath) async {
    ByteData data;

    if (Platform.isAndroid) {
      final file = File(imageAssetPath);
      final bytes = file.readAsBytesSync();
      data = bytes.buffer.asByteData();
    } else {
      data = await rootBundle.load(imageAssetPath);
    }
    final codec = await ui.instantiateImageCodec(
      data.buffer.asUint8List(),
      targetWidth: c.imageWidth.toInt(),
      targetHeight: c.imageHeight.toInt(),
    );
    var frame = await codec.getNextFrame();
    c.image = frame.image;
    c.isLoaded = true;

    c.updateCanvas();

    return frame.image;
  }

  void _getWidgetInfo(_) {
    final RenderBox renderBox =
        _widgetKey.currentContext?.findRenderObject() as RenderBox;
    _widgetKey.currentContext?.size;

    final ui.Size size = renderBox.size;

    c.width = size.width;
    c.height = size.height;

    c.initZoom();
  }

  clickDrawMode() {
    c.setMode(Mode.draw);
    c.clearSelection();

    c.groupbox = false;
  }

  clickZoomMode() {
    c.setMode(Mode.zoom);
    c.clearSelection();

    c.groupbox = false;
  }

  clickRelocationMode() {
    c.setMode(Mode.move);
  }

  clickRelocationNumberMode() {
    c.setMode(Mode.moveNumber);
  }

  checkSelected() {
    if (c.mode == Mode.select ||
        c.mode == Mode.selectEnd ||
        c.mode == Mode.move ||
        c.mode == Mode.moveNumber) {
      for (var i = 0; i < c.points.length; i++) {
        var item = c.points[i];

        if (item.selected == true) {
          return true;
        }
      }
    }

    return false;
  }

  relocationButton() {
    Function()? onPressed;

    if (checkSelected() == true) {
      onPressed = clickRelocationMode;
    }

    return IconButton(
        color: c.mode == Mode.move ? config.selectColor : Colors.black,
        icon: const Icon(CupertinoIcons.move, size: 30.0),
        onPressed: onPressed);
  }

  relocationNumberButton() {
    Function()? onPressed;

    if (checkSelected() == true) {
      onPressed = clickRelocationNumberMode;
    }

    return IconButton(
        color: c.mode == Mode.moveNumber ? config.selectColor : Colors.black,
        icon: const Icon(CupertinoIcons.slash_circle, size: 30.0),
        onPressed: onPressed);
  }

  deleteButton() {
    Function()? onPressed;

    if (checkSelected() == true) {
      onPressed = c.deleteSelection;
    }

    return IconButton(
        icon: const Icon(CupertinoIcons.trash, size: 30.0),
        onPressed: onPressed);
  }

  deleteWithoutNumberButton() {
    Function()? onPressed;

    if (checkSelected() == true) {
      onPressed = c.deleteSelectionWithoutNumber;
    }

    return IconButton(
        icon: const Icon(Icons.cut_outlined, size: 30.0), onPressed: onPressed);
  }

  undoButton() {
    Function()? onPressed;

    if (c.isUndo()) {
      onPressed = c.undo;
    }

    return IconButton(
        icon: const Icon(Icons.undo, size: 30.0), onPressed: onPressed);
  }

  redoButton() {
    Function()? onPressed;

    if (c.isRedo()) {
      onPressed = c.redo;
    }

    return IconButton(
        icon: const Icon(Icons.redo, size: 30.0), onPressed: onPressed);
  }

  Widget _menu(context) {
    return Row(children: [
      const SizedBox(width: 10.0),
      IconButton(
          icon: Icon(Icons.palette_outlined,
              size: 30.0,
              color: c.mode == Mode.draw ? config.selectColor : Colors.black),
          onPressed: () => clickDrawMode()),
      IconButton(
          icon: Icon(Icons.zoom_in,
              size: 30.0,
              color: c.mode == Mode.zoom ? config.selectColor : Colors.black),
          onPressed: () => clickZoomMode()),
      IconButton(
          icon: Icon(Icons.highlight_alt,
              size: 30.0,
              color: c.mode == Mode.select || c.mode == Mode.selectEnd
                  ? config.selectColor
                  : Colors.black),
          onPressed: () => c.setMode(Mode.select)),
      relocationButton(),
      relocationNumberButton(),
      deleteButton(),
      deleteWithoutNumberButton(),
      IconButton(
          icon: const Icon(CupertinoIcons.add, size: 30.0),
          onPressed: () => c.setZoom(c.zoom + 0.1)),
      IconButton(
          icon: const Icon(CupertinoIcons.minus, size: 30.0),
          onPressed: () => c.setZoom(c.zoom - 0.1)),
      undoButton(),
      redoButton(),
      Expanded(
          child: Row(mainAxisAlignment: MainAxisAlignment.end, children: [
        Expanded(child: Container()),
        TextButton(
            onPressed: () => clickPrevFloor(context),
            child: const Text('이전 층',
                style: TextStyle(
                    fontSize: 15,
                    fontWeight: FontWeight.bold,
                    color: Colors.black))),
        TextButton(
            onPressed: () => clickNextFloor(context),
            child: const Text('다음 층',
                style: TextStyle(
                    fontSize: 15,
                    fontWeight: FontWeight.bold,
                    color: Colors.black))),
        const SizedBox(width: 50),
        c.iconset == 1
            ? TextButton(
                onPressed: () => c.groupbox = !c.groupbox,
                child: const Text('그룹',
                    style: TextStyle(
                        fontSize: 15,
                        fontWeight: FontWeight.bold,
                        color: Colors.black)))
            : const SizedBox.shrink(),
        const SizedBox(width: 10),
        TextButton(
            onPressed: () => clickDetail(),
            child: const Text('상세',
                style: TextStyle(
                    fontSize: 15,
                    fontWeight: FontWeight.bold,
                    color: Colors.black))),
        Expanded(child: Container()),
        Row(children: [
          TextButton(
              style: TextButton.styleFrom(
                padding: EdgeInsets.zero,
              ),
              onPressed: () => c.setIconset(1),
              child: Text('결함도',
                  style: TextStyle(
                      fontSize: 15,
                      fontWeight: FontWeight.bold,
                      color: c.iconset == 1 ? Colors.red : Colors.black))),
          TextButton(
              style: TextButton.styleFrom(
                padding: EdgeInsets.zero,
              ),
              onPressed: () => c.setIconset(2),
              child: Text('기울기',
                  style: TextStyle(
                      fontSize: 15,
                      fontWeight: FontWeight.bold,
                      color: c.iconset == 2 ? Colors.red : Colors.black))),
          TextButton(
              style: TextButton.styleFrom(
                padding: EdgeInsets.zero,
              ),
              onPressed: () => c.setIconset(3),
              child: Text('강도/탄산화',
                  style: TextStyle(
                      fontSize: 15,
                      fontWeight: FontWeight.bold,
                      color: c.iconset == 3 ? Colors.red : Colors.black))),
          TextButton(
              style: TextButton.styleFrom(
                padding: EdgeInsets.zero,
              ),
              onPressed: () => c.setIconset(4),
              child: Text('부재',
                  style: TextStyle(
                      fontSize: 15,
                      fontWeight: FontWeight.bold,
                      color: c.iconset == 4 ? Colors.red : Colors.black))),
        ]),
      ])),
    ]);
  }

  @override
  Widget build(BuildContext context) {
    WidgetsBinding.instance.addPostFrameCallback(_getWidgetInfo);

    c.initEvent();

    getImageSize().then((size) {
      c.imageWidth = size.width;
      c.imageHeight = size.height;

      c.initZoom();

      //_loadImage('assets/imgs/1.png').then((img) {});
      _loadImage(filename).then((img) {});
    });

    c.width = Get.width;
    c.height = Get.height;
    c.toolboxPositionInit();

    return Column(children: [
      Obx(() => _menu(context)),
      Expanded(
        key: _widgetKey,
        child: Stack(children: [
          StreamBuilder(
              stream: c.events.stream,
              builder: (context, snapshot) {
                if (snapshot.hasData) {
                  return RepaintBoundary(
                      child: ClipRect(
                          child: Obx(
                    () => Container(
                        color: Colors.transparent,
                        width: c.width,
                        height: c.height,
                        child: CustomPaint(
                          painter: PainterDrawer(),
                        )),
                  )));
                }

                return Container();
              }),
          XGestureDetector(
            longPressTimeConsider: 500,
            bypassTapEventOnDoubleTap: false,
            onDoubleTap: onDoubleTap,
            // onLongPress: onLongPress,
            onMoveStart: onMoveStart,
            onMoveEnd: onMoveEnd,
            onMoveUpdate: onMoveUpdate,
            onScaleStart: onScaleStart,
            onScaleEnd: onScaleEnd,
            onScaleUpdate: onScaleUpdate,
            onTap: onTap,
            child: Container(
              color: Colors.transparent,
            ),
          ),
          ToolBox(),
          GroupBox(),
          DataBox(),
          InclinationBox(),
          FiberBox(),
          MaterialBox(),
        ]),
      )
    ]);
  }

  void onScaleEnd() {}

  void onScaleUpdate(ScaleEvent event) {
    var diff = event.scale;
    var newZoom = c.zoom * diff;

    c.setZoom(newZoom);
  }

  void onScaleStart(event) {
    c.startSx = event.dx;
    c.startSy = event.dy;

    var zoom = c.zoom;
    c.perX = (c.sx + c.startSx) / (c.imageWidth * zoom);
    c.perY = (c.sy + c.startSy) / (c.imageHeight * zoom);

    c.updateCanvas();
  }

  void onMoveUpdate(MoveEvent event) {
    if (c.mode == Mode.draw) {
      onDrawUpdate(event);
    } else if (c.mode == Mode.zoom) {
      onZoomUpdate(event);
    } else if (c.mode == Mode.move) {
      onRelocationUpdate(event);
    } else if (c.mode == Mode.moveNumber) {
      onRelocationNumberUpdate(event);
    } else {
      onSelectUpdate(event);
    }

    c.updateCanvas();
  }

  void onMoveEnd(MoveEvent event) {
    if (c.mode == Mode.draw) {
      onDrawEnd(event);
    } else if (c.mode == Mode.zoom) {
      onZoomEnd(event);
    } else if (c.mode == Mode.move) {
      onRelocationEnd(event);
    } else if (c.mode == Mode.moveNumber) {
      onRelocationNumberEnd(event);
    } else {
      onSelectEnd(event);
    }

    c.updateCanvas();
  }

  void onMoveStart(MoveEvent event) {
    if (c.mode == Mode.draw) {
      onDrawStart(event);
    } else if (c.mode == Mode.zoom) {
      onZoomStart(event);
    } else if (c.mode == Mode.move) {
      onRelocationStart(event);
    } else if (c.mode == Mode.moveNumber) {
      onRelocationNumberStart(event);
    } else {
      onSelectStart(event);
    }

    c.updateCanvas();
  }

  void onRelocationStart(MoveEvent event) {
    c.startSx = event.localPos.dx;
    c.startSy = event.localPos.dy;
  }

  void onRelocationUpdate(MoveEvent event) {
    c.endSx = event.localPos.dx;
    c.endSy = event.localPos.dy;

    var zoom = c.zoom;
    var dx = (c.endSx - c.startSx) / zoom;
    var dy = (c.endSy - c.startSy) / zoom;

    for (var i = 0; i < c.points.length; i++) {
      var item = c.points[i];

      if (item.selected == false) {
        continue;
      }

      for (var j = 0; j < item.items.length; j++) {
        var item2 = c.points[i].items[j];

        c.points[i].items[j] = Offset(item2.dx + dx, item2.dy + dy);
      }
    }

    c.startSx = c.endSx;
    c.startSy = c.endSy;

    c.updatePoints();
  }

  void onRelocationEnd(MoveEvent event) {
    onRelocationUpdate(event);
    c.modified = true;
  }

  void onRelocationNumberStart(MoveEvent event) {
    c.startSx = event.localPos.dx;
    c.startSy = event.localPos.dy;
  }

  void onRelocationNumberUpdate(MoveEvent event) {
    c.endSx = event.localPos.dx;
    c.endSy = event.localPos.dy;

    var zoom = c.zoom;
    var dx = (c.endSx - c.startSx) / zoom;
    var dy = (c.endSy - c.startSy) / zoom;

    for (var i = 0; i < c.points.length; i++) {
      var item = c.points[i];

      if (item.selected == false) {
        continue;
      }

      var j = item.items.length - 1;
      var item2 = c.points[i].items[j];

      c.points[i].items[j] = Offset(item2.dx + dx, item2.dy + dy);
    }

    c.startSx = c.endSx;
    c.startSy = c.endSy;

    c.updatePoints();
  }

  void onRelocationNumberEnd(MoveEvent event) {
    onRelocationNumberUpdate(event);
    c.modified = true;
  }

  void onDrawStart(MoveEvent event) {
    if (c.type == DrawType.multiline) {
      return;
    }

    var x = event.localPos.dx / c.zoom + c.sx;
    var y = event.localPos.dy / c.zoom + c.sy;
    c.startPoint(Offset(x, y));

    if (c.type == DrawType.line || c.type == DrawType.numberLine) {
      c.addPoint(Offset(x, y));
    }

    DateTime time = DateTime.now();
    final current = time.millisecondsSinceEpoch;
    c.lastUpdateTime = current;
  }

  void onDrawUpdate(MoveEvent event) {
    if (c.type == DrawType.multiline) {
      return;
    }

    var x = event.localPos.dx / c.zoom + c.sx;
    var y = event.localPos.dy / c.zoom + c.sy;

    if (c.index == inclinationLine ||
        c.index == basicVerticalLine ||
        c.index == basicHorizontalLine ||
        c.index == basicVerticalBreak ||
        c.index == basicHorizontalBreak) {
      DateTime time = DateTime.now();
      final current = time.millisecondsSinceEpoch;
      if (c.lastUpdateTime != 0) {
        final diff = current - c.lastUpdateTime;
        if (diff > 2000) {
          c.addPoint(Offset(x, y));
          c.lastUpdateTime = current;
          return;
        }
      }

      if (c.points[c.points.length - 1].items.length >= 2) {
        var dx = c.points[c.points.length - 1]
                .items[c.points[c.points.length - 1].items.length - 1].dx -
            x;
        var dy = c.points[c.points.length - 1]
                .items[c.points[c.points.length - 1].items.length - 1].dy -
            y;

        if (dx.abs() > 10 * c.zoom || dy.abs() > 10 * c.zoom) {
          c.lastUpdateTime = current;
        }
      } else {
        c.lastUpdateTime = current;
      }
    }

    if (c.type == DrawType.line || c.type == DrawType.numberLine) {
      c.movePoint(Offset(x, y));
    } else if (c.type == DrawType.curve) {
      c.addPoint(Offset(x, y));
    }
  }

  void onDrawEnd(MoveEvent event) {
    var x = event.localPos.dx / c.zoom + c.sx;
    var y = event.localPos.dy / c.zoom + c.sy;

    if (c.type == DrawType.line || c.type == DrawType.numberLine) {
      c.movePoint(Offset(x, y));
    } else if (c.type == DrawType.multiline) {
      onTap(event);
    }

    c.modified = true;
  }

  void onZoomStart(MoveEvent event) {
    c.startSx = event.localPos.dx;
    c.startSy = event.localPos.dy;
  }

  void onZoomUpdate(MoveEvent event) {
    moveProcess(event);
  }

  void onZoomEnd(MoveEvent event) {
    moveProcess(event);

    c.sx = c.sx + c.currentSx;
    c.sy = c.sy + c.currentSy;
    c.currentSx = 0.0;
    c.currentSy = 0.0;
  }

  void onSelectStart(MoveEvent event) {
    c.startSx = event.localPos.dx;
    c.startSy = event.localPos.dy;

    c.setMode(Mode.select);
  }

  void onSelectUpdate(MoveEvent event) {
    c.endSx = event.localPos.dx;
    c.endSy = event.localPos.dy;
  }

  void onSelectEnd(MoveEvent event) {
    var zoom = c.zoom;
    var dx = -1 * (c.sx + c.currentSx) * zoom;
    var dy = -1 * (c.sy + c.currentSy) * zoom;

    var sx = c.startSx;
    var sy = c.startSy;
    var ex = c.endSx;
    var ey = c.endSy;

    if (sx > ex) {
      var temp = sx;
      sx = ex;
      ex = temp;
    }

    if (sy > ey) {
      var temp = sy;
      sy = ey;
      ey = temp;
    }

    for (var i = 0; i < c.points.length; i++) {
      var points = c.points[i];
      for (var j = 0; j < points.items.length; j++) {
        var point = points.items[j];
        var x = point.dx * zoom + dx;
        var y = point.dy * zoom + dy;

        if (c.iconset == 1) {
          if (points.icon >= 200) {
            continue;
          }
        } else if (c.iconset == 2) {
          if (points.icon < 200 || points.icon >= 300) {
            continue;
          }
        } else if (c.iconset == 3) {
          if (points.icon < 300 || points.icon >= 400) {
            continue;
          }
        } else if (c.iconset == 4) {
          if (points.icon < 400 || points.icon >= 500) {
            continue;
          }
        }

        if (points.type == DrawType.number ||
            points.type == DrawType.material ||
            points.type == DrawType.icon) {
          var step = (c.iconZoom / 2.0) * zoom;

          var csx = x - step;
          var cex = x + step;
          var csy = y - step;
          var cey = y + step;

          if ((((csx >= sx && csx <= ex) || (cex >= sx && cex <= ex)) &&
                  ((csy >= sy && csy <= ey) || (cey >= sy && cey <= ey))) ||
              (((sx >= csx && sx <= cex) || (ex >= csx && ex <= cex)) &&
                  ((sy >= csy && sy <= cey) || (ey >= csy && ey <= cey)))) {
            c.points[i].selected = true;
            c.points[i].grouped = true;
            break;
          } else {
            c.points[i].selected = false;
            c.points[i].grouped = false;
          }
        } else {
          if (x >= sx && x <= ex && y >= sy && y <= ey) {
            c.points[i].selected = true;
            c.points[i].grouped = true;
            break;
          } else {
            c.points[i].selected = false;
            c.points[i].grouped = false;
          }
        }
      }
    }

    c.setMode(Mode.selectEnd);

    var count = 0;
    var position = 0;
    var index = 0;
    for (var i = 0; i < c.points.length; i++) {
      var item = c.points[i];

      if (item.type != DrawType.number &&
          item.type != DrawType.numberLine &&
          item.icon < 200) {
        continue;
      }

      if (item.selected == true) {
        position = i;
        count++;
        index = item.icon;
      }
    }

    if (count == 1) {
      if (index >= 200 && index < 300) {
        c.viewDataboxInclination(position);
      } else if (index >= 300 && index < 400) {
        c.viewDataboxFiber(position);
      } else if (index >= 400 && index < 500) {
        c.viewDataboxMaterial(position);
      } else {
        c.viewDatabox(position);
      }
    }
  }

  void onLongPress(TapEvent event) {
    if (c.mode != Mode.draw) {
      return;
    }

    if (c.type == DrawType.multiline) {
      var x = event.localPos.dx / c.zoom + c.sx;
      var y = event.localPos.dy / c.zoom + c.sy;

      c.addPoint(Offset(x, y));

      c.lineStart = false;
    } else {
      c.toolboxToggle();
    }

    c.updateCanvas();
  }

  void onDoubleTap(event) {
    if (c.mode == Mode.select || c.mode == Mode.selectEnd) {
      return;
    }

    c.modeToggle();
  }

  void onTap(TapEvent event) {
    if (c.mode == Mode.select || c.mode == Mode.selectEnd) {
      c.clearSelection();
      MoveEvent moveEvent =
          MoveEvent(event.localPos, event.position, event.pointer);
      onSelectStart(moveEvent);
      onSelectUpdate(moveEvent);
      onSelectEnd(moveEvent);
      return;
    }

    if (c.mode != Mode.draw) {
      return;
    }

    if (c.type != DrawType.multiline &&
        c.type != DrawType.icon &&
        c.type != DrawType.number &&
        c.type != DrawType.numberLine &&
        c.type != DrawType.material) {
      return;
    }

    var x = event.localPos.dx / c.zoom + c.sx;
    var y = event.localPos.dy / c.zoom + c.sy;

    if (c.lineStart == false ||
        c.type == DrawType.icon ||
        c.type == DrawType.number ||
        c.type == DrawType.numberLine ||
        c.type == DrawType.material) {
      c.startPoint(Offset(x, y));
      c.lineStart = true;
    } else {
      c.addPoint(Offset(x, y));
    }

    c.updateCanvas();
  }

  moveProcess(MoveEvent event) {
    var zoom = c.zoom;

    var x = (event.localPos.dx - c.startSx) / c.zoom * -1;
    var y = (event.localPos.dy - c.startSy) / c.zoom * -1;

    if (c.imageWidth * zoom < c.width) {
      x = 0;
    } else {
      if (x + c.sx * zoom < 0) {
        x = c.sx * zoom * -1;
      }

      if (x + c.sx * zoom >= c.imageWidth * zoom - c.width) {
        x = (c.imageWidth - c.sx) * zoom - c.width;
      }
    }

    if (c.imageHeight * zoom < c.height) {
      y = 0;
    } else {
      if (y + c.sy * zoom < 0) {
        y = c.sy * zoom * -1;
      }

      if (y + c.sy * zoom >= c.imageHeight * zoom - c.height) {
        y = (c.imageHeight - c.sy) * zoom - c.height;
      }
    }

    c.currentSx = x / zoom;
    c.currentSy = y / zoom;
  }

  clickPrevFloor(context) {
    clickNext(context, 1);
  }

  clickNextFloor(context) async {
    clickNext(context, 2);
  }

  clickNext(context, mode) async {
    if (c.modified == true) {
      showDialog<void>(
        context: context,
        builder: (context2) {
          return AlertDialog(
            title: const Text('층 이동'),
            content: const Text(
                '작업내역이 저장되지 않았습니다.\n저장없이 이동하시겠습니까.\n저장없이 이동 선택시 작업한 내역이 모두 삭제됩니다'),
            actions: <Widget>[
              ElevatedButton(
                child: const Text('닫기'),
                onPressed: () {
                  navigator!.pop(context2);
                },
              ),
              ElevatedButton(
                child: const Text('저장없이 이동'),
                onPressed: () {
                  navigator!.pop(context2);
                  clickNextProcess(mode);
                },
              )
            ],
          );
        },
      );

      return;
    }

    clickNextProcess(mode);
  }

  clickNextProcess(mode) {
    c.databox = false;
    final blueprintController = Get.find<BlueprintController>();

    Blueprint prev = Blueprint();
    Blueprint current = Blueprint();

    var find = false;
    for (var i = 0; i < blueprintController.items.length; i++) {
      final item = blueprintController.items[i];
      if (find == true) {
        if (mode == 2) {
          if (item.upload == 1 && item.offlinefilename != '') {
            current = item;
            break;
          }
        } else {
          current = prev;
          break;
        }
      } else {
        if (item.id == c.blueprint.id) {
          find = true;
          if (mode == 1) {
            current = prev;
            break;
          }
        } else {
          if (item.upload == 1 && item.offlinefilename != '') {
            prev = item;
          }
        }
      }
    }

    if (current.id == 0) {
      return;
    }

    authController.setTitle(current);
    c.blueprint = current;
    c.reset();
    c.load();

    getImageSize().then((size) async {
      c.imageWidth = size.width;
      c.imageHeight = size.height;

      c.initZoom();

      await _loadImage(current.offlinefilename);
    });
  }

  clickDetail() {
    if (c.iconset == 1) {
      Get.toNamed('/data');
    } else if (c.iconset == 2) {
      Get.toNamed('/inclination');
    } else if (c.iconset == 3) {
      Get.toNamed('/fiber');
    } else if (c.iconset == 4) {
      Get.toNamed('/material');
    }
  }
}
