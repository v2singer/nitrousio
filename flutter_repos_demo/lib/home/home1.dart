import 'package:flutter/material.dart';
import 'package:learn3/main.dart';

class MyHomePageState extends State<MyHomePage> {
  var _currentIndex = 0;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title),
        //leading: const Icon(Icons.home),
        actions: const [
          Icon(Icons.share),
          Icon(Icons.shopping_bag),
        ],
      ),
      //body: ColorIndexContainer(currentIndex: _currentIndex),
      body: ListView(children: const [
        ListTile(title: Text("name1")),
        ListTile(title: Text("name2"))
      ]),
      drawer: Drawer(
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
            ListTile(title: Text("帐号")),
          ],
        ),
      ),
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
