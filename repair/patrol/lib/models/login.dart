import 'package:common_control/common_control.dart';
import 'package:flutter/foundation.dart';
import 'package:patrol/models/user.dart';

class LoginManager {
  static Future<User> login(String loginid, String passwd) async {
    try {
      // URL 파라미터 인코딩 (특수문자 처리)
      var encodedLoginid = Uri.encodeComponent(loginid);
      var encodedPasswd = Uri.encodeComponent(passwd);
      var url = '/api/jwt?loginid=$encodedLoginid&passwd=$encodedPasswd';
      if (kDebugMode) {
        print('DEBUG: Login URL: $url');
      }
      var result = await Http.get(url);

      if (kDebugMode) {
        print('DEBUG: Server response: $result');
      }

      if (result == null) {
        if (kDebugMode) {
          print('DEBUG: result is null');
        }
        return User();
      }

      if (result['code'] != 'ok') {
        if (kDebugMode) {
          print('DEBUG: Login failed, code=${result["code"]}');
        }
        return User();
      }

      final token = result['token'];
      final user = User.fromJson(result['user']);
      user.extra["token"] = token;
      if (kDebugMode) {
        print(
            'DEBUG: Login success! user.id=${user.id}, user.level=${user.level}');
      }
      return user;
    } catch (e) {
      if (kDebugMode) {
        print('DEBUG: Login exception: $e');
      }
    }

    return User();
  }
}
