import 'package:common_control/common_control.dart';
import 'package:patrol/models/periodic.dart';

class MainController extends InfiniteController {
  MainController() : super(reader: PeriodicManager.search, params: 'category=3');

  final _loading = true.obs;

  get loading => _loading.value;
  set loading(value) => _loading.value = value;

  final _search = ''.obs;

  get search => _search.value;
  set search(value) => _search.value = value;

  @override
  void onInit() async {
    super.onInit();

    debounce(_search, (_) {
      params = 'category=3&title=$search';
      reset();
    }, time: const Duration(milliseconds: 300));
  }
}
