import 'package:common_control/common_control.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:localstorage/localstorage.dart';
import 'package:periodic/components/default_app_bar.dart';
import 'package:periodic/controllers/auth_controller.dart';
import 'package:periodic/controllers/setting_controller.dart';

class SettingScreen extends CWidget {
  SettingScreen({super.key});

  final c = Get.find<SettingController>();
  
  @override
  build(BuildContext context) {
    return CScaffold(appBar: DefaultAppBar(), backgroundColor: Colors.white, body: Obx(() => body()));
  }

  body() {
    final titleStyle = Style(
        margin: const EdgeInsets.only(left: 10),
        textStyle: const TextStyle(fontSize: 20, fontWeight: FontWeight.w500));

    final contentStyle = Style(
        margin: const EdgeInsets.only(left: 10),
        textStyle: const TextStyle(fontSize: 16, color: Colors.grey));

    List<Widget> zoomlevels = [];

    for (var i = 0; i < 15; i++) {
      zoomlevels.add(Radio(
        value: i + 1,
        groupValue: c.zoomlevel,
        onChanged: (Object? value) {
          c.zoomlevel = i + 1;
        },
      ));

      zoomlevels.add(Text('${i + 1}', style: const TextStyle(fontSize: 16)));
    }

    return CColumn(padding: const EdgeInsets.all(20.0), gap: 10, children: [
      CText('자동 저장', style: titleStyle),
      CText('작업마다 자동으로 저장됩니다. 추가로 저장 버튼을 클릭할 필요가 없습니다.', style: contentStyle),
      CRow(children: [
        Radio(
          value: true,
          groupValue: c.autosave,
          onChanged: (Object? value) {
            c.autosave = true;
          },
        ),
        const Text('사용', style: TextStyle(fontSize: 16)),
        Radio(
          value: false,
          groupValue: c.autosave,
          onChanged: (Object? value) {
            c.autosave = false;
          },
        ),
        const Text('사용 안함', style: TextStyle(fontSize: 16)),
      ]),
      Container(
          decoration: BoxDecoration(
              border: Border(
                  bottom: BorderSide(width: 1.0, color: Colors.grey[400]!)))),
      const SizedBox(height: 10),
      CText('전송후 자동 창닫기', style: titleStyle),
      CText('서버로 데이터 전송후 오류가 없을 경우 자동으로 팝업창을 닫습니다.', style: contentStyle),
      CRow(children: [
        Radio(
          value: true,
          groupValue: c.autoclose,
          onChanged: (Object? value) {
            c.autoclose = true;
          },
        ),
        const Text('사용', style: TextStyle(fontSize: 16)),
        Radio(
          value: false,
          groupValue: c.autoclose,
          onChanged: (Object? value) {
            c.autoclose = false;
          },
        ),
        const Text('사용 안함', style: TextStyle(fontSize: 16)),
      ]),
      Container(
          decoration: BoxDecoration(
              border: Border(
                  bottom: BorderSide(width: 1.0, color: Colors.grey[400]!)))),
      const SizedBox(height: 10),
      CText('순번 확대 레벨', style: titleStyle),
      CText('순번 확대 수준을 결정합니다. 숫자가 클수록 한번에 크게 확대됩니다.', style: contentStyle),
      CRow(children: zoomlevels),
      Container(
          decoration: BoxDecoration(
              border: Border(
                  bottom: BorderSide(width: 1.0, color: Colors.grey[400]!)))),
      const SizedBox(height: 10),
      ElevatedButton(
          onPressed: () => clickSave(),
          child: const Text('저장', style: TextStyle(fontSize: 16)))
    ]);
  }

  clickSave() async {
    final LocalStorage storageBlueprint = LocalStorage('settings.json');
    await storageBlueprint.ready;
    await storageBlueprint.setItem('autosave', c.autosave);
    await storageBlueprint.setItem('zoomlevel', c.zoomlevel);
    await storageBlueprint.setItem('autoclose', c.autoclose);

    final authController = Get.find<AuthController>();
    authController.autosave = c.autosave;
    authController.zoomlevel = c.zoomlevel;
    authController.autoclose = c.autoclose;

    Fluttertoast.showToast(
        msg: '저장되었습니다',
        toastLength: Toast.LENGTH_SHORT,
        gravity: ToastGravity.CENTER,
        timeInSecForIosWeb: 1,
        backgroundColor: Colors.grey[700],
        textColor: Colors.white,
        fontSize: 16.0);
  }
}
