import 'dart:math';

import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:patrol/components/painter/painter_controller.dart';

class PainterDrawer extends CustomPainter {
  final PainterController c = Get.put(PainterController());

  Paint paintBlack = Paint()
    ..style = PaintingStyle.stroke
    ..color = Colors.black
    ..strokeWidth = 2;

  Paint paintRed = Paint()
    ..style = PaintingStyle.stroke
    ..color = Colors.red
    ..strokeWidth = 2;

  Paint paintGreen = Paint()
    ..style = PaintingStyle.stroke
    ..color = Colors.green
    ..strokeWidth = 2;

  Paint paintBlue = Paint()
    ..style = PaintingStyle.stroke
    ..color = const Color(0xff0000ff)
    ..strokeWidth = 2;

  Paint paintLightblue = Paint()
    ..style = PaintingStyle.stroke
    ..color = Colors.blue
    ..strokeWidth = 2;

  Paint paintSelectbox = Paint()
    ..style = PaintingStyle.stroke
    ..color = const Color(0xff7ad7f0)
    ..strokeWidth = 2;

  Paint paintSelect = Paint()
    ..style = PaintingStyle.stroke
    ..color = const Color(0xffa000a0)
    ..strokeWidth = 2;

  Paint paintFillSelect = Paint()
    ..style = PaintingStyle.fill
    ..color = const Color(0xffa000a0);

  Paint paintFillLightblue = Paint()
    ..style = PaintingStyle.fill
    ..color = Colors.blue;

  Paint paintFillRed = Paint()
    ..style = PaintingStyle.fill
    ..color = Colors.red;

  Paint paintFillWhite = Paint()
    ..style = PaintingStyle.fill
    ..color = Colors.white;

  @override
  void paint(Canvas canvas, Size size) {
    var zoom = c.zoom;
    var dx = -1 * (c.sx + c.currentSx) * zoom;
    var dy = -1 * (c.sy + c.currentSy) * zoom;

    if (c.isLoaded == true) {
      paintImage(
          canvas: canvas,
          rect:
              Rect.fromLTWH(dx, dy, c.imageWidth * zoom, c.imageHeight * zoom),
          image: c.image,
          fit: BoxFit.fill,
          repeat: ImageRepeat.noRepeat,
          scale: 1.0,
          alignment: Alignment.center,
          flipHorizontally: false,
          filterQuality: FilterQuality.high);
    }

    for (var i = 0; i < c.points.length; i++) {
      Path path = Path();

      var points = c.points[i];

      if (points.type == DrawType.icon || points.type == DrawType.number) {
        drawIcon(canvas, points);

        continue;
      }

      for (var j = 0; j < points.items.length; j++) {
        var point = points.items[j];
        var x = point.dx * zoom + dx;
        var y = point.dy * zoom + dy;
        if (j == 0) {
          path.moveTo(x, y);
        } else {
          path.lineTo(x, y);
        }
      }

      if (points.selected == true) {
        canvas.drawPath(path, paintSelect);
      } else if (points.color == LineColor.black) {
        canvas.drawPath(path, paintBlack);
      } else if (points.color == LineColor.red) {
        canvas.drawPath(path, paintRed);
      } else if (points.color == LineColor.blue) {
        canvas.drawPath(path, paintBlue);
      } else if (points.color == LineColor.green) {
        canvas.drawPath(path, paintGreen);
      } else if (points.color == LineColor.lightblue) {
        canvas.drawPath(path, paintLightblue);
      }
    }

    if (c.mode == Mode.select) {
      Path path = Path();

      path.moveTo(c.startSx, c.startSy);
      path.lineTo(c.endSx, c.startSy);
      path.lineTo(c.endSx, c.endSy);
      path.lineTo(c.startSx, c.endSy);
      path.lineTo(c.startSx, c.startSy);
      canvas.drawPath(path, paintSelectbox);
    }
  }

