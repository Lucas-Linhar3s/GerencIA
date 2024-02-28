// ignore_for_file: public_member_api_docs, sort_constructors_first
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class CustomSearch extends StatefulWidget {
  final bool? obscure;
  final IconData? eyeIcon;
  final IconData preffix;
  final TextEditingController? searchController;
  final Function()? iconOnPressed;
  final String? errorText;
  final Widget? preffixIconSearch;
  final String hintTextSearch;
  final Function(String) onChanged;
  final String? Function(String?)? validator;
  const CustomSearch({
    Key? key,
    this.obscure,
    this.eyeIcon,
    required this.preffix,
    this.searchController,
    this.iconOnPressed,
    this.errorText,
    this.preffixIconSearch,
    required this.hintTextSearch,
    required this.onChanged,
    this.validator,
  }) : super(key: key);

  @override
  State<CustomSearch> createState() => _CustomSearchState();
}

class _CustomSearchState extends State<CustomSearch> {
  @override
  Widget build(BuildContext context) {
    FocusNode myFocusNode = FocusNode();

    final screenSize = MediaQuery.of(context).size;
    final screenHeight = screenSize.height;
    final screenWidth = screenSize.width;
    final double fullHdBreakpointHeight = 769.00;

    return SizedBox(
      width: screenWidth < 600 ? 460 : 500,
      child: TextFormField(
        validator: widget.validator,
        obscureText: widget.obscure ?? false,
        controller: widget.searchController,
        textInputAction: TextInputAction.next,
        decoration: InputDecoration(
          errorText: widget.errorText,
          errorBorder: OutlineInputBorder(
              borderRadius: BorderRadius.circular(8),
              borderSide: BorderSide(
                color: Colors.red,
              )),
          prefixIcon: Padding(
            padding: EdgeInsets.only(left: 20, right: 20),
            child: Icon(widget.preffix, color: Colors.grey),
          ),
          fillColor: Color(0xfff2f2f2),
          contentPadding: EdgeInsets.only(left: 10, bottom: 20, top: 20),
          suffixIcon: Padding(
            padding: const EdgeInsets.only(right: 20.0),
            child: IconButton(
                onPressed: widget.iconOnPressed, icon: Icon(widget.eyeIcon)),
          ),
          border: OutlineInputBorder(borderRadius: BorderRadius.circular(8)),
          focusedBorder: OutlineInputBorder(
            borderSide: const BorderSide(color: Color(0xFF77B7CC), width: 2.0),
            borderRadius: BorderRadius.circular(8),
          ),
          hintText: widget.hintTextSearch,
          hintStyle: TextStyle(
              fontSize: screenHeight < fullHdBreakpointHeight ? 13 : 15),
          labelStyle: TextStyle(
              color:
                  myFocusNode.hasFocus ? Color(0xFF77B7CC) : Color(0xFF77B7CC)),
          filled: true,
        ),
        onChanged: widget.onChanged,
      ),
    );
  }
}
