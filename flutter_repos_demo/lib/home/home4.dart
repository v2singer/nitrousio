import 'package:flutter/material.dart';
import 'package:learn3/main.dart';
import 'package:learn3/model/project.dart';

int add(int x, int y) {
  return x + y;
}

List<GitProject> runFilter(
    String keyword, List<GitProject> allProjects, int filterKeyType) {
  List<GitProject> result = [];
  print(filterKeyType);
  if (allProjects.isNotEmpty) {
    switch (filterKeyType) {
      case 0:
        result = allProjects
            .where((project) =>
                project.url.toLowerCase().contains(keyword.toLowerCase()))
            .toList();
        break;
      case 1:
        result = allProjects
            .where((project) => project.isSync)
            .where((project) =>
                project.url.toLowerCase().contains(keyword.toLowerCase()))
            .toList();
    }
  } else {
    result = allProjects;
  }
  return result;
}

class MyHomePageState extends State<MyHomePage> {
  List<GitProject> allProjects = [];
  List<GitProject> gitProjects = [];
  var nums = 1000;
  var extendMaxWidth = 400;
  int selectedIndex = 0;
  String _gSearchKey = "";

  @override
  void initState() {
    super.initState();

    for (int i = 0; i < nums; i++) {
      allProjects.add(GitProject(
          id: (i + 1).toString(),
          name: "name$i",
          url: "github.com/author/name$i",
          author: "author",
          isSync: i % 2 == 0));
    }
    gitProjects = allProjects;
  }

  @override
  Widget build(BuildContext context) {
    Widget page;
    page = Padding(
      padding: const EdgeInsets.all(10),
      child: Column(children: [
        const SizedBox(
          height: 20,
        ),
        TextField(
            onChanged: (value) {
              setState(() {
                _gSearchKey = value;
                gitProjects = runFilter(value, allProjects, selectedIndex);
              });
            },
            decoration: const InputDecoration(
                labelText: "search....", suffixIcon: Icon(Icons.search))),
        const SizedBox(
          height: 10,
        ),
        Expanded(
          child: gitProjects.isNotEmpty
              ? ListView.builder(
                  itemCount: gitProjects.length,
                  itemBuilder: (context, index) => Card(
                      key: ValueKey(gitProjects[index].id.toString()),
                      color: Colors.white,
                      elevation: 4,
                      margin: const EdgeInsets.symmetric(vertical: 5),
                      child: ListTile(
                          leading: Text(gitProjects[index].id.toString(),
                              style: const TextStyle(fontSize: 24)),
                          title: Text(gitProjects[index].name),
                          subtitle: Text(gitProjects[index].url),
                          trailing: IconButton(
                              onPressed: () {
                                setState(() {
                                  gitProjects[index].isSync =
                                      !gitProjects[index].isSync;
                                });
                              },
                              icon: gitProjects[index].isSync
                                  ? const Icon(Icons.sync_lock)
                                  : const Icon(Icons.sync)))),
                )
              : const Text(
                  "No result found",
                  style: TextStyle(fontSize: 24),
                ),
        )
      ]),
    );

    return LayoutBuilder(builder: (context, constraints) {
      return Scaffold(
        body: Row(children: [
          SafeArea(
              child: NavigationRail(
            extended: constraints.maxWidth >= extendMaxWidth,
            destinations: const [
              NavigationRailDestination(
                  icon: Icon(Icons.home), label: Text("ALL")),
              NavigationRailDestination(
                  icon: Icon(Icons.sync), label: Text("SYNC"))
            ],
            selectedIndex: selectedIndex,
            onDestinationSelected: (value) {
              print(value);
              setState(() {
                selectedIndex = value;
                gitProjects =
                    runFilter(_gSearchKey, allProjects, selectedIndex);
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