  drawIcon(canvas, points) {
    var zoom = c.zoom;
    var dx = -1 * (c.sx + c.currentSx) * zoom;
    var dy = -1 * (c.sy + c.currentSy) * zoom;
    var point = points.items[0];
    var x = point.dx * zoom + dx;
    var y = point.dy * zoom + dy;

    Paint paint;

    Path path = Path();

    var step = (c.iconZoom / 2.0) * zoom;
    var stepMiddle = step * 0.66;
    var stepSmall = step * 0.33;

    if (points.icon == 1 || points.icon == 2) {
      Color color = Colors.red;

      if (points.icon == 1) {
        paint = paintRed;
        color = Colors.red;
      } else {
        paint = paintBlue;
        color = const Color.fromRGBO(0, 0, 255, 1.0);
      }

      if (points.selected == true) {
        paint = paintSelect;
      }

      var offset = Offset(x, y);
      var textSpan = TextSpan(
        text: points.number.toString(),
        style: TextStyle(
            fontSize: step * 1.2, fontWeight: FontWeight.w600, color: color),
      );

      if (points.selected == true) {
        canvas.drawCircle(offset, step, paintFillSelect);
        canvas.drawCircle(offset, step, paintSelect);

        textSpan = TextSpan(
          text: points.number.toString(),
          style: TextStyle(
              fontSize: step * 1.2,
              fontWeight: FontWeight.w600,
              color: Colors.white),
        );
      } else {
        canvas.drawCircle(offset, step, paintFillWhite);
        canvas.drawCircle(offset, step, paint);
      }

      final textPainter = TextPainter()
        ..text = textSpan
        ..textDirection = TextDirection.ltr
        ..textAlign = TextAlign.center
        ..layout();

      offset = Offset(x - textPainter.width / 2, y - textPainter.height / 2);

      textPainter.paint(canvas, offset);

      return;
    } else if (points.icon == 101) {
      paint = paintLightblue;

      if (points.selected == true) {
        paint = paintSelect;
      }

      path.moveTo(x - step, y - step);
      path.lineTo(x + step, y + step);

      path.moveTo(x + step, y - step);
      path.lineTo(x - step, y + step);
    } else if (points.icon == 102) {
      paint = paintLightblue;

      if (points.selected == true) {
        paint = paintSelect;
      }

      path.moveTo(x - step, y);
      path.lineTo(x, y - step);
      path.lineTo(x + step, y);
      path.lineTo(x, y + step);
      path.lineTo(x - step, y);
    } else if (points.icon == 103) {
      paint = paintFillLightblue;

      if (points.selected == true) {
        paint = paintFillSelect;
      }

      path.moveTo(x + cos(30.0 * 3.14 / 180.0) * step,
          y + sin(30.0 * 3.14 / 180.0) * step);
      path.lineTo(x + cos(150.0 * 3.14 / 180.0) * step,
          y + sin(150.0 * 3.14 / 180.0) * step);
      path.lineTo(x + cos(270.0 * 3.14 / 180.0) * step,
          y + sin(270.0 * 3.14 / 180.0) * step);
      path.lineTo(x + cos(30.0 * 3.14 / 180.0) * step,
          y + sin(30.0 * 3.14 / 180.0) * step);
    } else if (points.icon == 104) {
      paint = paintFillRed;

      if (points.selected == true) {
        paint = paintFillSelect;
      }

      path.moveTo(x - step, y - step);
      path.lineTo(x + step, y - step);
      path.lineTo(x + step, y + step);
      path.lineTo(x - step, y + step);
      path.lineTo(x - step, y - step);
    } else if (points.icon == 105) {
      paint = paintBlue;

      if (points.selected == true) {
        paint = paintSelect;
      }

      canvas.drawCircle(Offset(x, y), stepMiddle, paint);
    } else if (points.icon == 106) {
      paint = paintBlue;

      if (points.selected == true) {
        paint = paintSelect;
      }

      canvas.drawCircle(Offset(x, y), step, paintBlue);

      path.moveTo(x + cos(30.0 * 3.14 / 180.0) * step,
          y + sin(30.0 * 3.14 / 180.0) * step);
      path.lineTo(x + cos(150.0 * 3.14 / 180.0) * step,
          y + sin(150.0 * 3.14 / 180.0) * step);
      path.lineTo(x + cos(270.0 * 3.14 / 180.0) * step,
          y + sin(270.0 * 3.14 / 180.0) * step);
      path.lineTo(x + cos(30.0 * 3.14 / 180.0) * step,
          y + sin(30.0 * 3.14 / 180.0) * step);
    } else if (points.icon == 107 ||
        points.icon == 108 ||
        points.icon == 109 ||
        points.icon == 110) {
      paint = paintBlue;

      if (points.selected == true) {
        paint = paintSelect;
      }

      canvas.drawCircle(Offset(x, y), stepSmall, paint);

      var angles = [135.0, 90.0, 45.0, 0.0];
      var angle = angles[points.icon - 107];

      path.moveTo(x + cos(angle * 3.14 / 180.0) * step,
          y + sin(angle * 3.14 / 180.0) * step);
      angle += 180.0;
      path.lineTo(x + cos(angle * 3.14 / 180.0) * step,
          y + sin(angle * 3.14 / 180.0) * step);
    } else if (points.icon == 111) {
      paint = paintFillRed;

      if (points.selected == true) {
        paint = paintFillSelect;
      }

      canvas.drawCircle(Offset(x, y), stepMiddle, paint);
    } else {
      paint = paintBlue;

      if (points.selected == true) {
        paint = paintSelect;
      }
    }

    canvas.drawPath(path, paint);
  }

  @override
  bool shouldRepaint(PainterDrawer oldDelegate) {
    return true;
  }
}
