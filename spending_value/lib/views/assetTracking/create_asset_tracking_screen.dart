import 'dart:io';

import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:image_picker/image_picker.dart';
import 'package:provider/provider.dart';
import 'package:spending_value/controllers/assetTracking/create_asset_tracking_store.dart';
import 'package:spending_value/extensions/date_extension.dart';
import 'package:spending_value/models/camera_capture.dart';

class CreateAssetTrackingScreen extends StatefulWidget {
  const CreateAssetTrackingScreen({super.key});

  @override
  State<CreateAssetTrackingScreen> createState() =>
      _CreateAssetTrackingScreenState();
}

class _CreateAssetTrackingScreenState extends State<CreateAssetTrackingScreen> {
  List<Map<String, dynamic>> _foundUsers = [];
  SearchController controllerSearch = SearchController();
  TextEditingController controllerDateBought = TextEditingController();

  void _addressControllerListener() {
    setState(() {
      _foundUsers = DummyData().runFilter(controllerSearch.text);
    });
  }

  Future<DateTime?> _pickDateTime(BuildContext context) async {
    return await showDatePicker(
      context: context,
      initialDate: DateTime.now(),
      firstDate: DateTime(DateTime.now().year - 100, 1),
      lastDate: DateTime(DateTime.now().year + 100, 1),
    );
  }

  Widget _previewImages(List<XFile> mediaFileList) {
    if (mediaFileList.isNotEmpty) {
      return Semantics(
        label: 'image_picker_example_picked_images',
        child: ListView.builder(
          key: UniqueKey(),
          itemBuilder: (BuildContext context, int index) {
            return Semantics(
              label: 'image_picker_example_picked_image',
              child: Image.file(
                File(mediaFileList[index].path),
                errorBuilder: (BuildContext context, Object error,
                    StackTrace? stackTrace) {
                  return const Center(
                      child: Text('This image type is not supported'));
                },
              ),
            );
          },
          itemCount: mediaFileList.length,
        ),
      );
    } else {
      return Container(
        decoration: BoxDecoration(border: Border.all(color: Colors.black54)),
        child: const Icon(Icons.photo_library),
      );
    }
  }

  @override
  void initState() {
    controllerSearch.addListener(_addressControllerListener);
    controllerDateBought.text = DateTime.now().toDMY();
    super.initState();
  }

  @override
  void dispose() {
    controllerSearch.removeListener(_addressControllerListener);
    controllerSearch.dispose();
    controllerDateBought.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    TextEditingController controllerName = TextEditingController();
    TextEditingController controllerPriceBought = TextEditingController();
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: const Text("Create New Item"),
        actions: const [
          TextButton(onPressed: null, child: Text("Save")),
        ],
      ),
      body: SingleChildScrollView(
        child: Padding(
          padding: const EdgeInsets.all(8.0),
          child: Column(
            children: [
              TextField(
                controller: controllerName,
                decoration: const InputDecoration(
                  border: OutlineInputBorder(),
                  labelText: 'Name',
                ),
              ),
              const SizedBox(
                height: 16,
              ),
              TextField(
                controller: controllerDateBought,
                decoration: const InputDecoration(
                  border: OutlineInputBorder(),
                  labelText: 'Date Bought',
                  suffixIcon: Icon(Icons.calendar_month_outlined),
                ),
                readOnly: true,
                onTap: () async {
                  DateTime? dt = await _pickDateTime(context);
                  if (dt == null) {
                    return;
                  }
                  controllerDateBought.text = dt.toDMY();
                },
              ),
              const SizedBox(
                height: 16,
              ),
              TextField(
                controller: controllerPriceBought,
                inputFormatters: [
                  FilteringTextInputFormatter.allow(
                      RegExp(r'(^\d*\.?\d{0,2})')),
                ],
                keyboardType: TextInputType.number,
                decoration: const InputDecoration(
                  border: OutlineInputBorder(),
                  prefixText: "RM ",
                  labelText: 'Price Bought',
                ),
              ),
              const SizedBox(
                height: 16,
              ),
              TextField(
                controller: controllerName,
                decoration: const InputDecoration(
                  border: OutlineInputBorder(),
                  labelText: 'Shop',
                ),
              ),
              const SizedBox(
                height: 16,
              ),
              TextField(
                controller: controllerName,
                decoration: const InputDecoration(
                  border: OutlineInputBorder(),
                  labelText: 'Location',
                ),
              ),
              const SizedBox(
                height: 16,
              ),
              Row(
                children: [
                  const Expanded(
                    child: SizedBox(
                      child: Text("Uploaded 0 receipt"),
                    ),
                  ),
                  SizedBox(
                    width: 70,
                    height: 70,
                    child: Consumer<CreateAssetTrackingStore>(
                      builder: (context, store, child) {
                        return _previewImages(store.receiptImages);
                      },
                    ),
                  ),
                  const SizedBox(
                    width: 8,
                  ),
                  Container(
                      width: 70,
                      height: 70,
                      decoration: const BoxDecoration(
                          color: Color(0xffFDCF09),
                          borderRadius: BorderRadius.all(Radius.circular(16)),
                          boxShadow: [BoxShadow(offset: Offset(0, 1))]),
                      child: const CameraCapture())
                ],
              ),
              const SizedBox(
                height: 16,
              ),
              TextField(
                controller: controllerName,
                decoration: const InputDecoration(
                  border: OutlineInputBorder(),
                  labelText: 'Pictures',
                ),
              ),
              const SizedBox(
                height: 16,
              ),
              SearchAnchor.bar(
                searchController: controllerSearch,
                isFullScreen: false,
                barLeading: const SizedBox(),
                barHintText: "Add Tags",
                barBackgroundColor:
                    const MaterialStatePropertyAll(Colors.white),
                viewHintText: "Add Tags",
                viewBackgroundColor: Colors.white,
                viewTrailing: const [
                  TextButton(onPressed: null, child: Text("Add New"))
                ],
                suggestionsBuilder:
                    (BuildContext context, SearchController controller) {
                  return _foundUsers
                      .map((e) => ListTile(
                            title: Text(e["name"] ?? ""),
                            trailing: const Icon(Icons.cancel_outlined),
                            onTap: () {
                              setState(() {
                                controller.closeView(e["name"] ?? "");
                              });
                            },
                          ))
                      .toList();
                },
              ),
              const SizedBox(
                height: 16,
              ),
              Chip(
                label: const Text("Test"),
                deleteIcon: const Icon(Icons.cancel_outlined),
                onDeleted: () {},
              ),
            ],
          ),
        ),
      ),
    );
  }
}

