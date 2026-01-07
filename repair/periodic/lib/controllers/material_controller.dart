import 'dart:convert';

import 'package:common_control/common_control.dart';
import 'package:periodic/components/painter/painter_controller.dart';

class Materialitem {
  String floor;
  String t;

  String floor2;
  String t2;

  Materialitem(
      {required this.floor,
      required this.t,
      required this.floor2,
      required this.t2});

  factory Materialitem.fromJson(Map<String, dynamic> json) {
    return Materialitem(
        floor: json['floor'] as String,
        t: json['t'] as String,
        floor2: json['floor2'] as String,
        t2: json['t2'] as String);
  }

  Map<String, dynamic> toJson() =>
      {'floor': floor, 't': t, 'floor2': floor2, 't2': t2};
}

class MaterialController extends GetxController {
  final _items = <List<Materialitem>>[].obs;
  final _x = 0.obs;
  final _y = 0.obs;
  final _z = 0.obs;

  final _colors = [].obs;

  get colors => _colors;
  set colors(value) => _colors.value = value;

  List<List<Materialitem>> get items => _items;
  set items(value) => _items.value = value;

  int get x => _x.value;
  set x(int value) => _x.value = value;

  int get y => _y.value;
  set y(int value) => _y.value = value;

  int get z => _z.value;
  set z(int value) => _z.value = value;

  @override
  onInit() {
    super.onInit();

    List<List<Materialitem>> alllist = [];

    var find = false;

    PainterController c = Get.find<PainterController>();
    for (var i = 0; i < c.points.length; i++) {
      var point = c.points[i];

      if (point.icon != 401 && point.icon != 402) {
        continue;
      }

      if (find == true) {
        find = false;

        colors[colors.length - 1] = point.icon;
        continue;
      }

      find = true;
      colors.add(point.icon);

      List<Materialitem> data = [];

      if (point.remark.length < 10) {
        for (var j = 0; j < 10; j++) {
          final item = Materialitem(floor: '', t: '', floor2: '', t2: '');
          data.add(item);
        }
      } else {
        var json = jsonDecode(point.remark);

        var list = json as List<dynamic>;

        for (var j = 0; j < list.length; j++) {
          final item = Materialitem.fromJson(list[j]);
          data.add(item);
        }
      }

      alllist.add(data);
    }

    items = alllist;
  }

  @override
  onClose() {
    super.onClose();

    var find = false;
    var pos = 0;
    PainterController c = Get.find<PainterController>();
    for (var i = 0; i < c.points.length; i++) {
      var point = c.points[i];

      if (point.icon != 401 && point.icon != 402) {
        continue;
      }

      if (find == true) {
        find = false;
        continue;
      }

      find = true;

      var json = jsonEncode(items[pos].map((item) => item.toJson()).toList());

      c.points[i].remark = json;
      pos++;
    }

    c.modified = true;
  }

  movePrev() {
    if (x > 0) {
      x--;
    } else {
      if (y > 0) {
        y--;
        x = 1;
      } else {
        if (z > 0) {
          z--;
          y = 9;
          x = 1;
        }
      }
    }
  }

  moveNext() {
    if (x < 1) {
      x++;
    } else {
      if (y < 9) {
        x = 0;
        y++;
      } else {
        if (z < items.length - 1) {
          z++;
          x = 0;
          y = 0;
        }
      }
    }
  }

  move(dz, dy, dx) {
    z = dz;
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
    for (var j = 0; j < items.length; j++) {
      for (var i = 0; i < items[j].length; i++) {
        items[j][i].floor2 = items[j][i].floor;

        if (items[j][i].t == '') {
          items[j][i].t2 = '';
        } else {
          items[j][i].t2 = '${int.parse(items[j][i].t) + 1}';
        }
      }
    }
  }
}
