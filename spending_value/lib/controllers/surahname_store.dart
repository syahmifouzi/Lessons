import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:flutter/material.dart';

class SurahnameStore extends ChangeNotifier {
  List<SurahDbDoc> listName = [];
  List<SurahDbDoc> filteredListName = [];
  SurahDbDoc selectedName = SurahDbDoc(id: "");

  void setSelectedName(SurahDbDoc x) {
    selectedName = x;
    notifyListeners();
  }

  void addListName(SurahDbDoc x) {
    listName.add(x);
  }

  void clearListName() {
    listName.clear();
  }

  void runFilter(String enteredKeyword) {
    List<SurahDbDoc> results = [];
    if (enteredKeyword.isEmpty) {
      // if the search field is empty or only contains white-space, we'll display all users
      results = listName;
    } else {
      results = listName
          .where((data) =>
              data.name.toLowerCase().contains(enteredKeyword.toLowerCase()))
          .toList();
      // we use the toLowerCase() method to make it case-insensitive
    }

    filteredListName = results;
  }

  void getSurahListName() async {
    clearListName();
    final db = FirebaseFirestore.instance;
    final dbRef = db.collection("surahName");
    final querySnapshot = await dbRef.get();
    if (querySnapshot.docs.isEmpty) {
      return;
    }
    for (QueryDocumentSnapshot<Map<String, dynamic>> element
        in querySnapshot.docs) {
      final name = element.get("name");
      addListName(SurahDbDoc(id: element.id, name: name));
    }
  }
}

class SurahDbDoc {
  SurahDbDoc({required this.id, this.name = ""});

  final String id;
  String name;

  factory SurahDbDoc.fromJson(String id1, Map<String, dynamic> data) {
    final id = id1;
    final name = data["name"];
    if (name is! String) {
      throw FormatException(
          'Invalid JSON: required "name" field of type String in $data');
    }
    return SurahDbDoc(id: id, name: name);
  }
}
