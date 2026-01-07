import 'package:common_control/common_control.dart';


enum UserLevel {
    none, normal, manager, admin, rootadmin
}
  
class User { 
  int id;
  String loginid;
  String passwd;
  String name;
  String email;
  UserLevel level;
  int apt;
  String date;
  bool checked;
  Map<String, dynamic> extra;  

  User({        
          this.id = 0,       
          this.loginid = '',       
          this.passwd = '',       
          this.name = '',       
          this.email = '',       
          this.level = UserLevel.none,       
          this.apt = 0,       
          this.date = '',
          this.extra = const{},
          this.checked = false}) ;
  

  factory User.fromJson(Map<String, dynamic> json) {
    return User(
        id: json['id'] as int,
        loginid: json['loginid'] as String,
        passwd: json['passwd'] as String,
        name: json['name'] as String,
        email: json['email'] as String,
        level: UserLevel.values[json['level'] as int],
        apt: json['apt'] as int,
        date: json['date'] as String,
        extra: json['extra'] == null ? <String, dynamic>{} : json['extra'] as Map<String, dynamic>);
  }

  Map<String, dynamic> toJson() =>
      { 'id': id,'loginid': loginid,'passwd': passwd,'name': name,'email': email,'level': level.index,'apt': apt,'date': date };
}

class UserManager {
  static const baseUrl = '/api/user';  

  static Future<List<User>> find(
      {int page = 0, int pagesize = 20, String? params}) async {
    var result =
        await Http.get(baseUrl, {'page': page, 'pagesize': pagesize}, params);
    if (result == null || result['items'] == null) {
      return List<User>.empty(growable: true);
    }

    return result['items']
        .map<User>((json) => User.fromJson(json))
        .toList();
  }

  static Future<User> get(int id) async {
    var result = await Http.get('$baseUrl/$id');
    if (result == null || result['item'] == null) {
      return User();
    }

    return User.fromJson(result['item']);
  }

  static Future<int> insert(User item) async {
    var result = await Http.insert(baseUrl, item.toJson());
    return result;
  }

  static update(User item) async {
    await Http.put(baseUrl, item.toJson());
  }

  static delete(User item) async {
    await Http.delete(baseUrl, item.toJson());
  }
}
