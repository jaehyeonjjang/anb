import 'package:common_control/common_control.dart';
import 'package:flutter/services.dart';
// import 'package:flutter_downloader/flutter_downloader.dart';
import 'package:periodic/components/painter/painter.dart';
import 'package:periodic/config/config.dart' as config;
import 'package:periodic/controllers/auth_controller.dart';
import 'package:periodic/controllers/blueprint_controller.dart';
import 'package:periodic/controllers/fiber_controller.dart';
import 'package:periodic/controllers/image_controller.dart';
import 'package:periodic/controllers/main_controller.dart';
import 'package:periodic/controllers/material_controller.dart';
import 'package:periodic/controllers/other_controller.dart';
import 'package:periodic/controllers/setting_controller.dart';
import 'package:periodic/controllers/write_controller.dart';
import 'package:periodic/screens/blueprint_screen.dart';
import 'package:periodic/screens/data_screen.dart';
import 'package:periodic/screens/fiber_screen.dart';
import 'package:periodic/screens/image_screen.dart';
import 'package:periodic/screens/inclination_screen.dart';
import 'package:periodic/screens/login_screen.dart';
import 'package:periodic/screens/main_screen.dart';
import 'package:periodic/screens/material_screen.dart';
import 'package:periodic/screens/other_screen.dart';
import 'package:periodic/screens/setting_screen.dart';
import 'package:periodic/screens/write_screen.dart';
import 'package:periodic/services/auth_service.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await SystemChrome.setPreferredOrientations([
    DeviceOrientation.landscapeLeft,
    DeviceOrientation.landscapeRight,
  ]);

  // await FlutterDownloader.initialize(debug: true, ignoreSsl: true);

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
          var item = Get.arguments['item'];
          final c = Get.find<BlueprintController>();
          c.periodic = item;
          c.id = item.id;
          c.init();
        }),
      ),
      GetPage(
        name: '/data',
        page: () => DataScreen(),
      ),
      GetPage(
        name: '/inclination',
        page: () => InclinationScreen(),
      ),
      GetPage(
        name: '/fiber',
        page: () => FiberScreen(),
        binding: BindingsBuilder(() {
          Get.put(FiberController());
        }),
      ),
      GetPage(
        name: '/material',
        page: () => MaterialScreen(),
        binding: BindingsBuilder(() {
          Get.put(MaterialController());
        }),
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
        name: '/other',
        page: () => OtherScreen(),
        binding: BindingsBuilder(() {
          Get.put(OtherController());
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
