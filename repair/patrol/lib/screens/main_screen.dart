import 'package:common_control/common_control.dart';
import 'package:patrol/components/default_app_bar.dart';
import 'package:patrol/components/search_input.dart';
import 'package:patrol/controllers/main_controller.dart';
import 'package:patrol/models/periodic.dart';

import 'package:patrol/controllers/auth_controller.dart';
import 'package:patrol/controllers/blueprint_controller.dart';
import 'package:patrol/models/user.dart';

class MainScreen extends StatelessWidget {
  MainScreen({Key? key}) : super(key: key);

  final AuthController authController = Get.find<AuthController>();
  final BlueprintController blueprintController =
      Get.find<BlueprintController>();
  final c = Get.find<MainController>();

  final textController = TextEditingController();

  Widget _renderItem(Periodic item, int index) {
    var date = '${item.resulttext4} ~ ${item.resulttext5}';
    if (date.length < 23) {
      date = '';
    }

    var statuss = ['준비', '준비', '착수', '완료', '중단'];

    Widget widget;

    if (authController.user.level == UserLevel.admin ||
        authController.user.level == UserLevel.rootadmin) {
      widget = CRow(gap: 10, children: [
        CText(item.apt.name,
            width: 500,
            textStyle:
                const TextStyle(fontSize: 25, fontWeight: FontWeight.w500)),
        Text(item.name, style: const TextStyle(fontSize: 20)),
      ]);
    } else {
      widget = Text(item.name, style: const TextStyle(fontSize: 20));
    }

    return CColumn(
        margin: const EdgeInsets.all(10),
        padding: const EdgeInsets.only(top: 10, bottom: 20),
        border: Border(bottom: BorderSide(color: Colors.grey[400]!, width: 1)),
        crossAxisAlignment: CrossAxisAlignment.start,
        onTap: () => Get.toNamed('/blueprint', arguments: {'item': item}),
        children: [
          Row(mainAxisAlignment: MainAxisAlignment.spaceBetween, children: [
            widget,
            CText(
                margin: const EdgeInsets.only(right: 10),
                '',
                textStyle: const TextStyle(fontSize: 16, color: Colors.blue)),
          ]),
          const SizedBox(height: 20),
          CRow(gap: 10, children: [
            Text(date,
                style: const TextStyle(fontSize: 16, color: Colors.black45)),
          ]),
        ]);
  }

  @override
  Widget build(BuildContext context) {
    Future.microtask(() {
      if (authController.periodic.id > 0) {
        Get.toNamed('/blueprint', arguments: {'item': authController.periodic});
      }
    });

    return CScaffold(
      appBar: DefaultAppBar(
        actions: [
          IconButton(
              icon: const Icon(Icons.settings),
              color: Colors.black,
              onPressed: () {
                Get.toNamed('/setting');
              }),
          IconButton(
            icon: const Icon(Icons.logout),
            color: Colors.black,
            onPressed: () {
              showDialog<void>(
                context: context,
                builder: (context2) {
                  return AlertDialog(
                    title: const Text('로그아웃'),
                    content: const Text('로그아웃 하시겠습니까'),
                    actions: <Widget>[
                      ElevatedButton(
                        child: const Text('취소'),
                        onPressed: () {
                          navigator!.pop(context2);
                        },
                      ),
                      ElevatedButton(
                        child: const Text('로그아웃'),
                        onPressed: () {
                          navigator!.pop(context2);

                          authController.logout();

                          Get.offAllNamed('/');
                        },
                      )
                    ],
                  );
                },
              );
            },
          ),
        ],
      ),
      body: Obx(() => body()),
      floatingActionButton: FloatingActionButton.extended(
        onPressed: () {
          Get.toNamed('/blueprint', arguments: {'item': null});
        },
        backgroundColor: Colors.red,
        label: const Text('순찰 시작', style: TextStyle(fontSize: 20)),
        icon: const Icon(Icons.add),
      ),
    );
  }

  body() {
    if (c.loading == false) {
      return const Center(
        child: SizedBox(
            width: 200,
            height: 250,
            child: Column(children: [
              SizedBox(
                  width: 150, height: 150, child: CircularProgressIndicator()),
              SizedBox(height: 20),
              Text('도면 데이터를 전송받는 중입니다'),
            ])),
      );
    }
    return CColumn(children: [
      SearchInput(
        controller: textController,
        hintText: '검색할 내용을 입력해주세요',
        onChanged: (value) {
          c.search = value;
        },
      ),
      Container(
          height: 10,
          decoration: BoxDecoration(
            border:
                Border(bottom: BorderSide(color: Colors.grey[400]!, width: 1)),
          )),
      const SizedBox(height: 10),
      Expanded(
          child: InfiniteScroll<Periodic>(builder: _renderItem, controller: c)),
    ]);
  }
}
