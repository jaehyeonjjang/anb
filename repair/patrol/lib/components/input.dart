import 'package:flutter/material.dart';
import 'package:flutter/cupertino.dart';

class Input extends StatelessWidget {
  final String params = '';
  final TextEditingController? controller;
  final String hintText;
  final bool password;
  final void Function(String)? onChanged;

  const Input(
      {Key? key,
      this.controller,
      this.password = false,
      this.hintText = '',
      this.onChanged})
      : super(key: key);

  @override
  Widget build(context) {
    return TextField(
        onChanged: onChanged,
        controller: controller,
        obscureText: password,
        decoration: InputDecoration(
            filled: true,
            fillColor: Colors.white,
            suffixIcon: IconButton(
              icon: const Icon(CupertinoIcons.clear_circled,
                  color: Colors.black45),
              onPressed: () {
                if (controller != null) {
                  controller!.clear();
                }
              },
            ),
            hintText: hintText,
            isDense: true,
            focusedBorder: const OutlineInputBorder(
              borderSide: BorderSide(color: Colors.black87),
            ),
            border: const OutlineInputBorder(
                borderSide:
                    BorderSide(color: Color.fromARGB(255, 31, 150, 243)))));
  }
}
