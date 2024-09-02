import 'package:flutter/material.dart';
import 'package:flutter_quran/modules/debug_print/debug_print.dart';
import 'package:flutter_quran/modules/list_recording.dart/list_recording.dart';
import 'package:flutter_quran/modules/voice_recording/view.dart';

class MainNavigation extends StatefulWidget {
  const MainNavigation({super.key});

  @override
  State<MainNavigation> createState() => _MainNavigationState();
}

class _MainNavigationState extends State<MainNavigation> {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: DefaultTabController(
        length: 3,
        child: Scaffold(
          appBar: AppBar(
            backgroundColor: Colors.blueGrey[400],
            bottom: const TabBar(tabs: [
              Tab(icon: Icon(Icons.list_rounded)),
              Tab(icon: Icon(Icons.voice_chat)),
              Tab(icon: Icon(Icons.bug_report)),
            ]),
          ),
          body: const TabBarView(
              children: [ListRecording(), VoiceRecording(), DebugprintView()]),
        ),
      ),
    );
  }
}
