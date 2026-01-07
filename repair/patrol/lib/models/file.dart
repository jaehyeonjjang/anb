import 'package:common_control/common_control.dart';


  
class File { 
  int id;
  String title;
  String filename;
  String originalfilename;
  int apt;
  String date;
  bool checked;
  Map<String, dynamic> extra;  

  File({        
          this.id = 0,       
          this.title = '',       
          this.filename = '',       
          this.originalfilename = '',       
          this.apt = 0,       
          this.date = '',
          this.extra = const{},
          this.checked = false}) ;
  

  factory File.fromJson(Map<String, dynamic> json) {
    return File(
        id: json['id'] as int,
        title: json['title'] as String,
        filename: json['filename'] as String,
        originalfilename: json['originalfilename'] as String,
        apt: json['apt'] as int,
        date: json['date'] as String, 
        extra: json['extra'] == null ? <String, dynamic>{} : json['extra'] as Map<String, dynamic>        
    );
  }

  Map<String, dynamic> toJson() =>
      { 'id': id,'title': title,'filename': filename,'originalfilename': originalfilename,'apt': apt,'date': date };
}

class FileManager {
  static const baseUrl = '/api/file';  

  static Future<List<File>> find(
      {int page = 0, int pagesize = 20, String? params}) async {
    var result =
        await Http.get(baseUrl, {'page': page, 'pagesize': pagesize}, params);
    if (result == null || result['items'] == null) {
      return List<File>.empty(growable: true);
    }

    return result['items']
        .map<File>((json) => File.fromJson(json))
        .toList();
  }

  static Future<File> get(int id) async {
    var result = await Http.get('$baseUrl/$id');
    if (result == null || result['item'] == null) {
      return File();
    }

    return File.fromJson(result['item']);
  }

  static Future<int> insert(File item) async {
    var result = await Http.insert(baseUrl, item.toJson());
    return result;
  }

  static update(File item) async {
    await Http.put(baseUrl, item.toJson());
  }

  static delete(File item) async {
    await Http.delete(baseUrl, item.toJson());
  }
}
