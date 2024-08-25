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
    return Center(
      child: Column(children: [
        Consumer<StopwatchStore>(builder: (context, cart, child) {
          return Text(
            cart.timeNow(),
            style: const TextStyle(fontSize: 48, fontWeight: FontWeight.bold),
          );
        }),
        Consumer<RecordingButtonStore>(builder: (context, cart, child) {
          switch (cart.buttonState) {
            case RecordingButtonState.initial:
              return const Text("Start New Recording");
            case RecordingButtonState.recording:
              return const Text("Now Recording");
            case RecordingButtonState.pausing:
              return const Text("Paused");
            default:
              return const Text("Saving");
          }
        }),
        const RecordingButton()
      ]),
    );
  }
}
