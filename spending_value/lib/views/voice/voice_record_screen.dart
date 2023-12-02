import 'dart:async';
import 'dart:io';
import 'dart:isolate';

import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:firebase_storage/firebase_storage.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:just_audio/just_audio.dart';
import 'package:path_provider/path_provider.dart';
import 'package:permission_handler/permission_handler.dart';
import 'package:provider/provider.dart';
import 'package:record/record.dart';
import 'package:spending_value/controllers/surahname_store.dart';

class VoiceRecordScreen extends StatefulWidget {
  const VoiceRecordScreen({super.key});

  @override
  State<VoiceRecordScreen> createState() => _VoiceRecordScreenState();
}

class _VoiceRecordScreenState extends State<VoiceRecordScreen> {
  final _audioRecorder = AudioRecorder();
  String? _recordedPath;
  String _progress = "Standby";
  String _title = "";
  DateTime _datetime = DateTime.now();
  final TextEditingController _controllerTitle = TextEditingController();
  final TextEditingController _controllerSurah = TextEditingController();
  final TextEditingController _controllerAyatStart = TextEditingController();
  final TextEditingController _controllerAyatEnd = TextEditingController();
  final SearchController _controllerSearch = SearchController();

  // Isolate? _isolateAudioRecorder;

  final testIsolateClass = IsolateAudioRecorder();
  int testi = 0;

  @override
  void initState() {
    super.initState();
    context.read<SurahnameStore>().getSurahListName();
  }

  @override
  void dispose() {
    // testDispose();
    isolateDispose();
    _audioRecorder.dispose();
    super.dispose();
  }

  void isolateDispose() {
    // if (_isolateAudioRecorder != null) {
    //   _isolateAudioRecorder!.kill(priority: Isolate.immediate);
    // }
  }

