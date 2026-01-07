import 'dart:convert';

import 'package:get/get.dart';
import 'package:localstorage/localstorage.dart';
import 'package:periodic/controllers/blueprint_controller.dart';
import 'package:periodic/models/periodicimage.dart';

class ImageController extends GetxController {
  final _item = Periodicimage().obs;
  final _images = <Periodicimage>[].obs;

  final _type = 2.obs;

  Periodicimage get item => _item.value;
  set item(value) => _item.value = value;

  List<Periodicimage> get images => _images;

  @override
  void onInit() async {
    super.onInit();

    final LocalStorage storage = LocalStorage('periodic.json');
    await storage.ready;
    final str = await storage.getItem('periodicimages');

    if (str == null || str == '') {
      return;
    }
    _images.value = json
        .decode(str)
        .map<Periodicimage>((json) => Periodicimage.fromJson(json))
        .toList();
  }

  set images(value) {
    _images.value = value;

    saveImage();
  }

  saveImage() async {
    final LocalStorage storage = LocalStorage('periodic.json');
    final str = json.encode(_images);

    await storage.ready;
    await storage.setItem('periodicimages', str);

    final BlueprintController blueprintController =
        Get.find<BlueprintController>();
    blueprintController.modified = true;
    blueprintController.modifiedImage = true;
  }

  int get type => _type.value;
  set type(value) => _type.value = value;

  removeImage(pos) {
    _images.removeAt(pos);
    _images.refresh();

    saveImage();
  }
}
