import 'package:flutter/material.dart';
import 'package:flutter_quran/modules/voice_recording/recording_button.dart';
import 'package:flutter_quran/modules/voice_recording/recording_button_store.dart';
import 'package:flutter_quran/modules/voice_recording/stopwatch_store.dart';
import 'package:provider/provider.dart';

class VoiceRecording extends StatefulWidget {
  const VoiceRecording({super.key});

  @override
  State<VoiceRecording> createState() => _VoiceRecordingState();
}

class _VoiceRecordingState extends State<VoiceRecording> {
  @override
  Widget build(BuildContext context) {
    return Container(
      color: Colors.blueGrey[900],
      child: Center(
        child: Column(children: [
          Consumer<StopwatchStore>(builder: (context, cart, child) {
            return Text(
              cart.timeNow(),
              style: TextStyle(
                  fontSize: 48,
                  fontWeight: FontWeight.bold,
                  color: Colors.blue[50]),
            );
          }),
          Consumer<RecordingButtonStore>(builder: (context, cart, child) {
            switch (cart.buttonState) {
              case RecordingButtonState.initial:
                return Text("Start New Recording",
                    style: TextStyle(color: Colors.blue[50]));
              case RecordingButtonState.recording:
                return Text("Now Recording",
                    style: TextStyle(color: Colors.blue[50]));
              case RecordingButtonState.pausing:
                return Text("Paused", style: TextStyle(color: Colors.blue[50]));
              default:
                return Text("Saving", style: TextStyle(color: Colors.blue[50]));
            }
          }),
          const RecordingButton()
        ]),
      ),
    );
  }
}
