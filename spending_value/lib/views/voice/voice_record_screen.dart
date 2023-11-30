import 'dart:io';

import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:firebase_storage/firebase_storage.dart';
import 'package:flutter/material.dart';
import 'package:just_audio/just_audio.dart';
import 'package:path_provider/path_provider.dart';
import 'package:permission_handler/permission_handler.dart';
import 'package:record/record.dart';

class VoiceRecordScreen extends StatefulWidget {
  const VoiceRecordScreen({super.key});

  @override
  State<VoiceRecordScreen> createState() => _VoiceRecordScreenState();
}

class _VoiceRecordScreenState extends State<VoiceRecordScreen> {
  final _audioRecorder = AudioRecorder();
  String? recordedPath;
  String progress = "Standby";
  String title = "";
  DateTime datetime = DateTime.now();

  @override
  void initState() {
    super.initState();
  }

  @override
  void dispose() {
    _audioRecorder.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Record New Voice"),
        actions: [
          TextButton(
              onPressed: () {
                saveRecording();
              },
              child: Text("Save"))
        ],
      ),
      body: Column(
        children: [
          Text(progress),
          Row(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Expanded(child: SizedBox()),
              ElevatedButton(
                  onPressed: () {
                    startRecording();
                  },
                  child: Text('Start Record')),
              Expanded(child: SizedBox()),
              ElevatedButton(
                  onPressed: () {
                    stopRecording();
                  },
                  child: Text('Stop Record')),
              Expanded(child: SizedBox()),
            ],
          ),
        ],
      ),
    );
  }

  void startRecording() async {
    recordedPath = null;
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

    datetime = DateTime.now();
    title = datetime.millisecondsSinceEpoch.toString();

    final filepath = appFolder.path + '/' + title + '.m4a';

    setState(() {
      progress = "Now Recording...";
    });

    await _audioRecorder.start(const RecordConfig(), path: filepath);
  }

  void stopRecording() async {
    recordedPath = await _audioRecorder.stop();
    setState(() {
      progress = "Recording Stopped";
    });
  }

  void saveRecording() async {
    String? result = await renameDialog();
    if (result == "Cancel" || recordedPath == null) {
      return;
    }
    if (title.isEmpty) {
      return;
    }
    if (result == null || result.isEmpty) {
      result = title;
    }
    AudioPlayer audioPlayer = AudioPlayer();
    final duration = await audioPlayer.setFilePath(recordedPath!);
    final db = FirebaseFirestore.instance;
    final newIdRef = db.collection("audio").doc();
    final storageRef = FirebaseStorage.instance.ref();
    final audioStorageRef = storageRef.child("audio/${newIdRef.id}.m4a");
    File file = File(recordedPath!);
    try {
      await audioStorageRef.putFile(file);
      String downloadUrl = await audioStorageRef.getDownloadURL();
      final audioDB = {
        "title": result,
        "url": downloadUrl,
        "timestamp": datetime,
        "duration": duration.toString()
      };
      newIdRef.set(audioDB);
    } on FirebaseException catch (e) {
      print(e);
    }
  }

  Future<String?> renameDialog() async {
    TextEditingController controllerTitle = TextEditingController();
    final result = await showDialog(
        context: context,
        builder: (context) => AlertDialog(
              title: Text("Rename Title"),
              content: TextField(
                  autofocus: true,
                  controller: controllerTitle,
                  decoration: InputDecoration(
                      border: OutlineInputBorder(),
                      labelText: 'Title',
                      hintText: "$title (default)")),
              actions: [
                TextButton(
                    onPressed: () => Navigator.pop(context, 'Cancel'),
                    child: Text('Cancel')),
                TextButton(
                    onPressed: () {
                      Navigator.pop(context, controllerTitle.text);
                    },
                    child: Text('Save')),
              ],
            ));

    return result;
  }
}

class Paths {
  static String recording = 'RapidNote/recordings';
}
