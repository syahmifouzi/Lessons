import 'dart:io';

import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:firebase_storage/firebase_storage.dart';
import 'package:flutter/material.dart';
import 'package:just_audio/just_audio.dart';
import 'package:path_provider/path_provider.dart';
import 'package:permission_handler/permission_handler.dart';
import 'package:record/record.dart';

class RecordAudioScreen extends StatefulWidget {
  const RecordAudioScreen({super.key});

  @override
  State<RecordAudioScreen> createState() => _RecordAudioScreenState();
}

class _RecordAudioScreenState extends State<RecordAudioScreen> {
  final _audioRecorder = AudioRecorder();

  @override
  void initState() {
    super.initState();
  }

  @override
  void dispose() {
    super.dispose();
  }

  void startRecording() async {
    Map<Permission, PermissionStatus> permissions =
        await [Permission.storage, Permission.microphone].request();

    bool permissionsGranted = permissions[Permission.storage]!.isGranted &&
        permissions[Permission.microphone]!.isGranted;

    if (permissionsGranted) {
      Directory appFolder = await getApplicationDocumentsDirectory();

      bool appFolderExists = await appFolder.exists();

      if (!appFolderExists) {
        final created = await appFolder.create(recursive: true);
        print(created.path);
      } else {
        print('appfolder exist');
      }

      final filepath = appFolder.path +
          '/' +
          DateTime.now().millisecondsSinceEpoch.toString() +
          '.m4a';

      print(filepath);

      await _audioRecorder.start(const RecordConfig(), path: filepath);

      // emit(RecordOn());
    } else {
      print('Permission not granted');
    }
  }

  void stopRecording() async {
    String? path = await _audioRecorder.stop();
    print('Output path $path');
    AudioPlayer _audioPlayer = AudioPlayer();
    final duration = await _audioPlayer.setFilePath(path!);
    print('duration is: $duration');
    final storageRef = FirebaseStorage.instance.ref();
    final testRef = storageRef.child("audio/test.m4a");
    File file = File(path);
    try {
      await testRef.putFile(file);
      String downloadUrl = await testRef.getDownloadURL();
      final db = FirebaseFirestore.instance;
      final audioDB = {"name": "test", "url": downloadUrl};
      db
          .collection("audio")
          .add(audioDB)
          .then((doc) => print("added with ID: ${doc.id}"));
    } on FirebaseException catch (e) {
      print(e);
    }
    await _audioPlayer.play();
    await _audioPlayer.stop();
    await _audioPlayer.dispose();
  }

  void playRecording() async {
    AudioPlayer _audioPlayer = AudioPlayer();
    final duration = await _audioPlayer.setUrl(
        'https://firebasestorage.googleapis.com/v0/b/spending-value.appspot.com/o/audio%2Ftest.m4a?alt=media&token=40b761bd-85c1-42f7-abb0-99fa89729532');
    await _audioPlayer.play();
    await _audioPlayer.stop();
    await _audioPlayer.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Column(
        children: [
          Text("Hello"),
          ElevatedButton(
              onPressed: () {
                startRecording();
              },
              child: Text('Start')),
          ElevatedButton(
              onPressed: () {
                stopRecording();
              },
              child: Text('Stop')),
          ElevatedButton(
              onPressed: () {
                playRecording();
              },
              child: Text('Play'))
        ],
      ),
    );
  }
}

class Paths {
  static String recording = 'RapidNote/recordings';
}
