import 'package:common_control/common_control.dart';
import 'package:periodic/components/painter/painter_controller.dart';
import 'package:periodic/controllers/blueprint_controller.dart';

class MaterialBox extends StatelessWidget {
  MaterialBox({super.key});

  final c = Get.find<PainterController>();
  final blueprintController = Get.find<BlueprintController>();

  final TextEditingController inputController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Obx(() => body(context));
  }

  Widget body(context) {
    if (c.materialbox == false) {
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
            alignment: Alignment.center,
            padding: const EdgeInsets.all(5),
            child: Text(item.number.toString())),
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

    return Positioned(
        bottom: 0.0,
        left: 0.0,
        child: Container(
          width: Get.width - 20,
          color: Colors.white,
          child: Column(children: [
            Table(
              columnWidths: const {
                0: FixedColumnWidth(100),
              },
              border: TableBorder.all(color: Colors.black),
              children: [
                TableRow(children: [
                  Container(
                      alignment: Alignment.center,
                      padding: const EdgeInsets.all(10),
                      child: const Text('측정위치')),
                  Container(
                      alignment: Alignment.center,
                      padding: const EdgeInsets.all(10),
                      child: const Text('도면표기 (mm)')),
                  Container(
                      alignment: Alignment.center,
                      padding: const EdgeInsets.all(10),
                      child: const Text('측정치수 (mm)')),
                  Container(
                      alignment: Alignment.center,
                      padding: const EdgeInsets.all(10),
                      child: const Text('마감재 유/무')),
                ]),
                TableRow(children: [
                  Container(
                    alignment: Alignment.center,
                    padding: const EdgeInsets.all(10),
                    child: Text(item.number.toString()),
                  ),
                  GestureDetector(
                    behavior: HitTestBehavior.translucent,
                    onTap: () => clickDataNumber(0, 1, context),
                    child: Container(
                      alignment: Alignment.center,
                      padding: const EdgeInsets.all(10),
                      child: Text(item.part),
                    ),
                  ),
                  GestureDetector(
                    behavior: HitTestBehavior.translucent,
                    onTap: () => clickDataNumber(0, 2, context),
                    child: Container(
                      alignment: Alignment.center,
                      padding: const EdgeInsets.all(10),
                      child: Text(item.member),
                    ),
                  ),
                  GestureDetector(
                    behavior: HitTestBehavior.translucent,
                    onTap: () => clickData(3, context),
                    child: Container(
                      alignment: Alignment.center,
                      padding: const EdgeInsets.all(10),
                      child: Text(item.shape),
                    ),
                  ),
                ]),
              ],
            )
          ]),
        ));
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

    if (pos == 3) {
      var titles = ['O', 'X'];

      for (var i = 0; i < titles.length; i++) {
        final item = titles[i];

        Widget btn;

        if (text == item) {
          btn = OutlinedButton(
              style: OutlinedButton.styleFrom(
                backgroundColor: Colors.grey[700],
              ),
              onPressed: () => clickButton(pos, item),
              child: Text(item, style: const TextStyle(color: Colors.white)));
        } else {
          btn = OutlinedButton(
              style: OutlinedButton.styleFrom(
                backgroundColor: Colors.white,
              ),
              onPressed: () => clickButton(pos, item),
              child: Text(item, style: const TextStyle(color: Colors.black)));
        }

        final widget = SizedBox(width: 190, child: btn);

        items.add(widget);
      }
    }

    inputController.text = text;
    Widget input = Container(
      margin: const EdgeInsets.only(top: 10),
      width: 1000,
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
            onPressed: () => clickInput(pos), child: const Text('직접 입력')),
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

  clickInput(pos) {
    final text = inputController.text;
    clickButton(pos, text);
  }

  clickButton(int pos, String name) {
    c.setData(pos, name);

    Get.back();
  }

  clickData(int pos, context) {
    showDialog(
        context: context,
        barrierDismissible: true, // 바깥 영역 터치시 닫을지 여부
        builder: (BuildContext context) {
          return AlertDialog(
            backgroundColor: Colors.white,
            content: SizedBox(
              width: 1000,
              child:
                  Wrap(spacing: 10, runSpacing: 0, children: drawButtons(pos)),
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
    return Expanded(child : CContainer(
        onTap: () => clickNumber(text, index, pos),
                  padding:const EdgeInsets.all(10),
                  decoration: BoxDecoration(
                    shape: BoxShape.rectangle,
                    border: const Border(
                      right: BorderSide(color: Colors.black),
                      bottom: BorderSide(color: Colors.black),
                    )),
                  child: Text(text, style: const TextStyle(fontSize: 20, fontWeight: FontWeight.bold), textAlign: TextAlign.center),
              ),);
  }

  drawButtonsNumber(int index, int pos) {
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

    c.tempText = text;

    return Column(children: [
        const SizedBox(height: 20),
        Row(children: [
        Expanded(child: Container(
          decoration: BoxDecoration(
            color: Colors.white,
            border: Border.all(color: Colors.black),
          ),
          padding: const EdgeInsets.all(10),
          child: Obx(() => Text(c.tempText,
            textAlign: TextAlign.end,
            style: const TextStyle(fontSize: 20, fontWeight: FontWeight.bold)),
          ),),),
          ]),
          const SizedBox(height: 10),
        Container(
          decoration: BoxDecoration(
            color: Colors.white,
                    shape: BoxShape.rectangle,
                    border: const Border(
                      top: BorderSide(color: Colors.black),
                      left: BorderSide(color: Colors.black)
                    )),
          child : Column(children: [
            Row(children: [
                numberButton('1', index, pos),
                numberButton('2', index, pos),
                numberButton('3', index, pos),
            ],),
            Row(children: [
                numberButton('4', index, pos),
                numberButton('5', index, pos),
                numberButton('6', index, pos),
            ],),
            Row(children: [
                numberButton('7', index, pos),
                numberButton('8', index, pos),
                numberButton('9', index, pos),
            ],),
            Row(children: [
                numberButton('0', index, pos),
                numberButton('<', index, pos),
                numberButton('CLEAR', index, pos),
            ],)
        ],),),
        const SizedBox(height: 20),
        Row(
          mainAxisAlignment: MainAxisAlignment.end,
          children: [
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
    clickButton(pos, text);
  }
}
