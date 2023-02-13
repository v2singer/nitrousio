import 'package:flutter/material.dart';
import 'package:learn3/main.dart';
import 'package:learn3/model/project.dart';
import 'package:learn3/views/listview.dart';

class MyHomePageState extends State<MyHomePage> {
  List<GitProject> gitProjects = [];
  var nums = 1000;
  var extendMaxWidth = 400;
  int selectedIndex = 0;

  @override
  void initState() {
    super.initState();

    for (int i = 0; i < nums; i++) {
      gitProjects.add(GitProject(
          id: (i + 1).toString(),
          name: "name$i",
          url: "github.com/author/name$i",
          author: "author",
          isSync: i % 2 == 0));
    }
  }

  @override
  Widget build(BuildContext context) {
    Widget page = ListViewBuilder(gitProjects: gitProjects);

    return LayoutBuilder(builder: (context, constraints) {
      return Scaffold(
        body: Row(children: [
          SafeArea(
              child: NavigationRail(
            extended: constraints.maxWidth >= extendMaxWidth,
            destinations: const [
              NavigationRailDestination(
                  icon: Icon(Icons.home), label: Text("HOME")),
              NavigationRailDestination(
                  icon: Icon(Icons.sync), label: Text("SYNC"))
            ],
            selectedIndex: selectedIndex,
            onDestinationSelected: (value) {
              setState(() {
                selectedIndex = value;
              });
            },
          )),
          Expanded(
              child: Container(
                  color: Theme.of(context).colorScheme.primaryContainer,
                  child: page)),
        ]),
      );
    });
  }
}

class UserDrawer extends StatelessWidget {
  const UserDrawer({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Drawer(
      child: ListView(
        children: const [
          DrawerHeader(
              decoration: BoxDecoration(color: Colors.lightBlueAccent),
              child: Center(
                child: SizedBox(
                  width: 60,
                  height: 60,
                  child: CircleAvatar(child: Icon(Icons.person)),
                ),
              )),
          ListTile(title: Text("account")),
        ],
      ),
    );
  }
}
