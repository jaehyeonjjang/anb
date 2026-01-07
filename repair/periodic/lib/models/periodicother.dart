import 'package:common_control/common_control.dart';

class Periodicother {
  int id;
  String name;
  int type;
  int result;
  String status;
  String position;
  String filename;
  String offlinefilename;
  int change;
  int category;
  int order;
  int periodic;
  String date;
  bool checked;
  Map<String, dynamic> extra;

  Periodicother(
      {this.id = 0,
      this.name = '',
      this.type = 0,
      this.result = 0,
      this.status = '',
      this.position = '',
      this.filename = '',
      this.offlinefilename = '',
      this.change = 0,
      this.category = 0,
      this.order = 0,
      this.periodic = 0,
      this.date = '',
      this.extra = const {},
      this.checked = false});

  factory Periodicother.fromJson(Map<String, dynamic> json) {
    return Periodicother(
        id: json['id'] as int,
        name: json['name'] as String,
        type: json['type'] as int,
        result: json['result'] as int,
        status: json['status'] as String,
        position: json['position'] as String,
        filename: json['filename'] as String,
        offlinefilename: json['offlinefilename'] as String,
        change: json['change'] as int,
        category: json['category'] as int,
        order: json['order'] as int,
        periodic: json['periodic'] as int,
        date: json['date'] as String,
        extra: json['extra'] == null
            ? <String, dynamic>{}
            : json['extra'] as Map<String, dynamic>);
  }

  Map<String, dynamic> toJson() => {
        'id': id,
        'name': name,
        'type': type,
        'result': result,
        'status': status,
        'position': position,
        'filename': filename,
        'offlinefilename': offlinefilename,
        'change': change,
        'category': category,
        'order': order,
        'periodic': periodic,
        'date': date
      };
}

class PeriodicotherManager {
  static const baseUrl = '/api/periodicother';

  static Future<List<Periodicother>> find(
      {int page = 0, int pagesize = 20, String? params}) async {
    var result =
        await Http.get(baseUrl, {'page': page, 'pagesize': pagesize}, params);
    if (result == null || result['items'] == null) {
      return List<Periodicother>.empty(growable: true);
    }

    return result['items']
        .map<Periodicother>((json) => Periodicother.fromJson(json))
        .toList();
  }

  static Future<Periodicother> get(int id) async {
    var result = await Http.get('$baseUrl/$id');
    if (result == null || result['item'] == null) {
      return Periodicother();
    }

    return Periodicother.fromJson(result['item']);
  }

  static Future<int> insert(Periodicother item) async {
    var result = await Http.insert(baseUrl, item.toJson());
    return result;
  }

  static update(Periodicother item) async {
    await Http.put(baseUrl, item.toJson());
  }

  static delete(Periodicother item) async {
    await Http.delete(baseUrl, item.toJson());
  }
}
