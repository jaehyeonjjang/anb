import 'package:common_control/common_control.dart';
import 'package:localstorage/localstorage.dart';
import 'package:patrol/controllers/blueprint_controller.dart';
import 'package:patrol/models/apt.dart';
import 'package:patrol/models/periodic.dart';
import 'package:patrol/models/user.dart';

class AuthController extends GetxController {
  final _authenticated = false.obs;
  final _user = User().obs;
  final _apt = Apt().obs;
  final _periodic = Periodic().obs;

  final _reboot = false.obs;

  final _title = ''.obs;

  final _autosave = true.obs;

  bool get autosave => _autosave.value;
  set autosave(value) => _autosave.value = value;

  final _autoclose = false.obs;

  bool get autoclose => _autoclose.value;
  set autoclose(value) => _autoclose.value = value;

  final _zoomlevel = 10.obs;

  int get zoomlevel => _zoomlevel.value;
  set zoomlevel(value) => _zoomlevel.value = value;

  bool get authenticated => _authenticated.value;
  set authenticated(value) => _authenticated.value = value;

  User get user => _user.value;
  set user(value) => _user.value = value;

  Apt get apt => _apt.value;
  set apt(value) => _apt.value = value;

  Periodic get periodic => _periodic.value;
  set periodic(value) => _periodic.value = value;

  String get title => _title.value;
  set title(value) => _title.value = value;

  bool get reboot => _reboot.value;
  set reboot(value) => _reboot.value = value;

  login() async {
    final LocalStorage storageSetting = LocalStorage('settings.json');
    await storageSetting.ready;
    final autosaveValue = await storageSetting.getItem('autosave');
    final zoomlevelValue = await storageSetting.getItem('zoomlevel');
    final autocloseValue = await storageSetting.getItem('autoclose');

    if (autosaveValue != null) {
      autosave = autosaveValue;
    }

    if (zoomlevelValue != null) {
      zoomlevel = zoomlevelValue;
    }

    if (autocloseValue != null) {
      autoclose = autocloseValue;
    }

    final storage = LocalStorage('login.json');
    await storage.ready;
    final userItem = await storage.getItem('user');
    final token = await storage.getItem('token');
    final aptItem = await storage.getItem('apt');

    if (userItem != null && userItem['id'] > 0) {
      authenticated = true;
      user = User.fromJson(userItem);

      CConfig().token = token;

      apt = Apt.fromJson(aptItem);

      final periodicItem = await storage.getItem('periodic');

      if (periodicItem != null && periodicItem['id'] > 0) {
        periodic = Periodic.fromJson(periodicItem);
      }
    }
  }

  logout() async {
    authenticated = false;
    user = User();
    apt = Apt();

    final LocalStorage storage = LocalStorage('login.json');
    await storage.ready;
    await storage.clear();
  }

  setTitle(item) {
    var top = '';

    if (periodic.id > 0) {
      top = periodic.apt.name;
      if (top == '') {
        top = apt.name;
      }
    } else {
      top = apt.name;
    }
    
    final blueprintController = Get.find<BlueprintController>();
    if (item.level == 1) {
      title = '$top - ${item.name}';
    } else {
      for (var i = 0; i < blueprintController.items.length; i++) {
        final parent = blueprintController.items[i];
        if (parent.id == item.parent) {
          title = '$top - ${parent.name} ${item.name}';
          break;
        }
      }
    }
  }
}
