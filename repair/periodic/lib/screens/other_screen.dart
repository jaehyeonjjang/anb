import 'dart:io';

import 'package:common_control/common_control.dart';
import 'package:flutter/cupertino.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:image_picker/image_picker.dart';
import 'package:periodic/components/default_app_bar.dart';
import 'package:periodic/controllers/auth_controller.dart';
import 'package:periodic/controllers/other_controller.dart';
import 'package:periodic/models/periodicother.dart';

class OtherScreen extends CWidget {
  OtherScreen({super.key});

  final AuthController authController = Get.find<AuthController>();
  final c = Get.find<OtherController>();

  clickBack(context) {
    if (c.modified == false) {
      Get.back();
      return true;
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
              Obx(() => authController.autosave == false
                  ? IconButton(
                      icon: const Icon(Icons.save),
                      color: Colors.black,
                      onPressed: () {
                        c.save();

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
          body: Obx(() => body(context)),
        ));
  }

  body(context) {
    final textStyle = Style(
      margin: const EdgeInsets.all(10),
      padding: const EdgeInsets.all(10),
      textStyle: const TextStyle(fontSize: 20),
    );

    final textStyleSelected = Style(
      margin: const EdgeInsets.all(10),
      padding: const EdgeInsets.all(10),
      textStyle: const TextStyle(fontSize: 20, fontWeight: FontWeight.w700),
      decoration: const BoxDecoration(
        border: Border(
          bottom: BorderSide(width: 1.5, color: Colors.blue),
        ),
      ),
    );

    return CColumn(children: [
      CRow(
        children: [
          CText('추락방지시설',
              style: c.tab == 10 ? textStyleSelected : textStyle,
              onTap: () => c.tab = 10),
          CText('도로포장',
              style: c.tab == 11 ? textStyleSelected : textStyle,
              onTap: () => c.tab = 11),
          CText('도로부 신축 이음부',
              style: c.tab == 12 ? textStyleSelected : textStyle,
              onTap: () => c.tab = 12),
          CText('환기구 등의 덮개',
              style: c.tab == 13 ? textStyleSelected : textStyle,
              onTap: () => c.tab = 13),
          CText('외벽 마감제',
              style: c.tab == 14 ? textStyleSelected : textStyle,
              onTap: () => c.tab = 14),
          CText('강재구조 노후',
              style: c.tab == 15 ? textStyleSelected : textStyle,
              onTap: () => c.tab = 15),
          CText('부대 점검사항',
              style: c.tab == 3 ? textStyleSelected : textStyle,
              onTap: () => c.tab = 3)
        ],
      ),
      CScroll(expanded: true, children: [checklist(context), list(context)])
    ]);
  }

  final picker = ImagePicker();

  Future getImage(int index, ImageSource imageSource) async {
    final image = await picker.pickImage(source: imageSource);

    if (image == null) {
      return;
    }

    var path = image.path;

    var offlinefilename = c.periodicothers[index].offlinefilename;
    if (offlinefilename == '') {
      offlinefilename = path;
    } else {
      var filenames = offlinefilename.split(',');
      filenames.add(path);

      offlinefilename = filenames.join(',');
    }

    c.periodicothers[index].offlinefilename = offlinefilename;
    c.periodicothers[index].change = 1;
    c.modified = true;
    c.updatePeriodicothers();
  }

  imageWidget(String filename, int index, int pos, context) {
    return CContainer(
        width: 20,
        height: 20,
        child: Image.file(File(filename)),
        onTap: () {
          showGeneralDialog(
              barrierDismissible: false,
              context: context,
              pageBuilder: (popContext, __, ___) {
                return Scaffold(
                    body: InkWell(
                  onTap: () {
                    navigator!.pop(popContext);
                  },
                  child: Stack(children: [
                    SizedBox(
                        width: double.infinity,
                        height: double.infinity,
                        child: Image.file(File(filename))),
                    Positioned(
                      top: 30.0,
                      right: 10.0,
                      child: CRow(children: [
                        IconButton(
                            icon: const Icon(CupertinoIcons.trash, size: 30.0),
                            onPressed: () {
                              showDialog<void>(
                                context: context,
                                builder: (context2) {
                                  return AlertDialog(
                                    title: const Text('이미지 삭제'),
                                    backgroundColor: Colors.white,
                                    content: const Text('이미지를 삭제하시겠습니까'),
                                    actions: <Widget>[
                                      ElevatedButton(
                                        child: const Text('취소'),
                                        onPressed: () {
                                          navigator!.pop(context2);
                                        },
                                      ),
                                      ElevatedButton(
                                        child: const Text('삭제'),
                                        onPressed: () {
                                          var offlinefilename = c
                                              .periodicothers[index]
                                              .offlinefilename;
                                          if (offlinefilename != '') {
                                            var filenames =
                                                offlinefilename.split(',');

                                            if (pos < filenames.length) {
                                              filenames.removeAt(pos);
                                            }

                                            offlinefilename =
                                                filenames.join(',');

                                            c.periodicothers[index]
                                                    .offlinefilename =
                                                offlinefilename;
                                            c.periodicothers[index].change = 1;
                                            c.modified = true;
                                            c.updatePeriodicothers();
                                          }

                                          navigator!.pop(context2);
                                          navigator!.pop(popContext);
                                        },
                                      )
                                    ],
                                  );
                                },
                              );
                            }),
                        IconButton(
                            icon: const Icon(Icons.close, size: 30.0),
                            onPressed: () {
                              navigator!.pop(popContext);
                            }),
                      ]),
                    ),
                  ]),
                ));
              });
        });
  }

  makeRow(index, Periodicother item, context) {
    List<Widget> images = [];

    if (item.offlinefilename != '') {
      final filenames = item.offlinefilename.split(',');
      for (var i = 0; i < filenames.length; i++) {
        var image = filenames[i];

        if (File(image).existsSync() != true) {
          continue;
        }

        images.add(imageWidget(image, index, i, context));
      }
    }

    var txt1 = '양호';
    var txt2 = '보통';
    if (item.type == 2) {
      txt1 = '없음';
      txt2 = '있음';
    }
    return TableRow(children: [
      Container(
        alignment: Alignment.centerLeft,
        color: Colors.white,
        padding: const EdgeInsets.all(10),
        margin: const EdgeInsets.only(top: 5),
        child: Text(item.name, style: const TextStyle(fontSize: 16)),
      ),
      Container(
        alignment: Alignment.center,
        color: Colors.white,
        child: CRow(children: [
          Radio(
            value: 1,
            groupValue: item.result,
            onChanged: (Object? value) {
              c.periodicothers[index].result = 1;
              c.periodicothers[index].change = 1;
              c.modified = true;
              c.updatePeriodicothers();
            },
          ),
          Text(txt1, style: const TextStyle(fontSize: 16)),
          Radio(
            value: 2,
            groupValue: item.result,
            onChanged: (Object? value) {
              c.periodicothers[index].result = 2;
              c.periodicothers[index].change = 1;
              c.modified = true;
              c.updatePeriodicothers();
            },
          ),
          Text(txt2, style: const TextStyle(fontSize: 16)),
        ]),
      ),
      Container(
        alignment: Alignment.center,
        color: Colors.white,
        child: TextField(
            controller: c.statusController[index],
            onChanged: (value) {
              c.periodicothers[index].status = value;
              c.periodicothers[index].change = 1;
              c.modified = true;
              c.updatePeriodicothers();
            },
            decoration: const InputDecoration(
                contentPadding: EdgeInsets.only(left: 10, right: 10),
                border: InputBorder.none)),
      ),
      Container(
        alignment: Alignment.center,
        color: Colors.white,
        child: TextField(
            controller: c.positionController[index],
            onChanged: (value) {
              c.periodicothers[index].position = value;
              c.periodicothers[index].change = 1;
              c.modified = true;
              c.updatePeriodicothers();
            },
            decoration: const InputDecoration(
                contentPadding: EdgeInsets.only(left: 10, right: 10),
                border: InputBorder.none)),
      ),
      Container(
        alignment: Alignment.center,
        padding: const EdgeInsets.all(10),
        child: CRow(children: [
          CContainer(
            child: const Icon(Icons.add_a_photo),
            onTap: () => getImage(index, ImageSource.camera),
          ),
          const SizedBox(width: 10),
          CContainer(
            child: const Icon(Icons.add_photo_alternate),
            onTap: () => getImage(index, ImageSource.gallery),
          ),
          CRow(
            margin: const EdgeInsets.only(left: 20),
            gap: 5,
            children: images,
          )
        ]),
      ),
    ]);
  }

  Widget list(context) {
    if (c.tab == 10 || c.tab == 11 || c.tab == 12 || c.tab == 15) {
      return Container();
    }

    var category = c.tab;

    if (c.tab == 13) {
      category = 2;
    } else if (c.tab == 14) {
      category = 1;
    }
    List<TableRow> items = [];

    TableRow title = TableRow(children: [
      Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: const Text('점검내용')),
      Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: const Text('점검결과')),
      Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: const Text('상태')),
      Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: const Text('해당위치')),
      Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: const Text('')),
    ]);

    items.add(title);

    for (var i = 0; i < c.periodicothers.length; i++) {
      Periodicother item = c.periodicothers[i];

      if (item.category != category) {
        continue;
      }

      items.add(makeRow(i, item, context));
    }

    return Container(
        padding: const EdgeInsets.all(10),
        child: SingleChildScrollView(
            child: Table(
          columnWidths: const {
            0: FixedColumnWidth(340),
            1: FixedColumnWidth(180),
            2: FixedColumnWidth(200),
            3: FixedColumnWidth(200),
          },
          border: TableBorder.all(color: Colors.black),
          children: items,
        )));
  }

  makeCheckRow(index, Periodicother item, context) {
    final tab = c.tab;
    List<Widget> images = [];

    if (item.offlinefilename != '') {
      final filenames = item.offlinefilename.split(',');
      for (var i = 0; i < filenames.length; i++) {
        var image = filenames[i];

        if (File(image).existsSync() != true) {
          continue;
        }

        images.add(imageWidget(image, index, i, context));
      }
    }

    // var txt1 = '양호';
    // var txt2 = '보통';
    // if (item.type == 2) {
    //   txt1 = '없음';
    //   txt2 = '있음';
    // }

    final names = item.name.split(',');

    List<Widget> widgets = [];

    if (item.type == 1 || item.type == 3) {
      for (var i = 0; i < names.length; i++) {
        widgets.add(Radio(
          value: names[i],
          groupValue: item.status,
          onChanged: (Object? value) {
            c.periodicothers[index].status = names[i];
            c.periodicothers[index].change = 1;
            c.modified = true;
            c.updatePeriodicothers();
          },
        ));

        widgets.add(Text(names[i], style: const TextStyle(fontSize: 16)));
      }
    } else {
      for (var i = 0; i < names.length; i++) {
        widgets.add(Checkbox(
          value: getCheckboxValue(tab, index, i, names),
          onChanged: (Object? value) {
            setCheckboxValue(names[i], tab, index, i, names);
            c.modified = true;
            c.updatePeriodicothers();
          },
        ));

        widgets.add(Text(names[i], style: const TextStyle(fontSize: 16)));
      }
    }

    return TableRow(children: [
      Container(
        alignment: Alignment.centerLeft,
        color: Colors.white,
        padding: const EdgeInsets.all(10),
        margin: const EdgeInsets.only(top: 5),
        child: Text(item.position, style: const TextStyle(fontSize: 16)),
      ),
      Container(
        alignment: Alignment.centerLeft,
        color: Colors.white,
        margin: const EdgeInsets.only(top: 5),
        child: CRow(children: widgets),
      ),
      item.name == "a,b,c,d,e"
          ? Container()
          : Container(
              alignment: Alignment.center,
              padding: const EdgeInsets.all(10),
              child: CRow(children: [
                CContainer(
                  child: const Icon(Icons.add_a_photo),
                  onTap: () => getImage(index, ImageSource.camera),
                ),
                const SizedBox(width: 10),
                CContainer(
                  child: const Icon(Icons.add_photo_alternate),
                  onTap: () => getImage(index, ImageSource.gallery),
                ),
                CRow(
                  margin: const EdgeInsets.only(left: 20),
                  gap: 5,
                  children: images,
                )
              ]),
            ),
    ]);
  }

  Widget checklist(context) {
    if (c.tab < 10) {
      return Container();
    }

    List<TableRow> items = [];

    TableRow title = TableRow(children: [
      Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: const Text('점검내용')),
      Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: const Text('상태')),
      Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: const Text('')),
    ]);

    items.add(title);

    for (var i = 0; i < c.periodicothers.length; i++) {
      Periodicother item = c.periodicothers[i];

      if (item.category != c.tab) {
        continue;
      }

      items.add(makeCheckRow(i, item, context));
    }

    return Container(
        padding: const EdgeInsets.all(10),
        child: SingleChildScrollView(
            child: Table(
          columnWidths: const {
            0: FixedColumnWidth(220),
            1: FixedColumnWidth(750),
          },
          border: TableBorder.all(color: Colors.black),
          children: items,
        )));
  }

  getCheckboxValue(category, index, i, names) {
    final str = names[i];

    final values = c.periodicothers[index].status.split(',');

    for (var i = 0; i < values.length; i++) {
      if (values[i] == str) {
        return true;
      }
    }

    return false;
  }

  setCheckboxValue(name, category, index, i, names) {
    final check = getCheckboxValue(category, index, i, names);

    final values = c.periodicothers[index].status.split(',');

    List<String> newValue = [];

    if (check == true) {
      for (var i = 0; i < values.length; i++) {
        if (values[i] == name) {
          continue;
        }

        newValue.add(values[i]);
      }
    } else {
      values.add(name);

      for (var i = 0; i < names.length; i++) {
        for (var j = 0; j < values.length; j++) {
          if (names[i] == values[j]) {
            newValue.add(values[j]);
          }
        }
      }
    }
    c.periodicothers[index].status = newValue.join(',');
    c.periodicothers[index].change = 1;
  }
}
