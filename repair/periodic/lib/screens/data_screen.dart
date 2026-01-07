import 'dart:io';

import 'package:common_control/common_control.dart';
import 'package:flutter/cupertino.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:image_picker/image_picker.dart';
import 'package:periodic/components/default_app_bar.dart';
import 'package:periodic/components/painter/painter_controller.dart';
import 'package:periodic/controllers/auth_controller.dart';
import 'package:periodic/controllers/blueprint_controller.dart';

class DataScreen extends CWidget {
  DataScreen({super.key});

  final AuthController authController = Get.find<AuthController>();
  final c = Get.find<PainterController>();
  final blueprintController = Get.find<BlueprintController>();

  final TextEditingController inputController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return CScaffold(
      appBar: DefaultAppBar(
        leading: IconButton(
            icon: const Icon(Icons.arrow_back_ios),
            onPressed: () {
              Get.back();
            }),
        actions: [
          IconButton(
            icon: const Icon(Icons.photo),
            color: Colors.black,
            onPressed: () => Get.toNamed('/image'),
          ),
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
    );
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
                                          c.removeDataimageIndex(index, pos);

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

  makeRow(int index, Point item, context) {
    List<Widget> images = [];

    for (var i = 0; i < item.images.length; i++) {
      var image = item.images[i];

      images.add(imageWidget(image, index, i, context));
    }

    var style = const TextStyle(color: Colors.red);
    if (item.icon == 2 ||
        item.icon == basicHorizontalLine ||
        item.icon == basicHorizontalBreak) {
      style = const TextStyle(color: Color.fromRGBO(0, 0, 255, 1.0));
    }

    showPopup(context, details, index) async {
      Offset tapPosition = Offset.zero;

      final RenderBox referenceBox = context.findRenderObject() as RenderBox;
      tapPosition = referenceBox.globalToLocal(details.globalPosition);

      final RenderObject? overlay =
          Overlay.of(context).context.findRenderObject();

      final result = await showMenu(
          context: context,

          // Show the context menu at the tap location
          position: RelativeRect.fromRect(
              Rect.fromLTWH(tapPosition.dx, tapPosition.dy, 30, 30),
              Rect.fromLTWH(0, 0, overlay!.paintBounds.size.width,
                  overlay.paintBounds.size.height)),

          // set a list of choices for the context menu
          items: [
            const PopupMenuItem(
              value: 'change',
              child: Text('색상 변경'),
            ),
            const PopupMenuItem(
              value: 'delete',
              child: Text('삭제'),
            ),
          ]);

      // Implement the logic for each choice here
      switch (result) {
        case 'change':
          if (c.points[index].icon == 1) {
            c.points[index].icon = 2;
          } else if (c.points[index].icon == 2) {
            c.points[index].icon = 1;
          } else if (c.points[index].icon == 3) {
            c.points[index].icon = 4;
          } else if (c.points[index].icon == 4) {
            c.points[index].icon = 3;
          } else if (c.points[index].icon == 5) {
            c.points[index].icon = 6;
          } else if (c.points[index].icon == 6) {
            c.points[index].icon = 5;
          }

          c.updatePoints();
          c.modified = true;
          break;
        case 'delete':
          c.removePoint(index);
          break;
      }
    }

    return TableRow(children: [
      GestureDetector(
        onTapDown: (details) => showPopup(context, details, index),
        child: Container(
          alignment: Alignment.center,
          color: Colors.white,
          padding: const EdgeInsets.all(10),
          child: Text(item.number.toString(), style: style),
        ),
      ),
      GestureDetector(
        behavior: HitTestBehavior.translucent,
        onTap: () => clickData(index, 1, context),
        child: Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: Text(item.part),
        ),
      ),
      GestureDetector(
        behavior: HitTestBehavior.translucent,
        onTap: () => clickData(index, 2, context),
        child: Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: Text(item.member),
        ),
      ),
      GestureDetector(
        behavior: HitTestBehavior.translucent,
        onTap: () => clickData(index, 3, context),
        child: Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: Text(item.shape),
        ),
      ),
      GestureDetector(
        behavior: HitTestBehavior.translucent,
        onTap: () => clickData(index, 4, context),
        child: Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: Text(item.weight),
        ),
      ),
      GestureDetector(
        behavior: HitTestBehavior.translucent,
        onTap: () => clickData(index, 5, context),
        child: Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: Text(item.length),
        ),
      ),
      GestureDetector(
        behavior: HitTestBehavior.translucent,
        onTap: () => clickData(index, 6, context),
        child: Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: Text(item.count.toString()),
        ),
      ),
      GestureDetector(
        behavior: HitTestBehavior.translucent,
        onTap: () => clickData(index, 7, context),
        child: Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: Text(item.progress),
        ),
      ),
      GestureDetector(
        behavior: HitTestBehavior.translucent,
        onTap: () => clickData(index, 8, context),
        child: Container(
          alignment: Alignment.center,
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

  Widget body(context) {
    List<TableRow> items = [];

    TableRow title = TableRow(children: [
      Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: const Text('그룹')),
      Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: const Text('부위')),
      Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: const Text('부재')),
      Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: const Text('결함종류')),
      Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: const Text('폭')),
      Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: const Text('길이')),
      Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: const Text('개소')),
      Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: const Text('진행사항')),
      Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: const Text('비고')),
      Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: const Text('사진자료')),
    ]);

    items.add(title);

    for (var i = 0; i < c.points.length; i++) {
      Point item = c.points[i];

      if (item.type != DrawType.number && item.type != DrawType.numberLine) {
        continue;
      }

      items.add(makeRow(i, item, context));
    }

    return Container(
        padding: const EdgeInsets.all(10),
        child: SingleChildScrollView(
            child: Table(
          columnWidths: const {
            0: FixedColumnWidth(50),
            1: FixedColumnWidth(120),
            2: FixedColumnWidth(70),
            3: FixedColumnWidth(160),
            4: FixedColumnWidth(60),
            5: FixedColumnWidth(60),
            6: FixedColumnWidth(60),
            7: FixedColumnWidth(90),
            8: FixedColumnWidth(190),
          },
          border: TableBorder.all(color: Colors.black),
          children: items,
        )));
  }

  final picker = ImagePicker();

  Future getImage(int index, ImageSource imageSource) async {
    final image = await picker.pickImage(source: imageSource);

    if (image == null) {
      return;
    }

    var path = image.path;

    c.points[index].images.add(path);
    c.updatePoints();
    c.modified = true;
  }

  drawButtons(index, pos) {
    List<Widget> items = [];

    Point point = c.points[index];

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

    if (point.icon == 2 ||
        point.icon == basicHorizontalLine ||
        point.icon == basicHorizontalBreak) {
      if (pos == 2) {
        category = 10;
      } else if (pos == 3) {
        category = 11;
      }
    }

    for (var i = 0; i < blueprintController.datacategorys.length; i++) {
      final item = blueprintController.datacategorys[i];

      if (item.category != category) {
        continue;
      }

      Widget btn;

      if (text == item.name) {
        btn = OutlinedButton(
            style: OutlinedButton.styleFrom(
              backgroundColor: Colors.grey[700],
            ),
            onPressed: () => clickButton(index, pos, item.name),
            child:
                Text(item.name, style: const TextStyle(color: Colors.white)));
      } else {
        btn = OutlinedButton(
            style: OutlinedButton.styleFrom(
              backgroundColor: Colors.white,
            ),
            onPressed: () => clickButton(index, pos, item.name),
            child:
                Text(item.name, style: const TextStyle(color: Colors.black)));
      }

      final widget = SizedBox(width: 180, child: btn);

      items.add(widget);
    }

    inputController.text = text;
    Widget input = Container(
      margin: const EdgeInsets.only(top: 10),
      width: 950,
      child: Row(children: [
        SizedBox(
          width: 300,
          child: TextField(controller: inputController),
        ),
        const SizedBox(width: 10),
        ElevatedButton(
            onPressed: () => clickInput(index, pos),
            child: const Text('직접 입력')),
        Expanded(child: Container()),
        ElevatedButton(
            style: ElevatedButton.styleFrom(
              backgroundColor: Colors.grey[700], // Background color
            ),
            onPressed: () => Get.back(),
            child: const Text('닫기', style: TextStyle(color: Colors.white))),
      ]),
    );

    items.add(input);
    return items;
  }

  clickInput(index, pos) {
    final text = inputController.text;
    clickButton(index, pos, text);
  }

  clickButton(int index, int pos, String name) {
    c.current = index;

    if (pos == 3) {
      var point = c.points[index];

      point.shape = name;
      for (var i = 0; i < blueprintController.datacategorys.length; i++) {
        final item = blueprintController.datacategorys[i];

        if (pos == item.category && item.name == name) {
          point.remark = item.remark;
        }
      }

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

  clickData(int index, int pos, context) {
    showDialog(
        context: context,
        barrierDismissible: true, // 바깥 영역 터치시 닫을지 여부
        builder: (BuildContext context) {
          return AlertDialog(
            backgroundColor: Colors.white,
            content: SizedBox(
              width: 950,
              child: Wrap(
                  spacing: 10,
                  runSpacing: 0,
                  children: drawButtons(index, pos)),
            ),
          );
        });
  }
}
