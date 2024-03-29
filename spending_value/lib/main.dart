import 'dart:isolate';

import 'package:firebase_messaging/firebase_messaging.dart';
import 'package:flutter/material.dart';
import 'package:firebase_core/firebase_core.dart';
import 'package:flutter/services.dart';
import 'package:path_provider/path_provider.dart';
import 'package:provider/provider.dart';
import 'package:record/record.dart';
// import 'package:record/record.dart';
import 'package:spending_value/routes/pages.dart';
import 'package:spending_value/controllers/providers.dart';
// import 'package:workmanager/workmanager.dart';
import 'firebase_options.dart';

// const recordAudio = "recordAudio";
// final audioRecorder = AudioRecorder();

// @pragma('vm:entry-point')
// void callbackDispatcher() {
//   Workmanager().executeTask((task, inputData) async {
//     switch (task) {
//       case recordAudio:
//         // record audio process
//         break;
//       default:
//     }

//     return Future.value(true);
//   });
// }

// IMPORTANT
// Must declare this function as static
// or declare as top level function
@pragma('vm:entry-point')
void setupFn(List<dynamic> args) async {
  SendPort sendPort = args[0];
  RootIsolateToken token = args[1];
  String appFolderPath = args[2];
  BackgroundIsolateBinaryMessenger.ensureInitialized(token);
  final audioRecorder = AudioRecorder();
  final datetimeL = await startRecorderIsolated(audioRecorder, appFolderPath);

  ReceivePort receivePort = ReceivePort();
  // Link to #1
  sendPort.send(receivePort.sendPort);
  // SendPort thispostport = await receivePort.first;
  SendPort? thispostport;
  await for (var msg in receivePort) {
    if (msg is SendPort) {
      thispostport = msg;
      final tosend = {
        "cmd": "start",
      };
      thispostport.send(tosend);
      continue;
    }
    switch (msg) {
      case "stop":
        if (thispostport != null) {
          final recordedPathL = await stopRecorderIsolated(audioRecorder);
          final tosend = {
            "cmd": "stop",
            "path": recordedPathL,
            "datetime": datetimeL
          };
          thispostport.send(tosend);
        }
        break;
      default:
        if (thispostport != null) {
          final tosend = {
            "cmd": "error",
          };
          thispostport.send(tosend);
        }
    }
  }
}

Future<DateTime> startRecorderIsolated(
    AudioRecorder audioRecorder, String appFolderPath) async {
  final datetime = DateTime.now();
  final title = datetime.millisecondsSinceEpoch.toString();

  final filepath = '$appFolderPath/$title.m4a';

  audioRecorder.start(
      const RecordConfig(autoGain: true, echoCancel: true, noiseSuppress: true),
      path: filepath);

  return datetime;
}

Future<String?> stopRecorderIsolated(AudioRecorder audioRecorder) async {
  final res = await audioRecorder.stop();
  audioRecorder.dispose();
  return res;
}

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await Firebase.initializeApp(
    options: DefaultFirebaseOptions.currentPlatform,
  );
  runApp(const InitializeApp());
}

class InitializeApp extends StatelessWidget {
  const InitializeApp({super.key});

  @override
  Widget build(BuildContext context) {
    final pushNotificationService = PushNotificationService();
    pushNotificationService.initialise();
    return MultiProvider(
      providers: Providers.providers,
      child: const MyApp(),
    );
  }
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Spending Value',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.amber),
        useMaterial3: true,
      ),
      initialRoute: '/',
      routes: Pages.routes,
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({super.key, required this.title});

  // This widget is the home page of your application. It is stateful, meaning
  // that it has a State object (defined below) that contains fields that affect
  // how it looks.

  // This class is the configuration for the state. It holds the values (in this
  // case the title) provided by the parent (in this case the App widget) and
  // used by the build method of the State. Fields in a Widget subclass are
  // always marked "final".

  final String title;

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  int _counter = 0;

  void _incrementCounter() {
    setState(() {
      // This call to setState tells the Flutter framework that something has
      // changed in this State, which causes it to rerun the build method below
      // so that the display can reflect the updated values. If we changed
      // _counter without calling setState(), then the build method would not be
      // called again, and so nothing would appear to happen.
      _counter++;
    });
  }

  // In this example, suppose that all messages contain a data field with the key 'type'.
  Future<void> setupInteractedMessage() async {
    // Get any messages which caused the application to open from
    // a terminated state.
    RemoteMessage? initialMessage =
        await FirebaseMessaging.instance.getInitialMessage();

    // If the message also contains a data property with a "type" of "chat",
    // navigate to a chat screen
    if (initialMessage != null) {
      _handleMessage(initialMessage);
    }

    // Also handle any interaction when the app is in the background via a
    // Stream listener
    FirebaseMessaging.onMessageOpenedApp.listen(_handleMessage);
  }

  void _handleMessage(RemoteMessage message) {
    if (message.data['type'] == 'chat') {
      print("Navigate to chat page");
      // Navigator.pushNamed(
      //   context,
      //   '/chat',
      //   arguments: ChatArguments(message),
      // );
    }
  }

  @override
  void initState() {
    super.initState();

    // Run code required to handle interacted messages in an async function
    // as initState() must not be async
    setupInteractedMessage();
  }

  @override
  Widget build(BuildContext context) {
    // This method is rerun every time setState is called, for instance as done
    // by the _incrementCounter method above.
    //
    // The Flutter framework has been optimized to make rerunning build methods
    // fast, so that you can just rebuild anything that needs updating rather
    // than having to individually change instances of widgets.
    return Scaffold(
      appBar: AppBar(
        // TRY THIS: Try changing the color here to a specific color (to
        // Colors.amber, perhaps?) and trigger a hot reload to see the AppBar
        // change color while the other colors stay the same.
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        // Here we take the value from the MyHomePage object that was created by
        // the App.build method, and use it to set our appbar title.
        title: Text(widget.title),
      ),
      body: Center(
        // Center is a layout widget. It takes a single child and positions it
        // in the middle of the parent.
        child: Column(
          // Column is also a layout widget. It takes a list of children and
          // arranges them vertically. By default, it sizes itself to fit its
          // children horizontally, and tries to be as tall as its parent.
          //
          // Column has various properties to control how it sizes itself and
          // how it positions its children. Here we use mainAxisAlignment to
          // center the children vertically; the main axis here is the vertical
          // axis because Columns are vertical (the cross axis would be
          // horizontal).
          //
          // TRY THIS: Invoke "debug painting" (choose the "Toggle Debug Paint"
          // action in the IDE, or press "p" in the console), to see the
          // wireframe for each widget.
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            const Text(
              'You have pushed the button this many times:',
            ),
            Text(
              '$_counter',
              style: Theme.of(context).textTheme.headlineMedium,
            ),
          ],
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: _incrementCounter,
        tooltip: 'Increment',
        child: const Icon(Icons.add),
      ), // This trailing comma makes auto-formatting nicer for build methods.
    );
  }
}

class PushNotificationService {
  PushNotificationService();

  Future initialise() async {
    final fcmToken2 = await FirebaseMessaging.instance.getToken();
    print("FirebaseMessaging token: $fcmToken2");

    FirebaseMessaging.instance.onTokenRefresh.listen((fcmToken) {
      // TODO: If necessary send token to application server.

      // Note: This callback is fired at each app startup and whenever a new
      // token is generated.
      print("token generated $fcmToken");
    }).onError((err) {
      // Error getting token.
      print("error $err");
    });
  }
}