  void testDispose() {
    final db = FirebaseFirestore.instance;
    final logmsg = {"message": "dispose called", "timestamp": DateTime.now()};
    db.collection("log").add(logmsg);
    print('dispose called');
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("Record New Voice"),
        actions: [
          TextButton(
              onPressed: () {
                saveRecording();
                // Isolate.spawn(controlPort)
              },
              child: const Text("Save"))
        ],
      ),
      body: SingleChildScrollView(
        child: Column(
          children: [
            ElevatedButton(
                onPressed: () {
                  testIsolateClass.spawn();
                },
                child: Text("spawn")),
            ElevatedButton(
                onPressed: () {
                  testIsolateClass.sendmsg("stop");
                },
                child: Text("send stop")),
            ElevatedButton(
                onPressed: () {
                  testi += 1;
                  testIsolateClass.sendmsg("$testi");
                },
                child: Text("send i++")),
            ElevatedButton(
                onPressed: () {
                  testIsolateClass.kill();
                },
                child: Text("kill")),
            Text(_progress),
            Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                const Expanded(child: SizedBox()),
                ElevatedButton(
                    onPressed: () {
                      startRecording();
                    },
                    child: const Text('Start Record')),
                const Expanded(child: SizedBox()),
                ElevatedButton(
                    onPressed: () {
                      stopRecording();
                    },
                    child: const Text('Stop Record')),
                const Expanded(child: SizedBox()),
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

  void startRecording() async {
    _recordedPath = null;
    Map<Permission, PermissionStatus> permissions =
        await [Permission.storage, Permission.microphone].request();

    bool permissionsGranted = permissions[Permission.storage]!.isGranted &&
        permissions[Permission.microphone]!.isGranted;

    if (!permissionsGranted) {
      print('Permission not granted');
      return;
    }
    Directory appFolder = await getTemporaryDirectory();

    // bool appFolderExists = await appFolder.exists();

    // if (!appFolderExists) {
    //   final created = await appFolder.create(recursive: true);
    //   print(created.path);
    // } else {
    //   print('appfolder exist');
    // }

    _datetime = DateTime.now();
    _title = _datetime.millisecondsSinceEpoch.toString();

    final filepath = '${appFolder.path}/$_title.m4a';

    setState(() {
      _progress = "Now Recording...";
    });

    await _audioRecorder.start(
        const RecordConfig(
            autoGain: true, echoCancel: true, noiseSuppress: true),
        path: filepath);
  }

  // Future<void> isolateRecording() async {
  //   ReceivePort port = ReceivePort();
  //   _isolateAudioRecorder =
  //       await Isolate.spawn<SendPort>(startIsolateRecording, port.sendPort);
  // }

  // void startIsolateRecording(SendPort sendPort) {
  //   ReceivePort receivePort = ReceivePort();
  //   StreamSubscription? isolateSubs;
  //   sendPort.send(receivePort.sendPort);
  //   isolateSubs = receivePort.listen((message) {});
  // }

  Future<void> stopRecording() async {
    _recordedPath = await _audioRecorder.stop();
    setState(() {
      _progress = "Recording Stopped";
    });
  }

  void saveRecording() async {
    if (_progress != "Recording Stopped") {
      await stopRecording();
    }
    AudioPlayer audioPlayer = AudioPlayer();
    final duration = await audioPlayer.setFilePath(_recordedPath!);
    final db = FirebaseFirestore.instance;
    final newIdRef = db.collection("audio").doc();
    final storageRef = FirebaseStorage.instance.ref();
    final audioStorageRef = storageRef.child("audio/${newIdRef.id}.m4a");
    File file = File(_recordedPath!);
    bool hasError = false;
    try {
      await audioStorageRef.putFile(file);
      String downloadUrl = await audioStorageRef.getDownloadURL();
      final audioDB = {
        "title": _controllerTitle.text,
        "url": downloadUrl,
        "timestamp": _datetime,
        "ayatStart": _controllerAyatStart.text,
        "ayatEnd": _controllerAyatEnd.text,
        "status": "Pending",
        "surah": _controllerSurah.text,
        "duration": duration.toString()
      };
      newIdRef.set(audioDB);
    } on FirebaseException catch (e) {
      hasError = true;
      print(e);
    }
    if (hasError) {
      printSnack("Failed save");
    }
    printSnack("Success save");
    navigateBack();
  }

  Future<String?> renameDialog() async {
    TextEditingController controllerTitle = TextEditingController();
    TextEditingController controllerSurah = TextEditingController();
    final result = await showDialog(
        context: context,
        builder: (context) => AlertDialog(
              title: const Text("Rename Title"),
              content: Column(
                children: [
                  TextField(
                      autofocus: true,
                      controller: controllerTitle,
                      decoration: InputDecoration(
                          border: const OutlineInputBorder(),
                          labelText: 'Title',
                          hintText: "$_title (default)")),
                  TextField(
                      controller: controllerSurah,
                      decoration: InputDecoration(
                          border: const OutlineInputBorder(),
                          labelText: 'Surah',
                          hintText: "$_title (default)")),
                ],
              ),
              actions: [
                TextButton(
                    onPressed: () => Navigator.pop(context, 'Cancel'),
                    child: const Text('Cancel')),
                TextButton(
                    onPressed: () {
                      Navigator.pop(context, controllerTitle.text);
                    },
                    child: const Text('Save')),
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
}

class Paths {
  static String recording = 'RapidNote/recordings';
}

class IsolateAudioRecorder {
  // https://blog.codemagic.io/understanding-flutter-isolates/
  // https://blog.logrocket.com/multithreading-flutter-using-dart-isolates/

  // ReceivePort? port;
  SendPort? postport;
  // ReceivePort? responseport;
  Isolate? isolateAudioRecorder;

  void spawn() async {
    print('isolateAudioRecorder is spawned');
    ReceivePort port = ReceivePort();
    isolateAudioRecorder =
        await Isolate.spawn<SendPort>(setupFn, port.sendPort);
    // Link to #1
    postport = await port.first;
    ReceivePort responseport = ReceivePort();
    postport!.send(responseport.sendPort);
    await for (var msg in responseport) {
      print(msg);
    }
  }

  void kill() {
    if (isolateAudioRecorder == null) {
      print('isolateAudioRecorder is null');
      return;
    }
    isolateAudioRecorder!.kill(priority: Isolate.immediate);
    print('isolateAudioRecorder is killed');
  }

  void sendmsg(String msg) async {
    if (postport == null) {
      return;
    }
    postport!.send(msg);
  }

  void setupFn(SendPort sendPort) async {
    ReceivePort receivePort = ReceivePort();
    // Link to #1
    sendPort.send(receivePort.sendPort);
    // SendPort thispostport = await receivePort.first;
    SendPort? thispostport;
    await for (var msg in receivePort) {
      if (msg is SendPort) {
        thispostport = msg;
        continue;
      }
      switch (msg) {
        case "stop":
          if (thispostport != null) {
            thispostport.send("Ok");
          } else {
            print('thispostport is null');
          }
          break;
        default:
          if (thispostport != null) {
            thispostport.send(msg);
          } else {
            print('thispostport is null');
          }
      }
    }
  }
}