class DummyData {
  final List<Map<String, dynamic>> allUsers = [
    {"id": 1, "name": "Andy", "age": 29},
    {"id": 2, "name": "Aragon", "age": 40},
    {"id": 3, "name": "Bob", "age": 5},
    {"id": 4, "name": "Barbara", "age": 35},
    {"id": 5, "name": "Candy", "age": 21},
    {"id": 6, "name": "Colin", "age": 55},
    {"id": 7, "name": "Audra", "age": 30},
    {"id": 8, "name": "Banana", "age": 14},
    {"id": 9, "name": "Caversky", "age": 100},
    {"id": 10, "name": "Becky", "age": 32},
  ];

  List<Map<String, dynamic>> runFilter(String enteredKeyword) {
    List<Map<String, dynamic>> results = [];
    if (enteredKeyword.isEmpty) {
      // if the search field is empty or only contains white-space, we'll display all users
      results = allUsers;
    } else {
      results = allUsers
          .where((user) =>
              user["name"].toLowerCase().contains(enteredKeyword.toLowerCase()))
          .toList();
      // we use the toLowerCase() method to make it case-insensitive
    }

    return results;
  }
}

// EXAMPLE SEARCH ANCHOR

// class _CreateAssetTrackingScreenState extends State<CreateAssetTrackingScreen> {
//   List<Map<String, dynamic>> _foundUsers = [];
//   SearchController controllerSearch = SearchController();

//   void _addressControllerListener() {
//     print("LISTENER TRIGGER: ${controllerSearch.text}");
//     setState(() {
//       _foundUsers = DummyData().runFilter(controllerSearch.text);
//     });
//     print("DEBUG $_foundUsers");
//   }

//   @override
//   void initState() {
//     controllerSearch.addListener(_addressControllerListener);
//     super.initState();
//   }

//   @override
//   void dispose() {
//     controllerSearch.removeListener(_addressControllerListener);
//     controllerSearch.dispose();
//     super.dispose();
//   }

//   @override
//   Widget build(BuildContext context) {
//     TextEditingController controllerName = TextEditingController();
//     return Scaffold(
//       appBar: AppBar(
//         backgroundColor: Theme.of(context).colorScheme.inversePrimary,
//         title: Text("Create New Item"),
//         actions: [
//           TextButton(onPressed: null, child: Text("Save")),
//         ],
//       ),
//       body: SingleChildScrollView(
//         child: Padding(
//           padding: const EdgeInsets.all(8.0),
//           child: Column(
//             children: [
//               TextField(
//                 controller: controllerName,
//                 decoration: InputDecoration(
//                   border: OutlineInputBorder(),
//                   labelText: 'Name',
//                 ),
//               ),
//               SearchAnchor(
//                 searchController: controllerSearch,
//                 isFullScreen: false,
//                 viewHintText: "Add Tags",
//                 viewBackgroundColor: Colors.white,
//                 viewTrailing: [
//                   TextButton(onPressed: null, child: Text("Add New"))
//                 ],
//                 builder: (BuildContext context, SearchController controller) {
//                   return SearchBar(
//                     controller: controllerSearch,
//                     hintText: "Add Tags",
//                     onTap: () => controllerSearch.openView(),
//                     onChanged: (_) {
//                       // This Fn Currently Bugged
//                       // https://github.com/flutter/flutter/issues/132915
//                     },
//                   );
//                 },
//                 suggestionsBuilder:
//                     (BuildContext context, SearchController controller) {
//                   return _foundUsers
//                       .map((e) => ListTile(
//                             title: Text(e["name"] ?? ""),
//                             trailing: Icon(Icons.cancel_outlined),
//                             onTap: () {
//                               setState(() {
//                                 controller.closeView(e["name"] ?? "");
//                               });
//                             },
//                           ))
//                       .toList();
//                 },
//               ),
//             ],
//           ),
//         ),
//       ),
//     );
//   }
// }
