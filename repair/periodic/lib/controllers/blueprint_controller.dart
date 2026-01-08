import 'dart:convert';
import 'dart:io';
import 'dart:typed_data';

import 'package:common_control/common_control.dart';
import 'package:dio/dio.dart';
import 'package:http/http.dart' as http;
import 'package:gal/gal.dart';
import 'package:localstorage/localstorage.dart';
import 'package:path/path.dart' as p;
import 'package:path_provider/path_provider.dart';
import 'package:periodic/components/painter/painter_controller.dart';
import 'package:periodic/config/config.dart' as config;
import 'package:periodic/controllers/auth_controller.dart';
import 'package:periodic/models/blueprint.dart';
import 'package:periodic/models/datacategory.dart';
import 'package:periodic/models/periodic.dart';
import 'package:periodic/models/periodicblueprintzoom.dart';
import 'package:periodic/models/periodicdata.dart';
import 'package:periodic/models/periodicother.dart';
import 'package:periodic/models/user.dart';

class BlueprintController extends GetxController {
  final _periodic = Periodic().obs;
  final _id = 0.obs;
  final _items = <Blueprint>[].obs;
  final _datacategorys = <Datacategory>[].obs;
  final _periodicothers = <Periodicother>[].obs;
  final _percent = 0.0.obs;
  final _cancel = false.obs;
  final _loading = false.obs;
  final _sendError = false.obs;

  final _total = 0.obs;

  final _modifiedOther = false.obs;
  final _modifiedImage = false.obs;
  bool get modifiedOther => _modifiedOther.value;
  set modifiedOther(value) => _modifiedOther.value = value;
  bool get modifiedImage => _modifiedImage.value;
  set modifiedImage(value) => _modifiedImage.value = value;

  Periodic get periodic => _periodic.value;
  set periodic(value) => _periodic.value = value;

  int get id => _id.value;
  set id(value) => _id.value = value;

  List<Blueprint> get items => _items;
  set items(value) => _items.value = value;

  List<Datacategory> get datacategorys => _datacategorys;
  set datacategorys(value) => _datacategorys.value = value;

  List<Periodicother> get periodicothers => _periodicothers;
  set periodicothers(value) => _periodicothers.value = value;

  double get percent => _percent.value;
  set percent(double value) {
    if (value < 0.0 || value > 1.0) {
      return;
    }
    _percent.value = value;
  }

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

  setModified(target) {
    for (var i = 0; i < items.length; i++) {
      var item = items[i];

      if (item.id == target) {
        items[i].extra["modified"] = true;
        break;
      }
    }

    _items.refresh();
  }

  setCheck(index, value) {
    items[index].checked = value;
    _items.refresh();
  }

  setCheckAll() {
    for (var i = 0; i < items.length; i++) {
      if (items[i].extra['modified'] == null) {
        items[i].checked = false;
      } else {
        items[i].checked = true;
      }
    }
  }

  Future<String?> findSavedFilePath(String filename) async {
    Directory? directory = await getExternalStorageDirectory();
    String baseDir =
        directory!.path.split("/Android")[0]; // `/storage/emulated/0/`

    // 일반적으로 `Pictures` 또는 `DCIM` 폴더에 저장됨
    List<String> possibleDirs = [
      "$baseDir/Pictures/",
      "$baseDir/DCIM/Camera/",
      "$baseDir/Download/"
    ];

    for (String dir in possibleDirs) {
      File file = File("$dir$filename");
      if (await file.exists()) {
        return file.path;
      }
    }
    return null; // 파일을 찾지 못한 경우
  }

