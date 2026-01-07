import 'dart:io';

import 'package:common_control/common_control.dart';
import 'package:flutter/cupertino.dart';
import 'package:image_picker/image_picker.dart';
import 'package:patrol/components/painter/painter_controller.dart';
import 'package:patrol/controllers/blueprint_controller.dart';

class DataBox extends StatelessWidget {
  DataBox({super.key});

  final c = Get.find<PainterController>();
  final blueprintController = Get.find<BlueprintController>();

  final TextEditingController inputController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Obx(() => body(context));
  }

  imageWidget(String filename, int index, context) {
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
                                          c.removeDataimage(index);

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

  Widget body(context) {
    if (c.databox == false) {
      return Container();
    }

    List<TableRow> lists = [];

    final widget = TableRow(children: [
      Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(5),
          child: const Text('순번')),
      Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(5),
          child: const Text('그룹/번호')),
      Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(5),
          child: const Text('지정')),
    ]);
    lists.add(widget);

    for (var i = 0; i < c.points.length; i++) {
      final item = c.points[i];

      if (item.type != DrawType.number) {
        continue;
      }

      if (!item.selected) {
        continue;
      }

      final widget = TableRow(children: [
        Container(
            alignment: Alignment.center,
            padding: const EdgeInsets.all(5),
            child: Text((i + 1).toString())),
        Container(
            alignment: Alignment.centerLeft,
            padding: const EdgeInsets.all(5),
            child: Text(item.remark)),
        Container(
          alignment: Alignment.center,
          height: 20,
          padding: const EdgeInsets.only(top: 6),
          child: Checkbox(
            value: c.points[i].grouped,
            onChanged: (value) {
              c.points[i].grouped = value;
              c.points.refresh();
            },
          ),
        ),
      ]);
      lists.add(widget);
    }

    if (c.current < 0 || c.current >= c.points.length) {
      return Container();
    }

    Point item = c.points[c.current];

    List<Widget> images = [];

    for (var i = 0; i < item.images.length; i++) {
      var image = item.images[i];

      images.add(imageWidget(image, i, context));
    }

    var width = Get.width / 3 * 2;

    return Positioned(
        bottom: 0.0,
        left: 0.0,
        child: Container(
          width: Get.width - 20,
          color: Colors.white,
          child: Column(children: [
            Table(
              columnWidths: {
                0: const FixedColumnWidth(50),
                1: FixedColumnWidth(width),                
              },
              border: TableBorder.all(color: Colors.black),
              children: [
                TableRow(children: [
                  Container(
                      alignment: Alignment.center,
                      padding: const EdgeInsets.all(10),
                      child: const Text('그룹')),
                  Container(
                      alignment: Alignment.center,
                      padding: const EdgeInsets.all(10),
                      child: const Text('내용')),
                  Container(
                      alignment: Alignment.center,
                      padding: const EdgeInsets.all(10),
                      child: const Text('사진자료')),
                ]),
                TableRow(children: [
                  Container(
                    alignment: Alignment.center,
                    padding: const EdgeInsets.all(10),
                    child: Text(item.number.toString()),
                  ),
                  GestureDetector(
                    behavior: HitTestBehavior.translucent,
                    onTap: () => clickData(8, context),
                    child: Container(
                      alignment: Alignment.centerLeft,
                      padding: const EdgeInsets.all(10),
                      child: Text(item.remark),
                    ),
                  ),
                  Container(
                    alignment: Alignment.center,
                    padding: const EdgeInsets.all(10),
                    child: CRow(children: [
                      CContainer(
                        child: const Icon(Icons.add_a_photo),
                        onTap: () => getImage(ImageSource.camera),
                      ),
                      const SizedBox(width: 10),
                      CContainer(
                        child: const Icon(Icons.add_photo_alternate),
                        onTap: () => getImage(ImageSource.gallery),
                      ),
                      CRow(
                        margin: const EdgeInsets.only(left: 20),
                        gap: 5,
                        children: images,
                      )
                    ]),
                  ),
                ]),
              ],
            )
          ]),
        ));
  }

  final picker = ImagePicker();

  Future getImage(ImageSource imageSource) async {
    final image = await picker.pickImage(source: imageSource);

    if (image == null) {
      return;
    }

    var path = image.path;

    final point = c.getCurrent();
    point.images.add(path);
    c.updatePoints();
    c.modified = true;
  }

  drawButtons(pos) {
    List<Widget> items = [];

    final point = c.getCurrent();
    var text = '';

    if (pos == 1) {
      text = point.part;
    } else if (pos == 2) {
      text = point.member;
    } else if (pos == 3) {
      text = point.shape;
    } else if (pos == 4) {
      text = point.weight;
    } else if (pos == 5) {
      text = point.length;
    } else if (pos == 6) {
      text = point.count;
    } else if (pos == 7) {
      text = point.progress;
    } else if (pos == 8) {
      text = point.remark;
    }

    var category = pos;

    if (point.icon == 2) {
      if (pos == 2) {
        category = 10;
      } else if (pos == 3) {
        category = 11;
      }
    }

    print('number : ${point.icon}');
    print('category : $category');

    inputController.text = text;
    Widget input = Container(
      margin: const EdgeInsets.only(top: 10),
      width: 950,
      child: Row(children: [
        Expanded(
          child: TextField(controller: inputController),
        ),
        const SizedBox(width: 10),
        ElevatedButton(
            onPressed: () => clickInput(pos), child: const Text('입력')),        
        const SizedBox(width: 10),
        ElevatedButton(
            style: ElevatedButton.styleFrom(
              backgroundColor: Colors.grey[700], // Background color
            ),
            onPressed: () => Get.back(),
            child: const Text('닫기')),
      ]),
    );

    items.add(input);
    return items;
  }

  clickInput(pos) {
    final text = inputController.text;
    clickButton(pos, text);
  }

  clickButton(int pos, String name) {
    if (pos == 3) {
      var point = c.getCurrent();

      point.shape = name;
      
      var strs = name.split('/');

      point.progress = 'X';
      for (var i = 0; i < strs.length; i++) {
        if (strs[i] == '누수') {
          point.progress = 'O';
          break;
        }
      }

      c.setCurrent(point);
      c.modified = true;
    } else {
      c.setData(pos, name);
    }
    Get.back();
  }

  clickData(int pos, context) {
    showDialog(
        context: context,
        barrierDismissible: true, // 바깥 영역 터치시 닫을지 여부
        builder: (BuildContext context) {
          return AlertDialog(
            content: SizedBox(
              width: 950,
              child:
                  Wrap(spacing: 10, runSpacing: 0, children: drawButtons(pos)),
            ),
          );
        });
  }
}
