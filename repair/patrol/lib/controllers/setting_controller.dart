import 'package:common_control/common_control.dart';
import 'package:patrol/controllers/auth_controller.dart';

class SettingController extends GetxController {
  final _autosave = true.obs;

  bool get autosave => _autosave.value;
  set autosave(value) => _autosave.value = value;

  final _zoomlevel = 10.obs;

  int get zoomlevel => _zoomlevel.value;
  set zoomlevel(value) => _zoomlevel.value = value;

  final _autoclose = false.obs;

  bool get autoclose => _autoclose.value;
  set autoclose(value) => _autoclose.value = value;

  @override
  onInit() {
    super.onInit();

    final c = Get.find<AuthController>();

    autosave = c.autosave;
    zoomlevel = c.zoomlevel;
    autoclose = c.autoclose;
  }
}
