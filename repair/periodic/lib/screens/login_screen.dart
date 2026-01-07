import 'package:common_control/common_control.dart';
import 'package:localstorage/localstorage.dart';
import 'package:periodic/components/input.dart';
import 'package:periodic/config/config.dart' as config;
import 'package:periodic/controllers/auth_controller.dart';
import 'package:periodic/models/apt.dart';
import 'package:periodic/models/login.dart';
import 'package:permission_handler/permission_handler.dart';
import 'package:rflutter_alert/rflutter_alert.dart';

class LoginController extends GetxController {
  final txtLoginid = TextEditingController();
  final txtPasswd = TextEditingController();
}

class LoginScreen extends StatelessWidget {
  final authController = Get.find<AuthController>();
  final LoginController c = Get.put(LoginController());

  LoginScreen({super.key});

  clickLogin(context) async {
    if (c.txtLoginid.text.isEmpty) {
      Alert(
        context: context,
        type: AlertType.error,
        desc: '아이디를 입력하세요',
        buttons: [
          DialogButton(
            onPressed: () => Navigator.pop(context),
            width: 120,
            child: const Text(
              "확인",
              style: TextStyle(color: Colors.white, fontSize: 20),
            ),
          )
        ],
      ).show();

      return;
    }

    if (c.txtPasswd.text.isEmpty) {
      Alert(
        context: context,
        type: AlertType.error,
        desc: '패스워드를 입력하세요',
        buttons: [
          DialogButton(
            onPressed: () => Navigator.pop(context),
            width: 120,
            child: const Text(
              "확인",
              style: TextStyle(color: Colors.white, fontSize: 20),
            ),
          )
        ],
      ).show();

      return;
    }

    var user = await LoginManager.login(c.txtLoginid.text, c.txtPasswd.text);
    if (user.id > 0) {
      CConfig().token = user.extra['token'];

      authController.authenticated = true;
      authController.user = user;

      var apt = await AptManager.get(user.apt);
      authController.apt = apt;

      final storage = LocalStorage('login.json');
      await storage.ready;
      await storage.setItem('user', user.toJson());
      await storage.setItem('token', CConfig().token);
      await storage.setItem('apt', apt.toJson());

      Get.offAllNamed('/');
    } else {
      Alert(
        context: context,
        type: AlertType.error,
        desc: '로그인 정보가 정확하지 않습니다',
        style: const AlertStyle(backgroundColor: Colors.white),
        buttons: [
          DialogButton(
            onPressed: () => Navigator.pop(context),
            width: 120,
            child: const Text(
              "확인",
              style: TextStyle(color: Colors.white, fontSize: 20),
            ),
          )
        ],
      ).show();
    }
  }

  grant() async {
    if (!await Permission.mediaLibrary.request().isGranted) {
      final status = await Permission.mediaLibrary.request();

      if (status == PermissionStatus.granted) {
      } else if (status == PermissionStatus.denied) {
      } else if (status == PermissionStatus.permanentlyDenied) {
        await openAppSettings();
      }
    }

    if (!await Permission.camera.request().isGranted) {
      final status = await Permission.camera.request();

      if (status == PermissionStatus.granted) {
      } else if (status == PermissionStatus.denied) {
      } else if (status == PermissionStatus.permanentlyDenied) {
        await openAppSettings();
      }
    }

    if (!await Permission.photos.request().isGranted) {
      final status = await Permission.photos.request();

      if (status == PermissionStatus.granted) {
      } else if (status == PermissionStatus.denied) {
      } else if (status == PermissionStatus.permanentlyDenied) {
        await openAppSettings();
      }
    }
  }

  @override
  Widget build(context) {
    if (config.platform() == 'android') {
      Future.microtask(() => grant());
    }

    return Scaffold(
      backgroundColor: config.backgroundColor,
      resizeToAvoidBottomInset: true,
      body: Container(
          alignment: Alignment.center,
          child: SizedBox(
            width: 300,
            child: Column(
                mainAxisAlignment: MainAxisAlignment.center,
                crossAxisAlignment: CrossAxisAlignment.stretch,
                children: [
                  Padding(
                      padding: const EdgeInsets.all(10.0),
                      child: Image.asset('assets/imgs/logo.png',
                          width: 273, height: 122, fit: BoxFit.contain)),
                  Input(
                    controller: c.txtLoginid,
                    hintText: '아이디를 입력해주세요',
                  ),
                  const SizedBox(height: 8),
                  Input(
                      controller: c.txtPasswd,
                      hintText: '패스워드를 입력해주세요',
                      password: true),
                  const SizedBox(height: 10),
                  ElevatedButton(
                      onPressed: () => clickLogin(context),
                      child: const Padding(
                          padding: EdgeInsets.symmetric(vertical: 15),
                          child: Text('로그인', style: TextStyle(fontSize: 17)))),
                ]),
          )),
    );
  }
}
