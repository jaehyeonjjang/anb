import 'package:flutter/material.dart';
import 'package:flutter/cupertino.dart';

class SearchInput extends StatelessWidget {
  final String params = '';
  final TextEditingController? controller;
  final String hintText;
  final bool password;
  final void Function(String)? onChanged;

  const SearchInput(
      {Key? key,
      this.controller,
      this.password = false,
      this.hintText = '',
      this.onChanged})
      : super(key: key);

  @override
  Widget build(context) {
    return Container(
      padding: const EdgeInsets.only(left: 10, right: 10, top: 10, bottom: 5),
      color: Colors.white,
      child: Container(
        padding: const EdgeInsets.all(5),
        decoration: BoxDecoration(
            color: Colors.grey[200], borderRadius: BorderRadius.circular(20)),
        child: TextField(
            style: const TextStyle(fontSize: 20),
            onChanged: onChanged,
            controller: controller,
            obscureText: password,
            decoration: InputDecoration(
                filled: true,
                fillColor: Colors.grey[200],
                suffixIcon: IconButton(
                  icon: const Icon(CupertinoIcons.clear_circled,
                      color: Colors.black45),
                  onPressed: () {
                    if (controller != null) {
                      controller!.clear();
                    }

                    if (onChanged != null) {
                      onChanged!('');
                    }
                  },
                ),
                hintText: hintText,
                isDense: true,
                contentPadding: const EdgeInsets.only(
                    left: 10, right: 10, top: 10, bottom: 5),
                border: InputBorder.none)),
      ),
    );
  }
}
