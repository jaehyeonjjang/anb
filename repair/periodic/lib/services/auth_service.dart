import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:periodic/controllers/auth_controller.dart';

class AuthService extends GetMiddleware {
  final c = Get.find<AuthController>();

  @override
  RouteSettings? redirect(String? route) {
    if (!c.authenticated) {
      return const RouteSettings(name: '/login');
    }

    return null;
  }
}
