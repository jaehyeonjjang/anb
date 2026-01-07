import 'dart:convert';
import 'dart:io' as io;
import 'dart:typed_data';

import 'package:common_control/common_control.dart';
import 'package:dio/dio.dart';
import 'package:gal/gal.dart';
import 'package:path_provider/path_provider.dart';
import 'package:intl/intl.dart';
import 'package:localstorage/localstorage.dart';
import 'package:patrol/components/painter/painter_controller.dart';
import 'package:patrol/controllers/auth_controller.dart';
import 'package:patrol/models/blueprint.dart';
import 'package:patrol/models/periodic.dart';
import 'package:patrol/models/periodicblueprintzoom.dart';

import 'package:patrol/models/periodicdata.dart';
import 'package:patrol/models/user.dart';

class BlueprintController extends GetxController {
  final _periodic = Periodic().obs;
  final _id = 0.obs;
  final _items = <Blueprint>[].obs;
  final _percent = 0.0.obs;
  final _cancel = false.obs;
  final _loading = false.obs;
  final _sendError = false.obs;

  final _total = 0.obs;

  Periodic get periodic => _periodic.value;
  set periodic(value) => _periodic.value = value;

  int get id => _id.value;
  set id(value) => _id.value = value;

  List<Blueprint> get items => _items;
  set items(value) => _items.value = value;

  double get percent => _percent.value;
  set percent(value) => _percent.value = value;

  bool get cancel => _cancel.value;
  set cancel(value) => _cancel.value = value;

  get loading => _loading.value;
  set loading(value) => _loading.value = value;

  final _modified = false.obs;
  bool get modified => _modified.value;
  set modified(value) => _modified.value = value;

  int get total => _total.value;
  set total(value) => _total.value = value;

  bool get sendError => _sendError.value;
  set sendError(value) => _sendError.value = value;

  final TextEditingController txtLocation = TextEditingController();
  final TextEditingController txtContent = TextEditingController();
  final TextEditingController txtProcess = TextEditingController();
  final TextEditingController txtOpinion = TextEditingController();

  Future<String> downloadFile(String filename) async {
    try {
      var response = await Dio()
          .get(filename, options: Options(responseType: ResponseType.bytes));

      if (response.data == null) {
        return '';
      }

      // 임시 파일로 저장
      final tempDir = await getTemporaryDirectory();
      final tempFile = io.File('${tempDir.path}/${DateTime.now().millisecondsSinceEpoch}.jpg');
      await tempFile.writeAsBytes(response.data);

      // 갤러리에 저장
      await Gal.putImage(tempFile.path);

      return tempFile.path;
    } catch (e) {
      return '';
    }
  }

