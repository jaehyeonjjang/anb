import 'package:common_control/common_control.dart';


  
class Periodicblueprintzoom { 
  int id;
  double iconzoom;
  double zoom;
  int status;
  int blueprint;
  int periodic;
  String date;
  bool checked;
  Map<String, dynamic> extra;  

  Periodicblueprintzoom({        
          this.id = 0,       
          this.iconzoom = 0.0,       
          this.zoom = 0.0,       
          this.status = 0,       
          this.blueprint = 0,       
          this.periodic = 0,       
          this.date = '',
          this.extra = const{},
          this.checked = false}) ;
  

  factory Periodicblueprintzoom.fromJson(Map<String, dynamic> json) {
    return Periodicblueprintzoom(
        id: json['id'] as int,
        iconzoom: json['iconzoom'] as double,
        zoom: json['zoom'] as double,
        status: json['status'] as int,
        blueprint: json['blueprint'] as int,
        periodic: json['periodic'] as int,
        date: json['date'] as String, 
        extra: json['extra'] == null ? <String, dynamic>{} : json['extra'] as Map<String, dynamic>        
    );
  }

  Map<String, dynamic> toJson() =>
      { 'id': id,'iconzoom': iconzoom,'zoom': zoom,'status': status,'blueprint': blueprint,'periodic': periodic,'date': date };
}

class PeriodicblueprintzoomManager {
  static const baseUrl = '/api/periodicblueprintzoom';  

  static Future<List<Periodicblueprintzoom>> find(
      {int page = 0, int pagesize = 20, String? params}) async {
    var result =
        await Http.get(baseUrl, {'page': page, 'pagesize': pagesize}, params);
    if (result == null || result['items'] == null) {
      return List<Periodicblueprintzoom>.empty(growable: true);
    }

    return result['items']
        .map<Periodicblueprintzoom>((json) => Periodicblueprintzoom.fromJson(json))
        .toList();
  }

  static Future<Periodicblueprintzoom> get(int id) async {
    var result = await Http.get('$baseUrl/$id');
    if (result == null || result['item'] == null) {
      return Periodicblueprintzoom();
    }

    return Periodicblueprintzoom.fromJson(result['item']);
  }

  static Future<int> insert(Periodicblueprintzoom item) async {
    var result = await Http.insert(baseUrl, item.toJson());
    return result;
  }

  static update(Periodicblueprintzoom item) async {
    await Http.put(baseUrl, item.toJson());
  }

  static delete(Periodicblueprintzoom item) async {
    await Http.delete(baseUrl, item.toJson());
  }
}
