import 'package:flutter/material.dart';
import 'package:learn3/main.dart';
import 'package:learn3/model/project.dart';

class MyHomePageState extends State<MyHomePage> {
  var _currentIndex = 0;
  List<GitProject> gitProjects = [];
  var nums = 5;

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
    return Scaffold(
      appBar: AppBar(
        title: const Text(""),
        //leading: const Icon(Icons.home),
        // actions: const [
        //   Icon(Icons.share),
        //   Icon(Icons.shopping_bag),
        // ],
      ),
      //body: ColorIndexContainer(currentIndex: _currentIndex),
      body: ListView.builder(
          itemCount: gitProjects.length,
          itemBuilder: (context, index) {
            return ListTile(title: Text(gitProjects[index].name));
          }),
      drawer: const UserDrawer(),
      bottomNavigationBar: BottomNavigationBar(
        currentIndex: _currentIndex,
        items: const [
          BottomNavigationBarItem(icon: Icon(Icons.home), label: "home"),
          BottomNavigationBarItem(icon: Icon(Icons.add), label: "add"),
          BottomNavigationBarItem(icon: Icon(Icons.person), label: "user")
        ],
        onTap: (index) {
          setState(() {
            _currentIndex = index;
          });
        },
      ),
    );
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
