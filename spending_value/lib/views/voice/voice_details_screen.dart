import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:firebase_storage/firebase_storage.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:provider/provider.dart';
import 'package:spending_value/controllers/audio_store.dart';
import 'package:spending_value/controllers/surahname_store.dart';

class VoiceDetailsScreen extends StatefulWidget {
  const VoiceDetailsScreen({super.key});

  @override
  State<VoiceDetailsScreen> createState() => _VoiceDetailsScreenState();
}

class _VoiceDetailsScreenState extends State<VoiceDetailsScreen> {
  // final _formKey = GlobalKey<FormState>();
  late AudioDbDoc _audiodata;
  late AudioStore _ctxDispose;

  final TextEditingController _controllerTitle = TextEditingController();
  final TextEditingController _controllerSurah = TextEditingController();
  final TextEditingController _controllerAyatStart = TextEditingController();
  final TextEditingController _controllerAyatEnd = TextEditingController();
  final TextEditingController _controllerStatus = TextEditingController();
  final SearchController _controllerSearch = SearchController();

  @override
  void initState() {
    super.initState();
    context.read<SurahnameStore>().getSurahListName();
    _audiodata = context.read<AudioStore>().audioDetails;
    context.read<AudioStore>().setinit();
    _controllerTitle.text = _audiodata.title;
    _controllerSurah.text = _audiodata.surah;
    _controllerAyatStart.text = _audiodata.ayatStart;
    _controllerAyatEnd.text = _audiodata.ayatEnd;
    _controllerStatus.text = _audiodata.status;
    _controllerSearch.text = _audiodata.surah;
  }

  @override
  void didChangeDependencies() {
    _ctxDispose = context.read<AudioStore>();
    super.didChangeDependencies();
  }

