import 'dart:convert';
import 'dart:io';

import 'package:common_control/common_control.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:localstorage/localstorage.dart';
import 'package:percent_indicator/percent_indicator.dart';
import 'package:periodic/components/default_app_bar.dart';
import 'package:periodic/components/painter/painter_controller.dart';
import 'package:periodic/controllers/auth_controller.dart';
import 'package:periodic/controllers/blueprint_controller.dart';
import 'package:periodic/models/blueprint.dart';
import 'package:periodic/models/periodic.dart';
import 'package:periodic/models/periodicimage.dart';
import 'package:periodic/models/periodicother.dart';
import 'package:periodic/models/upload.dart';

class BlueprintScreen extends CWidget {
  BlueprintScreen({super.key});

  final AuthController authController = Get.find<AuthController>();
  final c = Get.find<BlueprintController>();

  endProcess() async {
    final LocalStorage storageLogin = LocalStorage('login.json');
    await storageLogin.ready;
    await storageLogin.deleteItem('periodic');

    authController.periodic = Periodic();
    authController.title = '';
  }

  clickBack(context) {
    if (c.modified == false) {
      endProcess();
      Get.back();
      return true;
    }

    final ret = showDialog<void>(
      context: context,
      builder: (context2) {
        return AlertDialog(
          title: const Text('데이터 전송'),
          backgroundColor: Colors.white,
          content: const Text(
              '작업내역이 서버로 전송되지 않았습니다.\n전송 없이 종료하시겠습니까.\n저장없이 종료 선택시 작업한 내역이 모두 삭제됩니다'),
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
                endProcess();
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
                icon: const Icon(Icons.apartment),
                color: Colors.black,
                onPressed: () {
                  if (c.loading == false) {
                    return;
                  }

                  Get.toNamed('/other');
                },
              ),
              IconButton(
                icon: const Icon(Icons.photo),
                color: Colors.black,
                onPressed: () {
                  if (c.loading == false) {
                    return;
                  }

                  Get.toNamed('/image');
                },
              ),
            ],
          ),
          backgroundColor: Colors.white,
          body: Obx(() => body(context)),
        ));
  }

  body(context) {
    if (c.loading == false) {
      return Column(crossAxisAlignment: CrossAxisAlignment.center, children: [
        Expanded(child: Container()),
        CColumn(
          height: 200,
          padding: const EdgeInsets.only(left: 20, right: 20),
          crossAxisAlignment: CrossAxisAlignment.center,
          children: [
            const SizedBox(height: 10),
            LinearPercentIndicator(
              animation: true,
              lineHeight: 20.0,
              animationDuration: 10,
              percent: c.percent,
              center: Text('${(c.percent * 100.0).toInt()}%',
                  style: const TextStyle(color: Colors.white)),
              barRadius: const Radius.circular(10),
              progressColor: Colors.green,
            ),
            const SizedBox(height: 50),
            const Text('도면 데이터를 전송받는 중입니다', style: TextStyle(fontSize: 16)),
            const SizedBox(height: 10),
            const Text('시간이 소요되니 잠시 기다려 주세요', style: TextStyle(fontSize: 16)),
          ],
        ),
        Expanded(child: Container()),
      ]);
    }

    var items = c.items.map((item) => renderItem(item)).toList();
    items.insert(
        0,
        CContainer(
          border: Border.all(color: Colors.black),
          backgroundColor: c.modifiedOther == true
              ? const Color.fromRGBO(255, 200, 200, 1.0)
              : Colors.white,
          width: double.infinity,
          margin: const EdgeInsets.only(bottom: 10),
          child: CText(
            '부대시설',
            margin: const EdgeInsets.all(20),
          ),
          onTap: () {
            Get.toNamed('/other');
          },
        ));

    items.insert(
        1,
        CContainer(
          border: Border.all(color: Colors.black),
          backgroundColor: c.modifiedImage == true
              ? const Color.fromRGBO(255, 200, 200, 1.0)
              : Colors.white,
          width: double.infinity,
          margin: const EdgeInsets.only(bottom: 10),
          child: CText(
            '사진 자료',
            margin: const EdgeInsets.all(20),
          ),
          onTap: () {
            Get.toNamed('/image');
          },
        ));

    return CFixedBottom(
        padding: const EdgeInsets.all(10),
        bottom: ElevatedButton(
          style: ElevatedButton.styleFrom(
            minimumSize: const Size.fromHeight(60), // NEW
          ),
          onPressed: () => clickSend(context),
          child: const Text(
            '전송',
            style: TextStyle(fontSize: 20),
          ),
        ),
        children: items);
  }

  Widget renderItem(Blueprint item) {
    return CContainer(
      border: Border.all(color: Colors.black),
      backgroundColor: item.extra["modified"] != null
          ? const Color.fromRGBO(255, 200, 200, 1.0)
          : Colors.white,
      width: double.infinity,
      margin: EdgeInsets.only(bottom: 10, left: (item.level - 1) * 50),
      child: CText(
        item.name,
        margin: const EdgeInsets.all(20),
      ),
      onTap: () {
        if (item.upload != 1) {
          return;
        }

        if (item.filename == '') {
          return;
        }

        authController.setTitle(item);
        Get.toNamed('/write', arguments: {'item': item});
      },
    );
  }

  Widget renderItemSend(int index) {
    final item = c.items[index];

    Color color = Colors.white;

    if (item.extra['modified'] != null) {
      color = const Color.fromRGBO(255, 200, 200, 1.0);

      if (item.checked == false) {
        color = const Color.fromRGBO(220, 220, 220, 1.0);
      }
    }
    return CContainer(
      border: Border.all(color: Colors.black),
      backgroundColor: color,
      width: double.infinity,
      margin: EdgeInsets.only(bottom: 10, left: (item.level - 1) * 50),
      child: CBothSide(children: [
        CText(item.name, margin: const EdgeInsets.all(20)),
        item.extra['modified'] != null
            ? Checkbox(
                onChanged: (value) => c.setCheck(index, value),
                value: item.checked)
            : Container()
      ]),
    );
  }

  clickSend(context) async {
    var find = false;
    for (var i = 0; i < c.items.length; i++) {
      final item = c.items[i];

      if (item.extra['modified'] != null) {
        find = true;
        break;
      }
    }

    if (c.modified == true) {
      find = true;
    }

    if (find == false) {
      showDialog<void>(
        context: context,
        builder: (context2) {
          return AlertDialog(
            title: const Text('데이터 전송'),
            backgroundColor: Colors.white,
            content: const Text('작업 내역이 없습니다',
                style: TextStyle(color: Colors.red, fontSize: 20)),
            actions: <Widget>[
              ElevatedButton(
                child: const Text('닫기'),
                onPressed: () {
                  navigator!.pop(context2);
                },
              ),
            ],
          );
        },
      );

      return;
    }

    c.setCheckAll();

    final h = Get.height;
    showDialog<void>(
      context: context,
      builder: (context2) {
        return AlertDialog(
          title: const Text('데이터 전송'),
          backgroundColor: Colors.white,
          content: CColumn(width: 800, height: h - 320, children: [
            CContainer(
                height: h - 400,
                child: SingleChildScrollView(
                    child: Obx(() => CColumn(
                        children: List.generate(c.items.length,
                            (index) => renderItemSend(index)))))),
            const SizedBox(height: 40),
            const Text('온라인 상태에서만 전송이 가능합니다. 데이터를 전송하시겠습니까.',
                style: TextStyle(color: Colors.red, fontSize: 20)),
          ]),
          actions: <Widget>[
            ElevatedButton(
              child: const Text('취소'),
              onPressed: () {
                navigator!.pop(context2);
              },
            ),
            ElevatedButton(
              child: const Text('데이터 전송'),
              onPressed: () {
                c.cancel = true;
                navigator!.pop(context2);

                sendData(context);
              },
            )
          ],
        );
      },
    );
  }

  sendData(context) {
    Future.microtask(() => sendDataProcess());

    showGeneralDialog(
        barrierDismissible: false,
        context: context,
        pageBuilder: (popContext, __, ___) {
          return Obx(() => AlertDialog(
                title: const Text('데이터 전송'),
                backgroundColor: Colors.white,
                content: SizedBox(
                  width: 950,
                  height: 70,
                  child: CColumn(
                      crossAxisAlignment: CrossAxisAlignment.center,
                      children: [
                        const SizedBox(height: 10),
                        LinearPercentIndicator(
                          animation: true,
                          lineHeight: 20.0,
                          animationDuration: 10,
                          percent: c.percent,
                          center: Text('${(c.percent * 100.0).toInt()}%',
                              style: const TextStyle(color: Colors.white)),
                          barRadius: const Radius.circular(10),
                          progressColor: Colors.green,
                        ),
                        const SizedBox(height: 10),
                        c.sendError == true
                            ? const Text('전송중 오류가 발생했습니다',
                                style: TextStyle(color: Colors.red))
                            : c.percent == 1.0
                                ? const Text('전송이 완료되었습니다')
                                : Container()
                      ]),
                ),
                actions: [
                  c.percent != 1.0
                      ? ElevatedButton(
                          onPressed: () {
                            if (c.sendError == true) {
                              navigator!.pop(popContext);
                              return;
                            }

                            showDialog<void>(
                              context: context,
                              builder: (context2) {
                                return AlertDialog(
                                  title: const Text('작업 취소'),
                                  backgroundColor: Colors.white,
                                  content: const Text('작업을 취소하시겠습니까'),
                                  actions: <Widget>[
                                    ElevatedButton(
                                      child: const Text('닫기'),
                                      onPressed: () {
                                        navigator!.pop(context2);
                                      },
                                    ),
                                    ElevatedButton(
                                      child: const Text('작업 취소'),
                                      onPressed: () {
                                        c.cancel = true;
                                        navigator!.pop(context2);
                                        navigator!.pop(popContext);

                                        Fluttertoast.showToast(
                                            msg: '취소되었습니다',
                                            toastLength: Toast.LENGTH_SHORT,
                                            gravity: ToastGravity.CENTER,
                                            timeInSecForIosWeb: 1,
                                            backgroundColor: Colors.red,
                                            textColor: Colors.white,
                                            fontSize: 16.0);
                                      },
                                    )
                                  ],
                                );
                              },
                            );
                          },
                          child: const Text('취소'))
                      : ElevatedButton(
                          onPressed: () => clickSendFinish(),
                          child: const Text('닫기'))
                ],
              ));
        });
  }

  sendDataProcess() async {
    c.cancel = false;
    c.percent = 0.0;
    c.sendError = false;

    final LocalStorage storage = LocalStorage('periodic.json');

    var total = 1;
    var current = 0;

    List<dynamic> datas = [];

    for (var i = 0; i < c.items.length; i++) {
      if (c.items[i].checked == false) {
        continue;
      }

      final blueprint = c.items[i].id;

      await storage.ready;
      final data = await storage.getItem('data_$blueprint');
      final save = await storage.getItem('save_$blueprint');

      if (save == null || save == '') {
        continue;
      }

      if (data == null || data == '') {
        continue;
      }

      final item = json.decode(data);
      datas.add(item);

      final pointsPtr = item['points'];
      if (pointsPtr == null) {
        continue;
      }

      List<dynamic> points = pointsPtr as List<dynamic>;
      for (var j = 0; j < points.length; j++) {
        Point point = Point.fromJson(points[j]);
        total += point.images.length;
      }
    }

    for (var i = 0; i < c.periodicothers.length; i++) {
      Periodicother item = c.periodicothers[i];

      if (item.offlinefilename == '') {
        continue;
      }

      var filenames = item.offlinefilename.split(',');
      total += filenames.length;
    }

    storage.ready;
    final str = await storage.getItem('periodicimages');

    List<Periodicimage> images = [];
    if (str != null && str != '') {
      images = json
          .decode(str)
          .map<Periodicimage>((json) => Periodicimage.fromJson(json))
          .toList();

      total += images.length;

      for (var i = 0; i < images.length; i++) {
        Periodicimage image = images[i];
        if (c.cancel == true) {
          return;
        }

        var ret = '';
        for (var k = 0; k < 5; k++) {
          if (File(image.offlinefilename).existsSync() != true) {
            break;
          }

          ret = await UploadManager.image(image.offlinefilename);

          if (ret != '') {
            break;
          }

          if (k == 4) {
            c.sendError = true;
            break;
          }

          for (var j = 0; j < k + 1; j++) {
            if (c.cancel == true) {
              return;
            }
            sleep(const Duration(seconds: 1));
          }
        }

        images[i].filename = ret;

        current++;
        c.percent = current / total;
      }
    }

    for (var i = 0; i < datas.length; i++) {
      List<dynamic> points = datas[i]['points'];
      for (var j = 0; j < points.length; j++) {
        Point point = Point.fromJson(points[j]);

        datas[i]['points'][j]['onlineimages'] = [];
        for (var k = 0; k < point.images.length; k++) {
          final image = point.images[k];

          var ret = '';
          for (var l = 0; l < 5; l++) {
            if (File(image).existsSync() != true) {
              break;
            }

            ret = await UploadManager.image(image);

            if (ret != '') {
              break;
            }

            if (l == 4) {
              c.sendError = true;
              break;
            }

            for (var m = 0; m < l + 1; m++) {
              if (c.cancel == true) {
                return;
              }
              sleep(const Duration(seconds: 1));
            }
          }

          datas[i]['points'][j]['onlineimages'].add(ret);

          current++;
          c.percent = current / total;
        }
      }
    }

    for (var i = 0; i < c.periodicothers.length; i++) {
      Periodicother item = c.periodicothers[i];

      if (item.offlinefilename == '') {
        continue;
      }

      var filenames = item.offlinefilename.split(',');

      List<String> onlinefilenames = [];

      for (var j = 0; j < filenames.length; j++) {
        var image = filenames[j];
        if (File(image).existsSync() != true) {
          break;
        }

        final ret = await UploadManager.image(image);
        onlinefilenames.add(ret);

        current++;
        c.percent = current / total;
      }

      c.periodicothers[i].filename = onlinefilenames.join(',');
    }

    Map<String, dynamic> ret = {
      'user': authController.user.id,
      'id': c.id,
      'datas': datas,
      'images': images,
      'periodicothers': c.periodicothers
    };

    if (c.cancel == true) {
      return;
    }

    for (var i = 0; i < 5; i++) {
      if (c.cancel == true) {
        return;
      }

      final res = await Http.post('/api/periodic/upload', ret);
      if (res['code'] == 'ok') {
        c.percent = 1.0;
        if (authController.autoclose == true) {
          clickSendFinish();
        }
        return;
      }

      for (var j = 0; j < i + 1; j++) {
        if (c.cancel == true) {
          return;
        }
        sleep(const Duration(seconds: 1));
      }
    }

    c.sendError = true;
  }

  clickSendFinish() async {
    if (c.sendError == true) {
      Get.back();
      c.sendError = false;
      return;
    }

    endProcess();

    c.modified = false;
    Get.back();
    Get.back();
  }
}
