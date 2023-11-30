import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:spending_value/controllers/audio_store.dart';
import 'package:spending_value/routes/routes.dart';

class VoiceHomeScreen extends StatelessWidget {
  const VoiceHomeScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          Navigator.pushNamed(context, Routes.voiceRecordRoute);
        },
        tooltip: 'Create New',
        child: const Icon(Icons.add),
      ),
      body: ListRecording(),
    );
  }
}

class ListRecording extends StatelessWidget {
  const ListRecording({super.key});

  @override
  Widget build(BuildContext context) {
    final List<AudioDbDoc> audioList = [];
    final db = FirebaseFirestore.instance;
    final dbRef = db.collection("audio");
    return FutureBuilder(
        future: dbRef.get(),
        builder: ((context, querySnapshot) {
          if (querySnapshot.hasError) {
            return errorWidget();
          } else if (querySnapshot.hasData) {
            for (var docSnapshot in querySnapshot.data!.docs) {
              print('${docSnapshot.id} => ${docSnapshot.data()}');
              audioList
                  .add(AudioDbDoc.fromJson(docSnapshot.id, docSnapshot.data()));
            }
            return successWidget(audioList);
          }
          return loadingWidget();
        }));
  }

  Widget errorWidget() {
    return Center(child: Text("Error reading from database"));
  }

  Widget loadingWidget() {
    return Center(
        child: SizedBox(
            width: 60, height: 60, child: CircularProgressIndicator()));
  }

  Widget successWidget(List<AudioDbDoc> audioList) {
    return ListView.builder(
        itemCount: audioList.length,
        prototypeItem:
            ListTileRecording(audiodata: AudioDbDoc("proto", "proto", "proto")),
        itemBuilder: (context, index) {
          return ListTileRecording(audiodata: audioList[index]);
        });
  }
}

class ListTileRecording extends StatelessWidget {
  const ListTileRecording({super.key, required this.audiodata});
  final AudioDbDoc audiodata;

  @override
  Widget build(BuildContext context) {
    return Card(
      child: ListTile(
        title: Text(audiodata.title),
        shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(8)),
        onTap: () {
          context.read<AudioStore>().setAudio(audiodata);
          Navigator.pushNamed(context, Routes.voiceDetailsRoute);
        },
      ),
    );
  }
}
