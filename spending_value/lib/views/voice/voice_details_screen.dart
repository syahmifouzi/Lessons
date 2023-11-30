import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:firebase_storage/firebase_storage.dart';
import 'package:flutter/material.dart';
import 'package:just_audio/just_audio.dart';
import 'package:provider/provider.dart';
import 'package:spending_value/controllers/audio_store.dart';

class VoiceDetailsScreen extends StatefulWidget {
  const VoiceDetailsScreen({super.key});

  @override
  State<VoiceDetailsScreen> createState() => _VoiceDetailsScreenState();
}

class _VoiceDetailsScreenState extends State<VoiceDetailsScreen> {
  final _formKey = GlobalKey<FormState>();
  late AudioDbDoc _audiodata;
  late AudioStore _ctxDispose;

  @override
  void initState() {
    super.initState();
    _audiodata = context.read<AudioStore>().audioDetails;
    context.read<AudioStore>().setinit();
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
              onPressed: () => renameDialog(_audiodata.title),
              child: Text("Rename")),
          TextButton(onPressed: () => onDelete(), child: Text("Delete"))
        ],
        elevation: 1,
      ),
      body: Column(
        children: [
          // playerContainer(),
          Row(
            children: [
              SizedBox(
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
              SizedBox(
                width: 8,
              ),
            ],
          ),
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceEvenly,
            children: [
              ElevatedButton(
                  onPressed: () => context.read<AudioStore>().setPlay(),
                  child: Text("Play")),
              ElevatedButton(
                  onPressed: () => context.read<AudioStore>().setStop(),
                  child: Text("Stop")),
            ],
          ),
        ],
      ),
    );
  }

  void onRenameSave(String id, String title) async {
    final db = FirebaseFirestore.instance;
    final dbRef = db.collection("audio").doc(id);
    try {
      await dbRef.update({"title": title});
      print('Success update doc');
    } catch (e) {
      print('Error updating doc: $e');
    }
  }

  Future<String> renameDialog(String oldTitle) async {
    TextEditingController controllerTitle = TextEditingController();
    controllerTitle.text = oldTitle;
    final result = await showDialog(
        context: context,
        builder: (context) => AlertDialog(
              title: Text("Rename Title"),
              content: Form(
                key: _formKey,
                child: TextFormField(
                    controller: controllerTitle,
                    decoration: const InputDecoration(
                      border: OutlineInputBorder(),
                      labelText: 'Title',
                    ),
                    validator: (value) {
                      if (value == null || value.isEmpty) {
                        return "Please enter a title";
                      }
                      return null;
                    }),
              ),
              actions: [
                TextButton(
                    onPressed: () => Navigator.pop(context, 'Cancel'),
                    child: Text('Cancel')),
                TextButton(
                    onPressed: () {
                      if (_formKey.currentState!.validate()) {
                        Navigator.pop(context, controllerTitle.text);
                      }
                    },
                    child: Text('Save')),
              ],
            ));

    if (result == "Cancel" || result == null) {
      return result ?? "Error";
    }

    setState(() {
      _audiodata.title = result;
    });
    onRenameSave(_audiodata.id, _audiodata.title);

    return result;
  }

  void onDelete() async {
    bool? confirmDelete = await confirmDeleteDialog(_audiodata.title);
    if (!(confirmDelete ?? false)) {
      return;
    }
    final storageRef = FirebaseStorage.instance.ref();
    final storageId = _audiodata.id;
    final audioRef = storageRef.child("audio/$storageId.m4a");
    bool deleteSuccess = false;
    try {
      await audioRef.delete();
      deleteSuccess = true;
    } catch (e) {
      print("Error deleting storage: $e");
    }
    if (!deleteSuccess) {
      return;
    }
    final db = FirebaseFirestore.instance;
    try {
      await db.collection("audio").doc(storageId).delete();
    } catch (e) {
      print("Error deleting database: $e");
    }
  }

  Future<bool?> confirmDeleteDialog(String title) async {
    final result = await showDialog<bool>(
        context: context,
        builder: (context) => AlertDialog(
              title: Text("Confirm Delete $title?"),
              content: Text("Action is irreversible"),
              actions: [
                TextButton(
                    onPressed: () => Navigator.pop(context, false),
                    child: Text('Cancel')),
                TextButton(
                    onPressed: () => Navigator.pop(context, true),
                    child: Text('Delete')),
              ],
            ));

    return result;
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
