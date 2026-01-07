import 'package:common_control/common_control.dart';
import 'package:flutter/cupertino.dart';
import 'package:periodic/components/painter/painter_controller.dart';
import 'package:periodic/controllers/auth_controller.dart';

class ToolBox extends StatelessWidget {
  ToolBox({super.key});

  final c = Get.find<PainterController>();
  final authController = Get.find<AuthController>();

  Widget drawButton(int index, String title) {
    if (index == c.index && c.mode == Mode.draw) {
      return InkWell(
          onTap: () {
            c.setMode(Mode.draw);
            c.setIndex(index);
          },
          child: Container(
              decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(5.0),
                color: Colors.yellow,
              ),
              padding: const EdgeInsets.all(5.0),
              child: Image.asset(title)));
    }

    return InkWell(
        onTap: () {
          c.setMode(Mode.draw);
          c.setIndex(index);
        },
        child: Container(
            padding: const EdgeInsets.all(5.0), child: Image.asset(title)));
  }

  Widget drawText(int index, String title, Color color) {
    if (index == c.index && c.mode == Mode.draw) {
      return InkWell(
          onTap: () {
            c.setMode(Mode.draw);
            c.setIndex(index);
          },
          child: Container(
              decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(5.0),
                color: Colors.yellow,
              ),
              padding: const EdgeInsets.all(5.0),
              child: CText(title, textStyle: TextStyle(color: color))));
    }

    return InkWell(
        onTap: () {
          c.setMode(Mode.draw);
          c.setIndex(index);
        },
        child: Container(
            padding: const EdgeInsets.all(5.0),
            child: CText(title, textStyle: TextStyle(color: color))));
  }

  Widget toolBoxNumber(int index, Color color) {
    Color background = Colors.white;

    if (index == c.index) {
      background = Colors.yellow;
    }

    return InkWell(
        onTap: () {
          c.setMode(Mode.draw);
          c.setIndex(index);
        },
        child: Container(
            decoration: BoxDecoration(
              borderRadius: BorderRadius.circular(5.0),
              color: background,
            ),
            padding: const EdgeInsets.symmetric(horizontal: 4.0, vertical: 1.0),
            child: Container(
                decoration: BoxDecoration(
                    color: Colors.white,
                    shape: BoxShape.circle,
                    border: Border.all(width: 1.5, color: color)),
                padding: const EdgeInsets.all(5),
                child:
                    Text('1', style: TextStyle(fontSize: 13, color: color)))));
  }

  Widget toolBoxNumberText(int index, Color color, String name) {
    Color background = Colors.white;

    if (index == c.index) {
      background = Colors.yellow;
    }

    return InkWell(
        onTap: () {
          c.setMode(Mode.draw);
          c.setIndex(index);
        },
        child: Container(
            decoration: BoxDecoration(
              borderRadius: BorderRadius.circular(5.0),
              color: background,
            ),
            padding: const EdgeInsets.symmetric(horizontal: 8.0, vertical: 1.0),
            child: Row(children: [
              Container(
                  decoration: BoxDecoration(
                      color: Colors.white,
                      shape: BoxShape.circle,
                      border: Border.all(width: 1.5, color: color)),
                  padding: const EdgeInsets.all(5),
                  child:
                      Text('1', style: TextStyle(fontSize: 13, color: color))),
              const SizedBox(width: 3),
              Text(name)
            ])));
  }

  @override
  Widget build(BuildContext context) {
    return Obx(() => body());
  }

  List<Widget> defect() {
    return [
      Row(children: [
        CText('수직부재',
            width: 65,
            textStyle: const TextStyle(fontSize: 14, color: Colors.black)),
        toolBoxNumber(basicVertical, Colors.red),
        drawText(basicVerticalLine, '직선', Colors.red),
        drawText(basicVerticalBreak, '꺾은선', Colors.red),
        Expanded(
            child: CContainer(
                onTap: () {
                  c.toolboxPositionToggle();
                },
                alignment: Alignment.centerRight,
                child: c.toolboxPosition == 0
                    ? const Icon(
                        Icons.arrow_right,
                        size: 40.0,
                      )
                    : const Icon(
                        Icons.arrow_left,
                        size: 40.0,
                      )))
      ]),
      Row(children: [
        CText('수평부재',
            width: 65,
            textStyle: const TextStyle(fontSize: 14, color: Colors.black)),
        toolBoxNumber(basicHorizontal, const Color.fromRGBO(0, 0, 255, 1.0)),
        drawText(basicHorizontalLine, '직선', Color.fromRGBO(0, 0, 255, 10.0)),
        drawText(basicHorizontalBreak, '꺾은선', Color.fromRGBO(0, 0, 255, 10.0))
      ]),
      const SizedBox(height: 5.0),
      Row(mainAxisAlignment: MainAxisAlignment.spaceBetween, children: [
        const SizedBox(width: 1),
        InkWell(
            onTap: () {
              c.setNumberZoom(
                  c.numberZoom + authController.zoomlevel.toDouble());
            },
            child: const Icon(CupertinoIcons.add, size: 24.0)),
        const SizedBox(height: 5.0),
        InkWell(
            onTap: () {
              c.setNumberZoom(
                  c.numberZoom - authController.zoomlevel.toDouble());
            },
            child: const Icon(CupertinoIcons.minus, size: 24.0)),
        const SizedBox(width: 2),
      ]),
      const SizedBox(height: 10.0),
      CContainer(
        height: 1.0,
        backgroundColor: Color.fromRGBO(200, 200, 200, 1.0),
      ),
      const SizedBox(height: 10.0),
      Row(children: [
        CText('균열누수',
            width: 65,
            textStyle: const TextStyle(fontSize: 14, color: Colors.black)),
        drawButton(crackLineRed, 'assets/imgs/i021.png'),
        drawButton(crackLineBlue, 'assets/imgs/i007.png'),
        drawButton(crackLineViolet, 'assets/imgs/i017.png'),
      ]),
      Row(children: [
        CText('',
            width: 65,
            textStyle: const TextStyle(fontSize: 14, color: Colors.black)),
        drawButton(crackCurveRed, 'assets/imgs/curve_red_os.png'),
        drawButton(crackCurveBlue, 'assets/imgs/curve_blue_os.png'),
        drawButton(crackCurveViolet, 'assets/imgs/curve_violet_os.png'),
      ]),
      const SizedBox(height: 5.0),
      Row(mainAxisAlignment: MainAxisAlignment.spaceBetween, children: [
        const SizedBox(width: 1),
        InkWell(
            onTap: () {
              c.setCrackZoom(c.crackZoom + authController.zoomlevel.toDouble());
            },
            child: const Icon(CupertinoIcons.add, size: 24.0)),
        const SizedBox(height: 5.0),
        InkWell(
            onTap: () {
              c.setCrackZoom(c.crackZoom - authController.zoomlevel.toDouble());
            },
            child: const Icon(CupertinoIcons.minus, size: 24.0)),
        const SizedBox(width: 2),
      ]),
      const SizedBox(height: 10.0),
      CContainer(
        height: 1.0,
        backgroundColor: Color.fromRGBO(200, 200, 200, 1.0),
      ),
      const SizedBox(height: 10.0),
      Row(children: [
        CText('곡선',
            width: 65,
            textStyle: const TextStyle(fontSize: 14, color: Colors.black)),
        drawButton(curveRed, 'assets/imgs/curve_red.png'),
        drawButton(curveBlue, 'assets/imgs/curve_blue.png'),
        drawButton(curveGreen, 'assets/imgs/curve_green.png'),
        drawButton(curveViolet, 'assets/imgs/curve_violet.png'),
      ]),
      Row(children: [
        CText('직선',
            width: 65,
            textStyle: const TextStyle(fontSize: 14, color: Colors.black)),
        drawButton(lineRed, 'assets/imgs/line_red.png'),
        drawButton(lineBlue, 'assets/imgs/line_blue.png'),
        drawButton(lineGreen, 'assets/imgs/line_green.png'),
        drawButton(lineViolet, 'assets/imgs/line_violet.png'),
      ]),
      Row(children: [
        CText('철근노출',
            width: 65,
            textStyle: const TextStyle(fontSize: 14, color: Colors.black)),
        drawButton(130, 'assets/imgs/i130.png'),
        drawButton(101, 'assets/imgs/i001.png'),
        drawButton(131, 'assets/imgs/i131.png'),
      ]),
      Row(children: [
        CText('부식',
            width: 65,
            textStyle: const TextStyle(fontSize: 14, color: Colors.black)),
        drawButton(132, 'assets/imgs/i132.png'),
        drawButton(102, 'assets/imgs/i002.png'),
      ]),
      Row(children: [
        CText('보',
            width: 65,
            textStyle: const TextStyle(fontSize: 14, color: Colors.black)),
        drawButton(133, 'assets/imgs/i133.png'),
        drawButton(103, 'assets/imgs/i003.png'),
      ]),
      Row(children: [
        CText('기타',
            width: 65,
            textStyle: const TextStyle(fontSize: 14, color: Colors.black)),
        drawButton(104, 'assets/imgs/i004.png'),
        drawButton(134, 'assets/imgs/i134.png'),
      ]),
      Row(children: [
        CText('배관누수',
            width: 65,
            textStyle: const TextStyle(fontSize: 14, color: Colors.black)),
        drawButton(111, 'assets/imgs/i011.png'),
        drawButton(112, 'assets/imgs/i302.png'),
      ]),
      Row(children: [
        CText('누수',
            width: 65,
            textStyle: const TextStyle(fontSize: 14, color: Colors.black)),
        drawButton(115, 'assets/imgs/i015.png'),
        drawButton(105, 'assets/imgs/i005.png'),
      ]),
      // Row(children: [
      //   CText('균열누수',
      //       width: 65,
      //       textStyle: const TextStyle(fontSize: 14, color: Colors.black)),
      //   drawButton(107, 'assets/imgs/i007.png'),
      //   drawButton(108, 'assets/imgs/i008.png'),
      //   drawButton(109, 'assets/imgs/i009.png'),
      //   drawButton(110, 'assets/imgs/i010.png'),
      // ]),
      // Row(children: [
      //   CText('',
      //       width: 65,
      //       textStyle: const TextStyle(fontSize: 14, color: Colors.black)),
      //   drawButton(117, 'assets/imgs/i017.png'),
      //   drawButton(118, 'assets/imgs/i018.png'),
      //   drawButton(119, 'assets/imgs/i019.png'),
      //   drawButton(120, 'assets/imgs/i020.png'),
      // ]),
    ];
  }

  List<Widget> inclination() {
    return [
      Row(children: [
        drawText(inclinationLine, '직선', Colors.black),
        drawText(inclinationHorizontal, '가로곡선', Colors.black),
        drawText(inclinationVertical, '세로곡선', Colors.black),
        Expanded(
            child: CContainer(
                onTap: () {
                  c.toolboxPositionToggle();
                },
                alignment: Alignment.centerRight,
                child: c.toolboxPosition == 0
                    ? const Icon(
                        Icons.arrow_right,
                        size: 40.0,
                      )
                    : const Icon(
                        Icons.arrow_left,
                        size: 40.0,
                      )))
      ]),
    ];
  }

  List<Widget> fiber() {
    return [
      Row(children: [
        CText('수직',
            width: 65,
            textStyle: const TextStyle(fontSize: 14, color: Colors.black)),
        drawButton(fiberVertical, 'assets/imgs/i301.png'),
        Expanded(
            child: CContainer(
                onTap: () {
                  c.toolboxPositionToggle();
                },
                alignment: Alignment.centerRight,
                child: c.toolboxPosition == 0
                    ? const Icon(
                        Icons.arrow_right,
                        size: 40.0,
                      )
                    : const Icon(
                        Icons.arrow_left,
                        size: 40.0,
                      )))
      ]),
      Row(children: [
        CText('수평',
            width: 65,
            textStyle: const TextStyle(fontSize: 14, color: Colors.black)),
        drawButton(fiberHorizontal, 'assets/imgs/i302.png'),
      ])
    ];
  }

  List<Widget> material() {
    return [
      Row(children: [
        CText('수직부재',
            width: 65,
            textStyle: const TextStyle(fontSize: 14, color: Colors.black)),
        toolBoxNumber(materialVertical, Colors.red),
        Expanded(
            child: CContainer(
                onTap: () {
                  c.toolboxPositionToggle();
                },
                alignment: Alignment.centerRight,
                child: c.toolboxPosition == 0
                    ? const Icon(
                        Icons.arrow_right,
                        size: 40.0,
                      )
                    : const Icon(
                        Icons.arrow_left,
                        size: 40.0,
                      )))
      ]),
      Row(children: [
        CText('수평부재',
            width: 65,
            textStyle: const TextStyle(fontSize: 14, color: Colors.black)),
        toolBoxNumber(materialHorizontal, const Color.fromRGBO(0, 0, 255, 1.0)),
      ])
    ];
  }

  Widget body() {
    if (c.toolbox == false) {
      return Container();
    }

    List<Widget> widgets = [];

    if (c.iconset == 1) {
      widgets = defect();
    } else if (c.iconset == 2) {
      widgets = inclination();
    } else if (c.iconset == 3) {
      widgets = fiber();
    } else if (c.iconset == 4) {
      widgets = material();
    }

    widgets.add(const SizedBox(height: 5));
    widgets
        .add(Row(mainAxisAlignment: MainAxisAlignment.spaceBetween, children: [
      const SizedBox(width: 1),
      InkWell(
          onTap: () {
            c.setIconZoom(c.iconZoom + authController.zoomlevel.toDouble());
          },
          child: const Icon(CupertinoIcons.add, size: 24.0)),
      const SizedBox(height: 5.0),
      InkWell(
          onTap: () {
            c.setIconZoom(c.iconZoom - authController.zoomlevel.toDouble());
          },
          child: const Icon(CupertinoIcons.minus, size: 24.0)),
      const SizedBox(width: 2),
    ]));

    return Column(children: [
      Container(
          margin: EdgeInsets.only(left: 5.0 + c.toolboxPosition, top: 5.0),
          padding: const EdgeInsets.only(
              left: 10.0, top: 10.0, bottom: 10.0, right: 5.0),
          width: 230,
          decoration: BoxDecoration(
              border: Border.all(width: 1.0),
              borderRadius: const BorderRadius.all(
                Radius.circular(5.0),
              ),
              color: Colors.white),
          child: Column(mainAxisSize: MainAxisSize.min, children: widgets)),
      // Container(
      //     margin: EdgeInsets.only(left: 5.0 + c.toolboxPosition, top: 20.0),
      //     width: 205,
      //     child: inputBox()),
    ]);
  }

  inputBox() {
    List<TableRow> widgets = [
      TableRow(children: [
        Container(
            alignment: Alignment.center,
            padding: const EdgeInsets.symmetric(vertical: 5),
            child: const Text('층')),
        Container(
            alignment: Alignment.center,
            padding: const EdgeInsets.symmetric(vertical: 5),
            child: const Text('SH')),
        Container(
            alignment: Alignment.center,
            padding: const EdgeInsets.symmetric(vertical: 5),
            child: const Text('N')),
        Container(
            alignment: Alignment.center,
            padding: const EdgeInsets.symmetric(vertical: 5),
            child: const Text('PS')),
      ])
    ];

    for (var i = 0; i < 10; i++) {
      var item = TableRow(children: [
        Container(
            height: 40,
            alignment: Alignment.center,
            padding: const EdgeInsets.all(5),
            child: TextField()),
        Container(
            height: 40,
            alignment: Alignment.center,
            padding: const EdgeInsets.all(5),
            child: TextField()),
        Container(
            height: 40,
            alignment: Alignment.center,
            padding: const EdgeInsets.all(5),
            child: TextField()),
        Container(
            height: 40,
            alignment: Alignment.center,
            padding: const EdgeInsets.all(5),
            child: TextField()),
      ]);

      widgets.add(item);
    }

    var height = Get.height - 350;

    return SizedBox(
      height: height,
      child: SingleChildScrollView(
        child: Container(
          color: const Color.fromRGBO(255, 204, 204, 1.0),
          child: Table(
            columnWidths: const {
              0: FixedColumnWidth(40),
              2: FixedColumnWidth(40),
              3: FixedColumnWidth(40),
            },
            border: TableBorder.all(color: Colors.black),
            children: widgets,
          ),
        ),
      ),
    );
  }
}
