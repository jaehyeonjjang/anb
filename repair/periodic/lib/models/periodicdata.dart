import 'package:common_control/common_control.dart';
import 'package:periodic/models/blueprint.dart';

class Periodicdata {
  int id;
  int group;
  int type;
  String part;
  String member;
  String shape;
  String width;
  String length;
  int count;
  int progress;
  String remark;
  int order;
  String content;
  int status;
  String filename;
  String offlinefilename;
  Blueprint blueprint = Blueprint();
  int user;
  int periodic;
  String date;
  bool checked;
  Map<String, dynamic> extra;

  Periodicdata(
      {this.id = 0,
      this.group = 0,
      this.type = 0,
      this.part = '',
      this.member = '',
      this.shape = '',
      this.width = '',
      this.length = '',
      this.count = 0,
      this.progress = 0,
      this.remark = '',
      this.order = 0,
      this.content = '',
      this.status = 0,
      this.filename = '',
      this.offlinefilename = '',
      Blueprint? blueprint,
      this.user = 0,
      this.periodic = 0,
      this.date = '',
      this.extra = const {},
      this.checked = false}) {
    if (blueprint != null) {
      this.blueprint = blueprint;
    }
  }

  factory Periodicdata.fromJson(Map<String, dynamic> json) {
    return Periodicdata(
        id: json['id'] as int,
        group: json['group'] as int,
        type: json['type'] as int,
        part: json['part'] as String,
        member: json['member'] as String,
        shape: json['shape'] as String,
        width: json['width'] as String,
        length: json['length'] as String,
        count: json['count'] as int,
        progress: json['progress'] as int,
        remark: json['remark'] as String,
        order: json['order'] as int,
        content: json['content'] as String,
        status: json['status'] as int,
        filename: json['filename'] as String,
        offlinefilename: json['offlinefilename'] as String,
        blueprint: Blueprint.fromJson(json['extra']['blueprint']),
        user: json['user'] as int,
        periodic: json['periodic'] as int,
        extra: json['extra'] == null
            ? <String, dynamic>{}
            : json['extra'] as Map<String, dynamic>);
  }

  Map<String, dynamic> toJson() => {
        'id': id,
        'group': group,
        'type': type,
        'part': part,
        'member': member,
        'shape': shape,
        'width': width,
        'length': length,
        'count': count,
        'progress': progress,
        'remark': remark,
        'order': order,
        'content': content,
        'status': status,
        'filename': filename,
        'offlinefilename': offlinefilename,
        'blueprint': blueprint.id,
        'user': user,
        'periodic': periodic,
        'date': date
      };
}

class PeriodicdataManager {
  static const baseUrl = '/api/periodicdata';

  static Future<List<Periodicdata>> find(
      {int page = 0, int pagesize = 20, String? params}) async {
    var result =
        await Http.get(baseUrl, {'page': page, 'pagesize': pagesize}, params);
    if (result == null || result['items'] == null) {
      return List<Periodicdata>.empty(growable: true);
    }

    return result['items']
        .map<Periodicdata>((json) => Periodicdata.fromJson(json))
        .toList();
  }

  static Future<Periodicdata> get(int id) async {
    var result = await Http.get('$baseUrl/$id');
    if (result == null || result['item'] == null) {
      return Periodicdata();
    }

    return Periodicdata.fromJson(result['item']);
  }

  static Future<int> insert(Periodicdata item) async {
    var result = await Http.insert(baseUrl, item.toJson());
    return result;
  }

  static update(Periodicdata item) async {
    await Http.put(baseUrl, item.toJson());
  }

  static delete(Periodicdata item) async {
    await Http.delete(baseUrl, item.toJson());
  }
}
