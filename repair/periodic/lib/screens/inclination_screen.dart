import 'package:common_control/common_control.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:periodic/components/default_app_bar.dart';
import 'package:periodic/components/painter/painter_controller.dart';
import 'package:periodic/controllers/auth_controller.dart';
import 'package:periodic/controllers/blueprint_controller.dart';

class InclinationScreen extends CWidget {
  InclinationScreen({super.key});

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

  makeRow(int index, Point item, context) {
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
              value: 'delete',
              child: Text('삭제'),
            ),
          ]);

      // Implement the logic for each choice here
      switch (result) {
        case 'change':
          if (c.points[index].icon == 1) {
            c.points[index].icon = 2;
          } else {
            c.points[index].icon = 1;
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
          child: Text(item.number.toString()),
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
        onTap: () => clickDataNumber(index, 2, context),
        child: Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: Text(item.member),
        ),
      ),
      GestureDetector(
        behavior: HitTestBehavior.translucent,
        onTap: () => clickDataNumber(index, 3, context),
        child: Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: Text(item.shape),
        ),
      ),
      GestureDetector(
        behavior: HitTestBehavior.translucent,
        child: Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: Text(item.weight),
        ),
      ),
      GestureDetector(
        behavior: HitTestBehavior.translucent,
        onTap: () => clickData(index, 8, context),
        child: Container(
          padding: const EdgeInsets.all(10),
          child: Text(item.remark),
        ),
      ),
    ]);
  }

  Widget body(context) {
    List<TableRow> items = [];

    TableRow title = TableRow(children: [
      Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: const Text('측정위치')),
      Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: const Text('변위방향')),
      Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: const Text('변위량 (mm)')),
      Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: const Text('측정높이 (mm)')),
      Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: const Text('기울기')),
      Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(10),
          child: const Text('비고')),
    ]);

    items.add(title);

    for (var i = 0; i < c.points.length; i++) {
      Point item = c.points[i];

      if (item.icon < 200 || item.icon >= 300) {
        continue;
      }

      items.add(makeRow(i, item, context));
    }

    return Container(
        padding: const EdgeInsets.all(10),
        child: SingleChildScrollView(
            child: Table(
          columnWidths: const {
            0: FixedColumnWidth(100),
            1: FixedColumnWidth(150),
            2: FixedColumnWidth(150),
            3: FixedColumnWidth(150),
            4: FixedColumnWidth(150),
          },
          border: TableBorder.all(color: Colors.black),
          children: items,
        )));
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

    if (pos == 1) {
      var titles = ['좌측', '우측'];

      for (var i = 0; i < titles.length; i++) {
        final item = titles[i];

        Widget btn;

        if (text == item) {
          btn = OutlinedButton(
              style: OutlinedButton.styleFrom(
                backgroundColor: Colors.grey[700],
              ),
              onPressed: () => clickButton(index, pos, item),
              child: Text(item, style: const TextStyle(color: Colors.white)));
        } else {
          btn = OutlinedButton(
              style: OutlinedButton.styleFrom(
                backgroundColor: Colors.white,
              ),
              onPressed: () => clickButton(index, pos, item),
              child: Text(item, style: const TextStyle(color: Colors.black)));
        }

        final widget = SizedBox(width: 180, child: btn);

        items.add(widget);
      }
    }

    if (pos == 8) {
      var category = 20;

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
    }

    inputController.text = text;
    Widget input = Container(
      margin: const EdgeInsets.only(top: 10),
      width: 950,
      child: Row(children: [
        SizedBox(
          width: 300,
          child: TextField(
            controller: inputController,
            decoration: const InputDecoration(
              filled: true,
              fillColor: Colors.white,
            ),
          ),
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

    c.setData(pos, name);

    var str = '';
    var current = c.getCurrent();
    if (current.member != '' && current.shape != '') {
      var a = int.parse(current.member);
      var b = int.parse(current.shape);
      var c = b ~/ a;
      str = '1/$c';
    }
    c.setData(4, str);
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

  clickDataNumber(int index, int pos, context) {
    showDialog(
        context: context,
        barrierDismissible: true, // 바깥 영역 터치시 닫을지 여부
        builder: (BuildContext context) {
          return AlertDialog(
            backgroundColor: Colors.white,
            content: SizedBox(
              width: 400,
              height: 350,
              child: drawButtonsNumber(index, pos),
            ),
          );
        });
  }

  numberButton(text, index, pos) {
    return Expanded(
      child: CContainer(
        onTap: () => clickNumber(text, index, pos),
        padding: const EdgeInsets.all(10),
        decoration: BoxDecoration(
            shape: BoxShape.rectangle,
            border: const Border(
              right: BorderSide(color: Colors.black),
              bottom: BorderSide(color: Colors.black),
            )),
        child: Text(text,
            style: const TextStyle(fontSize: 20, fontWeight: FontWeight.bold),
            textAlign: TextAlign.center),
      ),
    );
  }

  drawButtonsNumber(int index, int pos) {
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

    c.tempText = text;

    return Column(children: [
      const SizedBox(height: 20),
      Row(children: [
        Expanded(
          child: Container(
            decoration: BoxDecoration(
              color: Colors.white,
              border: Border.all(color: Colors.black),
            ),
            padding: const EdgeInsets.all(10),
            child: Obx(
              () => Text(c.tempText,
                  textAlign: TextAlign.end,
                  style: const TextStyle(
                      fontSize: 20, fontWeight: FontWeight.bold)),
            ),
          ),
        ),
      ]),
      const SizedBox(height: 10),
      Container(
        decoration: BoxDecoration(
            color: Colors.white,
            shape: BoxShape.rectangle,
            border: const Border(
                top: BorderSide(color: Colors.black),
                left: BorderSide(color: Colors.black))),
        child: Column(
          children: [
            Row(
              children: [
                numberButton('1', index, pos),
                numberButton('2', index, pos),
                numberButton('3', index, pos),
              ],
            ),
            Row(
              children: [
                numberButton('4', index, pos),
                numberButton('5', index, pos),
                numberButton('6', index, pos),
              ],
            ),
            Row(
              children: [
                numberButton('7', index, pos),
                numberButton('8', index, pos),
                numberButton('9', index, pos),
              ],
            ),
            Row(
              children: [
                numberButton('0', index, pos),
                numberButton('<', index, pos),
                numberButton('CLEAR', index, pos),
              ],
            )
          ],
        ),
      ),
      const SizedBox(height: 20),
      Row(mainAxisAlignment: MainAxisAlignment.end, children: [
        ElevatedButton(
            style: ElevatedButton.styleFrom(
              backgroundColor: Colors.white,
              foregroundColor: Colors.black,
            ),
            onPressed: () => Get.back(),
            child: const Text('취소', style: TextStyle(color: Colors.black))),
        const SizedBox(width: 10),
        ElevatedButton(
            style: ElevatedButton.styleFrom(
              backgroundColor: Colors.grey[700], // Background color
            ),
            onPressed: () => clickSubmit(index, pos),
            child: const Text('확인', style: TextStyle(color: Colors.white))),
      ]),
    ]);
  }

  clickNumber(text, index, pos) {
    if (text == 'CLEAR') {
      c.tempText = '';
    } else if (text == '<') {
      if (c.tempText.isNotEmpty) {
        c.tempText = c.tempText.substring(0, c.tempText.length - 1);
      }
    } else {
      c.tempText = c.tempText + text;
    }
  }

  clickSubmit(index, pos) {
    var text = c.tempText;
    clickButton(index, pos, text);

    /*
    if (pos == 1) {
      c.points[index].part = text;
    } else if (pos == 2) {
      c.points[index].member = text;
    } else if (pos == 3) {
      c.points[index].shape = text;
    } else if (pos == 4) {
      c.points[index].weight = text;
    } else if (pos == 5) {
      c.points[index].length = text;
    } else if (pos == 6) {
      c.points[index].count = text;
    } else if (pos == 7) {
      c.points[index].progress = text;
    } else if (pos == 8) {
      c.points[index].remark = text;
    }

    c.updatePoints();

    Get.back();
    */
  }
}