  @override
  void dispose() {
    _ctxDispose.setdispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(_audiodata.title),
        actions: [
          TextButton(
              onPressed: () => onRenameSave(_audiodata.id),
              child: const Text("Save")),
          TextButton(onPressed: () => onDelete(), child: const Text("Delete"))
        ],
        elevation: 1,
      ),
      body: SingleChildScrollView(
        child: Column(
          children: [
            // playerContainer(),
            Row(
              children: [
                const SizedBox(
                  width: 8,
                ),
                Text(context.watch<AudioStore>().getPositionLabel()),
                Expanded(
                  child: Consumer<AudioStore>(
                    builder: (context, store, child) {
                      return Slider(
                          min: 0,
                          max: store.getTotalDurationMs(),
                          value: store.getPositionMs(),
                          secondaryTrackValue: store.getBufferedPositionMs(),
                          label: store.getPositionMs().toString(),
                          onChanged: ((value) {
                            store.setSeekMs(value);
                            // setState(() {
                            //   _sliderPrimaryValue = value;
                            // });
                          }));
                    },
                  ),
                ),
                Text(context.watch<AudioStore>().getTotalDurationLabel()),
                const SizedBox(
                  width: 8,
                ),
              ],
            ),
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
              children: [
                ElevatedButton(
                    onPressed: () => context.read<AudioStore>().setPlay(),
                    child: const Text("Play")),
                ElevatedButton(
                    onPressed: () => context.read<AudioStore>().setPause(),
                    child: const Text("Pause")),
                ElevatedButton(
                    onPressed: () => context.read<AudioStore>().setStop(),
                    child: const Text("Stop")),
              ],
            ),

            Padding(
              padding: const EdgeInsets.all(8.0),
              child: TextField(
                  controller: _controllerTitle,
                  decoration: const InputDecoration(
                    border: OutlineInputBorder(),
                    labelText: 'Title',
                  )),
            ),
            Padding(
              padding: const EdgeInsets.all(8.0),
              child: TextField(
                  controller: _controllerStatus,
                  decoration: const InputDecoration(
                    border: OutlineInputBorder(),
                    labelText: 'Status',
                  )),
            ),
            Consumer<SurahnameStore>(
              builder: (context, store, child) {
                return SearchAnchor.bar(
                  searchController: _controllerSearch,
                  // isFullScreen: false,
                  barLeading: const SizedBox(),
                  barHintText: "Surah",
                  barBackgroundColor:
                      const MaterialStatePropertyAll(Colors.white),
                  viewHintText: "Surah",
                  viewBackgroundColor: Colors.white,
                  viewTrailing: [
                    TextButton(
                        onPressed: () => addNewSurah(_controllerSearch.text),
                        child: const Text("Add New"))
                  ],
                  suggestionsBuilder:
                      (BuildContext context, SearchController controller) {
                    store.runFilter(controller.text);
                    return store.filteredListName
                        .map((e) => ListTile(
                              title: Text(e.name),
                              trailing: IconButton(
                                  onPressed: () => removeSurah(e.id),
                                  icon: const Icon(Icons.cancel_outlined)),
                              onTap: () {
                                _controllerSurah.text = e.name;
                                setState(() {
                                  controller.closeView(e.name);
                                });
                              },
                            ))
                        .toList();
                  },
                );
              },
            ),
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
              children: [
                Expanded(
                  child: Padding(
                    padding: const EdgeInsets.all(8.0),
                    child: TextField(
                        controller: _controllerAyatStart,
                        keyboardType: TextInputType.number,
                        inputFormatters: [
                          FilteringTextInputFormatter.allow(RegExp(r'\d*')),
                        ],
                        decoration: const InputDecoration(
                          border: OutlineInputBorder(),
                          labelText: 'Ayat Start',
                        )),
                  ),
                ),
                Expanded(
                  child: Padding(
                    padding: const EdgeInsets.all(8.0),
                    child: TextField(
                        controller: _controllerAyatEnd,
                        keyboardType: TextInputType.number,
                        inputFormatters: [
                          FilteringTextInputFormatter.allow(RegExp(r'\d*')),
                        ],
                        decoration: const InputDecoration(
                          border: OutlineInputBorder(),
                          labelText: 'Ayat End',
                        )),
                  ),
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }

  void onRenameSave(String id) async {
    final audioDB = {
      "title": _controllerTitle.text,
      "ayatStart": _controllerAyatStart.text,
      "ayatEnd": _controllerAyatEnd.text,
      "status": _controllerStatus.text,
      "surah": _controllerSurah.text,
    };
    final db = FirebaseFirestore.instance;
    final dbRef = db.collection("audio").doc(id);
    bool hasError = false;
    try {
      await dbRef.update(audioDB);
      print('Success update doc');
    } catch (e) {
      hasError = true;
      print('Error updating doc: $e');
    }
    if (hasError) {
      snackMsg("Failed to update on server");
      return;
    }
    snackMsg("Success update on server");
  }

  // Future<String> renameDialog(String oldTitle) async {
  //   TextEditingController controllerTitle = TextEditingController();
  //   controllerTitle.text = oldTitle;
  //   final result = await showDialog(
  //       context: context,
  //       builder: (context) => AlertDialog(
  //             title: const Text("Rename Title"),
  //             content: Form(
  //               key: _formKey,
  //               child: TextFormField(
  //                   controller: controllerTitle,
  //                   decoration: const InputDecoration(
  //                     border: OutlineInputBorder(),
  //                     labelText: 'Title',
  //                   ),
  //                   validator: (value) {
  //                     if (value == null || value.isEmpty) {
  //                       return "Please enter a title";
  //                     }
  //                     return null;
  //                   }),
  //             ),
  //             actions: [
  //               TextButton(
  //                   onPressed: () => Navigator.pop(context, 'Cancel'),
  //                   child: const Text('Cancel')),
  //               TextButton(
  //                   onPressed: () {
  //                     if (_formKey.currentState!.validate()) {
  //                       Navigator.pop(context, controllerTitle.text);
  //                     }
  //                   },
  //                   child: const Text('Save')),
  //             ],
  //           ));

  //   if (result == "Cancel" || result == null) {
  //     return result ?? "Error";
  //   }

  //   setState(() {
  //     _audiodata.title = result;
  //   });
  //   onRenameSave(_audiodata.id, _audiodata.title);

  //   return result;
  // }

  void onDelete() async {
    bool? confirmDelete = await confirmDeleteDialog(_audiodata.title);
    if (!(confirmDelete ?? false)) {
      return;
    }
    final storageRef = FirebaseStorage.instance.ref();
    final storageId = _audiodata.id;
    final audioRef = storageRef.child("audio/$storageId.m4a");
    bool deleteSuccess = false;
    bool hasError = false;
    try {
      await audioRef.delete();
      deleteSuccess = true;
    } catch (e) {
      hasError = true;
      print("Error deleting storage: $e");
    }
    if (!deleteSuccess) {
      return;
    }
    final db = FirebaseFirestore.instance;
    try {
      await db.collection("audio").doc(storageId).delete();
    } catch (e) {
      hasError = true;
      print("Error deleting database: $e");
    }
    if (hasError) {
      snackMsg("Failed to delete on server");
      return;
    }
    snackMsg("Success deleted on server");
    navigateBack();
  }

  Future<bool?> confirmDeleteDialog(String title) async {
    final result = await showDialog<bool>(
        context: context,
        builder: (context) => AlertDialog(
              title: Text("Confirm Delete $title?"),
              content: const Text("Action is irreversible"),
              actions: [
                TextButton(
                    onPressed: () => Navigator.pop(context, false),
                    child: const Text('Cancel')),
                TextButton(
                    onPressed: () => Navigator.pop(context, true),
                    child: const Text('Delete')),
              ],
            ));

    return result;
  }

  void addNewSurah(String name) {
    final nameSurah = {
      "name": name,
    };
    final db = FirebaseFirestore.instance;
    bool hasError = false;
    try {
      db.collection("surahName").add(nameSurah);
    } catch (e) {
      hasError = true;
      print("Error add new surah: $e");
    }
    if (hasError) {
      snackMsg("Failed to add new surah");
      return;
    }
    snackMsg("Success added new surah");
    context.read<SurahnameStore>().getSurahListName();

    _controllerSearch.closeView(name);
  }

  void removeSurah(String id) {
    final db = FirebaseFirestore.instance;
    bool hasError = false;
    try {
      db.collection("surahName").doc(id).delete();
    } catch (e) {
      hasError = true;
      print("Error delete surah: $e");
    }
    if (hasError) {
      snackMsg("Failed to delete surah");
      return;
    }
    snackMsg("Success deleted surah");
    context.read<SurahnameStore>().getSurahListName();
  }

  SnackBar snackMsg(String msg) {
    return SnackBar(
      content: Text(msg),
      action: SnackBarAction(
        label: 'Undo',
        onPressed: () {
          // Some code to undo the change.
        },
      ),
    );
  }

  void printSnack(String msg) {
    ScaffoldMessenger.of(context).showSnackBar(snackMsg(msg));
  }

  void navigateBack() {
    Navigator.pop(context);
  }

  // void playRecording() async {
  //   // final duration = await _audioPlayer.setUrl(_audiodata.url);
  //   final duration = await _audioPlayer.setUrl(
  //       "https://www.soundhelix.com/examples/mp3/SoundHelix-Song-2.mp3");
  //   _audioPlayer.play();
  // }

  // void stopPlaying() async {
  //   _audioPlayer.stop();
  // }
}
