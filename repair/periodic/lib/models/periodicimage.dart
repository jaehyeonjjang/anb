import 'package:common_control/common_control.dart';


  
class Periodicimage { 
  int id;
  int type;
  String filename;
  String offlinefilename;
  String name;
  int use;
  int order;
  int periodic;
  String date;
  bool checked;
  Map<String, dynamic> extra;  

  Periodicimage({        
          this.id = 0,       
          this.type = 0,       
          this.filename = '',       
          this.offlinefilename = '',       
          this.name = '',       
          this.use = 0,       
          this.order = 0,       
          this.periodic = 0,       
          this.date = '',
          this.extra = const{},
          this.checked = false}) ;
  

  factory Periodicimage.fromJson(Map<String, dynamic> json) {
    return Periodicimage(
        id: json['id'] as int,
        type: json['type'] as int,
        filename: json['filename'] as String,
        offlinefilename: json['offlinefilename'] as String,
        name: json['name'] as String,
        use: json['use'] as int,
        order: json['order'] as int,
        periodic: json['periodic'] as int,
        date: json['date'] as String, 
        extra: json['extra'] == null ? <String, dynamic>{} : json['extra'] as Map<String, dynamic>        
    );
  }

  Map<String, dynamic> toJson() =>
      { 'id': id,'type': type,'filename': filename,'offlinefilename': offlinefilename,'name': name,'use': use,'order': order,'periodic': periodic,'date': date };
}

class PeriodicimageManager {
  static const baseUrl = '/api/periodicimage';  

  static Future<List<Periodicimage>> find(
      {int page = 0, int pagesize = 20, String? params}) async {
    var result =
        await Http.get(baseUrl, {'page': page, 'pagesize': pagesize}, params);
    if (result == null || result['items'] == null) {
      return List<Periodicimage>.empty(growable: true);
    }

    return result['items']
        .map<Periodicimage>((json) => Periodicimage.fromJson(json))
        .toList();
  }

  static Future<Periodicimage> get(int id) async {
    var result = await Http.get('$baseUrl/$id');
    if (result == null || result['item'] == null) {
      return Periodicimage();
    }

    return Periodicimage.fromJson(result['item']);
  }

  static Future<int> insert(Periodicimage item) async {
    var result = await Http.insert(baseUrl, item.toJson());
    return result;
  }

  static update(Periodicimage item) async {
    await Http.put(baseUrl, item.toJson());
  }

  static delete(Periodicimage item) async {
    await Http.delete(baseUrl, item.toJson());
  }
}
