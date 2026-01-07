import 'package:common_control/common_control.dart';
import 'package:flutter/services.dart';

import 'package:patrol/components/painter/painter.dart';
import 'package:patrol/config/config.dart' as config;
import 'package:patrol/controllers/auth_controller.dart';
import 'package:patrol/controllers/image_controller.dart';
import 'package:patrol/controllers/main_controller.dart';
import 'package:patrol/controllers/setting_controller.dart';
import 'package:patrol/models/periodic.dart';
import 'package:patrol/screens/blueprint_screen.dart';
import 'package:patrol/screens/login_screen.dart';
import 'package:patrol/screens/main_screen.dart';
import 'package:patrol/screens/setting_screen.dart';
import 'package:patrol/screens/write_screen.dart';
import 'package:patrol/screens/image_screen.dart';
import 'package:patrol/screens/data_screen.dart';
import 'package:patrol/services/auth_service.dart';

import 'package:patrol/controllers/blueprint_controller.dart';
import 'package:patrol/controllers/write_controller.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await SystemChrome.setPreferredOrientations([DeviceOrientation.portraitDown]);

  CConfig().serverUrl = const String.fromEnvironment('API_BASE_URL',
      defaultValue: config.serverUrl);

  final c = Get.put(AuthController(), permanent: true);
  Get.put(BlueprintController(), permanent: true);

  CWidget.initialize();
  await c.login();

  runApp(GetMaterialApp(
    debugShowCheckedModeBanner: false,
    initialRoute: '/',
    theme: ThemeData(
        primaryColor: config.primaryColor,
        colorScheme:
            ColorScheme.fromSwatch(backgroundColor: config.backgroundColor)),
    getPages: [
      GetPage(
        name: '/login',
        page: () => LoginScreen(),
      ),
      GetPage(
        name: '/',
        page: () => MainScreen(),
        middlewares: [AuthService()],
        binding: BindingsBuilder(() {
          Get.put(MainController());
        }),
      ),
      GetPage(
        name: '/blueprint',
        page: () => BlueprintScreen(),
        binding: BindingsBuilder(() {
          final c = Get.find<BlueprintController>();

          var item = Get.arguments['item'];

          if (item == null) {
            c.periodic = Periodic();
            c.id = 0;
          } else {
            c.periodic = item;
            c.id = item.id;
          }
          c.init();
        }),
      ),
      GetPage(
        name: '/data',
        page: () => DataScreen(),
      ),
      GetPage(
        name: '/setting',
        page: () => SettingScreen(),
        binding: BindingsBuilder(() {
          Get.put(SettingController());
        }),
      ),
      GetPage(
        name: '/image',
        page: () => ImageScreen(),
        binding: BindingsBuilder(() {
          Get.put(ImageController());
        }),
      ),      
      GetPage(
        name: '/write',
        page: () => WriteScreen(),
        binding: BindingsBuilder(() {
          Get.put(WriteController());
          final PainterController painterController =
              Get.put(PainterController());
          var item = Get.arguments['item'];
          //c.item = item;
          painterController.blueprint = item;
          painterController.load();
        }),
      ),
    ],
  ));
}
