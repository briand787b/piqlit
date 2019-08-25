import 'package:flutter/material.dart';

import './widgets/media/media_page.dart';
import './widgets/view/view_page.dart';
import './widgets/servers/servers_page.dart';
import './widgets/cast/cast_page.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext ctx) {
    return MaterialApp(
      home: DefaultTabController(
        length: 4,
        child: Scaffold(
          appBar: AppBar(
            backgroundColor: Colors.amberAccent,
            title: Container(
              width: double.infinity,
              child: Text(
                'piqlit',
                style: TextStyle(
                  color: Colors.black,
                ),
                textAlign: TextAlign.center,
              ),
            ),
            bottom: TabBar(
              tabs: [
                Tab(icon: Icon(Icons.cloud, color: Colors.black)),
                Tab(icon: Icon(Icons.movie, color: Colors.black)),
                Tab(icon: Icon(Icons.play_arrow, color: Colors.black)),
                Tab(icon: Icon(Icons.cast, color: Colors.black)),
              ],
            ),
          ),
          body: TabBarView(
            children: [
              ServersPage(),
              MediaPage(),
              ViewPage(),
              CastPage(),
            ],
          ),
        ),
      ),
    );
  }
}