  init() async {
    modified = false;
    loading = false;
    percent = 0.0;

    final AuthController authController = Get.find<AuthController>();


    if (id != 0 && authController.periodic.id == id) {
      final LocalStorage storageBlueprint = LocalStorage('blueprints.json');
      await storageBlueprint.ready;
      var str = await storageBlueprint.getItem('blueprints');
      if (str != null && str != '') {
        items = json
            .decode(str)
            .map<Blueprint>((item) => Blueprint.fromJson(item))
            .toList();
      } else {
        items = <Blueprint>[];
      }

      modified = true;
      loading = true;
      
      return;
    }

    txtLocation.text = periodic.name;
    txtContent.text = periodic.resulttext1;
    txtProcess.text = periodic.resulttext2;
    txtOpinion.text = periodic.resulttext3;

    if (periodic.id == 0) {
      final now = DateTime.now();
      periodic.resulttext4 = DateFormat('yyyy-MM-dd HH:mm:ss').format(now);          
    }

    authController.periodic = periodic;

    final LocalStorage storageBlueprint = LocalStorage('blueprints.json');
    await storageBlueprint.ready;
    var str = await storageBlueprint.getItem('blueprints');

    List<Blueprint> oldBlueprints = [];

    if (str != null && str != '') {
      List<dynamic> old = json.decode(str);
      for (var i = 0; i < old.length; i++) {
        var item = Blueprint.fromJson(old[i]);

        if (item.offlinefilename == '') {
          continue;
        }

        oldBlueprints.add(item);
      }
    }

    var apt = authController.user.apt;

    if (authController.user.level == UserLevel.admin ||
        authController.user.level == UserLevel.rootadmin) {
      apt = periodic.apt.id;
    }

    var blueprints = await BlueprintManager.find(
        page: 0,
        pagesize: 0,
        params:
            'apt=$apt&category=3&orderby=bp_parentorder,bp_parent,bp_order desc,bp_id');

    List<Periodicblueprintzoom> zooms = [];
    List<Periodicdata> periodicdatas = [];

    if (id > 0) {
      zooms = await PeriodicblueprintzoomManager.find(
          page: 0, pagesize: 0, params: 'periodic=$id');

      periodicdatas = await PeriodicdataManager.find(
          page: 0, pagesize: 0, params: 'periodic=$id&orderby=pd_order,pd_id');
    }

    var localtotal = 0;

    for (var i = 0; i < blueprints.length; i++) {
      var blueprint = blueprints[i];

      if (blueprint.upload != 1) {
        continue;
      }

      if (blueprint.filename == '') {
        continue;
      }

      var find = false;
      for (var j = 0; j < oldBlueprints.length; j++) {
        if (blueprint.filename == oldBlueprints[j].filename) {
          find = true;
          break;
        }
      }

      if (find == true) {
        continue;
      }

      localtotal++;
    }

    for (var i = 0; i < periodicdatas.length; i++) {
      final data = periodicdatas[i];

      final onlineimages = data.filename.split(',');

      for (var k = 0; k < onlineimages.length; k++) {
        if (onlineimages[k] == '') {
          continue;
        }

        localtotal++;
      }
    }

    total = localtotal;
    var current = 0;

    for (var i = 0; i < blueprints.length; i++) {
      var blueprint = blueprints[i];

      if (blueprint.upload != 1) {
        continue;
      }

      if (blueprint.filename == '') {
        continue;
      }

      var find = false;
      for (var j = 0; j < oldBlueprints.length; j++) {
        if (blueprint.filename == oldBlueprints[j].filename) {
          blueprints[i].offlinefilename = oldBlueprints[j].offlinefilename;
          find = true;
          break;
        }
      }

      if (find == true) {
        continue;
      }

      final filename = '${CConfig().serverUrl}/webdata/${blueprint.filename}';
      var path = await downloadFile(filename);
      current++;
      percent = current / total;
      blueprints[i].offlinefilename = path;
    }

    items = blueprints;

    final blueprintStr = json.encode(blueprints);

    await storageBlueprint.ready;
    await storageBlueprint.setItem('blueprints', blueprintStr);

    final LocalStorage storage = LocalStorage('periodic.json');
    await storage.ready;
    await storage.clear();

    Map<int, double> zoomMap = <int, double>{};
    Map<int, double> iconzoomMap = <int, double>{};
    for (var i = 0; i < zooms.length; i++) {
      zoomMap[zooms[i].blueprint] = zooms[i].zoom;
    }

    for (var i = 0; i < zooms.length; i++) {
      iconzoomMap[zooms[i].blueprint] = zooms[i].iconzoom;
    }

    Map<int, List<Periodicdata>> dataMap = <int, List<Periodicdata>>{};
    for (var i = 0; i < periodicdatas.length; i++) {
      final item = periodicdatas[i];

      if (!dataMap.containsKey(item.blueprint.id)) {
        dataMap[item.blueprint.id] = [];
      }

      dataMap[item.blueprint.id]?.add(item);
    }

    for (var i = 0; i < blueprints.length; i++) {
      final item = blueprints[i];

      if (!dataMap.containsKey(item.id)) {
        continue;
      }

      var iconZoom = 100.0;
      if (iconzoomMap.containsKey(item.id)) {
        iconZoom = iconzoomMap[item.id]!;
      }

      var zoom = 1.0;
      if (zoomMap.containsKey(item.id)) {
        zoom = zoomMap[item.id]!;
      }

      final datas = dataMap[item.id]!;

      List<Point> points = [];

      for (var j = 0; j < datas.length; j++) {
        final data = datas[j];

        List<String> images = [];
        final onlineimages = data.filename.split(',');

        for (var k = 0; k < onlineimages.length; k++) {
          if (onlineimages[k] == '') {
            continue;
          }

          final filename = '${CConfig().serverUrl}/webdata/${onlineimages[k]}';
          var path = await downloadFile(filename);
          current++;
          percent = current / total;
          images.add(path);
        }

        if (data.content == '') {
          continue;
        }

        List<Offset> content = [];
        List<dynamic> temp = json.decode(data.content);

        for (var k = 0; k < temp.length; k++) {
          var item = temp[k];
          final offset = Offset(item['dx'], item['dy']);

          content.add(offset);
        }

        DrawType type = DrawType.number;

        if (data.type >= 100) {
          type = DrawType.icon;
        } else if (data.type >= 30 && data.type < 40) {
          type = DrawType.curve;
        } else if (data.type >= 40 && data.type < 50) {
          type = DrawType.line;
        }

        final point = Point(
            items: content,
            color: LineColor.black,
            width: 1,
            type: type,
            icon: data.type,
            number: data.group,
            part: data.part,
            member: data.member,
            shape: data.shape,
            weight: data.width,
            length: data.length,
            count: data.count.toString(),
            progress: data.progress == 1 ? 'O' : 'X',
            remark: data.remark,
            order: data.order,
            images: images,
            onlineimages: []);

        points.add(point);
      }

      Map<String, dynamic> items = {
        'id': item.id,
        'points': points.map((item) => item.toJson()).toList(),
        'zoom': zoom,
        'iconzoom': iconZoom
      };

      final str = json.encode(items);

      await storage.ready;
      await storage.setItem('data_${item.id}', str);
    }

    final LocalStorage storageLogin = LocalStorage('login.json');
    await storageLogin.ready;
    await storageLogin.setItem('periodic', periodic);

    percent = 1.0;

    loading = true;
  }
}
