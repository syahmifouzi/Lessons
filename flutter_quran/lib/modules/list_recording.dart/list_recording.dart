import 'package:flutter/material.dart';
import 'package:flutter_quran/modules/audio_player/audio_player_store.dart';
import 'package:flutter_quran/modules/audio_player/audio_player_view.dart';
import 'package:flutter_quran/modules/list_recording.dart/list_recording_store.dart';
import 'package:provider/provider.dart';

class ListRecording extends StatelessWidget {
  const ListRecording({super.key});

  Future<void> _showDialog(context) async {
    return showDialog(
        context: context,
        barrierDismissible: true,
        builder: (context) {
          return AlertDialog(
            title: const Text("Title"),
            actions: [
              TextButton(
                  onPressed: () => Navigator.of(context).pop(),
                  child: const Text("OK"))
            ],
          );
        });
  }

  @override
  Widget build(BuildContext context) {
    Provider.of<ListRecordingStore>(context, listen: false)
        .updateListOfRecording();
    return Consumer<ListRecordingStore>(builder: (context, cart, child) {
      List<RecordingFile> recordingFileList = cart.recordingFileList;
      return ListView.builder(
          itemCount: recordingFileList.length,
          itemBuilder: (context, index) {
            return ListTile(
                title: Text(recordingFileList[index].name),
                onTap: () async {
                  Provider.of<AudioPlayerStore>(context, listen: false)
                      .setRecordingFile(recordingFileList[index]);
                  final errorCode = await Provider.of<AudioPlayerStore>(context,
                          listen: false)
                      .getDuration();
                  if (errorCode == 1) {
                    return _showDialog(context);
                  }
                  Navigator.push(
                      context,
                      MaterialPageRoute(
                          builder: (context) => const AudioPlayerView()));
                  // onTap: () => cart.playAudio(recordingFileList[index]),
                });
          });
    });
  }
}
