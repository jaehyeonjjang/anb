import 'package:common_control/common_control.dart';


  
class Category { 
  int id;
  String name;
  int level;
  int parent;
  int cycle;
  int percent;
  String unit;
  int elevator;
  String remark;
  int order;
  int apt;
  String date;
  bool checked;
  Map<String, dynamic> extra;  

  Category({        
          this.id = 0,       
          this.name = '',       
          this.level = 0,       
          this.parent = 0,       
          this.cycle = 0,       
          this.percent = 0,       
          this.unit = '',       
          this.elevator = 0,       
          this.remark = '',       
          this.order = 0,       
          this.apt = 0,       
          this.date = '',
          this.extra = const{},
          this.checked = false}) ;
  

  factory Category.fromJson(Map<String, dynamic> json) {
    return Category(
        id: json['id'] as int,
        name: json['name'] as String,
        level: json['level'] as int,
        parent: json['parent'] as int,
        cycle: json['cycle'] as int,
        percent: json['percent'] as int,
        unit: json['unit'] as String,
        elevator: json['elevator'] as int,
        remark: json['remark'] as String,
        order: json['order'] as int,
        apt: json['apt'] as int,
        date: json['date'] as String,
        extra: json['extra'] == null ? <String, dynamic>{} : json['extra'] as Map<String, dynamic>
    );
  }

  Map<String, dynamic> toJson() =>
      { 'id': id,'name': name,'level': level,'parent': parent,'cycle': cycle,'percent': percent,'unit': unit,'elevator': elevator,'remark': remark,'order': order,'apt': apt,'date': date };
}

class CategoryManager {
  static const baseUrl = '/api/category';  

  static Future<List<Category>> find(
      {int page = 0, int pagesize = 20, String? params}) async {
    var result =
        await Http.get(baseUrl, {'page': page, 'pagesize': pagesize}, params);
    if (result == null || result['items'] == null) {
      return List<Category>.empty(growable: true);
    }

    return result['items']
        .map<Category>((json) => Category.fromJson(json))
        .toList();
  }

  static Future<Category> get(int id) async {
    var result = await Http.get('$baseUrl/$id');
    if (result == null || result['item'] == null) {
      return Category();
    }

    return Category.fromJson(result['item']);
  }

  static Future<int> insert(Category item) async {
    var result = await Http.insert(baseUrl, item.toJson());
    return result;
  }

  static update(Category item) async {
    await Http.put(baseUrl, item.toJson());
  }

  static delete(Category item) async {
    await Http.delete(baseUrl, item.toJson());
  }
}
