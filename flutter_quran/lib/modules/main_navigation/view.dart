import 'package:flutter/material.dart';
import 'package:flutter_quran/modules/debug_print/debug_print.dart';
import 'package:flutter_quran/modules/tutorial_page/view.dart';
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
            bottom: const TabBar(tabs: [
              Tab(icon: Icon(Icons.directions_car)),
              Tab(icon: Icon(Icons.directions_transit)),
              Tab(icon: Icon(Icons.bug_report)),
            ]),
          ),
          body: const TabBarView(children: [
            TutorialPage(title: "Hello World"),
            VoiceRecording(),
            DebugprintView()
          ]),
        ),
      ),
    );
  }
}
