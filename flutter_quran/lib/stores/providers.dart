import 'package:flutter_quran/modules/debug_print/debug_print_store.dart';
import 'package:flutter_quran/modules/list_recording.dart/list_recording_store.dart';
import 'package:flutter_quran/modules/voice_recording/recording_button_store.dart';
import 'package:flutter_quran/modules/voice_recording/stopwatch_store.dart';
import 'package:flutter_quran/modules/voice_recording/voice_recording_store.dart';
import 'package:provider/provider.dart';
import 'package:provider/single_child_widget.dart';

class Providers {
  static List<SingleChildWidget> providers = [
    ChangeNotifierProvider(create: (context) => RecordingButtonStore()),
    ChangeNotifierProvider(create: (context) => DebugPrintStore()),
    ChangeNotifierProvider(create: (context) => StopwatchStore()),
    ChangeNotifierProvider(create: (context) => VoiceRecordingStore()),
    ChangeNotifierProvider(create: (context) => ListRecordingStore()),
  ];
}
