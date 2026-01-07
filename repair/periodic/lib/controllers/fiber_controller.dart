import 'dart:convert';

import 'package:common_control/common_control.dart';
import 'package:periodic/components/painter/painter_controller.dart';

class Fiber {
  String floor;
  String sh1;
  String sh2;
  String n;
  String ps;

  String bfloor;
  String bsh;
  String bn;
  String bps;

  Fiber(
      {required this.floor,
      required this.sh1,
      required this.sh2,
      required this.n,
      required this.ps,
      required this.bfloor,
      required this.bsh,
      required this.bn,
      required this.bps});

  factory Fiber.fromJson(Map<String, dynamic> json) {
    return Fiber(
        floor: json['floor'] as String,
        sh1: json['sh1'] as String,
        sh2: json['sh2'] as String,
        n: json['n'] as String,
        ps: json['ps'] as String,
        bfloor: json['bfloor'] as String,
        bsh: json['bsh'] as String,
        bn: json['bn'] as String,
        bps: json['bps'] as String);
  }

  Map<String, dynamic> toJson() => {
        'floor': floor,
        'sh1': sh1,
        'sh2': sh2,
        'n': n,
        'ps': ps,
        'bfloor': bfloor,
        'bsh': bsh,
        'bn': bn,
        'bps': bps
      };
}

class FiberController extends GetxController {
  final _items = <Fiber>[].obs;
  final _x = 0.obs;
  final _y = 0.obs;

  List<Fiber> get items => _items;
  set items(value) => _items.value = value;

  int get x => _x.value;
  set x(int value) => _x.value = value;

  int get y => _y.value;
  set y(int value) => _y.value = value;

  @override
  onInit() {
    super.onInit();

    var pos = -1;
    PainterController c = Get.find<PainterController>();
    for (var i = 0; i < c.points.length; i++) {
      var point = c.points[i];

      if (point.icon != 300) {
        continue;
      }

      pos = i;
      if (point.remark.length < 10) {
        pos = -1;
        break;
      }

      var json = jsonDecode(point.remark);

      var list = json as List<dynamic>;

      List<Fiber> data = [];

      for (var i = 0; i < list.length; i++) {
        final item = Fiber.fromJson(list[i]);
        data.add(item);
      }

      items = data;
      break;
    }

    if (pos == -1) {
      for (var i = 0; i < 10; i++) {
        items.add(Fiber(
            floor: '',
            sh1: '',
            sh2: '',
            n: '',
            ps: '',
            bfloor: '',
            bsh: '',
            bn: '',
            bps: ''));
      }
    }
  }

  @override
  onClose() {
    super.onClose();

    var pos = -1;
    PainterController c = Get.find<PainterController>();
    for (var i = 0; i < c.points.length; i++) {
      var point = c.points[i];

      if (point.icon != 300) {
        continue;
      }

      pos = i;
      break;
    }

    var json = jsonEncode(items.map((item) => item.toJson()).toList());

    if (pos == -1) {
      var point = Point(
          items: [],
          color: LineColor.black,
          width: 1,
          type: DrawType.txt,
          icon: 300,
          number: 0,
          part: '',
          member: '',
          shape: '',
          weight: '',
          length: '',
          count: '',
          progress: '',
          remark: json,
          order: 0,
          images: [],
          onlineimages: []);
      c.points.add(point);
    } else {
      c.points[pos].remark = json;
    }

    c.modified = true;
  }

  movePrev() {
    if (x > 0) {
      x--;
    } else {
      if (y > 0) {
        y--;
        x = 4;
      }
    }
  }

  moveNext() {
    if (x < 4) {
      x++;
    } else {
      if (y < 9) {
        x = 0;
        y++;
      }
    }
  }

  move(dy, dx) {
    y = dy;
    x = dx;
  }

  udpate() {
    _x.refresh();
    _y.refresh();
    _items.refresh();

    refresh();
  }

  makeData() {
    for (var i = 0; i < 10; i++) {
      items[i].bfloor = items[i].floor;

      if (items[i].n == '') {
        items[i].bn = '';
      } else {
        items[i].bn = '${int.parse(items[i].n) + 1}';
      }

      if (items[i].ps == '') {
        items[i].bps = '';
      } else {
        items[i].bps = '${int.parse(items[i].ps) + 1}';
      }

      var sh1 = '';
      var sh2 = '';
      if (items[i].sh1 != '') {
        sh1 = '${int.parse(items[i].sh1) + 1}';
      }

      if (items[i].sh2 != '') {
        sh2 = '${int.parse(items[i].sh2) + 1}';
      }

      if (sh1 == '' && sh2 == '') {
        items[i].bsh = '';
      } else {
        items[i].bsh = '$sh1, $sh2';
      }
    }
  }
}
