import 'package:common_control/common_control.dart';

class Periodicdataimage {
  int id;
  String filename;
  String offlinefilename;
  int order;
  int periodicdata;
  int periodic;
  String date;
  bool checked;
  Map<String, dynamic> extra;

  Periodicdataimage(
      {this.id = 0,
      this.filename = '',
      this.offlinefilename = '',
      this.order = 0,
      this.periodicdata = 0,
      this.periodic = 0,
      this.date = '',
      this.extra = const {},
      this.checked = false});

  factory Periodicdataimage.fromJson(Map<String, dynamic> json) {
    return Periodicdataimage(
        id: json['id'] as int,
        filename: json['filename'] as String,
        offlinefilename: json['offlinefilename'] as String,
        order: json['order'] as int,
        periodicdata: json['periodicdata'] as int,
        periodic: json['periodic'] as int,
        date: json['date'] as String,
        extra: json['extra'] == null
            ? <String, dynamic>{}
            : json['extra'] as Map<String, dynamic>);
  }

  Map<String, dynamic> toJson() => {
        'id': id,
        'filename': filename,
        'offlinefilename': offlinefilename,
        'order': order,
        'periodicdata': periodicdata,
        'periodic': periodic,
        'date': date
      };
}

class PeriodicdataimageManager {
  static const baseUrl = '/api/periodicdataimage';

  static Future<List<Periodicdataimage>> find(
      {int page = 0, int pagesize = 20, String? params}) async {
    var result =
        await Http.get(baseUrl, {'page': page, 'pagesize': pagesize}, params);
    if (result == null || result['items'] == null) {
      return List<Periodicdataimage>.empty(growable: true);
    }

    return result['items']
        .map<Periodicdataimage>((json) => Periodicdataimage.fromJson(json))
        .toList();
  }

  static Future<Periodicdataimage> get(int id) async {
    var result = await Http.get('$baseUrl/$id');
    if (result == null || result['item'] == null) {
      return Periodicdataimage();
    }

    return Periodicdataimage.fromJson(result['item']);
  }

  static Future<int> insert(Periodicdataimage item) async {
    var result = await Http.insert(baseUrl, item.toJson());
    return result;
  }

  static update(Periodicdataimage item) async {
    await Http.put(baseUrl, item.toJson());
  }

  static delete(Periodicdataimage item) async {
    await Http.delete(baseUrl, item.toJson());
  }
}
