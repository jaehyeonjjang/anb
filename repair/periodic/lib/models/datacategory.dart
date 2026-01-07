import 'package:common_control/common_control.dart';


  
class Datacategory { 
  int id;
  String name;
  int category;
  int type;
  String remark;
  int order;
  String date;
  bool checked;
  Map<String, dynamic> extra;  

  Datacategory({        
          this.id = 0,       
          this.name = '',       
          this.category = 0,       
          this.type = 0,       
          this.remark = '',       
          this.order = 0,       
          this.date = '',
          this.extra = const{},
          this.checked = false}) ;
  

  factory Datacategory.fromJson(Map<String, dynamic> json) {
    return Datacategory(
        id: json['id'] as int,
        name: json['name'] as String,
        category: json['category'] as int,
        type: json['type'] as int,
        remark: json['remark'] as String,
        order: json['order'] as int,
        date: json['date'] as String, 
        extra: json['extra'] == null ? <String, dynamic>{} : json['extra'] as Map<String, dynamic>        
    );
  }

  Map<String, dynamic> toJson() =>
      { 'id': id,'name': name,'category': category,'type': type,'remark': remark,'order': order,'date': date };
}

class DatacategoryManager {
  static const baseUrl = '/api/datacategory';  

  static Future<List<Datacategory>> find(
      {int page = 0, int pagesize = 20, String? params}) async {
    var result =
        await Http.get(baseUrl, {'page': page, 'pagesize': pagesize}, params);
    if (result == null || result['items'] == null) {
      return List<Datacategory>.empty(growable: true);
    }

    return result['items']
        .map<Datacategory>((json) => Datacategory.fromJson(json))
        .toList();
  }

  static Future<Datacategory> get(int id) async {
    var result = await Http.get('$baseUrl/$id');
    if (result == null || result['item'] == null) {
      return Datacategory();
    }

    return Datacategory.fromJson(result['item']);
  }

  static Future<int> insert(Datacategory item) async {
    var result = await Http.insert(baseUrl, item.toJson());
    return result;
  }

  static update(Datacategory item) async {
    await Http.put(baseUrl, item.toJson());
  }

  static delete(Datacategory item) async {
    await Http.delete(baseUrl, item.toJson());
  }
}
