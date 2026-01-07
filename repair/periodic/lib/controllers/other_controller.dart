import 'dart:convert';

import 'package:common_control/common_control.dart';
import 'package:localstorage/localstorage.dart';
import 'package:periodic/controllers/blueprint_controller.dart';
import 'package:periodic/models/periodicother.dart';

import 'auth_controller.dart';

class OtherController extends GetxController {
  final _modified = false.obs;
  final _tab = 10.obs;
  final _periodicothers = <Periodicother>[].obs;

  bool get modified => _modified.value;
  set modified(value) => _modified.value = value;

  int get tab => _tab.value;
  set tab(value) => _tab.value = value;

  List<Periodicother> get periodicothers => _periodicothers;
  set periodicothers(value) => _periodicothers.value = value;

  List<TextEditingController> statusController = [];
  List<TextEditingController> positionController = [];

  updatePeriodicothers() {
    final authController = Get.find<AuthController>();
    if (authController.autosave == true) {
      save();
    }

    _periodicothers.refresh();
  }

  @override
  onInit() {
    super.onInit();

    final BlueprintController blueprintController =
        Get.find<BlueprintController>();

    for (var i = 0; i < blueprintController.periodicothers.length; i++) {
      Periodicother item = blueprintController.periodicothers[i];

      periodicothers.add(Periodicother(
          id: item.id,
          name: item.name,
          type: item.type,
          result: item.result,
          status: item.status,
          position: item.position,
          filename: item.filename,
          offlinefilename: item.offlinefilename,
          change: item.change,
          category: item.category,
          order: item.order,
          periodic: item.periodic,
          date: item.date));
      statusController.add(TextEditingController(text: item.status));
      positionController.add(TextEditingController(text: item.position));
    }
  }

  save() async {
    final periodicotherStr = json.encode(periodicothers);

    final LocalStorage storage = LocalStorage('blueprints.json');
    await storage.ready;
    await storage.setItem('periodicothers', periodicotherStr);

    final BlueprintController blueprintController =
        Get.find<BlueprintController>();

    blueprintController.periodicothers.clear();

    for (var i = 0; i < periodicothers.length; i++) {
      Periodicother item = periodicothers[i];

      blueprintController.periodicothers.add(Periodicother(
          id: item.id,
          name: item.name,
          type: item.type,
          result: item.result,
          status: item.status,
          position: item.position,
          filename: item.filename,
          offlinefilename: item.offlinefilename,
          change: item.change,
          category: item.category,
          order: item.order,
          periodic: item.periodic,
          date: item.date));
    }

    blueprintController.modified = true;
    blueprintController.modifiedOther = true;
    modified = false;
  }
}
