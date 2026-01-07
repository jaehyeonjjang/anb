import 'dart:io';

import 'package:common_control/common_control.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:image_picker/image_picker.dart';
import 'package:periodic/components/default_app_bar.dart';
import 'package:periodic/components/painter/painter.dart';
import 'package:periodic/controllers/auth_controller.dart';
import 'package:periodic/controllers/write_controller.dart';

enum DrawType { none, image, draw }

class Draw {
  DrawType type;
  String filename;
  Draw({required this.type, required this.filename});
}

class WriteScreen extends StatelessWidget {
  WriteScreen({super.key});

  final AuthController authController = Get.find<AuthController>();
  final PainterController painterController = Get.find<PainterController>();
  final c = Get.find<WriteController>();

  final now = DateTime.now();

  clickBack(context) {
    if (painterController.modified == false) {
      authController.title = '';
      Get.back();
      return true;
    }

    if (authController.autosave == true) {
      sleep(const Duration(milliseconds: 400));

      if (painterController.modified == false) {
        authController.title = '';
        Get.back();
        return true;
      }
    }

    final ret = showDialog<void>(
      context: context,
      builder: (context2) {
        return AlertDialog(
          title: const Text('데이터 저장'),
          backgroundColor: Colors.white,
          content: const Text(
              '작업내역이 저장되지 않았습니다.\n저장없이 종료하시겠습니까.\n저장없이 종료 선택시 작업한 내역이 모두 삭제됩니다'),
          actions: <Widget>[
            ElevatedButton(
              child: const Text('닫기'),
              onPressed: () {
                Navigator.pop(context2, false);
              },
            ),
            ElevatedButton(
              child: const Text('저장없이 종료'),
              onPressed: () {
                authController.title = '';
                Navigator.pop(context2, true);
                Get.back();
              },
            )
          ],
        );
      },
    );

    return ret;
  }

  @override
  Widget build(BuildContext context) {
    return WillPopScope(
        onWillPop: () {
          return clickBack(context);
        },
        child: CScaffold(
          appBar: DefaultAppBar(
            leading: IconButton(
                icon: const Icon(Icons.arrow_back_ios),
                onPressed: () {
                  clickBack(context);
                }),
            actions: [
              IconButton(
                icon: const Icon(Icons.photo),
                color: Colors.black,
                onPressed: () => Get.toNamed('/image'),
              ),
              Obx(() => authController.autosave == false
                  ? IconButton(
                      icon: const Icon(Icons.save),
                      color: Colors.black,
                      onPressed: () {
                        painterController.save();

                        Fluttertoast.showToast(
                            msg: '저장되었습니다',
                            toastLength: Toast.LENGTH_SHORT,
                            gravity: ToastGravity.CENTER,
                            timeInSecForIosWeb: 1,
                            backgroundColor: Colors.grey[700],
                            textColor: Colors.white,
                            fontSize: 16.0);
                      })
                  : Container()),
            ],
          ),
          backgroundColor: Colors.white,
          body: body(context),
        ));
  }

  body(context) {
    return Padding(
        padding: const EdgeInsets.all(10),
        child: Painter(
          filename: painterController.blueprint.offlinefilename,
        ));
  }

  clickSave(context) async {}

  final picker = ImagePicker();
}