  Future<String> downloadFile(String filename) async {
    if (config.platform() == 'web') {
      return '';
    }

    if (config.platform() == 'android') {
      var onlyFilename = p.basename(filename);
      if (p.extension(onlyFilename).toLowerCase() != ".jpg") {
        // var appDocDir = await getTemporaryDirectory();
        var appDocDir = await getApplicationDocumentsDirectory();
        String savePath = appDocDir.path + onlyFilename;
        await Dio().download(filename, savePath);
        return savePath;
        // j
        // final result = await ImageGallerySaver.saveFile(savePath, isReturnPathOfIOS: true);
        // print(result);
        // return result['filePath'];
        // String? path = await findSavedFilePath(onlyFilename);
        // if (path == null) {
        //   print('path == null');
        //   final temp = result['filePath'].split('/');
        //   final id = temp[temp.length - 1];
        //   print('id = $id');
        //   print('$id${p.extension(onlyFilename)}');
        //   path = await findSavedFilePath('$id${p.extension(onlyFilename)}');
        //   print(path);
        // }

        // return path!;
      } else {
        var response = await Dio()
            .get(filename, options: Options(responseType: ResponseType.bytes));

        if (response.data == null) {
          return '';
        }

        try {
          // 임시 파일로 저장
          final tempDir = await getTemporaryDirectory();
          final tempFile = File('${tempDir.path}/$onlyFilename');
          await tempFile.writeAsBytes(response.data);

          // 갤러리에 저장
          await Gal.putImage(tempFile.path);

          String? path = await findSavedFilePath(onlyFilename);
          if (path == null) {
            path = tempFile.path;
          }
          return path!;
        } catch (e) {
          // error
        }
      }

      return '';
    } else {
      // try {
      //   var tempId = await ImageDownloader.downloadImage(filename);
      //   if (tempId == null) {
      //     return '';
      //   }

      //   imageId = tempId;
      // } catch (e) {
      //   return '';
      // }
      final response = await http.get(Uri.parse(filename));
      if (response.statusCode == 200) {
        final directory = await getApplicationDocumentsDirectory();
        final filePath = '${directory.path}/${p.basename(filename)}';
        final file = File(filePath);
        await file.writeAsBytes(response.bodyBytes);
        return filePath;
      }
    }

    // var path = await ImageDownloader.findPath(imageId);
    // return path!;

    return '';
  }

