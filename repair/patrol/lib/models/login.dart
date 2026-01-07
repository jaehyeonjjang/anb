import 'package:common_control/common_control.dart';
import 'package:flutter/foundation.dart';
import 'package:patrol/models/user.dart';

class LoginManager {
  static Future<User> login(String loginid, String passwd) async {
    try {
      var url = '/api/jwt?loginid=$loginid&passwd=$passwd';
      var result = await Http.get(url);

      if (result['code'] != 'ok') {
        return User();
      }        

      final token = result['token'];
      final user = User.fromJson(result['user']);
      user.extra["token"] = token;
      return user;
    } catch (e) {
      if (kDebugMode) {
        print(e);
      }
    }

    return User();
  }
}
