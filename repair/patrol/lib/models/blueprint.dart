import 'package:common_control/common_control.dart';

enum BlueprintFloortype { none, parking, underground, ground, top, roof, other }

class Blueprint {
  int id;
  String name;
  int level;
  int parent;
  BlueprintFloortype floortype;
  String filename;
  int upload;
  int parentorder;
  int order;
  String offlinefilename;
  int aptdong;
  int apt;
  String date;
  bool checked;
  Map<String, dynamic> extra;

  Blueprint(
      {this.id = 0,
      this.name = '',
      this.level = 0,
      this.parent = 0,
      this.floortype = BlueprintFloortype.none,
      this.filename = '',
      this.upload = 0,
      this.parentorder = 0,
      this.order = 0,
      this.offlinefilename = '',
      this.aptdong = 0,
      this.apt = 0,
      this.date = '',
      this.extra = const {},
      this.checked = false});

  factory Blueprint.fromJson(Map<String, dynamic> json) {
    return Blueprint(
        id: json['id'] as int,
        name: json['name'] as String,
        level: json['level'] as int,
        parent: json['parent'] as int,
        floortype: BlueprintFloortype.values[json['floortype'] as int],
        filename: json['filename'] as String,
        upload: json['upload'] as int,
        parentorder: json['parentorder'] as int,
        order: json['order'] as int,
        offlinefilename: json['offlinefilename'] as String,
        aptdong: json['aptdong'] as int,
        apt: json['apt'] as int,
        date: json['date'] as String,
        extra: json['extra'] == null ? <String, dynamic>{} : json['extra'] as Map<String, dynamic>);
  }

  Map<String, dynamic> toJson() => {
        'id': id,
        'name': name,
        'level': level,
        'parent': parent,
        'floortype': floortype.index,
        'filename': filename,
        'upload': upload,
        'parentorder': parentorder,
        'order': order,
        'offlinefilename': offlinefilename,
        'aptdong': aptdong,
        'apt': apt,
        'date': date
      };
}

class BlueprintManager {
  static const baseUrl = '/api/blueprint';

  static Future<List<Blueprint>> find(
      {int page = 0, int pagesize = 20, String? params}) async {
    var result =
        await Http.get(baseUrl, {'page': page, 'pagesize': pagesize}, params);
    if (result == null || result['items'] == null) {
      return List<Blueprint>.empty(growable: true);
    }

    return result['items']
        .map<Blueprint>((json) => Blueprint.fromJson(json))
        .toList();
  }

  static Future<Blueprint> get(int id) async {
    var result = await Http.get('$baseUrl/$id');
    if (result == null || result['item'] == null) {
      return Blueprint();
    }

    return Blueprint.fromJson(result['item']);
  }

  static Future<int> insert(Blueprint item) async {
    var result = await Http.insert(baseUrl, item.toJson());
    return result;
  }

  static update(Blueprint item) async {
    await Http.put(baseUrl, item.toJson());
  }

  static delete(Blueprint item) async {
    await Http.delete(baseUrl, item.toJson());
  }
}