  init() async {
    modified = false;
    loading = false;
    percent = 0.0;

    final AuthController authController = Get.find<AuthController>();

    if (authController.periodic.id == id) {
      final LocalStorage storageBlueprint = LocalStorage('blueprints.json');
      await storageBlueprint.ready;
      var str = await storageBlueprint.getItem('blueprints');
      if (str != null && str != '') {
        items = json
            .decode(str)
            .map<Blueprint>((item) => Blueprint.fromJson(item))
            .toList();

        final LocalStorage storage = LocalStorage('periodic.json');

        for (var i = 0; i < items.length; i++) {
          final blueprint = items[i].id;

          await storage.ready;
          final data = await storage.getItem('data_$blueprint');
          final save = await storage.getItem('save_$blueprint');

          if (save == null || save == '') {
            continue;
          }

          if (data == null || data == '') {
            continue;
          }

          items[i].extra['modified'] = true;
        }
      } else {
        items = <Blueprint>[];
      }

      str = await storageBlueprint.getItem('datacategorys');
      if (str != null && str != '') {
        datacategorys = json
            .decode(str)
            .map<Datacategory>((item) => Datacategory.fromJson(item))
            .toList();
      } else {
        datacategorys = <Datacategory>[];
      }

      str = await storageBlueprint.getItem('periodicothers');
      if (str != null && str != '') {
        periodicothers = json
            .decode(str)
            .map<Periodicother>((item) => Periodicother.fromJson(item))
            .toList();
      } else {
        periodicothers = <Periodicother>[];
      }

      modified = true;
      loading = true;

      return;
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

    var datacategoryItems = await DatacategoryManager.find(
        page: 0, pagesize: 0, params: 'orderby=dc_order,dc_id');

    var blueprints = await BlueprintManager.find(
        page: 0,
        pagesize: 0,
        params:
            'apt=$apt&category=1&orderby=bp_parentorder,bp_parent,bp_order desc,bp_id');

    var periodicotherItems = await PeriodicotherManager.find(
        page: 0, pagesize: 0, params: 'periodic=$id&orderby=po_order,po_id');

    var zooms = await PeriodicblueprintzoomManager.find(
        page: 0, pagesize: 0, params: 'periodic=$id');

    var periodicdatas = await PeriodicdataManager.find(
        page: 0, pagesize: 0, params: 'periodic=$id&orderby=pd_order,pd_id');

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

    for (var i = 0; i < periodicotherItems.length; i++) {
      var periodic = periodicotherItems[i];

      periodicotherItems[i].change = 0;

      if (periodic.filename == '') {
        continue;
      }

      final offlinefilenames = periodic.offlinefilename.split(',');

      for (var j = 0; j < offlinefilenames.length; j++) {
        if (config.isExistFile(offlinefilenames[j]) == true) {
          continue;
        }

        localtotal++;
      }
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

    for (var i = 0; i < periodicotherItems.length; i++) {
      var periodic = periodicotherItems[i];

      periodicotherItems[i].change = 0;

      if (periodic.filename == '') {
        continue;
      }

      final filenames = periodic.filename.split(',');
      final offlinefilenames = periodic.offlinefilename.split(',');

      List<String> newFilenames = <String>[];

      for (var j = 0; j < offlinefilenames.length; j++) {
        if (config.isExistFile(offlinefilenames[j]) == true) {
          newFilenames.add(offlinefilenames[j]);
          continue;
        }

        if (j >= filenames.length) {
          break;
        }

        final onlinefilename = filenames[j];
        final filename = '${CConfig().serverUrl}/webdata/$onlinefilename';
        var path = await downloadFile(filename);
        if (path != "") {
          current++;
          percent = current / total;
          newFilenames.add(path);
        } else {
          filenames[j] = '';
          newFilenames.add('');
        }
      }

      periodicotherItems[i].offlinefilename = newFilenames.join(',');
    }

    items = blueprints;
    datacategorys = datacategoryItems;
    periodicothers = periodicotherItems;

    final blueprintStr = json.encode(blueprints);
    final datacategorysStr = json.encode(datacategorys);
    final periodicotherStr = json.encode(periodicothers);

    await storageBlueprint.ready;
    await storageBlueprint.setItem('blueprints', blueprintStr);
    await storageBlueprint.setItem('datacategorys', datacategorysStr);
    await storageBlueprint.setItem('periodicothers', periodicotherStr);

    final LocalStorage storage = LocalStorage('periodic.json');
    await storage.ready;
    await storage.clear();

    Map<int, double> zoomMap = <int, double>{};
    Map<int, double> iconzoomMap = <int, double>{};
    Map<int, double> numberzoomMap = <int, double>{};
    Map<int, double> crackzoomMap = <int, double>{};
    for (var i = 0; i < zooms.length; i++) {
      try {
        zoomMap[zooms[i].blueprint] = zooms[i].zoom;
      } catch (e) {
        zoomMap[zooms[i].blueprint] = zooms[i].zoom.toDouble();
      }
    }

    for (var i = 0; i < zooms.length; i++) {
      try {
        iconzoomMap[zooms[i].blueprint] = zooms[i].iconzoom;
      } catch (e) {
        iconzoomMap[zooms[i].blueprint] = zooms[i].iconzoom.toDouble();
      }
    }

    for (var i = 0; i < zooms.length; i++) {
      try {
        numberzoomMap[zooms[i].blueprint] = zooms[i].numberzoom;
      } catch (e) {
        numberzoomMap[zooms[i].blueprint] = zooms[i].numberzoom.toDouble();
      }
    }

    for (var i = 0; i < zooms.length; i++) {
      try {
        crackzoomMap[zooms[i].blueprint] = zooms[i].crackzoom;
      } catch (e) {
        crackzoomMap[zooms[i].blueprint] = zooms[i].crackzoom.toDouble();
      }
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

      var numberZoom = 100.0;
      if (numberzoomMap.containsKey(item.id)) {
        numberZoom = numberzoomMap[item.id]!;
      }

      var crackZoom = 100.0;
      if (crackzoomMap.containsKey(item.id)) {
        crackZoom = crackzoomMap[item.id]!;
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
          double dx = 0.0;
          double dy = 0.0;

          try {
            dx = item['dx'];
          } catch (e) {
            dx = item['dx'].toDouble();
          }

          try {
            dy = item['dy'];
          } catch (e) {
            dy = item['dy'].toDouble();
          }
          final offset = Offset(dx, dy);

          content.add(offset);
        }

        DrawType type = DrawType.number;
        var color = LineColor.black;
        if (data.type == basicHorizontalLine ||
            data.type == basicVerticalLine ||
            data.type == basicHorizontalBreak ||
            data.type == basicVerticalBreak) {
          type = DrawType.numberLine;
        } else if (data.type == inclinationLine) {
          type = DrawType.line;
        } else if (data.type >= 100) {
          type = DrawType.icon;
        } else if (data.type >= 30 && data.type < 40) {
          type = DrawType.curve;
        } else if (data.type >= 40 && data.type < 50) {
          type = DrawType.line;
        }

        if (data.type == inclinationLine ||
            data.type == inclinationHorizontal ||
            data.type == inclinationVertical) {
          color = LineColor.red;
        }

        if (data.type == basicHorizontalLine) {
          color = LineColor.red;
        } else if (data.type == basicVerticalLine) {
          color = LineColor.blue;
        } else if (data.type == basicHorizontalBreak) {
          color = LineColor.red;
        } else if (data.type == basicVerticalBreak) {
          color = LineColor.blue;
        }

        final point = Point(
            items: content,
            color: color,
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
        'iconzoom': iconZoom,
        'numberzoom': numberZoom,
        'crackzoom': crackZoom,
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
